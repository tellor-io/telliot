package main

import (
	"context"
	"math/big"
	"os"
	"os/signal"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/alecthomas/kong"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	tellorCommon "github.com/tellor-io/telliot/pkg/common"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/ops"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/util"
)

type configPath string

func (c configPath) AfterApply(kong *kong.Context) error {
	err := config.ParseConfig(string(c))
	if err != nil {
		return errors.Wrapf(err, "parsing config")
	}

	cfg := config.GetConfig()
	ctx := context.Background()
	err = util.SetupLoggingConfig(cfg.Logger)
	if err != nil {
		return errors.Wrapf(err, "parsing log config")
	}

	logger := util.SetupLogger(cfg.LogLevel)

	client, contract, account, err := setup(ctx, cfg)
	if err != nil {
		return errors.Wrapf(err, "setting up variables")
	}

	kong.BindTo(logger, (*log.Logger)(nil))
	kong.BindTo(client, (*rpc.ETHClient)(nil))
	kong.BindTo(ctx, (*context.Context)(nil))
	kong.Bind(contract)
	kong.Bind(account)
	kong.Bind(cfg)

	return nil
}

type tokenCmd struct {
	Address string `arg`
	Amount  string `arg`
}

type transferCmd tokenCmd

func (c *transferCmd) Run(ctx context.Context, logger log.Logger, client rpc.ETHClient, contract *contracts.Tellor, account *rpc.Account) error {
	address := ETHAddress{}
	err := address.Set(c.Address)
	if err != nil {
		return errors.Wrapf(err, "parsing address argument")
	}
	amount := TRBAmount{}
	err = amount.Set(c.Amount)
	if err != nil {
		return errors.Wrapf(err, "parsing amount argument")
	}
	return ops.Transfer(ctx, logger, client, contract, account, address.addr, amount.Int)
}

type approveCmd tokenCmd

func (c *approveCmd) Run(ctx context.Context, logger log.Logger, client rpc.ETHClient, contract *contracts.Tellor, account *rpc.Account) error {
	address := ETHAddress{}
	err := address.Set(c.Address)
	if err != nil {
		return errors.Wrapf(err, "parsing address argument")
	}
	amount := TRBAmount{}
	err = amount.Set(c.Amount)
	if err != nil {
		return errors.Wrapf(err, "parsing amount argument")
	}
	return ops.Approve(ctx, logger, client, contract, account, address.addr, amount.Int)
}

type balanceCmd struct {
	Address string `arg optional`
}

func (b *balanceCmd) Run(ctx context.Context, client rpc.ETHClient, contract *contracts.Tellor) error {
	addr := ETHAddress{}
	var err error
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
	return ops.Balance(ctx, client, contract.Getter, addr.addr)
}

type stakeCmd struct {
	Operation string `arg required`
}

func (s *stakeCmd) Run(ctx context.Context, logger log.Logger, client rpc.ETHClient, contract *contracts.Tellor, account *rpc.Account) error {
	switch s.Operation {
	case "deposit":
		return ops.Deposit(ctx, logger, client, contract, account)
	case "withdraw":
		return ops.WithdrawStake(ctx, logger, client, contract, account)
	case "request":
		return ops.RequestStakingWithdraw(ctx, logger, client, contract, account)
	case "status":
		return ops.ShowStatus(ctx, logger, client, contract, account)
	default:
		return errors.New("unknown stake command")
	}
}

type newDisputeCmd struct {
	requestId  string `arg required help:"the request id to dispute it"`
	timestamp  string `arg required help:"the submitted timestamp to dispute"`
	minerIndex string `arg required help:"the miner index to dispute"`
}

func (n newDisputeCmd) Run(ctx context.Context, client rpc.ETHClient, contract *contracts.Tellor, account *rpc.Account) error {
	requestID := EthereumInt{}
	err := requestID.Set(n.requestId)
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
	return ops.Dispute(ctx, client, contract, account, requestID.Int, timestamp.Int, minerIndex.Int)
}

type voteCmd struct {
	disputeId string `arg required help:"the dispute id"`
	support   bool   `arg required help:"true or false"`
}

func (v voteCmd) Run(ctx context.Context, client rpc.ETHClient, contract *contracts.Tellor, account *rpc.Account) error {
	disputeID := EthereumInt{}
	err := disputeID.Set(v.disputeId)
	if err != nil {
		return errors.Wrapf(err, "parsing argument")
	}
	return ops.Vote(ctx, client, contract, account, disputeID.Int, v.support)
}

type showCmd struct{}

func (v showCmd) Run(ctx context.Context, logger log.Logger, client rpc.ETHClient, contract *contracts.Tellor, account *rpc.Account) error {
	return ops.List(ctx, logger, client, contract, account)
}

type dataserverCmd struct{}

func (d dataserverCmd) Run(ctx context.Context, logger log.Logger, cfg *config.Config, client rpc.ETHClient, contract *contracts.Tellor, account *rpc.Account) error {
	// Create os kill sig listener.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	var ds *ops.DataServerOps
	var err error
	DB, proxy, err := migrateAndOpenDB(cfg)
	if err != nil {
		return errors.Wrapf(err, "initializing database")
	}
	ch := make(chan os.Signal)
	ds, err = ops.CreateDataServerOps(ctx, logger, cfg, DB, &proxy, client, contract, account, ch)
	if err != nil {
		return errors.Wrapf(err, "creating data server")
	}
	// Start and wait for it to be ready
	if err := ds.Start(ctx); err != nil {
		return errors.Wrapf(err, "starting data server")
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
			level.Warn(logger).Log("msg", "taking longer than expected to stop operations", "waited", time.Since(start))
		} else if dsStopped {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	level.Info(logger).Log("msg", "main shutdown complete")
	return nil
}

type mineCmd struct{}

func (m mineCmd) Run(ctx context.Context, logger log.Logger, cfg *config.Config, client rpc.ETHClient, contract *contracts.Tellor, account *rpc.Account) error {
	// Create os kill sig listener.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	exitChannels := make([]*chan os.Signal, 0)

	// cfg := config.GetConfig()
	var ds *ops.DataServerOps
	if !cfg.EnablePoolWorker {
		var err error
		DB, proxy, err := migrateAndOpenDB(cfg)
		if err != nil {
			return errors.Wrapf(err, "initializing database")
		}
		if !cfg.RemoteMining {
			ch := make(chan os.Signal)
			exitChannels = append(exitChannels, &ch)

			ds, err = ops.CreateDataServerOps(ctx, logger, cfg, DB, &proxy, client, contract, account, ch)
			if err != nil {
				return errors.Wrapf(err, "creating data server")
			}
			// Start and wait for it to be ready.
			if err := ds.Start(ctx); err != nil {
				return errors.Wrapf(err, "starting data server")
			}
			<-ds.Ready()
		}
	}
	// Start miner
	DB := ctx.Value(tellorCommon.DataProxyKey).(db.DataServerProxy)
	v, err := DB.Get(db.DisputeStatusKey)
	if err != nil {
		level.Warn(logger).Log("msg", "getting dispute status. Check if staked")
	}
	status, _ := hexutil.DecodeBig(string(v))
	if status.Cmp(big.NewInt(1)) != 0 {
		return errors.New("miner is not able to mine with current status")
	}
	ch := make(chan os.Signal)
	exitChannels = append(exitChannels, &ch)
	miner, err := ops.CreateMiningManager(logger, ch, cfg, DB, contract, account)
	if err != nil {
		return errors.Wrapf(err, "creating miner")
	}
	go func() {
		miner.Start(ctx)
	}()

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
			level.Warn(logger).Log("msg", "taking longer than expected to stop operations", "waited", time.Since(start))
		} else if dsStopped && minerStopped {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	level.Info(logger).Log("msg", "main shutdown complete")
	return nil
}
