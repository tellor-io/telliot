// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package rpc

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/contracts/getter"
	"github.com/tellor-io/TellorMiner/pkg/contracts/tellor"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"golang.org/x/crypto/ripemd160"
)

var (
	GWEI = int64(1000000000)
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

func PrepareContractTxn(ctx context.Context, proxy db.DataServerProxy, ctxName string, callback tellorCommon.TransactionGeneratorFN) error {

	cfg := config.GetConfig()
	client := ctx.Value(tellorCommon.ClientContextKey).(ETHClient)

	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		fmt.Println("Problem decoding private key", err)
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.NonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println("Problem getting nonce for miner address", err)
		return err
	}
	IntNonce := int64(nonce)
	i := 1
	keys := []string{
		db.GasKey,
	}
	m, err := proxy.BatchGet(keys)
	if err != nil {
		return nil
	}
	gasPrice := getInt(m[db.GasKey])
	if gasPrice.Cmp(big.NewInt(0)) == 0 {
		fmt.Println("Missing gas price from DB, falling back to client suggested gas price")
		gasPrice, err = client.SuggestGasPrice(context.Background())
		if err != nil {
			fmt.Println("Could not determine gas price to submit txn", err)
			return err
		}
	}
	mul := cfg.GasMultiplier
	if mul > 0 {
		fmt.Println("using gas multiplier : ", mul)
		gasPrice = gasPrice.Mul(gasPrice, big.NewInt(int64(mul)))
	}
	for i < 5 {
		if err != nil {
			return err
		}

		balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
		if err != nil {
			return err
		}

		cost := big.NewInt(1)
		cost = cost.Mul(gasPrice, big.NewInt(200000))
		if balance.Cmp(cost) < 0 {
			//FIXME: notify someone that we're out of funds!
			return fmt.Errorf("Insufficient funds to send transaction: %v < %v", balance, cost)
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
			//first time, try base gas price
			auth.GasPrice = gasPrice
		}
		max := cfg.GasMax
		var maxGasPrice *big.Int
		gasPrice1 := big.NewInt(GWEI)
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
		// Цcreate a wrapper to callback the actual txn generator fn.
		instanceTellor := ctx.Value(tellorCommon.ContractsTellorContextKey).(*tellor.Tellor)
		instanceGetter := ctx.Value(tellorCommon.ContractsGetterContextKey).(*getter.TellorGetters)

		wrapper := contractWrapper{auth, fromAddress, instanceTellor, instanceGetter}
		tx, err := callback(ctx, wrapper)

		if err != nil {
			if strings.Contains(err.Error(), "nonce too low") {
				IntNonce = IntNonce + 1
			} else if strings.Contains(err.Error(), "replacement transaction underpriced") {
				fmt.Println("replacement transaction underpriced")
			} else {
				fmt.Println("Unspecified Request Data  Error ", err)
				return nil
			}
		} else {
			if tx != nil {
				fmt.Printf("%s tx sent: %s\n", ctxName, tx.Hash().Hex())
			}

			return nil
		}

		//wait a bit and try again
		time.Sleep(15 * time.Second)
		i++
	}
	fmt.Printf("%s Could not submit txn after 5 attempts\n", ctxName)
	return nil
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

func Keccak256(input) [32]byte {
	hash := solsha3.SoliditySHA3(
		// Types.
		[]string{"string"},
		// Values.
		[]interface{}{
			input,
		},
	)
	var hashed [32]byte
	copy(hashed[:], hash)
	return hashed
}

func HashFn(data []byte, result *big.Int) error {

	hash := solsha3.SoliditySHA3(data)

	// Consider moving hasher constructor outside loop and replacing with hasher.Reset()
	hasher := ripemd160.New()

	if _, err := hasher.Write(hash); err != nil {
		return err
	}
	hash1 := hasher.Sum(nil)
	n := sha256.Sum256(hash1)
	result.SetBytes(n[:])
	return nil
}
