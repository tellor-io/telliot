// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package mining

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/rpc"
)

type SolutionSink interface {
	Submit(context.Context, *Result) (*types.Transaction, error)
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
	submitters       []chan *Result
	database         db.DataServerProxy
	contractInstance *contracts.ITellor
	cfg              *config.Config

	toMineInput    chan *Work
	solutionOutput chan *Result
}

// CreateMiningManager is the MiningMgr constructor.
func CreateMiningManager(
	logger log.Logger,
	ctx context.Context,
	cfg *config.Config,
	database db.DataServerProxy,
	contractInstance *contracts.ITellor,
	taskerCh chan *Work,
) (*MiningMgr, error) {
	group, err := SetupMiningGroup(ctx, logger, cfg, contractInstance)
	if err != nil {
		return nil, errors.Wrap(err, "setup miners")
	}

	client, err := rpc.NewClient(logger, cfg, os.Getenv(config.NodeURLEnvName))
	if err != nil {
		return nil, errors.Wrap(err, "creating client")
	}

	//ops logging
	logger, err = logging.ApplyFilter(*cfg, ComponentName, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	ctx, close := context.WithCancel(ctx)
	mng := &MiningMgr{
		ctx:              ctx,
		close:            close,
		logger:           log.With(logger, "component", ComponentName),
		group:            group,
		taskerCh:         taskerCh,
		submitters:       []chan *Result{},
		contractInstance: contractInstance,
		cfg:              cfg,
		database:         database,
		ethClient:        client,
		toMineInput:      make(chan *Work),
		solutionOutput:   make(chan *Result),
	}
	return mng, nil
}

// Start will start the mining run loop.
func (mgr *MiningMgr) Start() error {
	// Start the mining group.
	go mgr.group.Mine(mgr.toMineInput, mgr.solutionOutput)

	for {
		select {
		// Boss wants us to quit for the day.
		case <-mgr.ctx.Done():
			return nil
		// Found a solution.
		case solution := <-mgr.solutionOutput:
			for _, submitter := range mgr.submitters {
				submitter <- solution
			}
			level.Info(mgr.logger).Log("msg", "sent solution to all subscribed submitters")
		// Listen for new work from the tasker and send for mining.
		case work := <-mgr.taskerCh:
			mgr.toMineInput <- work
			level.Info(mgr.logger).Log("msg", "sent new chalenge to the mining group",
				"addr", work.PublicAddr,
				"challenge", fmt.Sprintf("%x", work.Challenge.Challenge),
				"difficulty", work.Challenge.Difficulty,
				"requestIDs", fmt.Sprintf("%+v", work.Challenge.RequestIDs),
			)
		}
	}
}

func (mgr *MiningMgr) Subscribe(submitCh chan *Result) {
	mgr.submitters = append(mgr.submitters, submitCh)
}

// Stop will take care of stopping the miner component.
func (mgr *MiningMgr) Stop() {
	mgr.close()
	level.Info(mgr.logger).Log("msg", "miner shutdown complete")
}
