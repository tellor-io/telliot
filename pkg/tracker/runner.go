// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/contracts/tellor"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/rpc"
)

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
func (r *Runner) Start(ctx context.Context, logger log.Logger, exitCh chan int) error {
	cfg := config.GetConfig()
	trackerNames := cfg.Trackers
	var trackers []Tracker
	for name, activated := range trackerNames {
		if activated {
			t, err := createTracker(name, logger)
			if err != nil {
				return fmt.Errorf("problem creating tracker. Name: %s, err: %s", name, err)
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
	level.Info(logger).Log("msg", fmt.Sprintf("Created %d trackers", len(trackers)))

	var err error
	masterInstance := ctx.Value(tellorCommon.ContractsTellorContextKey)
	if masterInstance == nil {
		contractAddress := common.HexToAddress(cfg.ContractAddress)
		masterInstance, err = tellor.NewTellor(contractAddress, r.client)
		if err != nil {
			return fmt.Errorf("Problem creating tellor master instance: %s", err)
		}
		ctx = context.WithValue(ctx, tellorCommon.ContractsTellorContextKey, masterInstance)
	}

	level.Info(logger).Log("msg", "Trackers will run", "sleepCycle", cfg.TrackerSleepCycle)
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
	level.Info(logger).Log("msg", "Waiting for trackers to complete initial requests")

	// Run the trackers until sigterm.
	go func() {
		i := 0
		for {
			select {
			case <-exitCh:
				{
					level.Info(logger).Log("msg", "Exiting run loop")
					ticker.Stop()
					return
				}
			case <-ticker.C:
				{
					//runnerLog.Info("Running trackers...")
					level.Debug(logger).Log("msg", "Running trackers")
					go func(count int) {
						idx := count % len(trackers)
						err := trackers[idx].Exec(ctx)
						if err != nil {
							level.Warn(logger).Log("msg", "problem in traker", "tracker", trackers[idx].String(), "err", err)
						}
						// Only the first trackers round execution.
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
