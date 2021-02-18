// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package main

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-kit/kit/log/level"
	"github.com/oklog/run"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	tellorCommon "github.com/tellor-io/telliot/pkg/common"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/ops"
	"github.com/tellor-io/telliot/pkg/pow"
	"github.com/tellor-io/telliot/pkg/rest"
	"github.com/tellor-io/telliot/pkg/rpc"
)

var GitTag string
var GitHash string

const versionMessage = `
    The official Tellor cli tool %s (%s)
    -----------------------------------------
	Website: https://tellor.io
	Github:  https://github.com/tellor-io/telliot
`

type VersionCmd struct {
}

func (cmd *VersionCmd) Run() error {
	//lint:ignore faillint it should print to console
	fmt.Printf(versionMessage, GitTag, GitHash)
	return nil
}

type configPath string
type tokenCmd struct {
	Config  configPath `type:"existingfile" help:"path to config file"`
	Address string     `arg:""`
	Amount  string     `arg:""`
	Account int        `arg:"" optional:""`
}

type transferCmd tokenCmd

func (c *transferCmd) Run() error {
	cfg, err := parseConfig(string(c.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	logger := logging.NewLogger()

	ctx := context.Background()
	client, contract, accounts, err := createTellorVariables(ctx, logger, cfg)
	if err != nil {
		return errors.Wrapf(err, "creating tellor variables")
	}

	address := ETHAddress{}
	err = address.Set(c.Address)
	if err != nil {
		return errors.Wrapf(err, "parsing address argument")
	}
	amount := TRBAmount{}
	err = amount.Set(c.Amount)
	if err != nil {
		return errors.Wrapf(err, "parsing amount argument")
	}
	account, err := getAccountFor(accounts, c.Account)
	if err != nil {
		return err
	}
	return ops.Transfer(ctx, logger, client, contract, account, address.addr, amount.Int)

}

type approveCmd tokenCmd

func (c *approveCmd) Run() error {
	cfg, err := parseConfig(string(c.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	logger := logging.NewLogger()

	ctx := context.Background()
	client, contract, accounts, err := createTellorVariables(ctx, logger, cfg)
	if err != nil {
		return errors.Wrapf(err, "creating tellor variables")
	}

	address := ETHAddress{}
	err = address.Set(c.Address)
	if err != nil {
		return errors.Wrapf(err, "parsing address argument")
	}
	amount := TRBAmount{}
	err = amount.Set(c.Amount)
	if err != nil {
		return errors.Wrapf(err, "parsing amount argument")
	}
	account, err := getAccountFor(accounts, c.Account)
	if err != nil {
		return err
	}

	return ops.Approve(ctx, logger, client, contract, account, address.addr, amount.Int)
}

type accountsCmd struct {
	Config configPath `type:"existingfile" help:"path to config file"`
}

func (a *accountsCmd) Run() error {
	cfg, err := parseConfig(string(a.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	logger := logging.NewLogger()

	ctx := context.Background()
	_, _, accounts, err := createTellorVariables(ctx, logger, cfg)
	if err != nil {
		return errors.Wrapf(err, "creating tellor variables")
	}

	for i, account := range accounts {
		level.Info(logger).Log("msg", "account", "no", i, "address", account.Address.String())
	}

	return nil
}

type balanceCmd struct {
	Config  configPath `type:"existingfile" help:"path to config file"`
	Address string     `arg:"" optional:""`
}

func (b *balanceCmd) Run() error {
	cfg, err := parseConfig(string(b.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	logger := logging.NewLogger()

	ctx := context.Background()
	client, contract, _, err := createTellorVariables(ctx, logger, cfg)
	if err != nil {
		return errors.Wrapf(err, "creating tellor variables")
	}

	addr := ETHAddress{}
	if b.Address == "" {
		err = addr.Set(contract.Address.String())
		if err != nil {
			return errors.Wrapf(err, "parsing argument")
		}
	} else {
		err = addr.Set(b.Address)
		if err != nil {
			return errors.Wrapf(err, "parsing argument")
		}
	}
	return ops.Balance(ctx, logger, client, contract, addr.addr)
}

type depositCmd struct {
	Config  configPath `type:"existingfile" help:"path to config file"`
	Account int        `arg:"" optional:""`
}

func (d depositCmd) Run() error {
	cfg, err := parseConfig(string(d.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	logger := logging.NewLogger()

	ctx := context.Background()
	client, contract, accounts, err := createTellorVariables(ctx, logger, cfg)
	if err != nil {
		return errors.Wrapf(err, "creating tellor variables")
	}
	account, err := getAccountFor(accounts, d.Account)
	if err != nil {
		return err
	}
	return ops.Deposit(ctx, logger, client, contract, account)

}

type withdrawCmd struct {
	Config  configPath `type:"existingfile" help:"path to config file"`
	Address string     `arg:"" required:""`
	Account int        `arg:"" optional:""`
}

func (w withdrawCmd) Run() error {
	cfg, err := parseConfig(string(w.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	logger := logging.NewLogger()

	ctx := context.Background()
	client, contract, accounts, err := createTellorVariables(ctx, logger, cfg)
	if err != nil {
		return errors.Wrapf(err, "creating tellor variables")
	}

	addr := ETHAddress{}
	err = addr.Set(w.Address)
	if err != nil {
		return errors.Wrapf(err, "parsing argument")
	}
	account, err := getAccountFor(accounts, w.Account)
	if err != nil {
		return err
	}
	return ops.WithdrawStake(ctx, logger, client, contract, account)

}

type requestCmd struct {
	Config  configPath `type:"existingfile" help:"path to config file"`
	Account int        `arg:"" optional:""`
}

func (r requestCmd) Run() error {
	cfg, err := parseConfig(string(r.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	logger := logging.NewLogger()

	ctx := context.Background()
	client, contract, accounts, err := createTellorVariables(ctx, logger, cfg)
	if err != nil {
		return errors.Wrapf(err, "creating tellor variables")
	}
	account, err := getAccountFor(accounts, r.Account)
	if err != nil {
		return err
	}
	return ops.RequestStakingWithdraw(ctx, logger, client, contract, account)
}

type statusCmd struct {
	Config  configPath `type:"existingfile" help:"path to config file"`
	Account int        `arg:"" optional:""`
}

func (s statusCmd) Run() error {
	cfg, err := parseConfig(string(s.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	logger := logging.NewLogger()

	ctx := context.Background()
	client, contract, accounts, err := createTellorVariables(ctx, logger, cfg)
	if err != nil {
		return errors.Wrapf(err, "creating tellor variables")
	}
	account, err := getAccountFor(accounts, s.Account)
	if err != nil {
		return err
	}
	return ops.ShowStatus(ctx, logger, client, contract, account)
}

type migrateCmd struct {
	Config configPath `type:"existingfile" help:"path to config file"`
}

func (s migrateCmd) Run() error {
	cfg, err := parseConfig(string(s.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	logger := logging.NewLogger()

	ctx := context.Background()
	client, contract, account, err := createTellorVariables(ctx, logger, cfg)
	if err != nil {
		return errors.Wrapf(err, "creating tellor variables")
	}

	auth, err := ops.PrepareEthTransaction(ctx, client, account)
	if err != nil {
		return errors.Wrap(err, "prepare ethereum transaction")
	}

	tx, err := contract.Migrate(auth)
	if err != nil {
		return errors.Wrap(err, "contract failed")
	}
	level.Info(logger).Log("msg", "TRB migrated", "txHash", tx.Hash().Hex())
	return nil
}

type newDisputeCmd struct {
	Config     configPath `type:"existingfile" help:"path to config file"`
	requestId  string     `arg:""  help:"the request id to dispute it"`
	timestamp  string     `arg:""  help:"the submitted timestamp to dispute"`
	minerIndex string     `arg:""  help:"the miner index to dispute"`
	Account    int        `arg:"" optional:""`
}

func (n newDisputeCmd) Run() error {
	cfg, err := parseConfig(string(n.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	logger := logging.NewLogger()

	ctx := context.Background()
	client, contract, accounts, err := createTellorVariables(ctx, logger, cfg)
	if err != nil {
		return errors.Wrapf(err, "creating tellor variables")
	}

	requestID := EthereumInt{}
	err = requestID.Set(n.requestId)
	if err != nil {
		return errors.Wrapf(err, "parsing argument")
	}
	timestamp := EthereumInt{}
	err = timestamp.Set(n.timestamp)
	if err != nil {
		return errors.Wrapf(err, "parsing argument")
	}
	minerIndex := EthereumInt{}
	err = minerIndex.Set(n.minerIndex)
	if err != nil {
		return errors.Wrapf(err, "parsing argument")
	}
	account, err := getAccountFor(accounts, n.Account)
	if err != nil {
		return err
	}
	return ops.Dispute(ctx, logger, client, contract, account, requestID.Int, timestamp.Int, minerIndex.Int)

}

type voteCmd struct {
	Config    configPath `type:"existingfile" help:"path to config file"`
	disputeId string     `arg:""  help:"the dispute id"`
	support   bool       `arg:""  help:"true or false"`
	Account   int        `arg:"" optional:""`
}

func (v voteCmd) Run() error {
	cfg, err := parseConfig(string(v.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	logger := logging.NewLogger()

	ctx := context.Background()
	client, contract, accounts, err := createTellorVariables(ctx, logger, cfg)
	if err != nil {
		return errors.Wrapf(err, "creating tellor variables")
	}

	disputeID := EthereumInt{}
	err = disputeID.Set(v.disputeId)
	if err != nil {
		return errors.Wrapf(err, "parsing argument")
	}
	account, err := getAccountFor(accounts, v.Account)
	if err != nil {
		return err
	}
	return ops.Vote(ctx, logger, client, contract, account, disputeID.Int, v.support)
}

type showCmd struct {
	Config  configPath `type:"existingfile" help:"path to config file"`
	Account int        `arg:"" optional:""`
}

func (s showCmd) Run() error {
	cfg, err := parseConfig(string(s.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	logger := logging.NewLogger()

	ctx := context.Background()
	client, contract, accounts, err := createTellorVariables(ctx, logger, cfg)
	if err != nil {
		return errors.Wrapf(err, "creating tellor variables")
	}
	account, err := getAccountFor(accounts, s.Account)
	if err != nil {
		return err
	}
	return ops.List(ctx, cfg, logger, client, contract, account)
}

type dataserverCmd struct {
	Config configPath `type:"existingfile" help:"path to config file"`
}

func (d dataserverCmd) Run() error {
	cfg, err := parseConfig(string(d.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	logger := logging.NewLogger()

	ctx := context.Background()
	client, contract, account, err := createTellorVariables(ctx, logger, cfg)
	if err != nil {
		return errors.Wrapf(err, "creating tellor variables")
	}

	// Create os kill sig listener.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	DB, err := migrateAndOpenDB(logger, cfg)
	if err != nil {
		return errors.Wrapf(err, "initializing database")
	}
	proxy, err := db.OpenLocal(logger, cfg, DB)
	if err != nil {
		return errors.Wrapf(err, "open remote DB instance")
	}
	ch := make(chan os.Signal)
	ds, err := ops.CreateDataServerOps(ctx, logger, cfg, proxy, client, contract, account)
	if err != nil {
		return errors.Wrapf(err, "creating data server")
	}
	// Start and wait for it to be ready
	if err := ds.Start(); err != nil {
		return errors.Wrapf(err, "starting data server")
	}
	<-ds.Ready()

	http.Handle("/metrics", promhttp.Handler())
	srv, err := rest.Create(logger, cfg, ctx, proxy, cfg.DataServer.ListenHost, cfg.DataServer.ListenPort)
	if err != nil {
		return errors.Wrapf(err, "creating http data server")
	}
	srv.Start()

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
			level.Warn(logger).Log("msg", "taking longer than expected to stop operations", "waited", time.Since(start))
		} else if dsStopped {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	level.Info(logger).Log("msg", "main shutdown complete")
	return nil
}

type mineCmd struct {
	Config configPath `type:"existingfile" help:"path to config file"`
}

func (m mineCmd) Run() error {
	// Defining a global context for starting and stopping of components.
	ctx := context.Background()
	cfg, err := parseConfig(string(m.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}
	logger := logging.NewLogger()
	client, contract, accounts, err := createTellorVariables(ctx, logger, cfg)
	if err != nil {
		return errors.Wrapf(err, "creating tellor variables")
	}
	// Create os kill sig listener.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// We define our run groups here.
	var g run.Group
	// Run groups.
	{
		// Handle interupts.
		g.Add(run.SignalHandler(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM))

		// Metrics server.
		{
			http.Handle("/metrics", promhttp.Handler())
			srv := &http.Server{Addr: fmt.Sprintf("%s:%d", cfg.Mine.ListenHost, cfg.Mine.ListenPort)}
			g.Add(func() error {
				level.Info(logger).Log("msg", "starting metrics server", "addr", cfg.Mine.ListenHost, "port", cfg.Mine.ListenPort)
				// returns ErrServerClosed on graceful close
				var err error
				if err = srv.ListenAndServe(); err != http.ErrServerClosed {
					err = errors.Wrapf(err, "ListenAndServe")
				}
				return err
			}, func(error) {
				srv.Close()
			})
		}

		// DataServer is the Telliot data server.
		var proxy db.DataServerProxy
		{
			var ds *ops.DataServerOps
			DB, err := migrateAndOpenDB(logger, cfg)
			if err != nil {
				return errors.Wrapf(err, "initializing database")
			}
			if cfg.Mine.RemoteDBHost != "" {
				proxy, err = db.OpenRemote(logger, cfg, DB)
			} else {
				proxy, err = db.OpenLocal(logger, cfg, DB)
			}
			if err != nil {
				return errors.Wrapf(err, "open remote DB instance")

			}

			// Not using a remote DB so need to start the trackers.
			if cfg.Mine.RemoteDBHost == "" {
				ds, err = ops.CreateDataServerOps(ctx, logger, cfg, proxy, client, contract, accounts)
				if err != nil {
					return errors.Wrapf(err, "creating data server")
				}
				// Start and wait for it to be ready.
				g.Add(func() error {
					return errors.Wrapf(ds.Start(), "starting data server")
				}, func(error) {
					ds.Stop()
				})

				// We need to wait until the DataServer instance is ready.
				<-ds.Ready()
			}
		}

		// Run a miner manager for each of the accounts.
		{
			// Run a tasker instance.
			tasker := pow.CreateTasker(ctx, logger, cfg, proxy, client, contract, accounts)
			g.Add(func() error {
				return tasker.Start()
			}, func(error) {
				tasker.Stop()
			})
			// Initialize the solutionHandler.
			submitters := []tellorCommon.TransactionSubmitter{}
			for _, account := range accounts {
				submitters = append(submitters, ops.NewSubmitter(logger, cfg, client, contract, account))
			}
			solutionHandler := pow.CreateSolutionHandler(cfg, logger, submitters, proxy)
			for _, account := range accounts {
				var v []byte
				for i := 0; i < 10; i++ {
					// Start the miner.
					v, err = proxy.Get(db.DisputeStatusPrefix + account.Address.String())
					if err != nil {
						level.Warn(logger).Log("msg", "getting dispute status. Check if staked", "err", err)
					}
					if len(v) != 0 {
						break
					}
					select {
					case <-c: // Early exit from os.Interrupt.
						return nil
					default:
					}
					time.Sleep(1 * time.Second)
				}
				if len(v) == 0 {
					return errors.New("no status result after 10 attempts. this usually means no connection to the DB")
				}

				status, _ := hexutil.DecodeBig(string(v))
				if status.Cmp(big.NewInt(1)) != 0 {
					return errors.New("miner is not able to mine with current status")
				}

			}
			// the Miner component.
			miner, err := ops.CreateMiningManager(logger, ctx, cfg, proxy, contract, tasker, solutionHandler)
			if err != nil {
				return errors.Wrapf(err, "creating miner")
			}
			g.Add(func() error {
				return miner.Start()
			}, func(error) {
				miner.Stop()
			})
		}
	}
	level.Info(logger).Log("msg", "main shutdown complete")
	return nil
}

func getAccountFor(accounts []*rpc.Account, accountNo int) (*rpc.Account, error) {
	if accountNo < 0 || accountNo >= len(accounts) {
		return nil, errors.New("account not found")
	}
	return accounts[accountNo], nil
}
