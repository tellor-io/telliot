// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package main

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	cli "github.com/jawher/mow.cli"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	tellorCommon "github.com/tellor-io/telliot/pkg/common"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/ops"
	"github.com/tellor-io/telliot/pkg/rest"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/util"
)

var ctx context.Context
var cont contracts.Tellor
var acc rpc.Account
var clt rpc.ETHClient
var logLevel string
var database db.DB
var proxy db.DataServerProxy

func ExitOnError(err error, operation string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s failed: %s\n", operation, err.Error())
		cli.Exit(-1)
	}
}

func setup() error {
	cfg := config.GetConfig()

	err := util.SetupLoggingConfig(cfg.Logger)
	if err != nil {
		return errors.Wrapf(err, "parsing log config")
	}

	logLevel = cfg.LogLevel

	if !cfg.EnablePoolWorker {
		// Create an rpc client
		client, err := rpc.NewClient(os.Getenv(config.NodeURLEnvName))
		if err != nil {
			return errors.Wrap(err, "create rpc client instance")
		}

		contract, err := contracts.NewTellor(cfg, client)
		if err != nil {
			return errors.Wrap(err, "creating contract")
		}

		account, err := rpc.NewAccount(cfg)
		if err != nil {
			return errors.Wrap(err, "creating account")
		}
		ctx = context.Background()
		// Issue #55, halt if client is still syncing with Ethereum network
		s, err := client.IsSyncing(ctx)
		if err != nil {
			return errors.Wrap(err, "determining if Ethereum client is syncing")
		}
		if s {
			return errors.New("ethereum node is still syncing with the network")
		}

		clt = client
		acc = account
		cont = contract
	}
	return nil
}

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

func AddDBToCtx(remote bool) error {
	DB, err := migrateAndOpenDB()
	if err != nil {
		return errors.Wrap(err, "opening DB instance")
	}
	var dataProxy db.DataServerProxy
	if remote {
		proxy, err := db.OpenRemoteDB(DB)
		if err != nil {
			return errors.Wrapf(err, "open remote DB instance")

		}
		dataProxy = proxy
	} else {
		proxy, err := db.OpenLocalProxy(DB)
		if err != nil {
			return errors.Wrapf(err, "opening local DB instance:")

		}
		dataProxy = proxy
	}
	database = DB
	proxy = dataProxy
	return nil
}

var GitTag string
var GitHash string

const versionMessage = `
    The official Tellor cli tool %s (%s)
    -----------------------------------------
	Website: https://tellor.io
	Github:  https://github.com/tellor-io/telliot
`

func App() *cli.Cli {

	app := cli.App("telliot", "The tellor.io official cli tool")

	// App wide config options
	configPath := app.StringOpt("config", "configs/config.json", "Path to the primary JSON config file")

	// Leaving this logSetup here because will change soon with new CLI
	logSetup := util.SetupLogger()

	// This will get run before any of the commands
	app.Before = func() {
		ExitOnError(config.ParseConfig(*configPath), "parsing config file")
		ExitOnError(setup(), "setting up")
	}

	versionMessage := fmt.Sprintf(versionMessage, GitTag, GitHash)
	app.Version("version", versionMessage)

	app.Command("stake", "staking operations", stakeCmd(logSetup))
	app.Command("transfer", "send TRB to address", moveCmd(ops.Transfer, logSetup))
	app.Command("approve", "approve TRB to address", moveCmd(ops.Approve, logSetup))
	app.Command("balance", "check balance of address", balanceCmd)
	app.Command("dispute", "dispute operations", disputeCmd(logSetup))
	app.Command("mine", "mine for TRB", mineCmd(logSetup))
	app.Command("dataserver", "start an independent dataserver", dataserverCmd(logSetup))
	return app
}

func stakeCmd(logSetup func(string) log.Logger) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		cmd.Command("deposit", "deposit TRB stake", simpleCmd(ops.Deposit, logSetup))
		cmd.Command("withdraw", "withdraw TRB stake", simpleCmd(ops.WithdrawStake, logSetup))
		cmd.Command("request", "request to withdraw TRB stake", simpleCmd(ops.RequestStakingWithdraw, logSetup))
		cmd.Command("status", "show current staking status", simpleCmd(ops.ShowStatus, logSetup))
	}
}

func simpleCmd(
	f func(context.Context,
		log.Logger,
		rpc.ETHClient,
		contracts.Tellor,
		rpc.Account) error,
	logSetup func(string) log.Logger) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		cmd.Action = func() {
			ExitOnError(f(ctx, logSetup(logLevel), clt, cont, acc), "")
		}
	}
}

func moveCmd(
	f func(context.Context,
		log.Logger,
		rpc.ETHClient,
		contracts.Tellor,
		rpc.Account,
		common.Address,
		*big.Int) error,
	logSetup func(string) log.Logger) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		amt := TRBAmount{}
		addr := ETHAddress{}
		cmd.VarArg("AMOUNT", &amt, "amount to transfer")
		cmd.VarArg("ADDRESS", &addr, "ethereum public address")
		cmd.Action = func() {
			ExitOnError(f(ctx, logSetup(logLevel), clt, cont, acc, addr.addr, amt.Int), "move")
		}
	}
}

func balanceCmd(cmd *cli.Cmd) {
	addr := ETHAddress{}
	cmd.VarArg("ADDRESS", &addr, "ethereum public address")
	cmd.Spec = "[ADDRESS]"
	cmd.Action = func() {
		var zero [20]byte
		if bytes.Equal(addr.addr.Bytes(), zero[:]) {
			addr.addr = acc.Address
		}
		ExitOnError(ops.Balance(ctx, clt, cont.Getter, addr.addr), "checking balance")
	}

}

func disputeCmd(loggerSetup func(string) log.Logger) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		cmd.Command("vote", "vote on an active dispute", voteCmd)
		cmd.Command("new", "start a new dispute", newDisputeCmd)
		cmd.Command("show", "show existing disputes", simpleCmd(ops.List, loggerSetup))
	}
}

func voteCmd(cmd *cli.Cmd) {
	disputeID := EthereumInt{}
	cmd.VarArg("DISPUTE_ID", &disputeID, "dispute id")
	supports := cmd.BoolArg("SUPPORT", false, "do you support the dispute? (true|false)")
	cmd.Action = func() {
		ExitOnError(ops.Vote(ctx, clt, cont, acc, disputeID.Int, *supports), "vote")
	}
}

func newDisputeCmd(cmd *cli.Cmd) {
	requestID := EthereumInt{}
	timestamp := EthereumInt{}
	minerIndex := EthereumInt{}
	cmd.VarArg("REQUEST_ID", &requestID, "request id")
	cmd.VarArg("TIMESTAMP", &timestamp, "timestamp")
	cmd.VarArg("MINER_INDEX", &minerIndex, "miner to dispute (0-4)")
	cmd.Action = func() {
		ExitOnError(ops.Dispute(ctx, clt, cont, acc, requestID.Int, timestamp.Int, minerIndex.Int), "new dipsute")
	}
}

func mineCmd(logSetup func(string) log.Logger) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		remoteDS := cmd.BoolOpt("remote r", false, "connect to remote dataserver")
		cmd.Action = func() {
			logger := logSetup(logLevel)
			// Create os kill sig listener.
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt)
			exitChannels := make([]*chan os.Signal, 0)

			cfg := config.GetConfig()
			var ds *ops.DataServerOps
			if !cfg.EnablePoolWorker {
				ExitOnError(AddDBToCtx(*remoteDS), "initializing database")
				if !*remoteDS {
					ch := make(chan os.Signal)
					exitChannels = append(exitChannels, &ch)

					var err error
					ds, err = ops.CreateDataServerOps(ctx, logger, cfg, database, &proxy, clt, &cont, &acc, ch)
					ExitOnError(err, "creating data server")
					// Start and wait for it to be ready.
					ExitOnError(ds.Start(ctx), "starting data server")
					<-ds.Ready()
				}
			}

			level.Info(logger).Log("msg", "starting metrics server", "address", cfg.Mine.ListenHost+":"+strconv.Itoa(int(cfg.Mine.ListenPort)))
			http.Handle("/metrics", promhttp.Handler())
			srv, err := rest.Create(ctx, proxy, cfg.Mine.ListenHost, cfg.Mine.ListenPort)
			ExitOnError(err, "creating data server instance")
			srv.Start()

			// Start miner
			v, err := proxy.Get(db.DisputeStatusKey)
			if err != nil {
				level.Warn(logger).Log("msg", "getting dispute status. Check if staked")
			}
			status, _ := hexutil.DecodeBig(string(v))
			if status.Cmp(big.NewInt(1)) != 0 {
				ExitOnError(errors.New("miner is not able to mine with current status"), "checking miner")
			}
			ch := make(chan os.Signal)
			exitChannels = append(exitChannels, &ch)
			miner, err := ops.CreateMiningManager(logger, ch, cfg, proxy, cont, acc)
			ExitOnError(err, "creating miner")
			go func() {
				miner.Start(ctx)
			}()

			// Wait for kill sig.
			<-c
			// Then notify exit channels.
			for _, ch := range exitChannels {
				*ch <- os.Interrupt
			}

			if err := srv.Stop(); err != nil {
				level.Warn(logger).Log("msg", "shutting down the server", "err", err)
			}

			cnt := 0
			start := time.Now()
			for {
				cnt++
				dsStopped := false
				minerStopped := false

				if ds != nil {
					dsStopped = !ds.Running
				} else {
					dsStopped = true
				}

				if miner != nil {
					minerStopped = !miner.Running
				} else {
					minerStopped = true
				}

				if !dsStopped && !minerStopped && cnt > 60 {
					level.Warn(logger).Log("msg", "taking longer than expected to stop operations", "waited", time.Since(start))
				} else if dsStopped && minerStopped {
					break
				}
				time.Sleep(500 * time.Millisecond)
			}
			level.Info(logger).Log("msg", "main shutdown complete")
		}
	}
}

func dataserverCmd(logSetup func(string) log.Logger) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		cmd.Action = func() {
			logger := logSetup(logLevel)
			// Create os kill sig listener.
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt)

			var ds *ops.DataServerOps
			ExitOnError(AddDBToCtx(true), "initializing database")
			ch := make(chan os.Signal)
			var err error
			ds, err = ops.CreateDataServerOps(ctx, logger, config.GetConfig(), database, &proxy, clt, &cont, &acc, ch)
			ExitOnError(err, "creating data server")

			// Start and wait for it to be ready
			err = ds.Start(ctx)
			ExitOnError(err, "starting data server")

			<-ds.Ready()

			http.Handle("/metrics", promhttp.Handler())
			cfg := config.GetConfig()
			srv, err := rest.Create(ctx, proxy, cfg.DataServer.ListenHost, cfg.DataServer.ListenPort)
			ExitOnError(err, "creating data server instance")
			srv.Start()

			// Wait for kill sig.
			<-c
			// Notify exit channels.
			ch <- os.Interrupt

			if err := srv.Stop(); err != nil {
				level.Warn(logger).Log("msg", "shutting down the server", "err", err)
			}

			cnt := 0
			start := time.Now()
			for {
				cnt++
				dsStopped := false

				if ds != nil {
					dsStopped = !ds.Running
				} else {
					dsStopped = true
				}

				if !dsStopped && cnt > 60 {
					level.Warn(logger).Log("msg", "taking longer than expected to stop operations", "waited", time.Since(start))
				} else if dsStopped {
					break
				}
				time.Sleep(500 * time.Millisecond)
			}
			level.Info(logger).Log("msg", "main shutdown complete")
		}

	}
}

func main() {
	// Programming is easy. Just create an App() and run it!!!!!
	app := App()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "app.Run failed: %v\n", err)
	}
}
