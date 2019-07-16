package tracker

/*
import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	tellor "github.com/tellor-io/TellorMiner/contracts"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rpc"
)

type TributeTracker struct {
}

func (b *TributeTracker) Exec(ctx context.Context) error {
	//cast client using type assertion since context holds generic interface{}
	client := ctx.Value(ClientContextKey).(rpc.ETHClient)
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)

	//get the single config instance
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
		return err
	}

	//get address from config
	_fromAddress := cfg.PublicAddress

	instance, err := tellor.NewTellorMaster(contractAddress, client)
	if err != nil {
		log.Fatal(err)
		return err
	}

	balance, err := instance.GetBalance(fromAddress)
	if err != nil {
		log.Fatal(err)
		return err
	}
	enc := hexutil.EncodeBig(balance)
	log.Printf("Balance: %v", enc)
	return DB.Put(db.TributeBalanceKey, []byte(enc))

	return nil
}