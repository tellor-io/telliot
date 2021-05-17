// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package reward

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/aggregator"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/ethereum"
)

func NewReward(logger log.Logger, aggr *aggregator.Aggregator, contractInstance *contracts.ITellor) *Reward {
	return &Reward{
		aggr:             aggr,
		logger:           logger,
		contractInstance: contractInstance,
		gasUsed:          make(map[int64]*big.Int),
	}
}

type Reward struct {
	logger           log.Logger
	aggr             *aggregator.Aggregator
	contractInstance *contracts.ITellor
	gasUsed          map[int64]*big.Int
}

// Current returns the profit in percents based on the current TRB price.
func (self *Reward) Current(slot *big.Int, gasPrice *big.Int) (int64, error) {
	gasUsed, err := self.GasUsed(slot)
	if err != nil {
		return 0, err
	}
	trbAmount, err := self.contractInstance.CurrentReward(nil)
	if err != nil {
		return 0, errors.New("getting currentReward from the chain")
	}

	trbPrice, err := self.trbPrice()
	if err != nil {
		return 0, errors.New("getting trb current TRB price")
	}

	rewardEth := self.convertTRBtoETH(trbAmount, trbPrice)

	txCost := big.NewInt(0).Mul(gasPrice, gasUsed)
	profit := big.NewInt(0).Sub(rewardEth, txCost)
	profitPercentFloat := float64(profit.Int64()) / float64(txCost.Int64()) * 100
	profitPercent := int64(profitPercentFloat)

	level.Debug(self.logger).Log(
		"msg", "profit checking",
		"reward", fmt.Sprintf("%.2e", float64(rewardEth.Int64())),
		"txCost", fmt.Sprintf("%.2e", float64(txCost.Int64())),
		"slot", slot,
		"gasUsed", gasUsed,
		"gasPrice", gasPrice,
		"profit", fmt.Sprintf("%.2e", float64(profit.Int64())),
		"profitMargin", profitPercent,
	)

	return profitPercent, nil
}

func (self *Reward) GasUsed(slot *big.Int) (*big.Int, error) {
	if gas, ok := self.gasUsed[slot.Int64()]; ok {
		return gas, nil
	}

	return nil, ErrNoDataForSlot{slot: slot.String()}

}

type ErrNoDataForSlot struct {
	slot string
}

func (e ErrNoDataForSlot) Error() string {
	return "no data for gas used for slot:" + e.slot
}

// SaveGasUsed calculates the price for a given slot.
func (self *Reward) SaveGasUsed(_gasUsed uint64, slot *big.Int) {
	gasUsed := big.NewInt(int64(_gasUsed))

	self.gasUsed[slot.Int64()] = gasUsed
	level.Info(self.logger).Log("msg", "saved transaction gas used", "amount", gasUsed.Int64(), "slot", slot.Int64())
}

func (s *Reward) trbPrice() (*big.Int, error) {
	trbPrice, confidence, err := s.aggr.TimeWeightedAvg("TRB/ETH", time.Now(), time.Hour)
	if err != nil {
		return nil, errors.New("getting the trb price from the aggregator")
	}

	if confidence < 0.5 {
		return nil, errors.New("trb price confidence too low")

	}

	return big.NewInt(int64(trbPrice)), nil
}

func (s *Reward) convertTRBtoETH(trbAmount, trbPrice *big.Int) *big.Int {
	ether := big.NewInt(params.Ether)
	precisionUpscale := big.NewInt(0).Div(ether, big.NewInt(int64(aggregator.DefaultGranularity)))
	trbPrice.Mul(trbPrice, precisionUpscale)

	eth := big.NewInt(0).Mul(trbPrice, trbAmount)
	eth.Div(eth, big.NewInt(1e18))
	return eth
}

func (s *Reward) Slot() (*big.Int, error) {
	slot, err := s.contractInstance.GetUintVar(nil, ethereum.Keccak256([]byte("_SLOT_PROGRESS")))
	if err != nil {
		return nil, errors.Wrap(err, "getting _SLOT_PROGRESS")
	}
	return slot, nil
}
