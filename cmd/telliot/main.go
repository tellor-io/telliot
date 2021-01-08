// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package main

import (
	"context"
	"os"
	"strconv"

	"github.com/alecthomas/kong"
	"github.com/pkg/errors"
	tellorCommon "github.com/tellor-io/telliot/pkg/common"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
)

// migrateAndOpenDB migrates the tx costs and deletes the db.
// The DB is always deleted because the price avarages calculations
// is not calculated properly between restarts.
// TODO don't do this and just improve the price calculations.
func migrateAndOpenDB() (db.DB, error) {
	cfg := config.GetConfig()
	// Create a db instance
	DB, err := db.Open(cfg.DBFile)
	if err != nil {
		return nil, errors.Wrapf(err, "opening DB instance")
	}

	var txsGas [][]byte
	for i := 0; i <= 5; i++ {
		txID := tellorCommon.PriceTXs + strconv.Itoa(i)
		txGas, err := DB.Get(txID)
		if err == nil && len(txGas) > 0 {
			txsGas = append(txsGas, txGas)
		}
	}
	if err := DB.Close(); err != nil {
		return nil, errors.Wrapf(err, "closing DB instance for migration")
	}
	os.RemoveAll(cfg.DBFile)

	DB, err = db.Open(cfg.DBFile)
	if err != nil {
		return nil, errors.Wrapf(err, "opening DB instance")
	}

	for i, txGas := range txsGas {
		txID := tellorCommon.PriceTXs + strconv.Itoa(i)
		_ = DB.Put(txID, txGas)
	}

	return DB, nil
}

func setup(ctx context.Context, cfg *config.Config) (rpc.ETHClient, *contracts.Tellor, *rpc.Account, error) {

	if !cfg.EnablePoolWorker {

		// Create an rpc client
		client, err := rpc.NewClient(os.Getenv(config.NodeURLEnvName))
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "create rpc client instance")
		}

		contract, err := contracts.NewTellor(cfg, client)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "create tellor master instance")
		}

		account, err := rpc.NewAccount(cfg)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "getting private key to ECDSA")
		}

		// Issue #55, halt if client is still syncing with Ethereum network
		s, err := client.IsSyncing(ctx)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "determining if Ethereum client is syncing")
		}
		if s {
			return nil, nil, nil, errors.New("ethereum node is still syncing with the network")
		}

		return client, &contract, &account, nil
	}
	// Not sure why we need this case.
	return nil, nil, nil, nil
}

func AddDBToCtx(remote bool) (db.DataServerProxy, db.DB, error) {
	cfg := config.GetConfig()
	// Create a db instance
	os.RemoveAll(cfg.DBFile)
	DB, err := db.Open(cfg.DBFile)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "opening DB instance")
	}

	var dataProxy db.DataServerProxy
	if remote {
		proxy, err := db.OpenRemoteDB(DB)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "open remote DB instance")

		}
		dataProxy = proxy
	} else {
		proxy, err := db.OpenLocalProxy(DB)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "opening local DB instance:")

		}
		dataProxy = proxy
	}
	return dataProxy, DB, nil
}

var cli struct {
	Config   configPath  `required type:"existingfile" help:"path to config file"`
	Transfer transferCmd `cmd help:"Transfer tokens"`
	Approve  approveCmd  `cmd help:"Approve tokens"`
	Balance  balanceCmd  `cmd help:"Check the balance of an address"`
	Stake    stakeCmd    `cmd help:"Perform one of the stake operations"`
	Dispute  struct {
		New  newDisputeCmd `cmd help:"start a new dispute"`
		Vote voteCmd       `cmd help:"vote on a open dispute"`
		Show showCmd       `cmd help:"show open disputes"`
	} `cmd help:"Perform commands related to disputes"`
	Dataserver dataserverCmd `cmd help:"launch only a dataserver instance"`
	Mine       mineCmd       `cmd help:"mine TRB and submit values"`
}

func main() {
	ctx := kong.Parse(&cli, kong.Name("Telliot"),
		kong.Description("The official Tellor cli tool"),
		kong.UsageOnError())
	err := ctx.Run(*ctx)
	ctx.FatalIfErrorf(err)
}
