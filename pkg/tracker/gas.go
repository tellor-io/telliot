package tracker

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type GasTracker struct {
}

func (b *GasTracker) Exec(ctx context.Context) (error) {
	client, err := ctx.Get("ETHClient")
		log.Fatal(err)
		return err
	}
	DB := ctx.Get("DB")

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	DB.update("Gas",gasPrice)
	return nil
}
