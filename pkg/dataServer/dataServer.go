// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package dataServer

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"

	"github.com/hashicorp/go-multierror"
	"github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/rest"
	"github.com/tellor-io/TellorMiner/pkg/rpc"
	"github.com/tellor-io/TellorMiner/pkg/tracker"
)

// DataServer holds refs to primary stack of utilities for data retrieval and serving.
type DataServer struct {
	server       *rest.Server
	DB           db.DB
	runner       *tracker.Runner
	ethClient    rpc.ETHClient
	exitCh       chan int
	runnerExitCh chan int
	Stopped      bool
	readyChannel chan bool
	logger       log.Logger
}

// CreateServer creates a data server stack and kicks off all go routines to start retrieving and serving data.
func CreateServer(ctx context.Context, logger log.Logger) (*DataServer, error) {
	cfg := config.GetConfig()

	DB := ctx.Value(common.DBContextKey).(db.DB)
	client := ctx.Value(common.ClientContextKey).(rpc.ETHClient)
	run, err := tracker.NewRunner(client, DB, logger)
	if err != nil {
		return nil, errors.Wrapf(err, "creating  tracker runner instance")
	}
	srv, err := rest.Create(ctx, cfg.ServerHost, cfg.ServerPort)
	if err != nil {
		return nil, errors.Wrapf(err, "creating runner instance")
	}
	// Make sure channel buffer size 1 since there is no guarantee that anyone
	// Would be listening to the channel
	ready := make(chan bool, 1)
	return &DataServer{server: srv,
		DB:           DB,
		runner:       run,
		ethClient:    client,
		exitCh:       nil,
		Stopped:      true,
		runnerExitCh: nil,
		readyChannel: ready,
		logger:       log.With(logger, "component", "data server")}, nil

}

// Start the data server and all underlying resources.
func (ds *DataServer) Start(ctx context.Context, exitCh chan int) error {
	ds.exitCh = exitCh
	ds.runnerExitCh = make(chan int)
	ds.Stopped = false
	err := ds.runner.Start(ctx, ds.runnerExitCh)
	if err != nil {
		return errors.Wrap(err, "starting runner")
	}

	ds.server.Start()
	go func() {
		<-ds.runner.Ready()
		level.Info(ds.logger).Log("msg", "runner signaled it is ready")
		ds.readyChannel <- true
		level.Info(ds.logger).Log("msg", "dataServer ready for use")
		<-ds.exitCh
		level.Info(ds.logger).Log("msg", "dataServer received signal to stop")
		if err := ds.stop(); err != nil {
			level.Info(ds.logger).Log("msg", "stopping the data server", "err", err)
		}
	}()
	return nil
}

// Ready provides notification channel that data from trackers is ready for use.
func (ds *DataServer) Ready() chan bool {
	return ds.readyChannel
}

func (ds *DataServer) stop() error {
	var final error
	// Stop tracker run loop.
	ds.runnerExitCh <- 1

	// Stop REST erver.
	if err := ds.server.Stop(); err != nil {
		final = multierror.Append(final, err)
	}

	// Stop the DB.
	if err := ds.DB.Close(); err != nil {
		final = multierror.Append(final, err)
	}

	// Stop the eth RPC client.
	ds.ethClient.Close()

	// All done.
	ds.Stopped = true
	return final
}
