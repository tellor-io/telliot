// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ops

import (
	"context"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/dataServer"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
)

// DataServerOps is the driver for data server.
// It is the operational database component. Its purpose is to monitor/track various
// values and cache those values in a local data store for faster retrieval from a miner.
type DataServerOps struct {
	exitCh  chan os.Signal
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
	account *rpc.Account,
	exitCh chan os.Signal,
) (*DataServerOps, error) {
	ds, err := dataServer.CreateServer(ctx, logger, config, DB, client, contract, account)
	if err != nil {
		return nil, err
	}
	done := make(chan int)
	ops := &DataServerOps{exitCh: exitCh, server: ds, logger: log.With(logger, "component", "ops"), done: done, Running: false}

	return ops, nil
}

// Start the data server.
func (ops *DataServerOps) Start(ctx context.Context) error {
	if err := ops.server.Start(ctx, ops.done); err != nil {
		return err
	}
	ops.Running = true
	go func() {
		<-ops.exitCh
		level.Info(ops.logger).Log("msg", "shutting down data server...")
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
	}()
	return nil
}

// Ready signals that the data server has completed at least one tracker cycle and any external dependencies
// should be ready to use its initial output.
func (ops *DataServerOps) Ready() chan bool {
	return ops.server.Ready()
}
