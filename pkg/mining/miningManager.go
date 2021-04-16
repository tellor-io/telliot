// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package mining

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/logging"
)

type Config struct {
	LogLevel  string
	Address   common.Address
	Heartbeat time.Duration
}

type SolutionSink interface {
	Submit(context.Context, *Result) (*types.Transaction, error)
}

const NumProcessors = 1

func SetupMiningGroup(logger log.Logger, cfg Config, contractInstance *contracts.ITellor) (*MiningGroup, error) {
	var hashers []Hasher
	level.Info(logger).Log("msg", "starting CPU mining", "threads", NumProcessors)
	for i := 0; i < NumProcessors; i++ {
		hashers = append(hashers, NewCpuMiner(int64(i)))
	}
	miningGrp, err := NewMiningGroup(logger, cfg, hashers, contractInstance)
	if err != nil {
		return nil, errors.Wrap(err, "creating new mining group")
	}
	return miningGrp, nil
}

// MiningMgr manages mining, submiting a solution and requesting data.
// In the tellor contract a solution is saved in slots where a value is valid only when it has 5 confirmed slots.
// The manager tracks tx costs and profitThreshold is set it skips any transactions below the profit threshold.
// The profit is calculated the same way as in the Tellor contract.
// Transaction cost for submitting in each slot might be different so because of this
// the manager needs to complete few transaction to gather the tx cost for each slot.
type MiningMgr struct {
	ctx              context.Context
	close            context.CancelFunc
	logger           log.Logger
	ethClient        contracts.ETHClient
	group            *MiningGroup
	taskerCh         chan *Work
	submitterCh      chan *Result
	contractInstance *contracts.ITellor
	toMineInput      chan *Work
	solutionOutput   chan *Result
}

// NewMiningManager is the MiningMgr constructor.
func NewMiningManager(
	logger log.Logger,
	ctx context.Context,
	cfg Config,
	contractInstance *contracts.ITellor,
	taskerCh chan *Work,
	submitterCh chan *Result,
	client contracts.ETHClient,
) (*MiningMgr, error) {

	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(logger, "component", ComponentName, "addr", cfg.Address.String()[:6])

	group, err := SetupMiningGroup(logger, cfg, contractInstance)
	if err != nil {
		return nil, errors.Wrap(err, "setup MiningGroup")
	}

	ctx, close := context.WithCancel(ctx)
	mng := &MiningMgr{
		ctx:              ctx,
		close:            close,
		logger:           logger,
		group:            group,
		taskerCh:         taskerCh,
		submitterCh:      submitterCh,
		contractInstance: contractInstance,
		ethClient:        client,
		toMineInput:      make(chan *Work),
		solutionOutput:   make(chan *Result),
	}
	return mng, nil
}

// Start will start the mining run loop.
func (mgr *MiningMgr) Start() error {
	// Start the mining group.
	go mgr.group.Mine(mgr.ctx, mgr.toMineInput, mgr.solutionOutput)

	for {
		select {

		// Boss wants us to quit for the day.
		case <-mgr.ctx.Done():
			return mgr.ctx.Err()

		// Found a solution.
		case solution := <-mgr.solutionOutput:
			level.Info(mgr.logger).Log("msg", "sending the solution to the submitter")
			mgr.submitterCh <- solution

		// Listen for new work from the tasker and send for mining.
		case work := <-mgr.taskerCh:
			mgr.toMineInput <- work
			level.Info(mgr.logger).Log("msg", "sent new challenge to the mining group",
				"challenge", fmt.Sprintf("%x", work.Challenge.Challenge),
				"difficulty", work.Challenge.Difficulty,
				"requestIDs", fmt.Sprintf("%+v", work.Challenge.RequestIDs),
			)
		}
	}

}

// Stop will take care of stopping the miner component.
func (mgr *MiningMgr) Stop() {
	mgr.close()
}
