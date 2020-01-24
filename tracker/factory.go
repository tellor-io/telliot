package tracker

import "fmt"

//CreateTracker a tracker instance by its well-known name
func createTracker(name string) (Tracker, error) {
	switch name {
	case "test":
		{
			return &TestTracker{}, nil
		}
	case "balance":
		{
			return &BalanceTracker{}, nil
		}
	case "currentVariables":
		{
			return &CurrentVariablesTracker{}, nil
		}
	case "disputeStatus":
		{
			return &DisputeTracker{}, nil
		}
	case "gas":
		{
			return &GasTracker{}, nil
		}
	case "top50":
		{
			return &Top50Tracker{}, nil
		}
	case "tributeBalance":
		{
			return &TributeTracker{}, nil
		}
	case "fetchData":
		{
			return &RequestDataTracker{}, nil
		}
	case "psr":
		{
			return BuildPSRTracker()
		}
	case "disputeChecker":
		return &disputeChecker{}, nil
	default:
		return nil, fmt.Errorf("no tracker with the name %s", name)
	}
	return nil, nil
}
