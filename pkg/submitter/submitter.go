// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package submitter

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/mining"
	"github.com/tellor-io/telliot/pkg/reward"
	"github.com/tellor-io/telliot/pkg/tracker/index"
	"github.com/tellor-io/telliot/pkg/transactor"
	"github.com/tellor-io/telliot/pkg/util"
)

const ComponentName = "submitter"

/**
* The submitter has one purpose: to either submit the solution on-chain
* or to reject it if the miner has already submitted a solution for the challenge
* or the the solution's challenge does not match current challenge
 */

type Submitter struct {
	ctx              context.Context
	close            context.CancelFunc
	logger           log.Logger
	cfg              *config.Config
	proxy            db.DataServerProxy
	account          *config.Account
	client           contracts.ETHClient
	contractInstance *contracts.ITellor
	resultCh         chan *mining.Result
	submitCount      prometheus.Counter
	submitFailCount  prometheus.Counter
	lastSubmitCncl   context.CancelFunc
	transactor       transactor.Transactor
	reward           *reward.Reward
}

func NewSubmitter(
	ctx context.Context,
	cfg *config.Config,
	logger log.Logger,
	client contracts.ETHClient,
	contractInstance *contracts.ITellor,
	account *config.Account,
	proxy db.DataServerProxy,
	transactor transactor.Transactor,
) (*Submitter, chan *mining.Result, error) {
	logger, err := logging.ApplyFilter(*cfg, ComponentName, logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(logger, "component", ComponentName, "addr", account.Address.String()[:6])
	ctx, close := context.WithCancel(ctx)
	submitter := &Submitter{
		ctx:              ctx,
		close:            close,
		client:           client,
		proxy:            proxy,
		cfg:              cfg,
		resultCh:         make(chan *mining.Result),
		account:          account,
		logger:           logger,
		contractInstance: contractInstance,
		transactor:       transactor,
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

	if cfg.Mine.ProfitThreshold > 0 { // Profit check is enabled.
		submitter.reward = reward.NewReward(logger, contractInstance, proxy)
	}

	return submitter, submitter.resultCh, nil
}

func (s *Submitter) Start() error {
	for {
		select {
		case <-s.ctx.Done():
			if s.lastSubmitCncl != nil {
				s.lastSubmitCncl()
			}
			level.Info(s.logger).Log("msg", "submitter shutdown complete")
			return s.ctx.Err()
		case result := <-s.resultCh:
			if s.lastSubmitCncl != nil {
				s.lastSubmitCncl()
			}
			var ctx context.Context
			ctx, s.lastSubmitCncl = context.WithCancel(s.ctx)

			level.Info(s.logger).Log("msg", "received a solution",
				"challenge", fmt.Sprintf("%x", result.Work.Challenge),
				"solution", result.Nonce,
				"difficulty", result.Work.Challenge.Difficulty,
				"requestIDs", fmt.Sprintf("%+v", result.Work.Challenge.RequestIDs),
			)
			s.Submit(ctx, result)
		}
	}
}

func (s *Submitter) CancelPendingSubmit() {
	if s.lastSubmitCncl != nil {
		s.lastSubmitCncl()
	}
}

func (s *Submitter) Stop() {
	s.close()
}

func (s *Submitter) blockUntilTimeToSubmit(newChallengeReplace context.Context) {
	var (
		lastSubmit time.Duration
		timestamp  *time.Time
		err        error
	)
	for {
		select {
		case <-newChallengeReplace.Done():
			level.Info(s.logger).Log("msg", "canceled pending submit while gettting last submit time")
		default:
		}
		lastSubmit, timestamp, err = s.lastSubmit()
		if err != nil {
			level.Debug(s.logger).Log("msg", "checking last submit time", "err", err)
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
	if lastSubmit < s.cfg.Mine.MinSubmitPeriod.Duration {
		level.Info(s.logger).Log("msg", "min transaction submit threshold hasn't passed",
			"nextSubmit", time.Duration(s.cfg.Mine.MinSubmitPeriod.Nanoseconds())-lastSubmit,
			"lastSubmit", lastSubmit,
			"lastSubmitTimestamp", timestamp.Format("2006-01-02 15:04:05.000000"),
			"minSubmitPeriod", s.cfg.Mine.MinSubmitPeriod,
		)
		timeToSubmit, cncl := context.WithDeadline(newChallengeReplace, timestamp.Add(s.cfg.Mine.MinSubmitPeriod.Duration))
		defer cncl()
		select {
		case <-newChallengeReplace.Done():
			level.Info(s.logger).Log("msg", "canceled pending submit while waiting for the time to submit")
		case <-timeToSubmit.Done(): // 15min since last submit has passed so can unblock.
		}
	}
}

func (s *Submitter) canSubmit() error {
	if s.reward != nil { // Profit check is enabled.
		profitPercent, err := s.profitPercent()
		if _, ok := errors.Cause(err).(reward.ErrNoDataForSlot); ok {
			level.Warn(s.logger).Log("msg", "skipping profit check when the slot has no record for how much gas it uses", "err", err)
		} else if err != nil {
			return errors.Wrapf(err, "submit solution profit check")
		} else if profitPercent < int64(s.cfg.Mine.ProfitThreshold) {
			return errors.Errorf("profit:%v lower then the profit threshold:%v", profitPercent, s.cfg.Mine.ProfitThreshold)
		}
	}

	statusID, err := s.minerStatus()
	if err != nil {
		return errors.Wrap(err, "getting miner status")
	}
	if statusID != 1 {
		return errors.Errorf("miner is not in a status that can submit:%v", minerStatusName(statusID))
	}

	return nil
}

func (s *Submitter) profitPercent() (int64, error) {
	slot, err := s.reward.Slot()
	if err != nil {
		return 0, errors.Wrapf(err, "getting current slot")
	}
	gasPrice, err := s.reward.GasPrice()
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

	return s.reward.Current(slot, gasPrice)
}

func (s *Submitter) Submit(newChallengeReplace context.Context, result *mining.Result) {
	go func(newChallengeReplace context.Context, result *mining.Result) {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-newChallengeReplace.Done():
				level.Info(s.logger).Log("msg", "pending submit canceled")
				return
			default:
			}

			s.blockUntilTimeToSubmit(newChallengeReplace)
			if err := s.canSubmit(); err != nil {
				level.Info(s.logger).Log("msg", "can't submit and will retry later", "reason", err)
				<-ticker.C
				continue
			}
			for {
				select {
				case <-newChallengeReplace.Done():
					level.Info(s.logger).Log("msg", "pending submit canceled")
					return
				default:
				}

				reqVals, err := s.requestVals(result.Work.Challenge.RequestIDs)
				if err != nil {
					level.Error(s.logger).Log("msg", "adding the request ids, retrying", "err", err)
					<-ticker.C
					continue
				}
				level.Info(s.logger).Log(
					"msg", "sending solution to the chain",
					"solutionNonce", result.Nonce,
					"IDs", fmt.Sprintf("%+v", result.Work.Challenge.RequestIDs),
					"vals", fmt.Sprintf("%+v", reqVals),
				)
				tx, recieipt, err := s.transactor.Transact(newChallengeReplace, result.Nonce, result.Work.Challenge.RequestIDs, reqVals)
				if err != nil {
					s.submitFailCount.Inc()
					level.Error(s.logger).Log("msg", "submiting a solution", "err", err)
					return
				}

				if recieipt.Status != types.ReceiptStatusSuccessful {
					s.submitFailCount.Inc()
					level.Error(s.logger).Log("msg", "submiting solution status not success", "status", recieipt.Status)
					return
				}
				level.Info(s.logger).Log("msg", "successfully submited solution",
					"txHash", tx.Hash().String(),
					"nonce", tx.Nonce(),
					"gasPrice", tx.GasPrice(),
					"data", fmt.Sprintf("%x", tx.Data()),
					"value", tx.Value(),
				)
				s.submitCount.Inc()

				slot, err := s.reward.Slot()
				if err != nil {
					level.Error(s.logger).Log("msg", "getting _SLOT_PROGRESS for saving gas used", "err", err)
				} else {
					s.reward.SaveGasUsed(recieipt.GasUsed, slot)
				}

				return
			}
		}
	}(newChallengeReplace, result)
}

func (s *Submitter) requestVals(requestIDs [5]*big.Int) ([5]*big.Int, error) {
	var currentValues [5]*big.Int

	// The submit contains values for 5 data IDs so add them here.
	for i := 0; i < 5; i++ {
		valKey := fmt.Sprintf("%s%d", db.QueriedValuePrefix, requestIDs[i].Uint64())
		m, err := s.proxy.BatchGet([]string{valKey})
		if err != nil {
			return currentValues, errors.Wrapf(err, "retrieve pricing from db for data id:%v", requestIDs[i].Uint64())
		}
		val := m[valKey]
		var value *big.Int
		if len(val) == 0 {
			jsonFile, err := os.Open(s.cfg.ManualDataFile)
			if err != nil {
				return currentValues, errors.Wrapf(err, "manualData read Error")
			}
			defer jsonFile.Close()
			byteValue, _ := ioutil.ReadAll(jsonFile)
			var result map[string]map[string]uint
			_ = json.Unmarshal([]byte(byteValue), &result)
			_id := strconv.FormatUint(requestIDs[i].Uint64(), 10)
			val := result[_id]["VALUE"]
			if val == 0 {
				return currentValues, errors.Errorf("pricing data not available from db or the manual file for request id:%v", requestIDs[i].Uint64())
			}
			value = big.NewInt(int64(val))
		} else {
			value, err = hexutil.DecodeBig(string(val))
			if err != nil {
				if requestIDs[i].Uint64() > index.MaxPSRID() {
					level.Error(s.logger).Log(
						"msg", "decoding price value prior to submiting solution",
						"err", err,
					)
					if len(val) == 0 {
						level.Error(s.logger).Log("msg", "0 value being submitted")
						currentValues[i] = big.NewInt(0)
					}
					continue
				}
				return currentValues, errors.Errorf("no value in database,  reg id:%v", requestIDs[i].Uint64())
			}
		}
		currentValues[i] = value
	}
	return currentValues, nil
}

func (s *Submitter) minerStatus() (int64, error) {
	// Check if the staked account is in dispute before sending a transaction.
	statusID, _, err := s.contractInstance.GetStakerInfo(&bind.CallOpts{}, s.account.Address)
	if err != nil {
		return 0, errors.Wrapf(err, "getting staker info from contract addr:%v", s.account.Address)
	}
	return statusID.Int64(), nil
}

func (s *Submitter) lastSubmit() (time.Duration, *time.Time, error) {
	address := "000000000000000000000000" + s.account.Address.Hex()[2:]
	decoded, err := hex.DecodeString(address)
	if err != nil {
		return 0, nil, errors.Wrapf(err, "decoding address")
	}
	last, err := s.contractInstance.GetUintVar(nil, util.Keccak256(decoded))

	if err != nil {
		return 0, nil, errors.Wrapf(err, "getting last submit time for:%v", s.account.Address.String())
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
