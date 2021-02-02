// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
)

// CreateTracker a tracker instance by its well-known name.
func createTracker(name string, logger log.Logger, config *config.Config, db db.DataServerProxy, client contracts.ETHClient, contract *contracts.Tellor, account *rpc.Account) ([]Tracker, error) {
	switch name {
	case "timeOut":
		{
			return []Tracker{NewTimeOutTracker(logger, config, db, contract, account)}, nil
		}
	case "balance":
		{
			return []Tracker{NewBalanceTracker(logger, db, client, account)}, nil
		}
	case "disputeStatus":
		{
			return []Tracker{NewDisputeTracker(logger, config, db, contract, account)}, nil
		}
	case "gas":
		{
			return []Tracker{NewGasTracker(logger, db, client)}, nil
		}
	case "currentVariables":
		{
			return []Tracker{NewCurrentVariablesTracker(logger, db, contract, account)}, nil
		}
	case "tributeBalance":
		{
			return []Tracker{NewTributeTracker(logger, db, contract, account)}, nil
		}
	case "indexers":
		{
			return BuildIndexTrackers(config, db, client)
		}
	case "disputeChecker":
		return []Tracker{NewDisputeChecker(logger, config, client, contract, 0)}, nil
	default:
		return nil, errors.Errorf("no tracker with the name %s", name)
	}
}
