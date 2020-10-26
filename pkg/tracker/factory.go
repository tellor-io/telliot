// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"fmt"

	"github.com/go-kit/kit/log"
)

// CreateTracker a tracker instance by its well-known name.
func createTracker(name string, logger log.Logger) ([]Tracker, error) {
	switch name {
	case "timeOut":
		{
			return []Tracker{NewTimeOutTracker(logger)}, nil
		}
	case "balance":
		{
			return []Tracker{NewBalanceTracker(logger)}, nil
		}
	case "currentVariables":
		{
			return []Tracker{New_CurrentVariablesTracker(logger)}, nil
		}
	case "disputeStatus":
		{
			return []Tracker{NewDisputeTracker(logger)}, nil
		}
	case "gas":
		{
			return []Tracker{NewGasTracker(logger)}, nil
		}
	case "newCurrentVariables":
		{
			return []Tracker{NewNewCurrentVariablesTracker(logger)}, nil
		}
	case "tributeBalance":
		{
			return []Tracker{NewTributeTracker(logger)}, nil
		}
	case "indexers":
		{
			return BuildIndexTrackers()
		}
	case "disputeChecker":
		return []Tracker{NewDisputeChecker(logger)}, nil
	default:
		return nil, fmt.Errorf("no tracker with the name %s", name)
	}
}
