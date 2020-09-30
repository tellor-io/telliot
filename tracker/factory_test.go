package tracker

import (
	"testing"
)

func TestCreateTracker(t *testing.T) {

	balanceTracker, _ := createTracker("balance")
	if balanceTracker[0].String() != "BalanceTracker" {
		t.Fatalf("Expected BalanceTracker but got %s", balanceTracker[0].String())
	}

	currentVariablesTracker, _ := createTracker("currentVariables")
	if currentVariablesTracker[0].String() != "CurrentVariablesTracker" {
		t.Fatalf("Expected CurrentVariablesTracker but got %s", currentVariablesTracker[0].String())
	}

	disputeStatusTracker, _ := createTracker("disputeStatus")
	if disputeStatusTracker[0].String() != "DisputeTracker" {
		t.Fatalf("Expected DisputeTracker but got %s", disputeStatusTracker[0].String())
	}

	gasTracker, _ := createTracker("gas")
	if gasTracker[0].String() != "GasTracker" {
		t.Fatalf("Expected GasTracker but got %s", gasTracker[0].String())
	}

	tributeBalanceTracker, _ := createTracker("tributeBalance")
	if tributeBalanceTracker[0].String() != "TributeTracker" {
		t.Fatalf("Expected TributeTracker but got %s", tributeBalanceTracker[0].String())
	}

	indexersTracker, err := createTracker("indexers")
	if err != nil {
		t.Fatalf("Could not build IndexTracker")
	}
	if len(indexersTracker) == 0 {
		t.Fatalf("Could not build all IndexTrackers: only tracking %d indexes", len(indexersTracker))
	}

	disputeChecker, _ := createTracker("disputeChecker")
	if disputeChecker[0].String() != "DisputeChecker" {
		t.Fatalf("Expected DisputeChecker but got %s", disputeChecker[0].String())
	}

	badTracker, err := createTracker("badTracker")
	if err == nil {
		t.Fatalf("expected error but instead received this tracker: %s", badTracker[0].String())
	}

}
