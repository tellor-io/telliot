// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	cli "github.com/jawher/mow.cli"
	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/contracts/getter"
	"github.com/tellor-io/TellorMiner/pkg/contracts/tellor"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/ops"
	"github.com/tellor-io/TellorMiner/pkg/rpc"
)

var ctx context.Context

func setupLogger(logLevel string) log.Logger {
	var lvl level.Option
	switch logLevel {
	case "error":
		lvl = level.AllowError()
	case "warn":
		lvl = level.AllowWarn()
	case "info":
		lvl = level.AllowInfo()
	case "debug":
		lvl = level.AllowDebug()
	default:
		panic("unexpected log level")
	}

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, lvl)

	return log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
}

func ErrorHandler(err error, operation string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s failed: %s\n", operation, err.Error())
		cli.Exit(-1)
	}
}

func buildContext() error {
	cfg := config.GetConfig()

	if !cfg.EnablePoolWorker {
		// Create an rpc client
		client, err := rpc.NewClient(cfg.NodeURL)
		if err != nil {
			return fmt.Errorf("Couldn't create client instance")
		}
		// Create an instance of the tellor master contract for on-chain interactions
		contractAddress := common.HexToAddress(cfg.ContractAddress)
		contractTellorInstance, err := tellor.NewTellor(contractAddress, client)
		if err != nil {
			return fmt.Errorf("Couldn't create Tellor Master instance")
		}

		contractGetterInstance, err := getter.NewTellorGetters(contractAddress, client)

		if err != nil {
			return fmt.Errorf("Couldn't create New Tellor transactor instance")
		}

		ctx = context.WithValue(context.Background(), tellorCommon.ClientContextKey, client)
		ctx = context.WithValue(ctx, tellorCommon.ContractAddress, contractAddress)
		ctx = context.WithValue(ctx, tellorCommon.ContractsTellorContextKey, contractTellorInstance)
		ctx = context.WithValue(ctx, tellorCommon.ContractsGetterContextKey, contractGetterInstance)

		privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
		if err != nil {
			return fmt.Errorf("problem getting private key: %s", err.Error())
		}
		ctx = context.WithValue(ctx, tellorCommon.PrivateKey, privateKey)

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return fmt.Errorf("error casting public key to ECDSA")
		}

		publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		ctx = context.WithValue(ctx, tellorCommon.PublicAddress, publicAddress)

		// Issue #55, halt if client is still syncing with Ethereum network
		s, err := client.IsSyncing(ctx)
		if err != nil {
			return fmt.Errorf("could not determine if Ethereum client is syncing: %v\n", err)
		}
		if s {
			return fmt.Errorf("ethereum node is still syncing with the network")
		}
	}
	return nil
}

func AddDBToCtx(remote bool) error {
	cfg := config.GetConfig()
	// Create a db instance
	os.RemoveAll(cfg.DBFile)
	DB, err := db.Open(cfg.DBFile)
	if err != nil {
		return fmt.Errorf("Error opening DB: %v\n", err)
	}

	var dataProxy db.DataServerProxy
	if remote {
		proxy, err := db.OpenRemoteDB(DB)
		if err != nil {
			return fmt.Errorf("Error remote DB: %v\n", err)

		}
		dataProxy = proxy
	} else {
		proxy, err := db.OpenLocalProxy(DB)
		if err != nil {
			return fmt.Errorf("Error opening local DB: %v\n", err)

		}
		dataProxy = proxy
	}
	ctx = context.WithValue(ctx, tellorCommon.DataProxyKey, dataProxy)
	ctx = context.WithValue(ctx, tellorCommon.DBContextKey, DB)
	return nil
}

var GitTag string
var GitHash string

const versionMessage = `
    The official Tellor Miner %s (%s)
    -----------------------------------------
	Website: https://tellor.io
	Github:  https://github.com/tellor-io/TellorMiner
`

func App() *cli.Cli {

	app := cli.App("TellorMiner", "The tellor.io official miner")

	// App wide config options
	configPath := app.StringOpt("config", "configs/config.json", "Path to the primary JSON config file")
	logLevel := app.StringOpt("logLevel", "Error", "The level of log messages")

	logger := setupLogger(*logLevel)
	// This will get run before any of the commands
	app.Before = func() {
		ErrorHandler(config.ParseConfig(*configPath), "parsing config file")
		ErrorHandler(buildContext(), "building context")
	}

	versionMessage := fmt.Sprintf(versionMessage, GitTag, GitHash)
	app.Version("version", versionMessage)

	app.Command("stake", "staking operations", stakeCmd(logger))
	app.Command("transfer", "send TRB to address", moveCmd(ops.Transfer, logger))
	app.Command("approve", "approve TRB to address", moveCmd(ops.Approve, logger))
	app.Command("balance", "check balance of address", balanceCmd)
	app.Command("dispute", "dispute operations", disputeCmd(logger))
	app.Command("mine", "mine for TRB", mineCmd(logger))
	app.Command("dataserver", "start an independent dataserver", dataserverCmd(logger))
	return app
}

func stakeCmd(logger log.Logger) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		cmd.Command("deposit", "deposit TRB stake", simpleCmd(ops.Deposit, logger))
		cmd.Command("withdraw", "withdraw TRB stake", simpleCmd(ops.WithdrawStake, logger))
		cmd.Command("request", "request to withdraw TRB stake", simpleCmd(ops.RequestStakingWithdraw, logger))
		cmd.Command("status", "show current staking status", simpleCmd(ops.ShowStatus, logger))
	}
}

func simpleCmd(f func(context.Context, log.Logger) error, logger log.Logger) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		cmd.Action = func() {
			ErrorHandler(f(ctx, logger), "")
		}
	}
}

func moveCmd(f func(context.Context, log.Logger, common.Address, *big.Int) error, logger log.Logger) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		amt := TRBAmount{}
		addr := ETHAddress{}
		cmd.VarArg("AMOUNT", &amt, "amount to transfer")
		cmd.VarArg("ADDRESS", &addr, "ethereum public address")
		cmd.Action = func() {
			ErrorHandler(f(ctx, logger, addr.addr, amt.Int), "move")
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
			addr.addr = ctx.Value(tellorCommon.PublicAddress).(common.Address)
		}
		ErrorHandler(ops.Balance(ctx, addr.addr), "checking balance")
	}
}

func disputeCmd(logger log.Logger) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		cmd.Command("vote", "vote on an active dispute", voteCmd)
		cmd.Command("new", "start a new dispute", newDisputeCmd)
		cmd.Command("show", "show existing disputes", simpleCmd(ops.List, logger))
	}
}

func voteCmd(cmd *cli.Cmd) {
	disputeID := EthereumInt{}
	cmd.VarArg("DISPUTE_ID", &disputeID, "dispute id")
	supports := cmd.BoolArg("SUPPORT", false, "do you support the dispute? (true|false)")
	cmd.Action = func() {
		ErrorHandler(ops.Vote(disputeID.Int, *supports, ctx), "vote")
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
		ErrorHandler(ops.Dispute(requestID.Int, timestamp.Int, minerIndex.Int, ctx), "new dipsute")
	}
}

func mineCmd(logger log.Logger) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		remoteDS := cmd.BoolOpt("remote r", false, "connect to remote dataserver")
		cmd.Action = func() {
			// Create os kill sig listener.
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt)
			exitChannels := make([]*chan os.Signal, 0)

			cfg := config.GetConfig()
			var ds *ops.DataServerOps
			if !cfg.EnablePoolWorker {
				ErrorHandler(AddDBToCtx(*remoteDS), "initializing database")
				if !*remoteDS {
					ch := make(chan os.Signal)
					exitChannels = append(exitChannels, &ch)

					var err error
					ds, err = ops.CreateDataServerOps(ctx, logger, ch)
					if err != nil {
						level.Error(logger).Log("msg", "error creating data server", "err", err)
						os.Exit(1)
					}
					// Start and wait for it to be ready.
					if err := ds.Start(ctx, logger); err != nil {
						level.Error(logger).Log("msg", "error starting data server", "err", err)
					}
					<-ds.Ready()
				}
			}
			// Start miner
			DB := ctx.Value(tellorCommon.DataProxyKey).(db.DataServerProxy)
			v, err := DB.Get(db.DisputeStatusKey)
			if err != nil {
				level.Warn(logger).Log("info", "could not get dispute status. Check if staked")
			}
			status, _ := hexutil.DecodeBig(string(v))
			if status.Cmp(big.NewInt(1)) != 0 {
				level.Error(logger).Log("msg", "Miner is not able to mine with status", "status", status, "info", "Stopping all mining immediately")
				os.Exit(1)
			}
			ch := make(chan os.Signal)
			exitChannels = append(exitChannels, &ch)
			miner, err := ops.CreateMiningManager(ctx, ch, ops.NewSubmitter())
			if err != nil {
				level.Error(logger).Log("msg", "unable to create miner", "err", err)
				os.Exit(1)
			}
			miner.Start(ctx)

			// Wait for kill sig.
			<-c
			// Then notify exit channels.
			for _, ch := range exitChannels {
				*ch <- os.Interrupt
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
					level.Warn(logger).Log("msg", "Taking longer than expected to stop operations", "waited", time.Since(start))
				} else if dsStopped && minerStopped {
					break
				}
				time.Sleep(500 * time.Millisecond)
			}
			level.Info(logger).Log("msg", "Main shutdown complete")
		}
	}
}

func dataserverCmd(logger log.Logger) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		cmd.Action = func() {
			// Create os kill sig listener.
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt)

			var ds *ops.DataServerOps
			ErrorHandler(AddDBToCtx(true), "initializing database")
			ch := make(chan os.Signal)
			var err error
			ds, err = ops.CreateDataServerOps(ctx, logger, ch)
			if err != nil {
				level.Error(logger).Log("msg", "error creating data server", "err", err)
				os.Exit(1)
			}
			// Start and wait for it to be ready
			if err := ds.Start(ctx, logger); err != nil {
				//Should we do this here or pass it down to errorhandler func for consistency?
				level.Error(logger).Log("msg", "error starting data server", "err", err)
				os.Exit(1)
			}
			<-ds.Ready()

			// Wait for kill sig.
			<-c
			// Notify exit channels.
			ch <- os.Interrupt

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
					level.Warn(logger).Log("msg", "Taking longer than expected to stop operations", "waited", time.Since(start))
				} else if dsStopped {
					break
				}
				time.Sleep(500 * time.Millisecond)
			}
			level.Info(logger).Log("msg", "Main shutdown complete")
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
