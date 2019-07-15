package main

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
	server    *rest.Server
	DB        db.DB
	runner    *tracker.Runner
	ethClient rpc.ETHClient
	exitCh    chan int
}

//CreateServer creates a data server stack and kicks off all go routines to start retrieving and serving data
func CreateServer() (*DataServer, error) {

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	if len(cfg.DBFile) == 0 {
		log.Fatal("Missing dbFile config setting")
	}

	DB, err := db.Open(cfg.DBFile)
	if err != nil {
		log.Fatal(err)
	}
	client, err := rpc.NewClient(cfg.NodeURL)
	if err != nil {
		log.Fatal(err)
	}

	run, err := tracker.NewRunner(client, DB)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.WithValue(context.Background(), common.DBContextKey, DB)
	ctx = context.WithValue(ctx, common.ClientContextKey, client)
	srv, err := rest.Create(ctx, cfg.ServerHost, cfg.ServerPort)

	return &DataServer{srv, DB, run, client, nil}, nil
}

//Start the data server and all underlying resources
func (ds *DataServer) Start(ctx context.Context, exitCh chan int) error {
	ds.exitCh = exitCh
	ctx = context.WithValue(ctx, common.DBContextKey, ds.DB)
	ctx = context.WithValue(ctx, common.ClientContextKey, ds.ethClient)
	err := ds.runner.Start(ctx, ds.exitCh)
	if err != nil {
		return err
	}

	ds.server.Start()
	return nil
}

//Stop the data server and all underlying resources
func (ds *DataServer) Stop() error {
	ds.exitCh <- 1
	ds.server.Stop()
	return nil
}
