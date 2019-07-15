package tracker

import (
	"context"
)

type TributeTracker struct {
}

func (b *TributeTracker) Exec(ctx context.Context) error {
	/*
		client := ctx.Value("ETHClient")
		DB := ctx.Value("DB")
		fromAddress := common.HexToAddress(_fromAddress)

		instance, err := tellor.NewTellorMaster(contractAddress, client)
		if err != nil {
			log.Fatal(err)
		}

		balance, err := instance.GetBalance(fromAddress)
		if err != nil {
			log.Fatal(err)
		}
		DB.update("TributeBalance", balance)
		return nil
	*/
	return nil
}
