// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/testutil"
	"github.com/tellor-io/telliot/pkg/util"
)

func TestDisputeCheckerInRange(t *testing.T) {
	cfg := config.OpenTestConfig(t)
	logger := util.SetupLogger("debug")
	DB, cleanup := db.OpenTestDB(t)
	client := rpc.NewMockClient()
	defer t.Cleanup(cleanup)
	proxy, err := db.OpenLocal(cfg, DB)
	testutil.Ok(t, err)

	if _, err := BuildIndexTrackers(cfg, proxy, client); err != nil {
		testutil.Ok(t, err)
	}
	contract, err := contracts.NewTellor(client)
	testutil.Ok(t, err)
	ctx := context.Background()
	ethUSDPairs := indexes["ETH/USD"]
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	time.Sleep(2 * time.Second)
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	disputeChecker := &disputeChecker{lastCheckedBlock: 500, config: cfg, logger: logger, client: client, contract: &contract}
	testutil.Ok(t, disputeChecker.Exec(ctx))
}

func TestDisputeCheckerOutOfRange(t *testing.T) {
	cfg := config.OpenTestConfig(t)
	logger := util.SetupLogger("debug")
	client := rpc.NewMockClient()
	contract, err := contracts.NewTellor(client)
	testutil.Ok(t, err)
	DB, cleanup := db.OpenTestDB(t)
	defer t.Cleanup(cleanup)
	proxy, err := db.OpenLocal(cfg, DB)
	testutil.Ok(t, err)
	disputeChecker := NewDisputeChecker(logger, cfg, client, &contract, 500)
	if _, err := BuildIndexTrackers(cfg, proxy, client); err != nil {
		testutil.Ok(t, err)
	}
	ethUSDPairs := indexes["ETH/USD"]
	ctx := context.Background()
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	time.Sleep(2 * time.Second)
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	testutil.Ok(t, disputeChecker.Exec(ctx))

	files, err := filepath.Glob("possible-dispute-*.txt")
	if err != nil {
		panic(err)
	}
	testutil.Assert(t, len(files) >= 1, "expected a possible-dispute file")

	for _, f := range files {
		testutil.Ok(t, os.Remove(f))
	}
}

func execEthUsdPsrs(ctx context.Context, t *testing.T, psrs []*IndexTracker) {
	for _, psr := range psrs {
		err := psr.Exec(ctx)

		testutil.Ok(t, err)

	}
}
