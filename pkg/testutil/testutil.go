package testutil

import (
	"context"
	"log"
	"path/filepath"
	"testing"

	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/contracts"
	"github.com/tellor-io/TellorMiner/contracts2"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rpc"
)

func CreateContext(t *testing.T) (context.Context, *config.Config, func()) {
	cfg := config.OpenTestConfig(t)
	cfg.IndexFolder = filepath.Join("..")
	// Don't need any trackers for the tests.
	cfg.Trackers = make(map[string]bool)

	dbLocal, cleanup := db.OpenTestDB(t)
	ctx := context.WithValue(context.Background(), common.DBContextKey, dbLocal)

	client, err := rpc.NewClient(cfg.NodeURL)
	if err != nil {
		log.Fatal(err)
	}
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
