package tracker

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/rpc"
)

//BalanceTracker concrete tracker type
type BalanceTracker struct {
}

//Exec implementation for tracker
func (b *BalanceTracker) Exec(ctx context.Context) error {

	//cast client using type assertion since context holds generic interface{}
	client := ctx.Value("ETHClient").(rpc.ETHClient)

	//do this later
	//DB := ctx.Value("DB")

	//get the single config instance
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
		return err
	}

	//get address from config
	_fromAddress := cfg.PublicAddress

	//convert to address
	fromAddress := common.HexToAddress(_fromAddress)

	balance, err := client.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("Balance: %v", balance)
	//DB.update("Balance", balance)
	return nil
}
