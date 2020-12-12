// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/tellor-io/telliot/pkg/tcontext"
	"github.com/tellor-io/telliot/pkg/testutil"
	"github.com/tellor-io/telliot/pkg/util"
)

func TestDisputeCheckerInRange(t *testing.T) {
	logSetup := util.SetupLogger()
	logger := logSetup("debug")
	ctx, _, cleanup := tcontext.CreateTestContext(t)
	t.Cleanup(cleanup)

	if _, err := BuildIndexTrackers(); err != nil {
		testutil.Ok(t, err)
	}
	ethUSDPairs := indexes["ETH/USD"]
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	time.Sleep(2 * time.Second)
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	disputeChecker := &disputeChecker{lastCheckedBlock: 500, logger: logger}
	testutil.Ok(t, disputeChecker.Exec(ctx))
}

func TestDisputeCheckerOutOfRange(t *testing.T) {
	ctx, cfg, cleanup := tcontext.CreateTestContext(t)
	t.Cleanup(cleanup)
	cfg.DisputeThreshold = 0.000000001
	logSetup := util.SetupLogger()
	logger := logSetup("debug")
	disputeChecker := NewDisputeChecker(logger, 500)
	if _, err := BuildIndexTrackers(); err != nil {
		testutil.Ok(t, err)
	}
	ethUSDPairs := indexes["ETH/USD"]
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
