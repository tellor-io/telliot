package tracker

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	tellor "github.com/tellor-io/TellorMiner/contracts"
	"github.com/tellor-io/TellorMiner/rpc"
)

type IDSpecifications struct {
	RequestID   uint   `json:"requestId"`
	QueryString string `json:"queryString"`
	Granularity int    `json:"granularity"`
}

func GetSpecs(ctx context.Context, requestID uint) (*IDSpecifications, error) {
	client := ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)

	//get the single config instance
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	contractAddress := common.HexToAddress(cfg.ContractAddress)
	instance, err := tellor.NewTellorMaster(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	queryString, _, _, granularity, _, _, err := instance.GetRequestVars(nil, big.NewInt(int64(requestID)))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &IDSpecifications{requestID, queryString, int(granularity.Uint64())}, nil
}
