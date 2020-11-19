// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/tellor-io/TellorMiner/pkg/tcontext"
	"github.com/tellor-io/TellorMiner/pkg/util"
)

func TestDisputeCheckerInRange(t *testing.T) {
	logSetup := util.SetupLogger()
	logger := logSetup("debug")
	ctx, _, cleanup := tcontext.CreateTestContext(t)
	t.Cleanup(cleanup)

	if _, err := BuildIndexTrackers(); err != nil {
		t.Fatal(err)
	}
	ethUSDPairs := indexes["ETH/USD"]
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	time.Sleep(2 * time.Second)
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	disputeChecker := &disputeChecker{lastCheckedBlock: 500, logger: logger}
	err := disputeChecker.Exec(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDisputeCheckerOutOfRange(t *testing.T) {
	ctx, cfg, cleanup := tcontext.CreateTestContext(t)
	t.Cleanup(cleanup)
	cfg.DisputeThreshold = 0.000000001
	logSetup := util.SetupLogger()
	logger := logSetup("debug")
	disputeChecker := NewDisputeChecker(logger, 500)
	if _, err := BuildIndexTrackers(); err != nil {
		t.Fatal(err)
	}
	ethUSDPairs := indexes["ETH/USD"]
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	time.Sleep(2 * time.Second)
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	err := disputeChecker.Exec(ctx)

	if err != nil {
		t.Fatal(err)
	}

	files, err := filepath.Glob("possible-dispute-*.txt")
	if err != nil {
		panic(err)
	}
	if len(files) != 1 {
		t.Fatal("expected a possible-dispute file")
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			t.Fatal(err)

		}
	}
}

func execEthUsdPsrs(ctx context.Context, t *testing.T, psrs []*IndexTracker) {
	for _, psr := range psrs {
		err := psr.Exec(ctx)
		if err != nil {
			t.Fatalf("failed to execute psr:%v", err)
		}
	}
}
