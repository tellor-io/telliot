// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package dataServer

// import (
// 	"context"

// 	"github.com/go-kit/kit/log"
// 	"github.com/pkg/errors"
// 	"github.com/tellor-io/telliot/pkg/config"
// 	"github.com/tellor-io/telliot/pkg/contracts"
// 	"github.com/tellor-io/telliot/pkg/db"
// 	"github.com/tellor-io/telliot/pkg/logging"
// )

// // DataServerOps is the driver for data server.
// // It is the operational database component. Its purpose is to monitor/track various
// // values and cache those values in a local data store for faster retrieval from a miner.
// type DataServerOps struct {
// 	ctx    context.Context
// 	close  context.CancelFunc
// 	server *DataServer
// 	logger log.Logger
// }

// // NewDataServerOps creates a data server instance for runtime.
// func NewDataServerOps(
// 	ctx context.Context,
// 	logger log.Logger,
// 	config *config.Config,
// 	DB db.DataServerProxy,
// 	client contracts.ETHClient,
// 	contract *contracts.ITellor,
// 	accounts []*config.Account,
// ) (*DataServerOps, error) {
// 	ds, err := NewServer(logger, config, DB, client, contract, accounts)
// 	if err != nil {
// 		return nil, err
// 	}
// 	logger, err = logging.ApplyFilter(*config, ComponentName, logger)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "apply filter logger")
// 	}
// 	ctx, close := context.WithCancel(ctx)
// 	ops := &DataServerOps{
// 		ctx:    ctx,
// 		close:  close,
// 		server: ds,
// 		logger: log.With(logger, "component", ComponentName),
// 	}

// 	return ops, nil
// }

// // Start the data server.
// func (ops *DataServerOps) Start() error {
// 	if err := ops.server.Start(ops.ctx); err != nil {
// 		return err
// 	}
// 	return nil
// }

// // Ready signals that the data server has completed at least one tracker cycle and any external dependencies
// // should be ready to use its initial output.
// func (ops *DataServerOps) Ready() chan bool {
// 	return ops.server.Ready()
// }

// // Stop will take care of stopping the dataserver component.
// func (ops *DataServerOps) Stop() {
// 	ops.close()
// }
