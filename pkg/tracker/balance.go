package tracker

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BalanceTracker struct {
}

func (b *BalanceTracker) Exec(ctx context.Context) (error) {
	client, err := ctx.Get("ETHClient")
		log.Fatal(err)
		return err
	}
	DB := ctx.Get("DB")
	fromAddress := common.HexToAddress(_fromAddress)
	balance, err := mikesClient.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	DB.update("Balance",balance)
	return nil
}
