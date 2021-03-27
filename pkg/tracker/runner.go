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
	"github.com/tellor-io/telliot/pkg/logging"
)

const ComponentName = "tracker"

// Runner will execute all configured trackers.
type Runner struct {
	db           db.DataServerProxy
	client       contracts.ETHClient
	contract     *contracts.ITellor
	accounts     []*config.Account
	readyChannel chan bool
	logger       log.Logger
	config       *config.Config
	trackerErr   *prometheus.CounterVec
}

// NewRunner will create a new runner instance.
func NewRunner(logger log.Logger, config *config.Config, db db.DataServerProxy, client contracts.ETHClient, contract *contracts.ITellor, accounts []*config.Account) (*Runner, error) {
	logger, err := logging.ApplyFilter(*config, ComponentName, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	return &Runner{
		config:       config,
		db:           db,
		client:       client,
		contract:     contract,
		accounts:     accounts,
		readyChannel: make(chan bool, 1),
		logger:       log.With(logger, "component", ComponentName),
		trackerErr: promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: "telliot",
			Subsystem: "tracker",
			Name:      "errors_total",
			Help:      "The total number of tracker errors. Usually caused by API throtling.",
		}, []string{"id"}),
	}, nil
}

// Start will kick off the runner until the given exit channel selects.
func (r *Runner) Start(ctx context.Context) error {
	var trackers []Tracker
	for name, activated := range r.config.Trackers.Names {
		if activated {
			level.Info(r.logger).Log("msg", "starting tracker", "name", name)
			t, err := createTracker(name, r.logger, r.config, r.db, r.client, r.contract, r.accounts)
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
			<-ctx.Done()
		}()
		return nil
	}

	level.Info(r.logger).Log("msg", "starting trackers", "sleepCycle", r.config.Trackers.SleepCycle)
	ticker := time.NewTicker(r.config.Trackers.SleepCycle.Duration / time.Duration(len(trackers)))

	defer ticker.Stop()
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
			case <-ctx.Done():
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
							level.Warn(r.logger).Log("msg", "tracker exec", "tracker", trackers[idx].String(), "err", err)
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
