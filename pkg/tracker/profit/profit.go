// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package profit

import (
	"context"
	"math/big"
	"time"

	"github.com/bluele/gcache"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	tellorCommon "github.com/tellor-io/telliot/pkg/common"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/contracts/tellor"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/logging"
)

const ComponentName = "profitTracker"

type ProfitTracker struct {
	client           contracts.ETHClient
	logger           log.Logger
	contractInstance *contracts.ITellor
	proxy            db.DataServerProxy
	ctx              context.Context
	stop             context.CancelFunc
	addrs            []common.Address

	cacheTXsProfit gcache.Cache
	cacheTXsCost   gcache.Cache

	submitProfit *prometheus.GaugeVec
	submitCost   *prometheus.GaugeVec
	balanceTRB   *prometheus.GaugeVec
	balanceETH   *prometheus.GaugeVec
}

func NewProfitTracker(
	logger log.Logger,
	ctx context.Context,
	cfg *config.Config,
	client contracts.ETHClient,
	contractInstance *contracts.ITellor,
	proxy db.DataServerProxy,
	addrs []common.Address,
) (*ProfitTracker, error) {
	logger, err := logging.ApplyFilter(*cfg, ComponentName, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(logger, "component", ComponentName)
	ctx, cncl := context.WithCancel(ctx)

	return &ProfitTracker{
		client:           client,
		logger:           logger,
		contractInstance: contractInstance,
		proxy:            proxy,
		addrs:            addrs,
		ctx:              ctx,
		stop:             cncl,

		cacheTXsProfit: gcache.New(50).LRU().Build(),
		cacheTXsCost:   gcache.New(50).LRU().Build(),

		submitProfit: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      "submit_profit",
			Help:      "Accumulated TRB amount from rewards for all registered addresses",
		},
			[]string{"addr"},
		),
		submitCost: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      "submit_cost",
			Help:      "Accumulated ETH cost from the submits for all registered addresses",
		},
			[]string{"addr"},
		),
		balanceTRB: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      "balance_trb",
			Help:      "Current TRB balance for all registered addresses",
		},
			[]string{"addr"},
		),
		balanceETH: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      "balance_eth",
			Help:      "Current ETH balance for all registered addresses",
		},
			[]string{"addr"},
		),
	}, nil
}

func (self *ProfitTracker) Start() error {
	level.Info(self.logger).Log("msg", "starting")

	for _, addr := range self.addrs {
		balance, err := self.getTRBBalance(addr)
		if err != nil {
			level.Error(self.logger).Log("msg", "getting initial TRB balance", "addr", addr.String(), "err", err)
		}
		level.Info(self.logger).Log("msg", "initial TRB balance", "addr", addr.String(), "balance", balance)
		self.balanceTRB.With(prometheus.Labels{"addr": addr.String()}).(prometheus.Gauge).Set(balance)
	}

	for _, addr := range self.addrs {
		balance, err := self.getETHBalance(addr)
		if err != nil {
			level.Error(self.logger).Log("msg", "getting initial ETH balance", "addr", addr.String(), "err", err)
		}
		level.Info(self.logger).Log("msg", "initial ETH balance", "addr", addr.String(), "balance", balance)
		self.balanceETH.With(prometheus.Labels{"addr": addr.String()}).(prometheus.Gauge).Set(balance)
	}

	go self.monitorCost()
	go self.monitorReward()

	<-self.ctx.Done()
	return nil
}

func (self *ProfitTracker) Stop() {
	self.stop()
}

func (self *ProfitTracker) monitorReward() {
	var err error
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	var sub event.Subscription
	events := make(chan *tellor.TellorTransferred)

	logger := log.With(self.logger, "event", "Transfer")

	for {
		sub, err = self.transferSub(events)
		if err != nil {
			level.Error(logger).Log("msg", "initial subscribing to events", "err", err)
			select {
			case <-ticker.C:
				continue
			case <-self.ctx.Done():
				return
			}
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
				sub, err = self.transferSub(events)
				if err != nil {
					level.Error(logger).Log("msg", "re-subscribing to events failed")
					select {
					case <-ticker.C:
						continue
					case <-self.ctx.Done():
						return
					}
				}
				break
			}
			level.Info(logger).Log("msg", "re-subscribed to events")
		case event := <-events:
			logger := log.With(logger, "addr", event.To.String()[:6], "tx", event.Raw.TxHash)

			if event.Raw.Removed {
				val, err := self.cacheTXsProfit.Get(txIDTransfer(event))
				if err != nil {
					level.Error(logger).Log("msg", "getting cache amount for removed event", "err", err)
					continue
				}
				level.Debug(logger).Log("msg", "removed event", "amount", val.(float64))
				self.submitProfit.With(prometheus.Labels{"addr": event.To.String()}).(prometheus.Gauge).Sub(val.(float64))
				continue
			}

			self.setProfitWhenConfirmed(logger, event)
		}
	}
}

func (self *ProfitTracker) monitorCost() {
	var err error
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	logger := log.With(self.logger, "event", "NonceSubmitted")

	var sub event.Subscription
	events := make(chan *tellor.TellorNonceSubmitted)

	for {
		sub, err = self.nonceSubmittedSub(events)
		if err != nil {
			level.Error(logger).Log("msg", "initial subscribing to events failed")
			select {
			case <-ticker.C:
				continue
			case <-self.ctx.Done():
				return
			}
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
					"Transferred subscription error",
					"err", err)
			}

			// Trying to resubscribe until it succeeds.
			for {
				sub, err = self.nonceSubmittedSub(events)
				if err != nil {
					level.Error(logger).Log("msg", "re-subscribing to events failed", "err", err)
					select {
					case <-ticker.C:
						continue
					case <-self.ctx.Done():
						return
					}
				}
				break
			}
			level.Info(logger).Log("msg", "re-subscribed to events")
		case event := <-events:
			logger := log.With(logger, "addr", event.Miner.String()[:6], "tx", event.Raw.TxHash)

			if event.Raw.Removed {
				val, err := self.cacheTXsCost.Get(txIDNonceSubmit(event))
				if err != nil {
					level.Error(logger).Log("msg", "getting cache amount for removed event", "err", err)
					continue
				}
				level.Debug(logger).Log("msg", "removed event", "amount", val.(float64))
				self.submitCost.With(prometheus.Labels{"addr": event.Miner.String()}).(prometheus.Gauge).Sub(val.(float64))
				continue
			}

			self.setCostWhenConfirmed(logger, event)
		}
	}
}

func (self *ProfitTracker) setCostWhenConfirmed(logger log.Logger, event *tellor.TellorNonceSubmitted) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		receipt, err := self.client.TransactionReceipt(self.ctx, event.Raw.TxHash)
		if err != nil {
			level.Error(logger).Log("msg", "receipt retrieval failed", "err", err)
		} else if receipt != nil { // Non need to check TX status as even failed TXs cost ETH and need to be recorded.
			tx, _, err := self.client.TransactionByHash(self.ctx, event.Raw.TxHash)
			if err != nil {
				level.Error(logger).Log("msg", "get transaction by hash", "err", err)
				return
			}
			cost, _ := big.NewFloat(0).Mul(big.NewFloat(float64(tx.GasPrice().Int64())), big.NewFloat(float64(receipt.GasUsed))).Float64()
			cost = cost / 1e18
			level.Debug(logger).Log("msg", "adding cost", "amount", cost)
			self.submitCost.With(prometheus.Labels{"addr": event.Miner.String()}).(prometheus.Gauge).Add(cost)

			if err := self.cacheTXsCost.Set(txIDNonceSubmit(event), cost); err != nil {
				level.Error(logger).Log("msg", "adding amount to the cache", "err", err)
			}

			balance, err := self.getETHBalance(event.Miner)
			if err != nil {
				level.Error(logger).Log("msg", "getting ETH balance", "err", err)
				return
			}
			level.Debug(logger).Log("msg", "new ETH balance", "balance", balance)
			self.balanceETH.With(prometheus.Labels{"addr": event.Miner.String()}).(prometheus.Gauge).Set(balance)
			return
		}

		level.Debug(logger).Log("msg", "transaction not yet mined")

		select {
		case <-self.ctx.Done():
			level.Debug(logger).Log("msg", "transaction confirmation check canceled")
			return
		case <-ticker.C:
		}
	}
}

func (self *ProfitTracker) setProfitWhenConfirmed(logger log.Logger, event *tellor.TellorTransferred) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		receipt, err := self.client.TransactionReceipt(self.ctx, event.Raw.TxHash)
		if err != nil {
			level.Error(logger).Log("msg", "receipt retrieval failed", "err", err)
		} else if receipt != nil {
			if receipt.Status != types.ReceiptStatusSuccessful {
				level.Error(logger).Log("msg", "event status not success so no profit added", "status", receipt.Status)
				return
			}

			trb, _ := big.NewFloat(float64(event.Value.Int64())).Float64()
			trb = trb / 1e18
			level.Debug(logger).Log("msg", "adding profit", "amount", trb)
			self.submitProfit.With(prometheus.Labels{"addr": event.To.String()}).(prometheus.Gauge).Add(trb)

			if err := self.cacheTXsProfit.Set(txIDTransfer(event), trb); err != nil {
				level.Error(logger).Log("msg", "adding amount to the cache", "err", err)
			}

			balance, err := self.getTRBBalance(event.To)
			if err != nil {
				level.Error(logger).Log("msg", "getting TRB balance", "err", err)
				return
			}
			level.Debug(logger).Log("msg", "new TRB balance", "balance", balance)
			self.balanceTRB.With(prometheus.Labels{"addr": event.To.String()}).(prometheus.Gauge).Set(balance)
			return
		}

		level.Debug(logger).Log("msg", "transaction not yet mined")

		select {
		case <-self.ctx.Done():
			level.Debug(logger).Log("msg", "transaction confirmation check canceled")
			return
		case <-ticker.C:
		}
	}
}

func (self *ProfitTracker) nonceSubmittedSub(output chan *tellor.TellorNonceSubmitted) (event.Subscription, error) {
	var tellorFilterer *tellor.TellorFilterer
	tellorFilterer, err := tellor.NewTellorFilterer(self.contractInstance.Address, self.client)
	if err != nil {
		return nil, errors.Wrap(err, "getting instance")
	}
	sub, err := tellorFilterer.WatchNonceSubmitted(&bind.WatchOpts{}, output, self.addrs, nil)
	if err != nil {
		return nil, errors.Wrap(err, "getting channel")
	}
	return sub, nil
}

func (self *ProfitTracker) transferSub(output chan *tellor.TellorTransferred) (event.Subscription, error) {
	var tellorFilterer *tellor.TellorFilterer
	tellorFilterer, err := tellor.NewTellorFilterer(self.contractInstance.Address, self.client)
	if err != nil {
		return nil, errors.Wrap(err, "getting instance")
	}
	sub, err := tellorFilterer.WatchTransferred(
		&bind.WatchOpts{},
		output,
		[]common.Address{
			common.HexToAddress("0x0000000000000000000000000000000000000000"),
		},
		self.addrs,
	)
	if err != nil {
		return nil, errors.Wrap(err, "getting channel")
	}
	return sub, nil
}

func (self *ProfitTracker) GasUsed(slot *big.Int) (*big.Int, error) {
	txID := tellorCommon.PriceTXs + slot.String()
	gas, err := self.proxy.Get(txID)
	if err != nil {
		return nil, errors.New("getting the tx eth cost from the db")
	}
	if gas == nil {
		return nil, ErrNoDataForSlot{slot: slot.String()}
	}
	return big.NewInt(0).SetBytes(gas), nil
}

type ErrNoDataForSlot struct {
	slot string
}

func (e ErrNoDataForSlot) Error() string {
	return "no data for gas used for slot:" + e.slot
}

// SaveGasUsed calculates the price for a given slot.
func (self *ProfitTracker) SaveGasUsed(receipt *types.Receipt, slot *big.Int) {
	gasUsed := big.NewInt(int64(receipt.GasUsed))

	txID := tellorCommon.PriceTXs + slot.String()
	err := self.proxy.Put(txID, gasUsed.Bytes())
	if err != nil {
		level.Error(self.logger).Log("msg", "saving transaction cost", "err", err)
	}
	level.Info(self.logger).Log("msg", "saved transaction gas used", "txHash", receipt.TxHash.String(), "amount", gasUsed.Int64(), "slot", slot.Int64())
}

func (self *ProfitTracker) getTRBBalance(addr common.Address) (float64, error) {
	balance, err := self.contractInstance.BalanceOf(nil, addr)
	if err != nil {
		return 0, errors.Wrap(err, "retrieving trb balance")
	}
	_balanceH, _ := big.NewFloat(1).SetString(balance.String())
	decimals, _ := big.NewFloat(1).SetString("1000000000000000000")
	if decimals != nil {
		_balanceH = _balanceH.Quo(_balanceH, decimals)
	}
	balanceH, _ := _balanceH.Float64()
	return balanceH, nil
}

func (self *ProfitTracker) getETHBalance(addr common.Address) (float64, error) {
	balance, err := self.client.BalanceAt(self.ctx, addr, nil)
	if err != nil {
		return 0, errors.Wrap(err, "retrieving balance")
	}
	_balanceH, _ := big.NewFloat(1).SetString(balance.String())
	decimals, _ := big.NewFloat(1).SetString("1000000000000000000")
	if decimals != nil {
		_balanceH = _balanceH.Quo(_balanceH, decimals)
	}
	balanceH, _ := _balanceH.Float64()
	return balanceH, nil
}

func txIDTransfer(event *tellor.TellorTransferred) string {
	return event.Raw.TxHash.String() + event.To.String()
}

func txIDNonceSubmit(event *tellor.TellorNonceSubmitted) string {
	return event.Raw.TxHash.String() + event.Miner.String()
}
