package testutil

import (
	"context"
	"log"
	"math/big"
	"testing"

	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/tellor-io/TellorMiner/abi/contracts"
	"github.com/tellor-io/TellorMiner/abi/contracts2"
	"github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/rpc"
)

func CreateContext(t *testing.T) (context.Context, *config.Config, func()) {
	cfg := config.OpenTestConfig(t)
	// Don't need any trackers for the tests.
	cfg.Trackers = make(map[string]bool)

	dbLocal, cleanup := db.OpenTestDB(t)
	ctx := context.WithValue(context.Background(), common.DBContextKey, dbLocal)

	opts := &rpc.MockOptions{
		ETHBalance:    big.NewInt(300000),
		Nonce:         1,
		GasPrice:      big.NewInt(7000000000),
		TokenBalance:  big.NewInt(0),
		Top50Requests: []*big.Int{},
	}
	client := rpc.NewMockClientWithValues(opts)
	ctx = context.WithValue(ctx, common.ClientContextKey, client)

	contractAddress := eth_common.HexToAddress(cfg.ContractAddress)
	master, err := contracts.NewTellorMaster(contractAddress, client)
	if err != nil {
		t.Fatalf("Problem creating tellor master instance: %v\n", err)
	}
	ctx = context.WithValue(ctx, common.MasterContractContextKey, master)

	instance, err := contracts2.NewTellor(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	ctx = context.WithValue(ctx, common.NewTellorContractContextKey, instance)

	proxy, err := db.OpenLocalProxy(dbLocal)
	if err != nil {
		log.Fatal(err)
	}
	ctx = context.WithValue(ctx, common.DataProxyKey, proxy)

	return ctx, cfg, cleanup
}
