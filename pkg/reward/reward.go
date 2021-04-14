// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package reward

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/common"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/tracker/index"
	"github.com/tellor-io/telliot/pkg/util"
)

func NewReward(logger log.Logger, contractInstance *contracts.ITellor, proxy db.DataServerProxy) *Reward {
	return &Reward{
		logger:           logger,
		contractInstance: contractInstance,
		proxy:            proxy,
	}
}

type Reward struct {
	logger           log.Logger
	contractInstance *contracts.ITellor
	proxy            db.DataServerProxy
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
	txID := common.PriceTXs + slot.String()
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
func (self *Reward) SaveGasUsed(_gasUsed uint64, slot *big.Int) {
	gasUsed := big.NewInt(int64(_gasUsed))

	txID := common.PriceTXs + slot.String()
	err := self.proxy.Put(txID, gasUsed.Bytes())
	if err != nil {
		level.Error(self.logger).Log("msg", "saving transaction cost", "err", err)
	}
	level.Info(self.logger).Log("msg", "saved transaction gas used", "amount", gasUsed.Int64(), "slot", slot.Int64())
}

func (s *Reward) trbPrice() (*big.Int, error) {
	_trbPrice, err := s.proxy.Get(db.QueriedValuePrefix + strconv.Itoa(index.RequestID_TRB_ETH))
	if err != nil {
		return nil, errors.New("getting the trb price from the db")
	}
	if len(_trbPrice) == 0 {
		return nil, errors.New("the db doesn't have the trb price")
	}
	trbPrice, err := hexutil.DecodeBig(string(_trbPrice))
	if err != nil {
		return nil, errors.New("decoding trb price from the db")
	}
	return trbPrice, nil
}

func (s *Reward) convertTRBtoETH(trbAmount, trbPrice *big.Int) *big.Int {
	wei := big.NewInt(common.WEI)
	precisionUpscale := big.NewInt(0).Div(wei, big.NewInt(index.PSRs[index.RequestID_TRB_ETH].Granularity()))
	trbPrice.Mul(trbPrice, precisionUpscale)

	eth := big.NewInt(0).Mul(trbPrice, trbAmount)
	eth.Div(eth, big.NewInt(1e18))
	return eth
}

func (s *Reward) Slot() (*big.Int, error) {
	slot, err := s.contractInstance.GetUintVar(nil, util.Keccak256([]byte("_SLOT_PROGRESS")))
	if err != nil {
		return nil, errors.Wrap(err, "getting _SLOT_PROGRESS")
	}
	return slot, nil
}

func (s *Reward) GasPrice() (*big.Int, error) {
	_gasPrice, err := s.proxy.Get(db.GasKey)
	if err != nil {
		return nil, errors.Wrap(err, "getting gas price")
	}
	gasPrice, err := hexutil.DecodeBig(string(_gasPrice))
	if err != nil {
		return nil, errors.Wrap(err, "decode gas price")
	}
	return gasPrice, nil
}
