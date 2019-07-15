package tracker

import (
	"context"
)

type GasTracker struct {
}

func (b *GasTracker) Exec(ctx context.Context) error {
	/*
		client, err := ctx.Get("ETHClient")
		if err != nil {
			log.Fatal(err)
			return err
		}
		DB := ctx.Get("DB")

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		DB.update("Gas", gasPrice)
		return nil
	*/
	return nil
}
