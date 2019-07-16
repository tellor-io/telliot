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

//CurrentVariablesTracker concrete tracker type
type CurrentVariablesTracker struct {
}

//Exec implementation for tracker
func (b *CurrentVariablesTracker) Exec(ctx context.Context) error {

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

	currentChallenge, requestID, difficulty, queryString, granularity, totalTip, err := instance.GetCurrentVariables(nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	DB.Put(db.CurrentChallengeKey, currentChallenge[:])
	DB.Put(db.RequestIdKey, []byte(hexutil.EncodeBig(requestID)))
	DB.Put(db.DifficultyKey, []byte(hexutil.EncodeBig(difficulty)))
	DB.Put(db.QueryStringKey, []byte(queryString))
	DB.Put(db.GranularityKey, []byte(hexutil.EncodeBig(granularity)))
	DB.Put(db.TotalTipKey, []byte(hexutil.EncodeBig(totalTip)))
	return nil
}
