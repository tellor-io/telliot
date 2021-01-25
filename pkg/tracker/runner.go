// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
)

// Runner will execute all configured trackers.
type Runner struct {
	db           db.DataServerProxy
	client       contracts.ETHClient
	contract     *contracts.Tellor
	account      *rpc.Account
	readyChannel chan bool
	logger       log.Logger
	config       *config.Config
	trackerErr   *prometheus.CounterVec
}

// NewRunner will create a new runner instance.
func NewRunner(logger log.Logger, config *config.Config, db db.DataServerProxy, client contracts.ETHClient, contract *contracts.Tellor, account *rpc.Account) (*Runner, error) {
	return &Runner{
		config:       config,
		db:           db,
		client:       client,
		contract:     contract,
		account:      account,
		readyChannel: make(chan bool, 1),
		logger:       log.With(logger, "component", "runner"),
		trackerErr: promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: "telliot",
			Subsystem: "tracker",
			Name:      "errors_total",
			Help:      "The total number of tracker errors. Usually caused by API throtling.",
		}, []string{"id"}),
	}, nil
}

// Start will kick off the runner until the given exit channel selects.
func (r *Runner) Start(ctx context.Context, exitCh chan int) error {
	trackerNames := r.config.Trackers
	var trackers []Tracker
	for name, activated := range trackerNames {
		if activated {
			level.Info(r.logger).Log("msg", "starting tracker", "name", name)
			t, err := createTracker(name, r.logger, r.config, r.db, r.client, r.contract, r.account)
			if err != nil {
				return errors.Wrapf(err, "creating tracker. Name: %s", name)
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

	level.Info(r.logger).Log("msg", "starting trackers", "sleepCycle", r.config.TrackerSleepCycle)
	ticker := time.NewTicker(r.config.TrackerSleepCycle.Duration / time.Duration(len(trackers)))

	// after first run, let others know that tracker output data is ready for use.
	doneFirstExec := make(chan bool, len(trackers))
	go func(n int) {
		for i := 0; i < n; i++ {
			<-doneFirstExec
		}
		r.readyChannel <- true
	}(len(trackers))
	level.Info(r.logger).Log("msg", "waiting for trackers to complete initial requests")

	// Run the trackers until sigterm.
	go func() {
		i := 0
		for {
			select {
			case <-exitCh:
				{
					level.Info(r.logger).Log("msg", "exiting run loop")
					ticker.Stop()
					return
				}
			case <-ticker.C:
				{
					go func(count int) {
						idx := count % len(trackers)
						err := trackers[idx].Exec(ctx)
						if err != nil {
							r.trackerErr.With(prometheus.Labels{"id": trackers[idx].String()}).(prometheus.Counter).Inc()
							level.Warn(r.logger).Log("msg", "problem in tracker", "tracker", trackers[idx].String(), "err", err)
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
