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

	ticker *time.Ticker

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

		ticker: time.NewTicker(time.Second),

		submitProfit: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      "submit_profit",
			Help:      "Accumulated TRB amount from rewards for all registered addresses",
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
	defer self.ticker.Stop()

	for _, addr := range self.addrs {
		amountTRB, err := self.getTRBBalance(addr)
		if err != nil {
			level.Error(self.logger).Log("msg", "getting initial TRB balance", "addr", addr.String(), "err", err)
		}
		level.Info(self.logger).Log("msg", "initial TRB balance", "addr", addr.String(), "balance", amountTRB)
		self.balanceTRB.With(prometheus.Labels{"addr": addr.String()}).(prometheus.Gauge).Set(amountTRB)
	}

	// go self.monitorCost()
	go self.monitorReward()

	<-self.ctx.Done()
	return nil
}

func (self *ProfitTracker) Stop() {
	self.stop()
}

// func (self *ProfitTracker) monitorCost() {
// 	var err error
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()

// 	var sub event.Subscription
// 	events := make(chan *tellor.ITellorNonceSubmitted)

// 	for {
// 		sub, err = self.nonceSubmittedSub(events)
// 		if err != nil {
// 			level.Error(self.logger).Log("msg", "initial subscribing to NonceSubmitted events failed")
// 			select {
// 			case <-ticker.C:
// 				continue
// 			case <-self.ctx.Done():
// 				return
// 			}
// 		}
// 		break
// 	}

// 	for {
// 		select {
// 		case <-self.ctx.Done():
// 			return
// 		case err := <-sub.Err():
// 			if err != nil {
// 				level.Error(self.logger).Log(
// 					"msg",
// 					"NonceSubmitted subscription error",
// 					"err", err)
// 			}

// 			for {
// 				sub, err = self.nonceSubmittedSub(events)
// 				if err != nil {
// 					level.Error(self.logger).Log("msg", "re-subscribing to NonceSubmitted events failed")
// 					select {
// 					case <-ticker.C:
// 						continue
// 					case <-self.ctx.Done():
// 						return
// 					}
// 				}
// 				break
// 			}
// 			level.Info(self.logger).Log("msg", "re-subscribed to NonceSubmitted events")
// 		case event := <-events:
// 			logger := log.With(self.logger, "addr", event.Miner.String()[:6], "tx", event.Raw.TxHash)
// 			go self.setCostWhenConfirmed(logger, event)
// 			fmt.Println("vLog cost", event)

// 			amount, err := self.getETHBalance()
// 			if err != nil {
// 				level.Error(self.logger).Log("msg", "getting ETH balance", "err", err)
// 			} else {
// 				level.Info(self.logger).Log("msg", "ETH balance", "amount", amount)

// 			}
// 		}
// 	}
// }

// func (self *ProfitTracker) setCostWhenConfirmed(logger log.Logger, event *tellor.ITellorNonceSubmitted) {
// 	if event.Raw.Removed { // Ignore remove events due to reorg.
// 		level.Debug(logger).Log("msg", "reorg event ignored")
// 		return
// 	}
// 	receipt, err := self.waitMined(logger, event.Raw.TxHash)
// 	if err != nil {
// 		level.Error(logger).Log("msg", "wait confirmation for cost event", "err", err)
// 		return
// 	}
// 	tx, _, err := self.client.TransactionByHash(self.ctx, event.Raw.TxHash)
// 	if err != nil {
// 		level.Error(logger).Log("msg", "get transaction by hash", "err", err)
// 		return
// 	}
// 	cost, _ := big.NewFloat(0).Mul(big.NewFloat(float64(tx.GasPrice().Int64())), big.NewFloat(float64(receipt.GasUsed))).Float64()
// 	level.Debug(logger).Log("msg", "adding cost", "amount", cost/1e18)
// 	self.submitCost.With(prometheus.Labels{"addr": event.Miner.String()[:6]}).(prometheus.Gauge).Add(cost / 1e18)
// }

func (self *ProfitTracker) setProfitWhenConfirmed(ctx context.Context, logger log.Logger, event *tellor.ITellorTransferred) {
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

			if err := self.cacheTXsProfit.Set(txID(event), trb); err != nil {
				level.Error(self.logger).Log("msg", "adding profit amount to the cache", "err", err)
			}

			amountTRB, err := self.getTRBBalance(event.To)
			if err != nil {
				level.Error(self.logger).Log("msg", "getting TRB balance", "err", err)
				return
			}
			level.Debug(logger).Log("msg", "new TRB balance", "balance", amountTRB)
			self.balanceTRB.With(prometheus.Labels{"addr": event.To.String()}).(prometheus.Gauge).Set(amountTRB)
			return
		}

		level.Debug(logger).Log("msg", "transaction not yet mined")

		select {
		case <-ctx.Done():
			level.Debug(logger).Log("msg", "transaction confirmation check canceled")
			return
		case <-self.ticker.C:
		}
	}
}

func (self *ProfitTracker) nonceSubmittedSub(output chan *tellor.ITellorNonceSubmitted) (event.Subscription, error) {
	var tellorFilterer *tellor.ITellorFilterer
	tellorFilterer, err := tellor.NewITellorFilterer(self.contractInstance.Address, self.client)
	if err != nil {
		return nil, errors.Wrap(err, "getting ITellorFilterer instance")
	}
	sub, err := tellorFilterer.WatchNonceSubmitted(&bind.WatchOpts{}, output, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "getting NonceSubmitted channel")
	}
	return sub, nil
}

func (self *ProfitTracker) monitorReward() {
	var err error
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	var sub event.Subscription
	events := make(chan *tellor.ITellorTransferred)

	for {
		sub, err = self.transferSub(events)
		if err != nil {
			level.Error(self.logger).Log("msg", "initial subscribing to Transferred events failed")
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
				level.Error(self.logger).Log(
					"msg",
					"Transferred subscription error",
					"err", err)
			}

			// Trying to resubscribe until it succeeds.
			for {
				sub, err = self.transferSub(events)
				if err != nil {
					level.Error(self.logger).Log("msg", "re-subscribing to Transferred events failed")
					select {
					case <-ticker.C:
						continue
					case <-self.ctx.Done():
						return
					}
				}
				break
			}
			level.Info(self.logger).Log("msg", "re-subscribed to Transferred events")
		case event := <-events:
			logger := log.With(self.logger, "event", "transfer", "addr", event.To.String()[:6], "tx", event.Raw.TxHash)

			if event.Raw.Removed {
				trb, err := self.cacheTXsProfit.Get(txID(event))
				if err != nil {
					level.Error(logger).Log("msg", "getting cache amount for removed event", "err", err)
					continue
				}
				level.Debug(logger).Log("msg", "removed event", "amount", trb.(float64))
				self.submitProfit.With(prometheus.Labels{"addr": event.To.String()}).(prometheus.Gauge).Sub(trb.(float64))
				continue
			}

			// ctx, cncl := context.WithCancel(self.ctx)
			// self.cacheTXsProfit.Set(contextID(event), cncl)
			self.setProfitWhenConfirmed(self.ctx, logger, event)
		}
	}
}

func (self *ProfitTracker) transferSub(output chan *tellor.ITellorTransferred) (event.Subscription, error) {
	var tellorFilterer *tellor.ITellorFilterer
	tellorFilterer, err := tellor.NewITellorFilterer(self.contractInstance.Address, self.client)
	if err != nil {
		return nil, errors.Wrap(err, "getting ITellorFilterer instance")
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
		return nil, errors.Wrap(err, "getting Transferred channel")
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

func txID(event *tellor.ITellorTransferred) string {
	return event.Raw.TxHash.String() + event.To.String()
}
