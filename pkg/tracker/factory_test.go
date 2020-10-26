// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"testing"

	"github.com/tellor-io/TellorMiner/pkg/util"
)

func TestCreateTracker(t *testing.T) {

	logger := util.SetupLogger("debug")
	balanceTracker, _ := createTracker("balance", logger)
	if balanceTracker[0].String() != BalanceTrackerName {
		t.Fatalf("Expected BalanceTracker but got %s", balanceTracker[0].String())
	}

	currentVariablesTracker, _ := createTracker("currentVariables", logger)
	if currentVariablesTracker[0].String() != CurrentVariablesTrackerName {
		t.Fatalf("Expected CurrentVariablesTracker but got %s", currentVariablesTracker[0].String())
	}

	disputeStatusTracker, _ := createTracker("disputeStatus", logger)
	if disputeStatusTracker[0].String() != DisputeTrackerName {
		t.Fatalf("Expected DisputeTracker but got %s", disputeStatusTracker[0].String())
	}

	gasTracker, _ := createTracker("gas", logger)
	if gasTracker[0].String() != "GasTracker" {
		t.Fatalf("Expected GasTracker but got %s", gasTracker[0].String())
	}

	tributeBalanceTracker, _ := createTracker("tributeBalance", logger)
	if tributeBalanceTracker[0].String() != "TributeTracker" {
		t.Fatalf("Expected TributeTracker but got %s", tributeBalanceTracker[0].String())
	}

	indexersTracker, err := createTracker("indexers", logger)
	if err != nil {
		t.Fatalf("Could not build IndexTracker")
	}
	if len(indexersTracker) == 0 {
		t.Fatalf("Could not build all IndexTrackers: only tracking %d indexes", len(indexersTracker))
	}

	disputeChecker, _ := createTracker("disputeChecker", logger)
	if disputeChecker[0].String() != "DisputeChecker" {
		t.Fatalf("Expected DisputeChecker but got %s", disputeChecker[0].String())
	}

	badTracker, err := createTracker("badTracker", logger)
	if err == nil {
		t.Fatalf("expected error but instead received this tracker: %s", badTracker[0].String())
	}

}
