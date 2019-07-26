package tracker

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	tellor "github.com/tellor-io/TellorMiner/contracts"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rpc"
)

//DisputeTracker struct
type DisputeTracker struct {
}

func (b *DisputeTracker) String() string {
	return "DisputeTracker"
}

//Exec - Places the Dispute Status in the database
func (b *DisputeTracker) Exec(ctx context.Context) error {
	//cast client using type assertion since context holds generic interface{}
	client := ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)

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

	_conAddress := cfg.ContractAddress

	//convert to address
	contractAddress := common.HexToAddress(_conAddress)

	instance, err := tellor.NewTellorMaster(contractAddress, client)
	if err != nil {
		log.Fatal(err)
		return err
	}

	status, _, err := instance.GetStakerInfo(nil, fromAddress)
	if err != nil {
		log.Fatal(err)
		return err
	}
	enc := hexutil.EncodeBig(status)
	log.Printf("Staker Status: %v", enc)
	return DB.Put(db.DisputeStatusKey, []byte(enc))
}
