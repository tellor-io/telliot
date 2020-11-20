// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"testing"

	"github.com/pkg/errors"

	"github.com/tellor-io/TellorMiner/pkg/testutil"
	"github.com/tellor-io/TellorMiner/pkg/util"
)

func TestCreateTracker(t *testing.T) {

	logSetup := util.SetupLogger()
	logger := logSetup("debug")
	balanceTracker, _ := createTracker("balance", logger)
	if balanceTracker[0].String() != BalanceTrackerName {
		testutil.Ok(t, errors.Errorf("Expected BalanceTracker but got %s", balanceTracker[0].String()))
	}

	currentVariablesTracker, _ := createTracker("currentVariables", logger)
	if currentVariablesTracker[0].String() != "CurrentVariablesTracker" {
		testutil.Ok(t, errors.Errorf("Expected CurrentVariablesTracker but got %s", currentVariablesTracker[0].String()))
	}

	disputeStatusTracker, _ := createTracker("disputeStatus", logger)
	if disputeStatusTracker[0].String() != DisputeTrackerName {
		testutil.Ok(t, errors.Errorf("Expected DisputeTracker but got %s", disputeStatusTracker[0].String()))
	}

	gasTracker, _ := createTracker("gas", logger)
	if gasTracker[0].String() != "GasTracker" {
		testutil.Ok(t, errors.Errorf("Expected GasTracker but got %s", gasTracker[0].String()))
	}

	tributeBalanceTracker, _ := createTracker("tributeBalance", logger)
	if tributeBalanceTracker[0].String() != "TributeTracker" {
		testutil.Ok(t, errors.Errorf("Expected TributeTracker but got %s", tributeBalanceTracker[0].String()))
	}

	indexersTracker, err := createTracker("indexers", logger)
	testutil.Ok(t, err, "Could not build IndexTracker")
	if len(indexersTracker) == 0 {
		testutil.Ok(t, errors.Errorf("Could not build all IndexTrackers: only tracking %d indexes", len(indexersTracker)))
	}

	disputeChecker, _ := createTracker("disputeChecker", logger)
	if disputeChecker[0].String() != "DisputeChecker" {
		testutil.Ok(t, errors.Errorf("Expected DisputeChecker but got %s", disputeChecker[0].String()))
	}

	_, err = createTracker("badTracker", logger)
	testutil.Assert(t, err != nil, "expected error but instead received tracker")

}
