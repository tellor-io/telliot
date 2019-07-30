package util

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

type IDSpecifications struct {
	requestID   uint
	queryString string
	granularity int
}

func getSpecs(ctx context.Context, requestID uint) IDSpecifications {
	var thisId IDSpecifications
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
	queryString, _, _, granularity, _, err := instance.GetRequestVars(requestId, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	thisId.requestID = requestID
	thisId.queryString = queryString
	thisId.granularity = granularity

	return thisID
}
