package tracker

import "fmt"

//CreateTracker a tracker instance by its well-known name
func createTracker(name string) ([]Tracker, error) {
	switch name {
	case "timeOut":
		{
			return []Tracker{&TimeOutTracker{}}, nil
		}
	case "balance":
		{
			return []Tracker{&BalanceTracker{}}, nil
		}
	case "currentVariables":
		{
			return []Tracker{&CurrentVariablesTracker{}}, nil
		}
	case "disputeStatus":
		{
			return []Tracker{&DisputeTracker{}}, nil
		}
	case "gas":
		{
			return []Tracker{&GasTracker{}}, nil
		}	
	case "newCurrentVariables":
			{
				return []Tracker{&NewCurrentVariablesTracker{}}, nil
			}
	case "tributeBalance":
		{
			return []Tracker{&TributeTracker{}}, nil
		}
	case "indexers":
		{
			return BuildIndexTrackers()
		}
	case "disputeChecker":
		return []Tracker{&disputeChecker{}}, nil
	default:
		return nil, fmt.Errorf("no tracker with the name %s", name)
	}
}
