package testutil

import (
	"context"
	"math/big"
	"testing"

	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/contracts/getter"
	"github.com/tellor-io/TellorMiner/pkg/contracts/tellor"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/rpc"
)

func CreateContext(t *testing.T) (context.Context, *config.Config, func()) {
	cfg := config.OpenTestConfig(t)
	// Don't need any trackers for the tests.
	cfg.Trackers = make(map[string]bool)

	dbLocal, cleanup := db.OpenTestDB(t)
	ctx := context.WithValue(context.Background(), common.DBContextKey, dbLocal)

	hash := math.PaddedBigBytes(big.NewInt(256), 32)
	var b32 [32]byte
	for i, v := range hash {
		b32[i] = v
	}

	opts := &rpc.MockOptions{
		ETHBalance:    big.NewInt(300000),
		Nonce:         1,
		GasPrice:      big.NewInt(7000000000),
		TokenBalance:  big.NewInt(0),
		Top50Requests: []*big.Int{},
		CurrentChallenge: &rpc.CurrentChallenge{
			ChallengeHash: b32, RequestID: big.NewInt(1),
			Difficulty: big.NewInt(500), QueryString: "json(https://coinbase.com)",
			Granularity: big.NewInt(1000), Tip: big.NewInt(0),
		},
	}
	client := rpc.NewMockClientWithValues(opts)
	ctx = context.WithValue(ctx, common.ClientContextKey, client)

	contractAddress := eth_common.HexToAddress(cfg.ContractAddress)
	master, err := tellor.NewTellor(contractAddress, client)
	if err != nil {
		t.Fatalf("Problem creating tellor master instance: %v\n", err)
	}
	ctx = context.WithValue(ctx, common.ContractAddress, contractAddress)
	ctx = context.WithValue(ctx, common.ContractsTellorContextKey, master)

	instanceTellor, err := tellor.NewTellor(contractAddress, client)
	if err != nil {
		t.Fatalf("Problem creating tellor master instance: %v\n", err)
	}
	ctx = context.WithValue(ctx, common.ContractsTellorContextKey, instanceTellor)

	instanceGetter, err := getter.NewTellorGetters(contractAddress, client)
	if err != nil {
		t.Fatalf("Problem creating tellor master instance: %v\n", err)
	}
	ctx = context.WithValue(ctx, common.ContractsGetterContextKey, instanceGetter)

	proxy, err := db.OpenLocalProxy(dbLocal)
	if err != nil {
		t.Fatalf("Problem creating proxy: %v\n", err)
	}
	ctx = context.WithValue(ctx, common.DataProxyKey, proxy)

	return ctx, cfg, cleanup
}
