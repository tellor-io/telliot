package ops

import (
	"context"
	"os"
	"time"

	"github.com/tellor-io/TellorMiner/dataServer"
	"github.com/tellor-io/TellorMiner/util"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/ethereum/go-ethereum/common"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	tellor1 "github.com/tellor-io/TellorMiner/contracts1"
	tellor "github.com/tellor-io/TellorMiner/contracts"
)

/**
 * This is the operational deposit component. Its purpose is to deposit Tellor Tokens so you can mine
 */


 func Deposit(ctx context.Context) (error) {
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

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice


	contractAddress := common.HexToAddress(cfg.ContractAddress)

	instance, err := tellor.NewTellorMaster(contractAddress, client)
	if err != nil {
		log.Fatal(err)
		return err
	}

	balance, err := instance.BalanceOf(nil, _fromAddress)
	log.Printf("Balance: %v\n", balance)
	if err != nil {
		log.Fatal(err)
		return err
	}
	amt := big.NewInt(amount)
	if balance.Cmp(amt) < 0{
		fmt.PrintLn("You must have the amount you want to send")
		return nil
	}

	instance := ctx.Value(tellorCommon.TransactorContractContextKey).(*tellor1.TellorTransactor)

	tx, err := instance.DepositStake(auth)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())

	return nil
}