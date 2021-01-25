// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package main

import (
	"context"
	"os"
	"strconv"

	"github.com/alecthomas/kong"
	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
	tellorCommon "github.com/tellor-io/telliot/pkg/common"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/util"
)

func parseConfig(path string) (*config.Config, error) {
	err := config.ParseConfig(path)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing config")
	}
	return config.GetConfig(), nil
}

func createLogger(logConfig map[string]string, level string) (log.Logger, error) {
	err := util.SetupLoggingConfig(logConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing log config")
	}
	logger := util.SetupLogger(level)
	return logger, nil
}

func createTellorVariables(ctx context.Context, cfg *config.Config) (contracts.ETHClient, *contracts.Tellor, *rpc.Account, error) {

	if !cfg.EnablePoolWorker {

		// Create an rpc client
		client, err := rpc.NewClient(os.Getenv(config.NodeURLEnvName))
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "create rpc client instance")
		}

		contract, err := contracts.NewTellor(client)
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

// migrateAndOpenDB migrates the tx costs and deletes the db.
// The DB is always deleted because the price avarages calculations
// is not calculated properly between restarts.
// TODO don't do this and just improve the price calculations.
func migrateAndOpenDB(cfg *config.Config) (db.DB, error) {
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

var cli struct {
	Transfer transferCmd `cmd:"" help:"Transfer tokens"`
	Approve  approveCmd  `cmd:"" help:"Approve tokens"`
	Balance  balanceCmd  `cmd:"" help:"Check the balance of an address"`
	Stake    struct {
		Deposit  depositCmd  `cmd:"" help:"deposit a stake"`
		Request  requestCmd  `cmd:"" help:"request to withdraw stake"`
		Withdraw withdrawCmd `cmd:"" help:"withdraw stake"`
		Status   statusCmd   `cmd:"" help:"show stake status"`
	} `cmd:"" help:"Perform one of the stake operations"`
	Dispute struct {
		New  newDisputeCmd `cmd:"" help:"start a new dispute"`
		Vote voteCmd       `cmd:"" help:"vote on a open dispute"`
		Show showCmd       `cmd:"" help:"show open disputes"`
	} `cmd:"" help:"Perform commands related to disputes"`
	Dataserver dataserverCmd `cmd:"" help:"launch only a dataserver instance"`
	Mine       mineCmd       `cmd:"" help:"mine TRB and submit values"`
	Version    VersionCmd    `cmd:"" help:"Show the Docker version information"`
}

func main() {
	ctx := kong.Parse(&cli, kong.Name("Telliot"),
		kong.Description("The official Tellor cli tool"),
		kong.UsageOnError())
	err := ctx.Run(*ctx)
	ctx.FatalIfErrorf(err)
}
