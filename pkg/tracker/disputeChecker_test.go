// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/testutil"
)

func TestDisputeCheckerInRange(t *testing.T) {
	ctx, _, cleanup := testutil.CreateContext(t)
	t.Cleanup(cleanup)

	if _, err := BuildIndexTrackers(); err != nil {
		t.Fatal(err)
	}
	logger := testutil.SetupLogger()
	ethUSDPairs := indexes["ETH/USD"]
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	time.Sleep(2 * time.Second)
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	ctx = context.WithValue(ctx, tellorCommon.ContractAddress, common.Address{0x0000000000000000000000000000000000000000})
	disputeChecker := &disputeChecker{lastCheckedBlock: 500}
	err := disputeChecker.Exec(ctx, logger)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDisputeCheckerOutOfRange(t *testing.T) {
	ctx, cfg, cleanup := testutil.CreateContext(t)
	logger := testutil.SetupLogger()
	t.Cleanup(cleanup)
	cfg.DisputeThreshold = 0.000000001
	disputeChecker := &disputeChecker{lastCheckedBlock: 500}
	if _, err := BuildIndexTrackers(); err != nil {
		t.Fatal(err)
	}
	ethUSDPairs := indexes["ETH/USD"]
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	time.Sleep(2 * time.Second)
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	ctx = context.WithValue(ctx, tellorCommon.ContractAddress, common.Address{0x0000000000000000000000000000000000000000})
	err := disputeChecker.Exec(ctx, logger)
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
	logger := testutil.SetupLogger()
	for _, psr := range psrs {
		err := psr.Exec(ctx, logger)
		if err != nil {
			t.Fatalf("failed to execute psr: %v", err)
		}
	}
}
