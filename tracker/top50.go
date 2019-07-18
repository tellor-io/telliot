package tracker

/*
import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	tellor "github.com/tellor-io/TellorMiner/contracts"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rpc"
)

//CurrentVariablesTracker concrete tracker type
type Top50Tracker struct {
}

//Exec implementation for tracker
func (b *Top50Tracker) Exec(ctx context.Context) error {

	//cast client using type assertion since context holds generic interface{}
	client := ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)

	//get the single config instance
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
		return err
	}

	contractAddress := common.HexToAddress(cfg.ContractAddress)
	instance, err := tellor.NewTellorMaster(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	top50, err := instance.GetRequestQ(nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	rIDs := []byte{}
	for i, _ := range top50 {
		rIDs = append(rIDs[:], top50[i].Bytes()[:]...)

	}
	return DB.Put(db.TotalTipKey, rIDs)
}
*/
