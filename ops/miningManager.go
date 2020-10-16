// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ops

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/pow"
	"github.com/tellor-io/TellorMiner/util"
)

type WorkSource interface {
	GetWork(input chan *pow.Work) (*pow.Work, bool)
}

type SolutionSink interface {
	Submit(context.Context, *pow.Result) bool
}

// MiningMgr holds items for mining and requesting data.
type MiningMgr struct {
	//primary exit channel
	exitCh  chan os.Signal
	log     *util.Logger
	Running bool

	group      *pow.MiningGroup
	tasker     WorkSource
	solHandler SolutionSink
	solution   *pow.Result

	dataRequester *DataRequester
}

// CreateMiningManager creates a new manager that mananges mining and data requests.
func CreateMiningManager(ctx context.Context, exitCh chan os.Signal, submitter tellorCommon.TransactionSubmitter, logger log.Logger) (*MiningMgr, error) {
	cfg := config.GetConfig()

	group, err := pow.SetupMiningGroup(cfg)
	if err != nil {

		return nil, fmt.Errorf("failed to setup miners: %s", err.Error())
	}

	mng := &MiningMgr{
		exitCh:     exitCh,
		log:        util.NewLogger("ops", "MiningMgr"),
		Running:    false,
		group:      group,
		tasker:     nil,
		solution:   nil,
		solHandler: nil,
	}

	if cfg.EnablePoolWorker {
		pool := pow.CreatePool(cfg, group)
		mng.tasker = pool
		mng.solHandler = pool
	} else {
		proxy := ctx.Value(tellorCommon.DataProxyKey).(db.DataServerProxy)
		mng.tasker = pow.CreateTasker(cfg, proxy)
		mng.solHandler = pow.CreateSolutionHandler(cfg, submitter, proxy)
		if cfg.RequestData > 0 {
			level.Info(logger).Log("msg", "dataRequester created")
			mng.dataRequester = CreateDataRequester(exitCh, submitter, cfg.RequestDataInterval.Duration, proxy)
		}
	}
	return mng, nil
}

// Start will start the mining run loop.
func (mgr *MiningMgr) Start(ctx context.Context) {
	mgr.Running = true
	go func(ctx context.Context) {
		cfg := config.GetConfig()

		ticker := time.NewTicker(cfg.MiningInterruptCheckInterval.Duration)

		//if you make these buffered, think about the effects on synchronization!
		input := make(chan *pow.Work)
		output := make(chan *pow.Result)
		if cfg.RequestData > 0 {
			if err := mgr.dataRequester.Start(ctx); err != nil {
				level.Error(logger).Log("error starting the data requester", err)
			}
		}

		//start the mining group
		go mgr.group.Mine(input, output)

		// sends work to the mining group
		sendWork := func() {
			if cfg.EnablePoolWorker {
				mgr.tasker.GetWork(input)
			} else {
				work, instantSubmit := mgr.tasker.GetWork(input)
				if instantSubmit {
					if mgr.solution == nil {
						level.Info(logger).Log("msg", "Instant Submit Called! ")
						mgr.solution = &pow.Result{Work: work, Nonce: "1"}
						goodSubmit := mgr.solHandler.Submit(ctx, mgr.solution)
						if goodSubmit {
							level.Info(logger).Log("msg", "good submit")
						}
					}
				} else if work != nil {
					mgr.solution = nil
					input <- work
				}
			}
		}
		// Send the initial challenge.
		sendWork()
		for {
			select {
			// Boss wants us to quit for the day.
			case <-mgr.exitCh:
				//exit
				input <- nil

			//found a solution
			case result := <-output:
				if result == nil {
					mgr.Running = false
					return
				}
				mgr.solution = result
				goodSubmit := mgr.solHandler.Submit(ctx, mgr.solution)
				if goodSubmit {
					level.Info(logger).Log("msg", "solution submitted")
				}
				sendWork()

			//time to check for a new challenge
			case <-ticker.C:
				sendWork()
			}
		}
	}(ctx)
}
