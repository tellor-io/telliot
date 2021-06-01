// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tellorAccess

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
	"github.com/tellor-io/telliot/pkg/aggregator"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/format"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/mining"
	"github.com/tellor-io/telliot/pkg/transactor"
)

const ComponentName = "submitterTellorAccess"

type Config struct {
	Enabled         bool
	LogLevel        string
	MinSubmitPeriod format.Duration
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
	contractInstance *contracts.ITellorAccess
	transactor       transactor.Transactor
	submitCount      prometheus.Counter
	submitFailCount  prometheus.Counter
	lastSubmitCncl   context.CancelFunc
	aggregator       *aggregator.Aggregator
}

func New(
	ctx context.Context,
	logger log.Logger,
	cfg Config,
	client contracts.ETHClient,
	contractInstance *contracts.ITellorAccess,
	account *ethereum.Account,
	transactor transactor.Transactor,
	aggregator *aggregator.Aggregator,
) (*Submitter, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(logger, "component", ComponentName, "addr", account.Address.String()[:6])
	ctx, close := context.WithCancel(ctx)
	submitter := &Submitter{
		ctx:              ctx,
		close:            close,
		client:           client,
		cfg:              cfg,
		account:          account,
		logger:           logger,
		contractInstance: contractInstance,
		transactor:       transactor,
		aggregator:       aggregator,
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

	return submitter, nil
}

func (self *Submitter) Start() error {
	for {
		select {
		case <-self.ctx.Done():
			if self.lastSubmitCncl != nil {
				self.lastSubmitCncl()
			}
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

			canSubmit, err := self.contractInstance.IsReporter()(&bind.CallOpts{}, self.account.Address)
			if err != nil {
				level.Info(self.logger).Log("msg", "checking reporter status", "err", err)
				<-ticker.C
				continue
			}
			if !canSubmit {
				level.Info(self.logger).Log("msg", "addr not a reporter")
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
					self.reward.SaveGasUsed(recieipt.GasUsed, slot)
				}

				return
			}
		}
	}(newChallengeReplace, result)
}

func (self *Submitter) requestVals(requestIDs [5]*big.Int) ([5]*big.Int, error) {
	var currentValues [5]*big.Int
	for i, reqID := range requestIDs {
		val, err := self.aggregator.GetValueForIDWithDefaultGranularity(reqID.Int64(), time.Now())
		if err != nil {
			return currentValues, errors.Wrapf(err, "getting value for request ID:%v", reqID)
		}
		currentValues[i] = big.NewInt(int64(val))
	}
	return currentValues, nil
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
