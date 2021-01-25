// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/testutil"
	"github.com/tellor-io/telliot/pkg/util"
)

func TestRunner(t *testing.T) {
	cfg := config.OpenTestConfig(t)
	logger := util.SetupLogger("debug")

	exitCh := make(chan int)

	startBal := big.NewInt(356000)

	hash := math.PaddedBigBytes(big.NewInt(256), 32)
	var b32 [32]byte
	for i, v := range hash {
		b32[i] = v
	}

	top50 := make([]*big.Int, 51)
	mockQueryParams := &rpc.MockQueryMeta{QueryString: "json(https://api.binance.com/api/v1/klines?symbol=ETHBTC&interval=1d&limit=1).0.4", Granularity: 1000}
	paramsMap := make(map[uint]*rpc.MockQueryMeta)
	for i := range top50 {
		top50[i] = big.NewInt(int64(i + 51))
		paramsMap[uint(i+51)] = mockQueryParams
	}

	queryStr := "json(https://coinbase.com)"
	chal := &rpc.CurrentChallenge{ChallengeHash: b32, RequestID: big.NewInt(1),
		Difficulty: big.NewInt(500), QueryString: queryStr,
		Granularity: big.NewInt(1000), Tip: big.NewInt(0)}
	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: big.NewInt(0), MiningStatus: true, Top50Requests: top50, CurrentChallenge: chal, DisputeStatus: big.NewInt(1), QueryMetadata: paramsMap}
	client := rpc.NewMockClientWithValues(opts)

	DB, cleanup := db.OpenTestDB(t)
	defer t.Cleanup(cleanup)
	proxy, err := db.OpenLocal(cfg, DB)
	testutil.Ok(t, err)
	contract, err := contracts.NewTellor(client)
	testutil.Ok(t, err)
	account, err := rpc.NewAccount(cfg)
	testutil.Ok(t, err)

	runner, _ := NewRunner(logger, cfg, proxy, client, &contract, &account)

	runner.Ready()
	if err := runner.Start(context.Background(), exitCh); err != nil {
		testutil.Ok(t, err)
	}
	fmt.Println("runner done")
	time.Sleep(2 * time.Second)
	close(exitCh)
	time.Sleep(1 * time.Second)
}
