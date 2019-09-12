package rpc

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/contracts"
	tellor1 "github.com/tellor-io/TellorMiner/contracts1"
)

//contractWrapper is internal wrapper of contract instance for calling common contract functions
type contractWrapper struct {
	contract    *tellor1.TellorTransactor
	contract2   *contracts.TellorMaster
	options     *bind.TransactOpts
	fromAddress common.Address
}

func (c contractWrapper) AddTip(requestID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return c.contract.AddTip(c.options, requestID, amount)
}

func (c contractWrapper) SubmitSolution(solution string, requestID *big.Int, value *big.Int) (*types.Transaction, error) {
	return c.contract.SubmitMiningSolution(c.options, solution, requestID, value)
}

func (c contractWrapper) DidMine(challenge [32]byte) (bool, error) {
	return c.contract2.DidMine(nil, challenge, c.fromAddress)
}

func PrepareContractTxn(ctx context.Context, ctxName string, callback tellorCommon.TransactionGeneratorFN) error {

	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}
	client := ctx.Value(tellorCommon.ClientContextKey).(ETHClient)

	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
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
		return err
	}

	i := 0
	IntNonce := int64(nonce)
	for i < 5 {
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			return err
		}

		balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
		if err != nil {
			return err
		}

		cost := new(big.Int)
		cost.Mul(gasPrice, big.NewInt(200000))
		if balance.Cmp(cost) < 0 {
			//FIXME: notify someone that we're out of funds!
			return fmt.Errorf("Insufficient funds to send transaction: %v < %v", balance, cost)
		}

		auth := bind.NewKeyedTransactor(privateKey)
		auth.Nonce = big.NewInt(IntNonce)
		auth.Value = big.NewInt(0)     // in weiF
		auth.GasLimit = uint64(200000) // in units
		gasPrice1 := big.NewInt(0)
		gasPrice1.Mul(gasPrice, big.NewInt(int64(i*11)))
		gasPrice1.Div(gasPrice1, big.NewInt(int64(100)))
		auth.GasPrice = gasPrice.Add(gasPrice, gasPrice1)
		mult := cfg.GasMax
		res := big.NewInt(1000000000)
		if mult > 0 {
			res = res.Mul(res, big.NewInt(int64(mult)))
		} else {
			res = res.Mul(res, big.NewInt(int64(100)))
		}
		if auth.GasPrice.Cmp(res) < 0 {
			//create a wrapper to callback the actual txn generator fn
			instance := ctx.Value(tellorCommon.TransactorContractContextKey).(*tellor1.TellorTransactor)
			instance2 := ctx.Value(tellorCommon.MasterContractContextKey).(*contracts.TellorMaster)

			wrapper := contractWrapper{options: auth, contract: instance, contract2: instance2, fromAddress: fromAddress}
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
					fmt.Printf("%s tx sent: %s", ctxName, tx.Hash().Hex())
				}

				return nil
			}
		} else {
			fmt.Printf("%s Gas Prices Too high!!!\n", ctxName)
		}
		//wait a bit and try again
		time.Sleep(5 * time.Second)
		i++
	}
	fmt.Printf("%s Could not submit txn after 5 attempts\n", ctxName)
	return nil
}
