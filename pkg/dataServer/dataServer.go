// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package dataServer

const ComponentName = "dataServer"

type Config struct {
	LogLevel        string
	ListenHost      string
	ListenPort      uint
	ServerWhitelist []string
}

// // DataServer holds refs to primary stack of utilities for data retrieval and serving.
// type DataServer struct {
// 	DB           db.DataServerProxy
// 	tsdbDB       *tsdb.DB
// 	runner       *tracker.Runner
// 	ethClient    contracts.ETHClient
// 	Stopped      bool
// 	readyChannel chan bool
// 	logger       log.Logger
// }

// // NewServer creates a data server stack and kicks off all go routines to start retrieving and serving data.
// func NewServer(
// 	logger log.Logger,
// 	cfg *config.Config,
// 	DB db.DataServerProxy,
// 	client contracts.ETHClient,
// 	contract *contracts.ITellor,
// 	accounts []*ethereum.Account,
// ) (*DataServer, error) {

// 	if len(cfg.ServerWhitelist) == 0 {
// 		return errors.New("ServerWhitelist shouldn't be empty while running as dataserver")
// 	}
// 	logger, err = logging.ApplyFilter(*cfg, ComponentName, logger)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "apply filter logger")
// 	}
// 	// Make sure channel buffer size 1 since there is no guarantee that anyone
// 	// Would be listening to the channel
// 	ready := make(chan bool, 1)
// 	return &DataServer{
// 		DB:           DB,
// 		runner:       run,
// 		ethClient:    client,
// 		Stopped:      true,
// 		readyChannel: ready,
// 		logger:       log.With(logger, "component", ComponentName)}, nil

// }

// // Start the data server and all underlying resources.
// func (ds *DataServer) Start(ctx context.Context) error {
// 	ds.Stopped = false
// 	runnerCtx, runnercCloser := context.WithCancel(ctx)
// 	err := ds.runner.Start(runnerCtx)
// 	if err != nil {
// 		runnercCloser()
// 		return errors.Wrap(err, "starting runner data server")
// 	}

// 	go func() {
// 		<-ds.runner.Ready()
// 		level.Info(ds.logger).Log("msg", "runner signaled it is ready")
// 		ds.readyChannel <- true
// 		level.Info(ds.logger).Log("msg", "dataServer ready for use")
// 		<-ctx.Done()
// 		level.Info(ds.logger).Log("msg", "dataServer received signal to stop")
// 		runnercCloser()
// 		if err := ds.stop(); err != nil {
// 			level.Info(ds.logger).Log("msg", "stopping the data server", "err", err)
// 		}
// 	}()
// 	return nil
// }

// // Ready provides notification channel that data from trackers is ready for use.
// func (ds *DataServer) Ready() chan bool {
// 	return ds.readyChannel
// }

// func (ds *DataServer) stop() error {
// 	var final error

// 	// Stop the DB.
// 	if err := ds.DB.Close(); err != nil {
// 		final = multierror.Append(final, err)
// 	}

// 	// Stop the eth RPC client.
// 	ds.ethClient.Close()

// 	// All done.
// 	ds.Stopped = true
// 	return final
// }
