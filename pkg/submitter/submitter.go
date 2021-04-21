// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package submitter

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/mining"
	"github.com/tellor-io/telliot/pkg/reward"
	"github.com/tellor-io/telliot/pkg/tracker/gasPrice"
	"github.com/tellor-io/telliot/pkg/transactor"
	"github.com/tellor-io/telliot/pkg/util"
)

const ComponentName = "submitter"

type Config struct {
	LogLevel string
	// Minimum percent of profit when submitting a solution.
	// For example if the tx cost is 0.01 ETH and current reward is 0.02 ETH
	// a ProfitThreshold of 200% or more will wait until the reward is increased or
	// the gas cost is lowered.
	// a ProfitThreshold of 199% or less will submit
	ProfitThreshold uint64
	MinSubmitPeriod util.Duration
}

/**
* The submitter has one purpose: to either submit the solution on-chain
* or to reject it if the miner has already submitted a solution for the challenge
* or the the solution'self challenge does not match current challenge
 */

type Submitter struct {
	ctx              context.Context
	close            context.CancelFunc
	logger           log.Logger
	cfg              Config
	account          *ethereum.Account
	client           contracts.ETHClient
	contractInstance *contracts.ITellor
	resultCh         chan *mining.Result
	submitCount      prometheus.Counter
	submitFailCount  prometheus.Counter
	lastSubmitCncl   context.CancelFunc
	transactor       transactor.Transactor
	reward           *reward.Reward
	gasPriceTracker  *gasPrice.GasTracker
}

func NewSubmitter(
	ctx context.Context,
	cfg Config,
	logger log.Logger,
	client contracts.ETHClient,
	contractInstance *contracts.ITellor,
	account *ethereum.Account,
	reward *reward.Reward,
	transactor transactor.Transactor,
	gasPriceTracker *gasPrice.GasTracker,
) (*Submitter, chan *mining.Result, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(logger, "component", ComponentName, "addr", account.Address.String()[:6])
	ctx, close := context.WithCancel(ctx)
	submitter := &Submitter{
		ctx:              ctx,
		close:            close,
		client:           client,
		cfg:              cfg,
		resultCh:         make(chan *mining.Result),
		account:          account,
		reward:           reward,
		logger:           logger,
		contractInstance: contractInstance,
		transactor:       transactor,
		gasPriceTracker:  gasPriceTracker,
		submitCount: promauto.NewCounter(prometheus.CounterOpts{
			Namespace:   "telliot",
			Subsystem:   ComponentName,
			Name:        "submit_total",
			Help:        "The total number of submitted solutions",
			ConstLabels: prometheus.Labels{"account": account.Address.String()},
		}),
		submitFailCount: promauto.NewCounter(prometheus.CounterOpts{
			Namespace:   "telliot",
			Subsystem:   ComponentName,
			Name:        "submit_fails_total",
			Help:        "The total number of failed submission",
			ConstLabels: prometheus.Labels{"account": account.Address.String()},
		}),
	}

	return submitter, submitter.resultCh, nil
}

func (self *Submitter) Start() error {
	for {
		select {
		case <-self.ctx.Done():
			if self.lastSubmitCncl != nil {
				self.lastSubmitCncl()
			}
			level.Info(self.logger).Log("msg", "submitter shutdown complete")
			return self.ctx.Err()
		case result := <-self.resultCh:
			if self.lastSubmitCncl != nil {
				self.lastSubmitCncl()
			}
			var ctx context.Context
			ctx, self.lastSubmitCncl = context.WithCancel(self.ctx)

			level.Info(self.logger).Log("msg", "received a solution",
				"challenge", fmt.Sprintf("%x", result.Work.Challenge),
				"solution", result.Nonce,
				"difficulty", result.Work.Challenge.Difficulty,
				"requestIDs", fmt.Sprintf("%+v", result.Work.Challenge.RequestIDs),
			)
			self.Submit(ctx, result)
		}
	}
}

func (self *Submitter) CancelPendingSubmit() {
	if self.lastSubmitCncl != nil {
		self.lastSubmitCncl()
	}
}

func (self *Submitter) Stop() {
	self.close()
}

func (self *Submitter) blockUntilTimeToSubmit(newChallengeReplace context.Context) {
	var (
		lastSubmit time.Duration
		timestamp  *time.Time
		err        error
	)
	for {
		select {
		case <-newChallengeReplace.Done():
			level.Info(self.logger).Log("msg", "canceled pending submit while gettting last submit time")
		default:
		}
		lastSubmit, timestamp, err = self.lastSubmit()
		if err != nil {
			level.Debug(self.logger).Log("msg", "checking last submit time", "err", err)
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
	if lastSubmit < self.cfg.MinSubmitPeriod.Duration {
		level.Info(self.logger).Log("msg", "min transaction submit threshold hasn't passed",
			"nextSubmit", time.Duration(self.cfg.MinSubmitPeriod.Nanoseconds())-lastSubmit,
			"lastSubmit", lastSubmit,
			"lastSubmitTimestamp", timestamp.Format("2006-01-02 15:04:05.000000"),
			"minSubmitPeriod", self.cfg.MinSubmitPeriod,
		)
		timeToSubmit, cncl := context.WithDeadline(newChallengeReplace, timestamp.Add(self.cfg.MinSubmitPeriod.Duration))
		defer cncl()
		select {
		case <-newChallengeReplace.Done():
			level.Info(self.logger).Log("msg", "canceled pending submit while waiting for the time to submit")
		case <-timeToSubmit.Done(): // 15min since last submit has passed so can unblock.
		}
	}
}

func (self *Submitter) canSubmit() error {
	if self.cfg.ProfitThreshold > 0 { // Profit check is enabled.
		profitPercent, err := self.profitPercent()
		if _, ok := errors.Cause(err).(reward.ErrNoDataForSlot); ok {
			level.Warn(self.logger).Log("msg", "skipping profit check when the slot has no record for how much gas it uses", "err", err)
		} else if err != nil {
			return errors.Wrapf(err, "submit solution profit check")
		} else if profitPercent < int64(self.cfg.ProfitThreshold) {
			return errors.Errorf("profit:%v lower then the profit threshold:%v", profitPercent, self.cfg.ProfitThreshold)
		}
	}

	statusID, err := self.minerStatus()
	if err != nil {
		return errors.Wrap(err, "getting miner status")
	}
	if statusID != 1 {
		return errors.Errorf("miner is not in a status that can submit:%v", minerStatusName(statusID))
	}

	return nil
}

func (self *Submitter) profitPercent() (int64, error) {
	slot, err := self.reward.Slot()
	if err != nil {
		return 0, errors.Wrapf(err, "getting current slot")
	}
	gasPrice, err := self.gasPriceTracker.Query(self.ctx)
	if err != nil {
		return 0, errors.Wrapf(err, "getting current Gas price")
	}

	// Need the price for next slot transaction so increment by one.
	slot.Add(slot, big.NewInt(1))

	// Slots numbers are from 0 to 4 so
	// when next slot is 4+1=5 get the price for slot 0.
	if slot.Int64() == 5 {
		slot.SetInt64(0)
	}

	return self.reward.Current(slot, big.NewInt(int64(gasPrice)))
}

func (self *Submitter) Submit(newChallengeReplace context.Context, result *mining.Result) {
	go func(newChallengeReplace context.Context, result *mining.Result) {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-newChallengeReplace.Done():
				level.Info(self.logger).Log("msg", "pending submit canceled")
				return
			default:
			}

			self.blockUntilTimeToSubmit(newChallengeReplace)
			if err := self.canSubmit(); err != nil {
				level.Info(self.logger).Log("msg", "can't submit and will retry later", "reason", err)
				<-ticker.C
				continue
			}
			for {
				select {
				case <-newChallengeReplace.Done():
					level.Info(self.logger).Log("msg", "pending submit canceled")
					return
				default:
				}

				reqVals, err := self.requestVals(result.Work.Challenge.RequestIDs)
				if err != nil {
					level.Error(self.logger).Log("msg", "adding the request ids, retrying", "err", err)
					<-ticker.C
					continue
				}
				level.Info(self.logger).Log(
					"msg", "sending solution to the chain",
					"solutionNonce", result.Nonce,
					"IDs", fmt.Sprintf("%+v", result.Work.Challenge.RequestIDs),
					"vals", fmt.Sprintf("%+v", reqVals),
				)
				tx, recieipt, err := self.transactor.Transact(newChallengeReplace, result.Nonce, result.Work.Challenge.RequestIDs, reqVals)
				if err != nil {
					self.submitFailCount.Inc()
					level.Error(self.logger).Log("msg", "submiting a solution", "err", err)
					return
				}

				if recieipt.Status != types.ReceiptStatusSuccessful {
					self.submitFailCount.Inc()
					level.Error(self.logger).Log("msg", "submiting solution status not success", "status", recieipt.Status, "hash", tx.Hash())
					return
				}
				level.Info(self.logger).Log("msg", "successfully submited solution",
					"txHash", tx.Hash().String(),
					"nonce", tx.Nonce(),
					"gasPrice", tx.GasPrice(),
					"data", fmt.Sprintf("%x", tx.Data()),
					"value", tx.Value(),
				)
				self.submitCount.Inc()

				slot, err := self.reward.Slot()
				if err != nil {
					level.Error(self.logger).Log("msg", "getting _SLOT_PROGRESS for saving gas used", "err", err)
				} else {
					self.reward.SaveGasUsed(recieipt.GasUsed, slot.Sub(slot, big.NewInt(1)))
				}

				return
			}
		}
	}(newChallengeReplace, result)
}

func (self *Submitter) requestVals(requestIDs [5]*big.Int) ([5]*big.Int, error) {
	var currentValues [5]*big.Int

	// The submit contains values for 5 data IDs so add them here.
	for i := 0; i < 5; i++ {
		// Look back only 2 times the API tracker cycle to use only fresh values.
		// q, err := self.tsDB.Querier(self.ctx, timestamp.FromTime(time.Now().Add(-(2*self.cfg.Trackers.SleepCycle.Duration))), timestamp.FromTime(time.Now().Round(0)))
		// if err != nil {
		// 	return currentValues, err
		// }
		// defer q.Close()
		// s := q.Select(false, nil, labels.MustNewMatcher(labels.MatchEqual, "__name__", requestIDs[i].String()))

		// valKey := fmt.Sprintf("%s%d", db.QueriedValuePrefix, requestIDs[i].Uint64())
		// m, err := self.proxy.BatchGet([]string{valKey})
		// if err != nil {
		// 	return currentValues, errors.Wrapf(err, "retrieve pricing from db for data id:%v", requestIDs[i].Uint64())
		// }
		// val := m[valKey]
		// var value *big.Int
		// if len(val) == 0 {
		// 	jsonFile, err := os.Open(self.cfg.ManualDataFile)
		// 	if err != nil {
		// 		return currentValues, errors.Wrapf(err, "manualData read Error")
		// 	}
		// 	defer jsonFile.Close()
		// 	byteValue, _ := ioutil.ReadAll(jsonFile)
		// 	var result map[string]map[string]uint
		// 	_ = json.Unmarshal([]byte(byteValue), &result)
		// 	_id := strconv.FormatUint(requestIDs[i].Uint64(), 10)
		// 	val := result[_id]["VALUE"]
		// 	if val == 0 {
		// 		return currentValues, errors.Errorf("pricing data not available from db or the manual file for request id:%v", requestIDs[i].Uint64())
		// 	}
		// 	value = big.NewInt(int64(val))
		// } else {
		// 	value, err = hexutil.DecodeBig(string(val))
		// 	if err != nil {
		// 		if requestIDs[i].Uint64() > index.MaxPSRID() {
		// 			level.Error(self.logger).Log(
		// 				"msg", "decoding price value prior to submiting solution",
		// 				"err", err,
		// 			)
		// 			if len(val) == 0 {
		// 				level.Error(self.logger).Log("msg", "0 value being submitted")
		// 				currentValues[i] = big.NewInt(0)
		// 			}
		// 			continue
		// 		}
		// 		return currentValues, errors.Errorf("no value in database,  reg id:%v", requestIDs[i].Uint64())
		// 	}
		// }
		// currentValues[i] = value
	}
	return currentValues, nil
}

func (self *Submitter) minerStatus() (int64, error) {
	// Check if the staked account is in dispute before sending a transaction.
	statusID, _, err := self.contractInstance.GetStakerInfo(&bind.CallOpts{}, self.account.Address)
	if err != nil {
		return 0, errors.Wrapf(err, "getting staker info from contract addr:%v", self.account.Address)
	}
	return statusID.Int64(), nil
}

func (self *Submitter) lastSubmit() (time.Duration, *time.Time, error) {
	address := "000000000000000000000000" + self.account.Address.Hex()[2:]
	decoded, err := hex.DecodeString(address)
	if err != nil {
		return 0, nil, errors.Wrapf(err, "decoding address")
	}
	last, err := self.contractInstance.GetUintVar(nil, ethereum.Keccak256(decoded))

	if err != nil {
		return 0, nil, errors.Wrapf(err, "getting last submit time for:%v", self.account.Address.String())
	}
	// The Miner has never submitted so put a timestamp at the beginning of unix time.
	if last.Int64() == 0 {
		last.Set(big.NewInt(1))
	}

	lastInt := last.Int64()
	now := time.Now()
	var lastSubmit time.Duration
	var tm time.Time
	if lastInt > 0 {
		tm = time.Unix(lastInt, 0)
		lastSubmit = now.Sub(tm)
	}

	return lastSubmit, &tm, nil
}

func minerStatusName(statusID int64) string {
	// From https://github.com/tellor-io/tellor3/blob/7c2f38a0e3f96631fb0f96e0d0a9f73e7b355766/contracts/TellorStorage.sol#L41
	switch statusID {
	case 0:
		return "Not staked"
	case 1:
		return "Staked"
	case 2:
		return "LockedForWithdraw"
	case 3:
		return "OnDispute"
	case 4:
		return "ReadyForUnlocking"
	case 5:
		return "Unlocked"
	default:
		return "Unknown"
	}
}
