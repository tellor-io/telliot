// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ops

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/dataServer"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/rpc"
)

const ComponentName = "ops"

// DataServerOps is the driver for data server.
// It is the operational database component. Its purpose is to monitor/track various
// values and cache those values in a local data store for faster retrieval from a miner.
type DataServerOps struct {
	ctx     context.Context
	close   context.CancelFunc
	done    chan int
	server  *dataServer.DataServer
	logger  log.Logger
	Running bool
}

// CreateDataServerOps creates a data server instance for runtime.
func CreateDataServerOps(
	ctx context.Context,
	logger log.Logger,
	config *config.Config,
	DB db.DataServerProxy,
	client contracts.ETHClient,
	contract *contracts.Tellor,
	accounts []*rpc.Account,
) (*DataServerOps, error) {
	ctx, close := context.WithCancel(ctx)
	ds, err := dataServer.CreateServer(ctx, logger, config, DB, client, contract, accounts)
	if err != nil {
		close()
		return nil, err
	}
	done := make(chan int)
	logger, err = logging.ApplyFilter(*config, ComponentName, logger)
	if err != nil {
		close()
		return nil, errors.Wrap(err, "apply filter logger")
	}

	ops := &DataServerOps{
		ctx:     ctx,
		close:   close,
		server:  ds,
		logger:  log.With(logger, "component", ComponentName),
		done:    done,
		Running: false,
	}

	return ops, nil
}

// Start the data server.
func (ops *DataServerOps) Start() error {
	if err := ops.server.Start(ops.ctx, ops.done); err != nil {
		return err
	}
	ops.Running = true
	return nil
}

// Ready signals that the data server has completed at least one tracker cycle and any external dependencies
// should be ready to use its initial output.
func (ops *DataServerOps) Ready() chan bool {
	return ops.server.Ready()
}

// Stop will take care of stopping the dataserver component.
func (ops *DataServerOps) Stop() {
	level.Info(ops.logger).Log("msg", "shutting down data server...")
	ops.close()
	ops.done <- 1
	cnt := 0
	for {
		time.Sleep(500 * time.Millisecond)
		cnt++
		if ops.server.Stopped {
			break
		}
		if cnt > 60 {
			level.Warn(ops.logger).Log("msg", "expected data server to stop by now, Giving up...")
			return
		}
	}
	ops.Running = false
	level.Info(ops.logger).Log("msg", "data server shutdown complete")
}
