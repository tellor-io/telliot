package ops

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/contracts"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rpc"
)

func TestDataServerOps(t *testing.T) {
	exitCh := make(chan os.Signal)
	cfg := config.GetConfig()

	if len(cfg.DBFile) == 0 {
		log.Fatal("Missing dbFile config setting")
	}

	DB, err := db.Open(cfg.DBFile)
	if err != nil {
		log.Fatal(err)
	}
	client, err := rpc.NewClient(cfg.NodeURL)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(cfg.ContractAddress)
	masterInstance, err := contracts.NewTellorMaster(contractAddress, client)
	if err != nil {
		t.Fatalf("Problem creating tellor master instance: %v\n", err)
	}

	ctx := context.WithValue(context.Background(), tellorCommon.DBContextKey, DB)
	ctx = context.WithValue(ctx, tellorCommon.ClientContextKey, client)
	ctx = context.WithValue(ctx, tellorCommon.MasterContractContextKey, masterInstance)

	ops, err := CreateDataServerOps(ctx, exitCh)
	if err != nil {
		t.Fatal(err)
	}
	ops.Start(ctx)
	time.Sleep(2 * time.Second)
	exitCh <- os.Interrupt
	time.Sleep(1 * time.Second)
	if ops.Running {
		t.Fatal("data server is still running after stopping")
	}
}
