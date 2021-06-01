// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kit/kit/log/level"
	"github.com/oklog/run"
	"github.com/pkg/errors"
	promConfig "github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/storage/remote"
	"github.com/prometheus/prometheus/tsdb"
	"github.com/tellor-io/telliot/pkg/aggregator"
	"github.com/tellor-io/telliot/pkg/cli"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/mining"
	"github.com/tellor-io/telliot/pkg/reward"
	"github.com/tellor-io/telliot/pkg/submitter/tellor"
	"github.com/tellor-io/telliot/pkg/submitter/tellorAccess"
	"github.com/tellor-io/telliot/pkg/tasker"
	"github.com/tellor-io/telliot/pkg/tracker/gasPrice"
	"github.com/tellor-io/telliot/pkg/tracker/index"
	"github.com/tellor-io/telliot/pkg/tracker/profit"
	"github.com/tellor-io/telliot/pkg/transactor"
	"github.com/tellor-io/telliot/pkg/web"
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
	logger := logging.NewLogger()

	cfg, err := config.ParseConfig(logger, string(c.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	ctx := context.Background()
	client, contract, accounts, err := createTellorVariables(ctx, logger, cfg.Ethereum)
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
	return cli.Transfer(ctx, logger, client, contract, account, address.addr, amount.Int)

}

type approveCmd tokenCmd

func (c *approveCmd) Run() error {
	logger := logging.NewLogger()

	cfg, err := config.ParseConfig(logger, string(c.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	ctx := context.Background()
	client, contract, accounts, err := createTellorVariables(ctx, logger, cfg.Ethereum)
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
	return cli.Approve(ctx, logger, client, contract, account, address.addr, amount.Int)
}

type accountsCmd struct {
	Config configPath `type:"existingfile" help:"path to config file"`
}

func (a *accountsCmd) Run() error {
	logger := logging.NewLogger()

	cfg, err := config.ParseConfig(logger, string(a.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	ctx := context.Background()
	_, _, accounts, err := createTellorVariables(ctx, logger, cfg.Ethereum)
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
	logger := logging.NewLogger()

	cfg, err := config.ParseConfig(logger, string(b.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	ctx := context.Background()
	client, contract, _, err := createTellorVariables(ctx, logger, cfg.Ethereum)
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
	return cli.Balance(ctx, logger, client, contract, addr.addr)
}

type depositCmd struct {
	Config  configPath `type:"existingfile" help:"path to config file"`
	Account int        `arg:"" optional:""`
}

func (d depositCmd) Run() error {
	logger := logging.NewLogger()

	cfg, err := config.ParseConfig(logger, string(d.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	ctx := context.Background()
	client, contract, accounts, err := createTellorVariables(ctx, logger, cfg.Ethereum)
	if err != nil {
		return errors.Wrapf(err, "creating tellor variables")
	}
	account, err := getAccountFor(accounts, d.Account)
	if err != nil {
		return err
	}
	return cli.Deposit(ctx, logger, client, contract, account)

}

type withdrawCmd struct {
	Config  configPath `type:"existingfile" help:"path to config file"`
	Address string     `arg:"" required:""`
	Account int        `arg:"" optional:""`
}

func (w withdrawCmd) Run() error {
	logger := logging.NewLogger()

	cfg, err := config.ParseConfig(logger, string(w.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	ctx := context.Background()
	client, contract, accounts, err := createTellorVariables(ctx, logger, cfg.Ethereum)
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
	return cli.WithdrawStake(ctx, logger, client, contract, account)

}

type requestCmd struct {
	Config  configPath `type:"existingfile" help:"path to config file"`
	Account int        `arg:"" optional:""`
}

func (r requestCmd) Run() error {
	logger := logging.NewLogger()

	cfg, err := config.ParseConfig(logger, string(r.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	ctx := context.Background()
	client, contract, accounts, err := createTellorVariables(ctx, logger, cfg.Ethereum)
	if err != nil {
		return errors.Wrapf(err, "creating tellor variables")
	}
	account, err := getAccountFor(accounts, r.Account)
	if err != nil {
		return err
	}
	return cli.RequestStakingWithdraw(ctx, logger, client, contract, account)
}

type statusCmd struct {
	Config  configPath `type:"existingfile" help:"path to config file"`
	Account int        `arg:"" optional:""`
}

func (s statusCmd) Run() error {
	logger := logging.NewLogger()

	cfg, err := config.ParseConfig(logger, string(s.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	ctx := context.Background()
	client, contract, accounts, err := createTellorVariables(ctx, logger, cfg.Ethereum)
	if err != nil {
		return errors.Wrapf(err, "creating tellor variables")
	}
	account, err := getAccountFor(accounts, s.Account)
	if err != nil {
		return err
	}
	return cli.ShowStatus(ctx, logger, client, contract, account)
}

type migrateCmd struct {
	Config configPath `type:"existingfile" help:"path to config file"`
}

func (s migrateCmd) Run() error {
	logger := logging.NewLogger()

	cfg, err := config.ParseConfig(logger, string(s.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	ctx := context.Background()
	client, contract, accounts, err := createTellorVariables(ctx, logger, cfg.Ethereum)
	if err != nil {
		return errors.Wrapf(err, "creating tellor variables")
	}

	// Do migration for each account.
	for _, account := range accounts {
		level.Info(logger).Log("msg", "TRB migration", "account", account.Address.String())
		auth, err := ethereum.PrepareEthTransaction(ctx, client, account)
		if err != nil {
			return errors.Wrap(err, "prepare ethereum transaction")
		}

		tx, err := contract.Migrate(auth)
		if err != nil {
			return errors.Wrap(err, "contract failed")
		}
		level.Info(logger).Log("msg", "TRB migrated", "txHash", tx.Hash().Hex())
	}
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
	logger := logging.NewLogger()

	cfg, err := config.ParseConfig(logger, string(n.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	ctx := context.Background()
	client, contract, accounts, err := createTellorVariables(ctx, logger, cfg.Ethereum)
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
	return cli.Dispute(ctx, logger, client, contract, account, requestID.Int, timestamp.Int, minerIndex.Int)
}

type voteCmd struct {
	Config    configPath `type:"existingfile" help:"path to config file"`
	disputeId string     `arg:""  help:"the dispute id"`
	support   bool       `arg:""  help:"true or false"`
	Account   int        `arg:"" optional:""`
}

func (v voteCmd) Run() error {
	logger := logging.NewLogger()

	cfg, err := config.ParseConfig(logger, string(v.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	ctx := context.Background()
	client, contract, accounts, err := createTellorVariables(ctx, logger, cfg.Ethereum)
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
	return cli.Vote(ctx, logger, client, contract, account, disputeID.Int, v.support)
}

type showCmd struct {
	Config  configPath `type:"existingfile" help:"path to config file"`
	Account int        `arg:"" optional:""`
}

func (s showCmd) Run() error {
	logger := logging.NewLogger()

	cfg, err := config.ParseConfig(logger, string(s.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	ctx := context.Background()
	client, contract, accounts, err := createTellorVariables(ctx, logger, cfg.Ethereum)
	if err != nil {
		return errors.Wrapf(err, "creating tellor variables")
	}
	account, err := getAccountFor(accounts, s.Account)
	if err != nil {
		return err
	}

	// Open the TSDB database.
	var querable storage.SampleAndChunkQueryable
	if cfg.Db.RemoteHost != "" {
		querable, err = remoteDB(cfg.Db)
		if err != nil {
			return errors.Wrapf(err, "opening remote tsdb DB")
		}
	} else {
		if err := os.MkdirAll(cfg.Db.Path, 0777); err != nil {
			return errors.Wrapf(err, "creating tsdb DB folder")
		}
		tsdbOptions := tsdb.DefaultOptions()
		tsdbOptions.NoLockfile = true
		querable, err = tsdb.Open(cfg.Db.Path, nil, nil, tsdbOptions)
		if err != nil {
			return errors.Wrapf(err, "opening tsdb DB")
		}
	}

	aggregator, err := aggregator.New(logger, ctx, cfg.Aggregator, querable, client)
	if err != nil {
		return errors.Wrapf(err, "creating aggregator")
	}

	return cli.List(ctx, cfg.Disputer, logger, client, contract, account, aggregator)
}

type dataserverCmd struct {
	Config configPath `type:"existingfile" help:"path to config file"`
}

func (self dataserverCmd) Run() error {
	logger := logging.NewLogger()

	cfg, err := config.ParseConfig(logger, string(self.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	// Defining a global context for starting and stopping of components.
	ctx := context.Background()

	// We define our run groups here.
	var g run.Group
	// Run groups.
	{
		// Handle interupts.
		g.Add(run.SignalHandler(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM))

		// Open the TSDB database.
		tsdbOptions := tsdb.DefaultOptions()
		// 48h are enough as the aggregator needs data only 24 hours in the past.
		tsdbOptions.RetentionDuration = int64(2 * 24 * time.Hour / time.Millisecond)
		if err := os.MkdirAll(cfg.Db.Path, 0777); err != nil {
			return errors.Wrapf(err, "creating tsdb DB folder")
		}
		tsDB, err := tsdb.Open(cfg.Db.Path, nil, nil, tsdbOptions)
		if err != nil {
			return errors.Wrapf(err, "creating tsdb DB")
		}

		defer func() {
			if err := tsDB.Close(); err != nil {
				level.Error(logger).Log("msg", "closing the tsdb", "err", err)
			}
		}()

		// Index tracker.

		// The client is needed when the api requests data from the blockchain.
		// TODO create an eth client only if the api config file has eth address.
		client, err := ethereum.NewClient(logger, cfg.Ethereum, os.Getenv(config.NodeURLEnvName))
		if err != nil {
			return errors.Wrap(err, "create rpc client instance")
		}

		index, err := index.New(logger, ctx, cfg.IndexTracker, tsDB, client)
		if err != nil {
			return errors.Wrapf(err, "creating index tracker")
		}

		g.Add(func() error {
			err := index.Run()
			level.Info(logger).Log("msg", "index shutdown complete")
			return err
		}, func(error) {
			index.Stop()
		})

		// Web/Api server.
		{
			srv, err := web.New(logger, ctx, tsDB, cfg.Web)
			if err != nil {
				return errors.Wrap(err, "create web server")
			}
			g.Add(func() error {
				err := srv.Start()
				level.Info(logger).Log("msg", "web server shutdown complete")
				return err
			}, func(error) {
				srv.Stop()
			})
		}
	}

	if err := g.Run(); err != nil {
		level.Error(logger).Log("msg", "main exited with error", "err", err)
		return err
	}

	level.Info(logger).Log("msg", "main shutdown complete")
	return nil
}

type mineCmd struct {
	Config configPath `type:"existingfile" help:"path to config file"`
}

func (self mineCmd) Run() error {
	logger := logging.NewLogger()

	cfg, err := config.ParseConfig(logger, string(self.Config))
	if err != nil {
		return errors.Wrapf(err, "creating config")
	}

	// Defining a global context for starting and stopping of components.
	ctx := context.Background()

	client, contract, accounts, err := createTellorVariables(ctx, logger, cfg.Ethereum)
	if err != nil {
		return errors.Wrapf(err, "creating tellor variables")
	}

	// We define our run groups here.
	var g run.Group
	// Run groups.
	{
		// Handle interupts.
		g.Add(run.SignalHandler(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM))

		// Open a local or remote instance of the TSDB database.
		var tsDB storage.SampleAndChunkQueryable
		if cfg.Db.RemoteHost != "" {
			tsDB, err = remoteDB(cfg.Db)
			if err != nil {
				return errors.Wrapf(err, "opening remote tsdb DB")
			}
		} else {
			// Open the TSDB database.
			tsdbOptions := tsdb.DefaultOptions()
			// 2 days are enough as the aggregator needs data only 24 hours in the past.
			tsdbOptions.RetentionDuration = int64(2 * 24 * time.Hour)
			_tsDB, err := tsdb.Open(cfg.Db.Path, nil, nil, tsdbOptions)
			if err != nil {
				return errors.Wrapf(err, "opening local tsdb DB")
			}
			defer func() {
				if err := _tsDB.Close(); err != nil {
					level.Error(logger).Log("msg", "closing the tsdb", "err", err)
				}
			}()
			tsDB = _tsDB
		}

		// Web/Api server.
		{
			srv, err := web.New(logger, ctx, tsDB, cfg.Web)
			if err != nil {
				return errors.Wrap(err, "create web server")
			}
			g.Add(func() error {
				err := srv.Start()
				level.Info(logger).Log("msg", "web server shutdown complete")
				return err
			}, func(error) {
				srv.Stop()
			})
		}

		// Index tracker.
		// Run only when not using remote DB.
		if cfg.Db.RemoteHost == "" {
			_tsDB, ok := tsDB.(*tsdb.DB)
			if !ok {
				return errors.Wrapf(err, "tsdb is not a writable DB instance")
			}
			index, err := index.New(logger, ctx, cfg.IndexTracker, _tsDB, client)
			if err != nil {
				return errors.Wrapf(err, "creating index tracker")
			}

			g.Add(func() error {
				err := index.Run()
				level.Info(logger).Log("msg", "index shutdown complete")
				return err
			}, func(error) {
				index.Stop()
			})
		}

		// Aggregator.
		aggregator, err := aggregator.New(logger, ctx, cfg.Aggregator, tsDB, client)
		if err != nil {
			return errors.Wrapf(err, "creating aggregator")
		}

		// Profit tracker.
		var accountAddrs []common.Address
		for _, acc := range accounts {
			accountAddrs = append(accountAddrs, acc.Address)
		}
		profitTracker, err := profit.NewProfitTracker(logger, ctx, cfg.ProfitTracker, client, contract, accountAddrs)
		if err != nil {
			return errors.Wrapf(err, "creating profit tracker")
		}
		g.Add(func() error {
			err := profitTracker.Start()
			level.Info(logger).Log("msg", "profit shutdown complete")
			return err
		}, func(error) {
			profitTracker.Stop()
		})

		if cfg.SubmitterTellor.Enabled {
			// Event tasker.
			tasker, taskerChs, err := tasker.NewTasker(ctx, logger, cfg.Tasker, client, contract, accounts)
			if err != nil {
				return errors.Wrapf(err, "creating tasker")
			}
			g.Add(func() error {
				err := tasker.Start()
				level.Info(logger).Log("msg", "tasker shutdown complete")
				return err
			}, func(error) {
				tasker.Stop()
			})

			// Create a submitter for each account.
			gasPriceTracker := gasPrice.New(logger, client)
			for _, account := range accounts {
				transactor, err := transactor.NewTransactor(logger, cfg.Transactor, gasPriceTracker, client, account, contract)
				if err != nil {
					return errors.Wrapf(err, "creating transactor")
				}

				reward := reward.NewReward(logger, aggregator, contract)
				// Get a channel on which it listens for new data to submit.
				submitter, submitterCh, err := tellor.New(
					ctx,
					logger,
					cfg.SubmitterTellor,
					client,
					contract,
					account,
					reward,
					transactor,
					gasPriceTracker,
					aggregator,
				)
				if err != nil {
					return errors.Wrapf(err, "creating tellor submitter")
				}
				g.Add(func() error {
					err := submitter.Start()
					level.Info(logger).Log("msg", "tellor submitter shutdown complete",
						"addr", account.Address.String(),
					)
					return err
				}, func(error) {
					submitter.Stop()
				})

				// Will be used to cancel pending submissions.
				tasker.AddSubmitCanceler(submitter)

				// The Miner component.
				miner, err := mining.NewMiningManager(logger, ctx, cfg.Mining, contract, taskerChs[account.Address.String()], submitterCh, client)
				if err != nil {
					return errors.Wrapf(err, "creating miner")
				}
				g.Add(func() error {
					err := miner.Start()
					level.Info(logger).Log("msg", "miner shutdown complete",
						"addr", account.Address.String(),
					)
					return err
				}, func(error) {
					miner.Stop()
				})
			}
		}

		if cfg.SubmitterTellorAccess.Enabled {
			// Create a submitter for each account.
			for _, account := range accounts {
				submitter, err := tellorAccess.New(
					ctx,
					logger,
					cfg.SubmitterTellorAccess,
					client,
					contract,
					account,
					aggregator,
				)
				if err != nil {
					return errors.Wrapf(err, "creating tellor access submitter")
				}
				g.Add(func() error {
					err := submitter.Start()
					level.Info(logger).Log("msg", "tellor access submitter shutdown complete",
						"addr", account.Address.String(),
					)
					return err
				}, func(error) {
					submitter.Stop()
				})
			}

		}

	}

	if err := g.Run(); err != nil {
		level.Error(logger).Log("msg", "main exited with error", "err", err)
		return err
	}

	level.Info(logger).Log("msg", "main shutdown complete")
	return nil
}

func remoteDB(cfg db.Config) (storage.SampleAndChunkQueryable, error) {

	url, err := url.Parse("http://" + cfg.RemoteHost + ":" + strconv.Itoa(int(cfg.RemotePort)) + "/api/v1/read")
	if err != nil {
		return nil, err
	}
	client, err := remote.NewReadClient("", &remote.ClientConfig{
		URL:     &promConfig.URL{URL: url},
		Timeout: model.Duration(cfg.RemoteTimeout.Duration),
		HTTPClientConfig: promConfig.HTTPClientConfig{
			FollowRedirects: true,
		},
	})
	if err != nil {
		return nil, err
	}
	return remote.NewSampleAndChunkQueryableClient(
		client,
		labels.Labels{},
		[]*labels.Matcher{},
		true,
		func() (i int64, err error) { return 0, nil },
	), nil
}

func getAccountFor(accounts []*ethereum.Account, accountNo int) (*ethereum.Account, error) {
	if accountNo < 0 || accountNo >= len(accounts) {
		return nil, errors.New("account not found")
	}
	return accounts[accountNo], nil
}
