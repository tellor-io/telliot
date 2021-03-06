// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package dataServer

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/rpc"
)

// DataServerOps is the driver for data server.
// It is the operational database component. Its purpose is to monitor/track various
// values and cache those values in a local data store for faster retrieval from a miner.
type DataServerOps struct {
	ctx    context.Context
	close  context.CancelFunc
	server *DataServer
	logger log.Logger
}

// CreateDataServerOps creates a data server instance for runtime.
func CreateDataServerOps(
	ctx context.Context,
	logger log.Logger,
	config *config.Config,
	DB db.DataServerProxy,
	client contracts.ETHClient,
	contract *contracts.ITellor,
	accounts []*rpc.Account,
) (*DataServerOps, error) {
	ds, err := CreateServer(logger, config, DB, client, contract, accounts)
	if err != nil {
		return nil, err
	}
	logger, err = logging.ApplyFilter(*config, ComponentName, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	ctx, close := context.WithCancel(ctx)
	ops := &DataServerOps{
		ctx:    ctx,
		close:  close,
		server: ds,
		logger: log.With(logger, "component", ComponentName),
	}

	return ops, nil
}

// Start the data server.
func (ops *DataServerOps) Start() error {
	if err := ops.server.Start(ops.ctx); err != nil {
		return err
	}
	return nil
}

// Ready signals that the data server has completed at least one tracker cycle and any external dependencies
// should be ready to use its initial output.
func (ops *DataServerOps) Ready() chan bool {
	return ops.server.Ready()
}

// Stop will take care of stopping the dataserver component.
func (ops *DataServerOps) Stop() {
	ops.close()
	level.Info(ops.logger).Log("msg", "data server shutdown complete")
}
