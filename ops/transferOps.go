package ops

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	tellor "github.com/tellor-io/TellorMiner/contracts"
	tellor1 "github.com/tellor-io/TellorMiner/contracts1"
	"github.com/tellor-io/TellorMiner/util"
	"log"
	"math/big"
)

/**
 * This is the operational transfer component. Its purpose is to transfer tellor tokens
 */

func Transfer(toAddress string, amount string,ctx context.Context) error {

	instance := ctx.Value(tellorCommon.MasterContractContextKey).(tellor.TellorMaster)
	senderPubAddr := ctx.Value(tellorCommon.PublicAddress).(common.Address)

	balance, err := instance.BalanceOf(nil, senderPubAddr)
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

	auth, err := PrepareEthTransaction(ctx)
	if err != nil {
		return fmt.Errorf("failed to prepare ethereum transaction: %s", err.Error())
	}

	toAdd := common.HexToAddress(toAddress)
	tx, err := instance2.Transfer(auth, toAdd,amt)
	if err != nil {
		return fmt.Errorf("contract failed: %s", err.Error())
	}

	fmt.Printf("Transferred %s to %s... with tx: %s\n", util.FormatTRBBalance(amt), toAddress[:10], tx.Hash().Hex())

	return nil
}