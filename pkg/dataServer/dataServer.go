// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package dataServer

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/tracker"
)

// DataServer holds refs to primary stack of utilities for data retrieval and serving.
type DataServer struct {
	DB           db.DataServerProxy
	runner       *tracker.Runner
	ethClient    contracts.ETHClient
	exitCh       chan int
	runnerExitCh chan int
	Stopped      bool
	readyChannel chan bool
	logger       log.Logger
}

// CreateServer creates a data server stack and kicks off all go routines to start retrieving and serving data.
func CreateServer(
	ctx context.Context,
	logger log.Logger,
	config *config.Config,
	DB db.DataServerProxy,
	client contracts.ETHClient,
	contract *contracts.Tellor,
	account *rpc.Account,
) (*DataServer, error) {
	run, err := tracker.NewRunner(logger, config, DB, client, contract, account)
	if err != nil {
		return nil, errors.Wrapf(err, "creating data server tracker runner instance")
	}
	// Make sure channel buffer size 1 since there is no guarantee that anyone
	// Would be listening to the channel
	ready := make(chan bool, 1)
	return &DataServer{
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
		return errors.Wrap(err, "starting runner data server")
	}

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
