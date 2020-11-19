// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"errors"
	"fmt"
	"testing"

	"github.com/tellor-io/TellorMiner/pkg/testutil"
	"github.com/tellor-io/TellorMiner/pkg/util"
)

func TestCreateTracker(t *testing.T) {

	logSetup := util.SetupLogger()
	logger := logSetup("debug")
	balanceTracker, _ := createTracker("balance", logger)
	if balanceTracker[0].String() != BalanceTrackerName {
		testutil.Ok(t, errors.New(fmt.Sprintf("Expected BalanceTracker but got %s", balanceTracker[0].String())))
	}

	currentVariablesTracker, _ := createTracker("currentVariables", logger)
	if currentVariablesTracker[0].String() != "CurrentVariablesTracker" {
		testutil.Ok(t, errors.New(fmt.Sprintf("Expected CurrentVariablesTracker but got %s", currentVariablesTracker[0].String())))
	}

	disputeStatusTracker, _ := createTracker("disputeStatus", logger)
	if disputeStatusTracker[0].String() != DisputeTrackerName {
		testutil.Ok(t, errors.New(fmt.Sprintf("Expected DisputeTracker but got %s", disputeStatusTracker[0].String())))
	}

	gasTracker, _ := createTracker("gas", logger)
	if gasTracker[0].String() != "GasTracker" {
		testutil.Ok(t, errors.New(fmt.Sprintf("Expected GasTracker but got %s", gasTracker[0].String())))
	}

	tributeBalanceTracker, _ := createTracker("tributeBalance", logger)
	if tributeBalanceTracker[0].String() != "TributeTracker" {
		testutil.Ok(t, errors.New(fmt.Sprintf("Expected TributeTracker but got %s", tributeBalanceTracker[0].String())))
	}

	indexersTracker, err := createTracker("indexers", logger)
	testutil.Ok(t, err, "Could not build IndexTracker")
	if len(indexersTracker) == 0 {
		testutil.Ok(t, errors.New(fmt.Sprintf("Could not build all IndexTrackers: only tracking %d indexes", len(indexersTracker))))
	}

	disputeChecker, _ := createTracker("disputeChecker", logger)
	if disputeChecker[0].String() != "DisputeChecker" {
		testutil.Ok(t, errors.New(fmt.Sprintf("Expected DisputeChecker but got %s", disputeChecker[0].String())))
	}

	badTracker, err := createTracker("badTracker", logger)
	testutil.Ok(t, err, "expected error but instead received this tracker: %s", badTracker[0].String())

}
