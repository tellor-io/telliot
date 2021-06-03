// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package reward

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log"
	"github.com/tellor-io/telliot/pkg/testutil"
)

// TestProfitCalculation ensures correct calculation and pricision conversion of the reward tracker.
// With
// GAS usage of 1
// GAS price of 1ETH
// TRB price of 1ETH
// the rewards percent should equal the trb reward ammount.
// Example
// half TRB reward or 5e17 should equal in -50% profit.
// 1 TRB reward or 1e18 should equal in 0% profit.
// 2 TRB reward or 2e18 should equal in 100% profit.
func TestProfitCalculation(t *testing.T) {
	logger := log.NewNopLogger()

	gasUsed := float64(1)
	gasCost := float64(params.Ether)
	costTotal := gasUsed * gasCost

	trbPrice := float64(1)

	aggregator := &MockAggr{TRBPrice: trbPrice}
	slotNum := big.NewInt(1)

	for _, rewardAmmount := range []float64{5e17, 1e18, 2e18, 3e18} {
		contractCaller := &MockContractCaler{trbRewardAmount: big.NewInt(int64(rewardAmmount))}
		reward := New(logger, aggregator, contractCaller)
		reward.SaveGasUsed(slotNum, uint64(gasUsed))

		rewardAct, err := reward.Current(slotNum, big.NewInt(int64(gasCost)))
		testutil.Ok(t, err)
		profit := rewardAmmount*trbPrice - costTotal
		profitPercent := (profit / costTotal) * 100
		fmt.Println("profitPercent, float64(rewardAct)", profitPercent, float64(rewardAct))
		testutil.Equals(t, profitPercent, float64(rewardAct))
	}
}

type MockAggr struct {
	TRBPrice float64
}

func (self *MockAggr) TimeWeightedAvg(_ string, _ time.Time, _ time.Duration) (float64, float64, error) {
	return self.TRBPrice, 1, nil
}

type MockContractCaler struct {
	trbRewardAmount *big.Int
}

func (*MockContractCaler) GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error) {
	return nil, nil
}

func (self *MockContractCaler) CurrentReward(opts *bind.CallOpts) (*big.Int, error) {
	return self.trbRewardAmount, nil
}
