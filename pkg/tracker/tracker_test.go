// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"

	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/testutil"
)

func TestCreateTracker(t *testing.T) {
	logger := logging.NewLogger()
	cfg, err := config.OpenTestConfig("../..")
	testutil.Ok(t, err)
	DB, cleanup, err := db.OpenTestDB(cfg)
	testutil.Ok(t, err)
	defer func() {
		testutil.Ok(t, cleanup())
	}()
	client := rpc.NewMockClient()
	proxy, err := db.OpenLocal(logger, cfg, DB)
	testutil.Ok(t, err)

	gasTracker, _ := createTracker("gas", logger, cfg, proxy, client, nil)
	if gasTracker[0].String() != "GasTracker" {
		testutil.Ok(t, errors.Errorf("Expected GasTracker but got %s", gasTracker[0].String()))
	}

	indexersTracker, err := createTracker("indexers", logger, cfg, proxy, client, nil)
	testutil.Ok(t, err, "build IndexTracker")
	if len(indexersTracker) == 0 {
		testutil.Ok(t, errors.Errorf("build all IndexTrackers: only tracking %d indexes", len(indexersTracker)))
	}

	disputeChecker, _ := createTracker("disputeChecker", logger, cfg, proxy, client, nil)
	if disputeChecker[0].String() != "DisputeChecker" {
		testutil.Ok(t, errors.Errorf("Expected DisputeChecker but got %s", disputeChecker[0].String()))
	}

	_, err = createTracker("badTracker", logger, cfg, proxy, client, nil)
	testutil.Assert(t, err != nil, "expected error but instead received tracker")

}

func TestRunner(t *testing.T) {
	cfg, err := config.OpenTestConfig("../../")
	testutil.Ok(t, err)
	logger := logging.NewLogger()

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
		TokenBalance: big.NewInt(0), MiningStatus: true, Top50Requests: top50, CurrentChallenge: chal, QueryMetadata: paramsMap}
	client := rpc.NewMockClientWithValues(opts)

	DB, cleanup, err := db.OpenTestDB(cfg)
	testutil.Ok(t, err)
	defer func() {
		testutil.Ok(t, cleanup())
	}()

	proxy, err := db.OpenLocal(logger, cfg, DB)
	testutil.Ok(t, err)
	contract, err := contracts.NewITellor(client)
	testutil.Ok(t, err)
	testutil.Ok(t, err)

	runner, _ := NewRunner(logger, cfg, proxy, client, contract)

	runner.Ready()
	if err := runner.Start(context.Background()); err != nil {
		testutil.Ok(t, err)
	}
	level.Info(logger).Log("msg", "runner done")
	time.Sleep(2 * time.Second)
	close(exitCh)
	time.Sleep(1 * time.Second)
}
