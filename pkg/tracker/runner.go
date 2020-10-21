// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	tellor "github.com/tellor-io/TellorMiner/abi/contracts"
	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/rpc"
	"github.com/tellor-io/TellorMiner/pkg/util"
)

var runnerLog = util.NewLogger("tracker", "Runner")

// Runner will execute all configured trackers.
type Runner struct {
	client       rpc.ETHClient
	db           db.DB
	readyChannel chan bool
}

// NewRunner will create a new runner instance.
func NewRunner(client rpc.ETHClient, db db.DB) (*Runner, error) {
	return &Runner{client: client, db: db, readyChannel: make(chan bool, 1)}, nil
}

// Start will kick off the runner until the given exit channel selects.
func (r *Runner) Start(ctx context.Context, exitCh chan int) error {
	cfg := config.GetConfig()
	trackerNames := cfg.Trackers
	var trackers []Tracker
	for name, activated := range trackerNames {
		if activated {
			t, err := createTracker(name)
			if err != nil {
				runnerLog.Error("Problem creating tracker: %s\n", err.Error())
				continue
			}
			trackers = append(trackers, t...)
		}
	}
	if len(trackers) == 0 {
		// Set as ready and listen the exit channel to not block.
		r.readyChannel <- true
		go func() {
			<-exitCh
		}()
		return nil
	}
	runnerLog.Info("Created %d trackers", len(trackers))

	var err error
	masterInstance := ctx.Value(tellorCommon.MasterContractContextKey)
	if masterInstance == nil {
		contractAddress := common.HexToAddress(cfg.ContractAddress)
		masterInstance, err = tellor.NewTellorMaster(contractAddress, r.client)
		if err != nil {
			runnerLog.Error("Problem creating tellor master instance: %v\n", err)
			return err
		}
		ctx = context.WithValue(ctx, tellorCommon.MasterContractContextKey, masterInstance)
	}

	runnerLog.Info("Trackers will run every %v\n", cfg.TrackerSleepCycle)
	ticker := time.NewTicker(cfg.TrackerSleepCycle.Duration / time.Duration(len(trackers)))
	if ctx.Value(tellorCommon.ClientContextKey) == nil {
		ctx = context.WithValue(ctx, tellorCommon.ClientContextKey, r.client)
	}
	if ctx.Value(tellorCommon.DBContextKey) == nil {
		ctx = context.WithValue(ctx, tellorCommon.DBContextKey, r.db)
	}

	// after first run, let others know that tracker output data is ready for use.
	doneFirstExec := make(chan bool, len(trackers))
	go func(n int) {
		for i := 0; i < n; i++ {
			<-doneFirstExec
		}
		r.readyChannel <- true
	}(len(trackers))
	runnerLog.Info("Waiting for trackers to complete initial requests")

	//run the trackers until we quit
	go func() {
		i := 0
		for {
			select {
			case <-exitCh:
				{
					runnerLog.Info("Exiting run loop")
					ticker.Stop()
					return
				}
			case <-ticker.C:
				{
					//runnerLog.Info("Running trackers...")
					go func(count int) {
						idx := count % len(trackers)
						err := trackers[idx].Exec(ctx)
						if err != nil {
							runnerLog.Error("Problem in tracker %s: %v\n", trackers[idx].String(), err)
						}
						//only increment this the first time a tracker is run
						if count < len(trackers) {
							doneFirstExec <- true
						}
					}(i)
					i++
				}
			}
		}
	}()

	return nil
}

// Ready provides notification channel to know that the tracker data output is ready for use.
func (r *Runner) Ready() chan bool {
	return r.readyChannel
}