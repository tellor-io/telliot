// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package profitChecker

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	tellorCommon "github.com/tellor-io/telliot/pkg/common"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/tracker"
)

type ProfitChecker struct {
	client           contracts.ETHClient
	logger           log.Logger
	contractInstance *contracts.ITellor
	proxy            db.DataServerProxy

	submitProfit *prometheus.GaugeVec
	submitCost   *prometheus.GaugeVec
	submitReward *prometheus.GaugeVec
}

func NewProfitChecker(
	logger log.Logger,
	client contracts.ETHClient,
	contractInstance *contracts.ITellor,
	proxy db.DataServerProxy,
) *ProfitChecker {

	return &ProfitChecker{
		client:           client,
		logger:           logger,
		contractInstance: contractInstance,
		proxy:            proxy,

		submitProfit: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: "mining",
			Name:      "submit_profit",
			Help:      "The current submit profit in percents",
		},
			[]string{"slot"},
		),
		submitCost: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: "mining",
			Name:      "submit_cost",
			Help:      "The current submit cost in 1e18 eth",
		},
			[]string{"slot"},
		),
		submitReward: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: "mining",
			Name:      "submit_reward",
			Help:      "The current reward in 1e18 eth",
		},
			[]string{"slot"},
		),
	}
}

// Current returns the profit in percents.
func (self *ProfitChecker) Current(slot *big.Int, gasPrice *big.Int) (int64, error) {
	gasUsed, err := self.gasUsed(slot)
	if err != nil {
		return 0, err
	}
	reward, err := self.currentReward()
	if err != nil {
		return 0, errors.Wrap(err, "getting current rewards")
	}

	txCost := big.NewInt(0).Mul(gasPrice, gasUsed)
	profit := big.NewInt(0).Sub(reward, txCost)
	profitPercentFloat := float64(profit.Int64()) / float64(txCost.Int64()) * 100
	profitPercent := int64(profitPercentFloat)

	level.Debug(self.logger).Log(
		"msg", "profit checking",
		"reward", fmt.Sprintf("%.2e", float64(reward.Int64())),
		"txCost", fmt.Sprintf("%.2e", float64(txCost.Int64())),
		"slot", slot,
		"gasUsed", gasUsed,
		"gasPrice", gasPrice,
		"profit", fmt.Sprintf("%.2e", float64(profit.Int64())),
		"profitMargin", profitPercent,
	)

	self.submitProfit.With(prometheus.Labels{"slot": strconv.Itoa(int(slot.Int64()))}).(prometheus.Gauge).Set(float64(profitPercent))
	self.submitCost.With(prometheus.Labels{"slot": strconv.Itoa(int(slot.Int64()))}).(prometheus.Gauge).Set(float64(txCost.Int64()))
	self.submitReward.With(prometheus.Labels{"slot": strconv.Itoa(int(slot.Int64()))}).(prometheus.Gauge).Set(float64(reward.Int64()))

	return profitPercent, nil
}

func (self *ProfitChecker) gasUsed(slot *big.Int) (*big.Int, error) {
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

// currentReward returns the current TRB rewards converted to ETH.
func (self *ProfitChecker) currentReward() (*big.Int, error) {
	reward, err := self.contractInstance.CurrentReward(nil)
	if err != nil {
		return nil, errors.New("getting currentReward from the chain")
	}
	return self.convertTRBtoETH(reward)
}

func (self *ProfitChecker) convertTRBtoETH(trb *big.Int) (*big.Int, error) {
	val, err := self.proxy.Get(db.QueriedValuePrefix + strconv.Itoa(tracker.RequestID_TRB_ETH))
	if err != nil {
		return nil, errors.New("getting the trb price from the db")
	}
	if len(val) == 0 {
		return nil, errors.New("the db doesn't have the trb price")
	}
	priceTRB, err := hexutil.DecodeBig(string(val))
	if err != nil {
		return nil, errors.New("decoding trb price from the db")
	}
	wei := big.NewInt(tellorCommon.WEI)
	precisionUpscale := big.NewInt(0).Div(wei, big.NewInt(tracker.PSRs[tracker.RequestID_TRB_ETH].Granularity()))
	priceTRB.Mul(priceTRB, precisionUpscale)

	eth := big.NewInt(0).Mul(priceTRB, trb)
	eth.Div(eth, big.NewInt(1e18))
	return eth, nil
}

type ErrNoDataForSlot struct {
	slot string
}

func (e ErrNoDataForSlot) Error() string {
	return "no data for gas used for slot:" + e.slot
}

// SaveGasUsed calculates the price for a given slot.
func (self *ProfitChecker) SaveGasUsed(receipt *types.Receipt, slot *big.Int) {
	gasUsed := big.NewInt(int64(receipt.GasUsed))

	txID := tellorCommon.PriceTXs + slot.String()
	err := self.proxy.Put(txID, gasUsed.Bytes())
	if err != nil {
		level.Error(self.logger).Log("msg", "saving transaction cost", "err", err)
	}
	level.Info(self.logger).Log("msg", "saved transaction gas used", "txHash", receipt.TxHash.String(), "amount", gasUsed.Int64(), "slot", slot.Int64())
}
