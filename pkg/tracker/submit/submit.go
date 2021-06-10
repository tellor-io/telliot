// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package submit

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/pkg/timestamp"
	"github.com/prometheus/prometheus/tsdb"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/contracts/tellor"
	"github.com/tellor-io/telliot/pkg/logging"
)

const ComponentName = "submitTracker"

type Config struct {
	LogLevel string
}

type Submit struct {
	ctx                  context.Context
	close                context.CancelFunc
	cfg                  Config
	tsDB                 *tsdb.DB
	client               contracts.ETHClient
	contractTellor       *contracts.ITellor
	contractTellorAccess *contracts.ITellorAccess
	pendingTx            map[string]context.CancelFunc
	mtx                  sync.Mutex
	logger               log.Logger
}

func New(
	logger log.Logger,
	cfg Config,
	tsDB *tsdb.DB,
	client contracts.ETHClient,
	contractTellor *contracts.ITellor,
	contractTellorAccess *contracts.ITellorAccess,
) (*Submit, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(logger, "component", ComponentName)
	return &Submit{
		client:               client,
		contractTellor:       contractTellor,
		contractTellorAccess: contractTellorAccess,
		cfg:                  cfg,
		tsDB:                 tsDB,
		logger:               logger,
	}, nil
}

func (self *Submit) Start(ctx context.Context) error {
	go self.monitorTellor()
	// go self.monitorTellorAccess()
	return nil
}

func (self *Submit) Stop() {
	self.close()
}

func (self *Submit) monitorTellor() {
	var err error
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	logger := log.With(self.logger, "contract", "tellor")

	var sub event.Subscription
	events := make(chan *tellor.TellorNonceSubmitted)

	for {
		select {
		case <-self.ctx.Done():
			return
		default:
		}
		sub, err = self.newSubTellor(events)
		if err != nil {
			level.Error(logger).Log("msg", "initial subscribing to events failed")
			<-ticker.C
			continue
		}
		break
	}

	for {
		select {
		case <-self.ctx.Done():
			return
		case err := <-sub.Err():
			if err != nil {
				level.Error(logger).Log(
					"msg",
					"subscription error",
					"err", err)
			}

			// Trying to resubscribe until it succeeds.
			for {
				select {
				case <-self.ctx.Done():
					return
				default:
				}
				sub, err = self.newSubTellor(events)
				if err != nil {
					level.Error(logger).Log("msg", "re-subscribing to events failed", "err", err)
					<-ticker.C
					continue
				}
				break
			}
			level.Info(logger).Log("msg", "re-subscribed to events")
		case event := <-events:
			if event.Raw.Removed {
				self.mtx.Lock()
				pending, ok := self.pendingTx[event.Raw.TxHash.String()]
				if !ok {
					level.Error(self.logger).Log("msg", "missing pending TX for removed event")
					continue
				}
				pending()
				delete(self.pendingTx, event.Raw.TxHash.String())
				self.mtx.Unlock()
			}
			go func(ctx context.Context) {
				ticker := time.NewTicker(1 * time.Minute)
				defer ticker.Stop()

				select {
				case <-ticker.C:
					if err := self.setValTellor(event); err != nil {
						level.Error(logger).Log(
							"msg", "setting value",
							"event", fmt.Sprintf("%+v", event),
							"err", err,
						)
					}
				case <-ctx.Done():
					return
				}
				self.mtx.Lock()
				delete(self.pendingTx, event.Raw.TxHash.String())
				self.mtx.Unlock()
			}(self.ctx)
		}
	}
}

func (self *Submit) setValTellor(event *tellor.TellorNonceSubmitted) error {
	level.Debug(self.logger).Log(
		"msg", "adding event to the db",
		"contract", "tellor",
		"event", fmt.Sprintf("%+v", event),
	)
	ts := timestamp.FromTime(time.Now())

	for i, val := range event.Value {
		appender := self.tsDB.Appender(self.ctx)
		lbls := labels.Labels{
			labels.Label{Name: "__name__", Value: "submit_value"},
			labels.Label{Name: "contract", Value: "tellor"},
			labels.Label{Name: "id", Value: event.RequestId[i].String()},
		}

		sort.Sort(lbls) // This is important! The labels need to be sorted to avoid creating the same series with duplicate reference.

		if _, err := appender.Append(0,
			lbls,
			ts,
			float64(val.Int64()),
		); err != nil {
			return errors.Wrapf(err, "append values to the DB")
		}
	}
	return nil
}

func (self *Submit) newSubTellor(output chan *tellor.TellorNonceSubmitted) (event.Subscription, error) {
	tellorFilterer, err := tellor.NewTellorFilterer(self.contractTellor.Address, self.client)
	if err != nil {
		return nil, errors.Wrap(err, "getting instance")
	}
	sub, err := tellorFilterer.WatchNonceSubmitted(&bind.WatchOpts{Context: self.ctx}, output, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "getting channel")
	}
	return sub, nil
}
