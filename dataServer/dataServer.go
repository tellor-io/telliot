package dataServer

import (
	"context"
	"log"

	"github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rest"
	"github.com/tellor-io/TellorMiner/rpc"
	"github.com/tellor-io/TellorMiner/tracker"
)

//DataServer holds refs to primary stack of utilities for data retrieval and serving
type DataServer struct {
	server       *rest.Server
	DB           db.DB
	runner       *tracker.Runner
	ethClient    rpc.ETHClient
	exitCh       chan int
	runnerExitCh chan int
	Stopped      bool
}

//CreateServer creates a data server stack and kicks off all go routines to start retrieving and serving data
func CreateServer(ctx context.Context) (*DataServer, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	DB := ctx.Value(common.DBContextKey).(db.DB)
	client := ctx.Value(common.ClientContextKey).(rpc.ETHClient)
	run, err := tracker.NewRunner(client, DB)
	if err != nil {
		log.Fatal(err)
	}
	srv, err := rest.Create(ctx, cfg.ServerHost, cfg.ServerPort)

	return &DataServer{server: srv, DB: DB, runner: run, ethClient: client, exitCh: nil, Stopped: true, runnerExitCh: nil}, nil
}

//Start the data server and all underlying resources
func (ds *DataServer) Start(ctx context.Context, exitCh chan int) error {
	ds.exitCh = exitCh
	ds.runnerExitCh = make(chan int)
	err := ds.runner.Start(ctx, ds.runnerExitCh)
	if err != nil {
		return err
	}

	ds.server.Start()
	go func() {
		<-ds.exitCh
		ds.stop()
	}()
	return nil
}

func (ds *DataServer) stop() error {
	//stop tracker run loop
	ds.runnerExitCh <- 1

	//stop REST erver
	ds.server.Stop()

	//stop the DB
	ds.DB.Close()

	//stop the eth RPC client
	ds.ethClient.Close()

	//all done
	ds.Stopped = true
	return nil
}
