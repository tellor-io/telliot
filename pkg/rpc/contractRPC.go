// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package rpc

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/contracts/getter"
	"github.com/tellor-io/TellorMiner/pkg/contracts/tellor"
	"github.com/tellor-io/TellorMiner/pkg/db"
)

// contractWrapper is internal wrapper of contract instance for calling common contract functions.
type contractWrapper struct {
	options     *bind.TransactOpts
	fromAddress common.Address

	*tellor.Tellor
	*getter.TellorGetters
}

func (c contractWrapper) AddTip(requestID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return c.Tellor.AddTip(c.options, requestID, amount)
}

func (c contractWrapper) SubmitSolution(solution string, requestID [5]*big.Int, value [5]*big.Int) (*types.Transaction, error) {
	return c.Tellor.SubmitMiningSolution(c.options, solution, requestID, value)
}

func (c contractWrapper) DidMine(challenge [32]byte) (bool, error) {
	return c.TellorGetters.DidMine(nil, challenge, c.fromAddress)
}

func SubmitContractTxn(ctx context.Context, proxy db.DataServerProxy, ctxName string, callback tellorCommon.TransactionGeneratorFN) (*types.Transaction, error) {

	cfg := config.GetConfig()
	client := ctx.Value(tellorCommon.ClientContextKey).(ETHClient)

	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "decoding private key")
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.NonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, errors.Wrap(err, "getting nonce for miner address")
	}
	IntNonce := int64(nonce)
	keys := []string{
		db.GasKey,
	}
	m, err := proxy.BatchGet(keys)
	if err != nil {
		return nil, errors.Wrap(err, "getting data from the db")
	}
	gasPrice := getInt(m[db.GasKey])
	if gasPrice.Cmp(big.NewInt(0)) == 0 {
		fmt.Println("Missing gas price from DB, falling back to client suggested gas price")
		gasPrice, err = client.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, errors.Wrap(err, "determine gas price to submit txn")
		}
	}
	mul := cfg.GasMultiplier
	if mul > 0 {
		fmt.Println("using gas multiplier : ", mul)
		gasPrice = gasPrice.Mul(gasPrice, big.NewInt(int64(mul)))
	}

	var finalError error
	for i := 0; i <= 5; i++ {
		balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
		if err != nil {
			finalError = err
			continue
		}

		cost := big.NewInt(1)
		cost = cost.Mul(gasPrice, big.NewInt(200000))
		if balance.Cmp(cost) < 0 {
			// FIXME: notify someone that we're out of funds!
			finalError = errors.Wrapf(err, "Insufficient funds to send transaction: %v < %v", balance, cost)
			continue
		}

		auth := bind.NewKeyedTransactor(privateKey)
		auth.Nonce = big.NewInt(IntNonce)
		auth.Value = big.NewInt(0)      // in weiF
		auth.GasLimit = uint64(3000000) // in units
		if gasPrice.Cmp(big.NewInt(0)) == 0 {
			gasPrice = big.NewInt(100)
		}
		if i > 1 {
			gasPrice1 := new(big.Int).Set(gasPrice)
			gasPrice1.Mul(gasPrice1, big.NewInt(int64(i*11))).Div(gasPrice1, big.NewInt(int64(100)))
			auth.GasPrice = gasPrice1.Add(gasPrice, gasPrice1)
		} else {
			// First time, try base gas price.
			auth.GasPrice = gasPrice
		}
		max := cfg.GasMax
		var maxGasPrice *big.Int
		gasPrice1 := big.NewInt(tellorCommon.GWEI)
		if max > 0 {
			maxGasPrice = gasPrice1.Mul(gasPrice1, big.NewInt(int64(max)))
		} else {
			maxGasPrice = gasPrice1.Mul(gasPrice1, big.NewInt(int64(100)))
		}

		if auth.GasPrice.Cmp(maxGasPrice) > 0 {
			fmt.Printf("%s Gas Prices Too high! Attempted gas price: %v is higher than max: %v.  Will default to max\n", ctxName, auth.GasPrice, maxGasPrice)
			auth.GasPrice = maxGasPrice
		}

		fmt.Println("Using gas price", gasPrice)
		// Create a wrapper to callback the actual txn generator fn.
		instanceTellor := ctx.Value(tellorCommon.ContractsTellorContextKey).(*tellor.Tellor)
		instanceGetter := ctx.Value(tellorCommon.ContractsGetterContextKey).(*getter.TellorGetters)

		wrapper := contractWrapper{auth, fromAddress, instanceTellor, instanceGetter}
		tx, err := callback(ctx, wrapper)

		if err != nil {
			if errors.Is(err, core.ErrNonceTooLow) {
				IntNonce = IntNonce + 1
			} else if errors.Is(err, core.ErrReplaceUnderpriced) {
				finalError = err
				continue
			} else {
				fmt.Println("Unspecified Request Data  Error ", err)
				finalError = err
				continue
			}
		}

		if tx != nil {
			return tx, nil
		}

		time.Sleep(15 * time.Second)
	}

	return nil, errors.Wrapf(finalError, "could not submit txn after 5 attempts ctx:%v", ctxName)
}

func getInt(data []byte) *big.Int {
	if len(data) == 0 {
		return nil
	}

	val, err := hexutil.DecodeBig(string(data))
	if err != nil {
		return nil
	}
	return val
}

func Keccak256(input []byte) [32]byte {
	hash := crypto.Keccak256(input)
	var hashed [32]byte
	copy(hashed[:], hash)

	return hashed
}
