package tracker

import (
	"testing"
)

func TestCreateTracker(t *testing.T) {
	testTracker, _ := createTracker("test")
	if testTracker[0].String() != "TestTracker" {
		t.Fatalf("Expected TestTracker but got %s", testTracker[0].String())
	}

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

	//indexersTracker, _ := createTracker("indexers")
	//if indexersTracker[0].String() != "IndexersTracker" {
	//	t.Fatalf("Expected IndexersTracker but got %s", indexersTracker[0].String())
	//}

	disputeChecker, _ := createTracker("disputeChecker")
	if disputeChecker[0].String() != "DisputeChecker" {
		t.Fatalf("Expected DisputeChecker but got %s", disputeChecker[0].String())
	}

	badTracker, err := createTracker("badTracker")
	if err == nil {
		t.Fatalf("expected error but instead received this tracker: %s", badTracker[0].String())
	}

}
