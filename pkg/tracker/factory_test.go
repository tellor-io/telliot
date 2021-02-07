// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"testing"

	"github.com/pkg/errors"

	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/testutil"
)

func TestCreateTracker(t *testing.T) {

	logger := logging.NewLogger()
	cfg := config.OpenTestConfig(t)
	DB, cleanup := db.OpenTestDB(t)
	defer t.Cleanup(cleanup)
	client := rpc.NewMockClient()
	proxy, err := db.OpenLocal(logger, cfg, DB)
	testutil.Ok(t, err)
	accounts := []*rpc.Account{{}}
	balanceTracker, _ := createTracker("balance", logger, cfg, proxy, client, nil, accounts)
	if balanceTracker[0].String() != BalanceTrackerName {
		testutil.Ok(t, errors.Errorf("Expected BalanceTracker but got %s", balanceTracker[0].String()))
	}

	currentVariablesTracker, _ := createTracker("currentVariables", logger, cfg, proxy, client, nil, accounts)
	if currentVariablesTracker[0].String() != "CurrentVariablesTracker" {
		testutil.Ok(t, errors.Errorf("Expected CurrentVariablesTracker but got %s", currentVariablesTracker[0].String()))
	}

	disputeStatusTracker, _ := createTracker("disputeStatus", logger, cfg, proxy, client, nil, accounts)
	if disputeStatusTracker[0].String() != DisputeTrackerName {
		testutil.Ok(t, errors.Errorf("Expected DisputeTracker but got %s", disputeStatusTracker[0].String()))
	}

	gasTracker, _ := createTracker("gas", logger, cfg, proxy, client, nil, accounts)
	if gasTracker[0].String() != "GasTracker" {
		testutil.Ok(t, errors.Errorf("Expected GasTracker but got %s", gasTracker[0].String()))
	}

	tributeBalanceTracker, _ := createTracker("tributeBalance", logger, cfg, proxy, client, nil, accounts)
	if tributeBalanceTracker[0].String() != "TributeTracker" {
		testutil.Ok(t, errors.Errorf("Expected TributeTracker but got %s", tributeBalanceTracker[0].String()))
	}

	indexersTracker, err := createTracker("indexers", logger, cfg, proxy, client, nil, accounts)
	testutil.Ok(t, err, "build IndexTracker")
	if len(indexersTracker) == 0 {
		testutil.Ok(t, errors.Errorf("build all IndexTrackers: only tracking %d indexes", len(indexersTracker)))
	}

	disputeChecker, _ := createTracker("disputeChecker", logger, cfg, proxy, client, nil, accounts)
	if disputeChecker[0].String() != "DisputeChecker" {
		testutil.Ok(t, errors.Errorf("Expected DisputeChecker but got %s", disputeChecker[0].String()))
	}

	_, err = createTracker("badTracker", logger, cfg, proxy, client, nil, accounts)
	testutil.Assert(t, err != nil, "expected error but instead received tracker")

}
