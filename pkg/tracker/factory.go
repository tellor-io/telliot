// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
)

// CreateTracker a tracker instance by its well-known name.
func createTracker(name string, logger log.Logger, config *config.Config, db db.DataServerProxy, client contracts.ETHClient, contract *contracts.ITellor) ([]Tracker, error) {
	switch name {
	case "gas":
		{
			return []Tracker{NewGasTracker(logger, db, client)}, nil
		}
	case "indexers":
		{
			return BuildIndexTrackers(logger, config, db, client)
		}
	case "disputeChecker":
		return []Tracker{NewDisputeChecker(logger, config, client, contract, 0)}, nil
	default:
		return nil, errors.Errorf("no tracker with the name %s", name)
	}
}
