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

	"github.com/go-kit/kit/log/level"
	"github.com/oklog/run"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
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
	"github.com/tellor-io/telliot/pkg/tracker/index"
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
	var tsDBRead storage.SampleAndChunkQueryable
	if cfg.Db.RemoteHost != "" {
		tsDBRead, err = remoteDB(cfg.Db)
		if err != nil {
			return errors.Wrapf(err, "opening remote tsdb DB")
		}
	} else {
		if err := os.MkdirAll(cfg.Db.Path, 0777); err != nil {
			return errors.Wrapf(err, "creating tsdb DB folder")
		}
		tsDBRead, err = tsdb.OpenDBReadOnly(cfg.Db.Path, nil)
		if err != nil {
			return errors.Wrapf(err, "opening tsdb DB")
		}
	}

	aggregator, err := aggregator.New(logger, ctx, cfg.Aggregator, tsDBRead, client)
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
		// 2 days are enough as the aggregator needs data only 24 hours in the past.
		tsdbOptions.RetentionDuration = int64(2 * 24 * time.Hour)
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

		// Open a read only instance of TSDB database.
		tsDBRead, err := tsdb.OpenDBReadOnly(cfg.Db.Path, nil)
		if err != nil {
			return errors.Wrapf(err, "opening tsdb DB")
		}

		// Web/Api server.
		{
			srv, err := web.New(logger, ctx, tsDBRead, cfg.Web)
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

	client, _, _, err := createTellorVariables(ctx, logger, cfg.Ethereum)
	if err != nil {
		return errors.Wrapf(err, "creating tellor variables")
	}

	// We define our run groups here.
	var g run.Group
	// Run groups.
	{
		// Handle interupts.
		g.Add(run.SignalHandler(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM))

		// Open a read only instance of TSDB database.
		var tsDBRead storage.SampleAndChunkQueryable
		if cfg.Db.RemoteHost != "" {
			tsDBRead, err = remoteDB(cfg.Db)
			if err != nil {
				return errors.Wrapf(err, "opening remote tsdb DB")
			}
		} else {
			tsDBRead, err = tsdb.OpenDBReadOnly(cfg.Db.Path, nil)
			if err != nil {
				return errors.Wrapf(err, "opening tsdb DB")
			}
		}

		// Web/Api server.
		{
			srv, err := web.New(logger, ctx, tsDBRead, cfg.Web)
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
			// Open the TSDB database.
			tsdbOptions := tsdb.DefaultOptions()
			// 2 days are enough as the aggregator needs data only 24 hours in the past.
			tsdbOptions.RetentionDuration = int64(2 * 24 * time.Hour)
			tsDB, err := tsdb.Open(cfg.Db.Path, nil, nil, tsdbOptions)
			if err != nil {
				return errors.Wrapf(err, "creating tsdb DB")
			}

			defer func() {
				if err := tsDB.Close(); err != nil {
					level.Error(logger).Log("msg", "closing the tsdb", "err", err)
				}
			}()

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
		}

		// Aggregator.
		aggregator, err := aggregator.New(logger, ctx, cfg.Aggregator, tsDBRead, client)
		if err != nil {
			return errors.Wrapf(err, "creating aggregator")
		}

		value := promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: "aggregator",
			Name:      "value",
			Help:      "The aggregated value",
		},
			[]string{"id"},
		)

		ctx, cncl := context.WithCancel(ctx)

		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		g.Add(func() error {
			for {
				for i := int64(1); i <= 58; i++ {
					select {
					case <-ctx.Done():
						return nil
					default:
					}
					val, err := aggregator.GetValueForID(i, time.Now())
					if err != nil {
						level.Error(logger).Log("msg", "get value", "ID", i, "err", err)
						continue
					}
					value.With(
						prometheus.Labels{
							"id": strconv.Itoa(int(i)),
						},
					).(prometheus.Gauge).Set(val)

					level.Info(logger).Log("msg", "got value", "ID", i, "VAL", val)
				}
				select {
				case <-ctx.Done():
					return nil
				case <-ticker.C:
				}
			}
		}, func(error) {
			cncl()
		})
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
