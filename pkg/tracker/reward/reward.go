// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package reward

import (
	"context"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/tsdb"

	"github.com/tellor-io/telliot/pkg/aggregator"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/contracts/tellor"
	"github.com/tellor-io/telliot/pkg/db"
	eth "github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/logging"
)

const ComponentName = "rewardTracker"
const DefaultRetry = 30 * time.Second

type Config struct {
	LogLevel string
}

type RewardTracker struct {
	client           *ethclient.Client
	logger           log.Logger
	contractInstance *contracts.ITellor
	ctx              context.Context
	stop             context.CancelFunc
	addr             common.Address

	tsDB   *tsdb.DB
	aggr   aggregator.IAggregator
	engine *promql.Engine
}

func NewRewardTracker(
	logger log.Logger,
	ctx context.Context,
	cfg Config,
	tsDB *tsdb.DB,
	client *ethclient.Client,
	contractInstance *contracts.ITellor,
	addr common.Address,
	aggr aggregator.IAggregator,
) (*RewardTracker, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(logger, "component", ComponentName)

	opts := promql.EngineOpts{
		Logger:               logger,
		Reg:                  nil,
		MaxSamples:           30000,
		Timeout:              10 * time.Second,
		LookbackDelta:        5 * time.Minute,
		EnableAtModifier:     true,
		EnableNegativeOffset: true,
	}
	engine := promql.NewEngine(opts)

	ctx, cncl := context.WithCancel(ctx)
	return &RewardTracker{
		client:           client,
		logger:           logger,
		contractInstance: contractInstance,
		addr:             addr,
		ctx:              ctx,
		stop:             cncl,
		tsDB:             tsDB,
		engine:           engine,
		aggr:             aggr,
	}, nil
}

func (self *RewardTracker) Start() error {
	level.Info(self.logger).Log("msg", "starting")

	var err error
	ticker := time.NewTicker(DefaultRetry)
	defer ticker.Stop()

	logger := log.With(self.logger, "event", "NonceSubmitted")

	var sub event.Subscription
	events := make(chan *tellor.TellorNonceSubmitted)
	for {
		select {
		case <-self.ctx.Done():
			return errors.New("context canceled")
		default:
		}
		sub, err = self.nonceSubmittedSub(events)
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
			return errors.New("context canceled")
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
					return errors.New("context canceled")
				default:
				}
				sub, err = self.nonceSubmittedSub(events)
				if err != nil {
					level.Error(logger).Log("msg", "re-subscribing to events failed", "err", err)
					<-ticker.C
					continue
				}
				break
			}
			level.Info(logger).Log("msg", "re-subscribed to events")
		case event := <-events:
			err := self.recordGasUsage(event)
			if err != nil {
				level.Error(self.logger).Log("msg", "record gas usage", "err", err)
			}
			err = self.recordGasUsageEstimated()
			if err != nil {
				level.Error(self.logger).Log("msg", "record gas usage estimation", "err", err)
			}
		}
	}
}

func (self *RewardTracker) Stop() {
	self.stop()
}

func (self *RewardTracker) recordGasUsage(event *tellor.TellorNonceSubmitted) error {
	receipt, err := self.client.TransactionReceipt(self.ctx, event.Raw.TxHash)
	if err != nil {
		return errors.Wrap(err, "receipt retrieval")
	} else if receipt != nil && receipt.Status == types.ReceiptStatusSuccessful {
		tx, _, err := self.client.TransactionByHash(self.ctx, event.Raw.TxHash)
		if err != nil {
			return errors.Wrap(err, "tx retrieval")
		}
		level.Debug(self.logger).Log("msg", "adding gas used", "gasUsed", receipt.GasUsed)
		if err := self.addRefundValue(event.Slot.String(), tx.Gas()-receipt.GasUsed); err != nil {
			return errors.Wrap(err, "adding gas used")
		}
		return nil
	}
	return errors.New("transaction not yet mined")
}

func (self *RewardTracker) recordGasUsageEstimated() error {
	ctx, cncl := context.WithTimeout(self.ctx, 2*time.Second)
	defer cncl()
	abi, err := abi.JSON(strings.NewReader(tellor.TellorABI))
	if err != nil {
		return errors.Wrap(err, "getting abi")
	}

	// Getting current slot.
	slot, err := self.Slot()
	if err != nil {
		return errors.Wrap(err, "getting current slot")
	}

	// Getting current challenge.
	vars, err := self.contractInstance.GetNewCurrentVariables(&bind.CallOpts{Context: ctx})
	if err != nil {
		return errors.Wrap(err, "call GetNewCurrentVariables")
	}
	// Use arbitrary values as when tested didn't make a noticeable difference
	// between using real vs arbitrary values.
	reqVals := [5]*big.Int{
		big.NewInt(math.MaxInt64),
		big.NewInt(math.MaxInt64),
		big.NewInt(math.MaxInt64),
		big.NewInt(math.MaxInt64),
		big.NewInt(math.MaxInt64),
	}
	packed, err := abi.Pack("submitMiningSolution", "", vars.RequestIds, reqVals)
	if err != nil {
		return errors.Wrap(err, "packing submitMiningSolution args")
	}
	data := ethereum.CallMsg{
		From: self.addr,
		To:   &self.contractInstance.Address,
		// Hardcoded gas price. Looks like it doesn't matter in the calculation.
		GasPrice: big.NewInt(0).Mul(big.NewInt(30), big.NewInt(params.GWei)),
		Data:     packed,
	}
	gasEstimation, err := self.client.EstimateGas(ctx, data)
	if err != nil {
		return errors.Wrap(err, "call client.EstimateGas")
	}
	err = self.addGasEstimation(slot, gasEstimation)
	if err != nil {
		return errors.Wrap(err, "add gas estimation to the db")
	}
	return nil
}

func (self *RewardTracker) addGasEstimation(slot *big.Int, gasEstimation uint64) error {
	lbls := labels.Labels{
		labels.Label{Name: "__name__", Value: "gas_usage_estimation"},
		labels.Label{Name: "slot", Value: slot.String()},
	}
	if err := db.Add(self.ctx, self.tsDB, lbls, float64(gasEstimation)); err != nil {
		return errors.Wrap(err, "adding gasEstimation value to the db")
	}
	return nil
}

func (self *RewardTracker) addRefundValue(slot string, gasUsed uint64) error {
	lbls := labels.Labels{
		labels.Label{Name: "__name__", Value: "refund"},
		labels.Label{Name: "slot", Value: slot},
	}
	if err := db.Add(self.ctx, self.tsDB, lbls, float64(gasUsed)); err != nil {
		return errors.Wrap(err, "adding gasUsed value to the db")
	}
	return nil
}

func (self *RewardTracker) nonceSubmittedSub(output chan *tellor.TellorNonceSubmitted) (event.Subscription, error) {
	tellorFilterer, err := tellor.NewTellorFilterer(self.contractInstance.Address, self.client)
	if err != nil {
		return nil, errors.Wrap(err, "getting instance")
	}
	sub, err := tellorFilterer.WatchNonceSubmitted(&bind.WatchOpts{Context: self.ctx}, output, []common.Address{}, nil)
	if err != nil {
		return nil, errors.Wrap(err, "getting channel")
	}
	return sub, nil
}

func (self *RewardTracker) Slot() (*big.Int, error) {
	slot, err := self.contractInstance.GetUintVar(nil, eth.Keccak256([]byte("_SLOT_PROGRESS")))
	if err != nil {
		return nil, errors.Wrap(err, "getting _SLOT_PROGRESS")
	}
	return slot, nil
}
