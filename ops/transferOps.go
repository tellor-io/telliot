package ops

import (
	"context"
	"math/big"
	"fmt"
	"crypto/ecdsa"
	"log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/tellor-io/TellorMiner/rpc"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/ethereum/go-ethereum/common"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	tellor1 "github.com/tellor-io/TellorMiner/contracts1"
	tellor "github.com/tellor-io/TellorMiner/contracts"
)

/**
 * This is the operational transfer component. Its purpose is to transfer tellor tokens
 */

func Transfer(toAddress string, amount string,ctx context.Context) (error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}
	client := ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)

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
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		return err
	}

	cost := new(big.Int)
	cost.Mul(gasPrice, big.NewInt(700000))
	if balance.Cmp(cost) < 0 {
		//FIXME: notify someone that we're out of funds!
		return fmt.Errorf("Insufficient funds to send transaction: %v < %v", balance, cost)
	}
	fmt.Println(nonce)
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(nonce)
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice.Mul(gasPrice,big.NewInt(3))
	_fromAddress := cfg.PublicAddress


	contractAddress := common.HexToAddress(cfg.ContractAddress)
	toAdd := common.HexToAddress(toAddress)
	instance, err := tellor.NewTellorMaster(contractAddress, client)
	if err != nil {
		log.Fatal(err)
		return err
	}

	balance, err = instance.BalanceOf(nil, common.HexToAddress(_fromAddress))
	log.Printf("Balance: %v\n", balance)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("My balance",balance)
	amt := new(big.Int)
	amt,_ = amt.SetString(amount, 10)
	if balance.Cmp(amt) < 0{
		fmt.Println("You must have the amount you want to send")
		return nil
	}
	instance2 := ctx.Value(tellorCommon.TransactorContractContextKey).(*tellor1.TellorTransactor)

	tx, err := instance2.Transfer(auth, toAdd,amt)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())

	return nil
}