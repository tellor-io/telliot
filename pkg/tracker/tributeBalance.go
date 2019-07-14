package tracker

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TributeTracker struct {
}

func (b *TributeTracker) Exec(ctx context.Context) (error) {
	client, err := ctx.Get("ETHClient")
		log.Fatal(err)
		return err
	}
	DB := ctx.Get("DB")
	fromAddress := common.HexToAddress(_fromAddress)

	instance, err := tellor.NewTellorMaster(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	balance, err := instance.GetBalance(fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	DB.update("TributeBalance",balance)
	return nil
}
