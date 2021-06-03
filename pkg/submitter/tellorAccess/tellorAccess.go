// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tellorAccess

import (
	"context"
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
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/transactor"
)

const ComponentName = "submitterTellorAccess"

type Config struct {
	Enabled  bool
	LogLevel string
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
	self.Submit(1)
	self.Submit(2)

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-self.ctx.Done():
			return self.ctx.Err()
		case <-ticker.C:
			self.Submit(1)
			self.Submit(2)
		}
	}
}

func (self *Submitter) Stop() {
	self.close()
}

func (self *Submitter) Submit(reqID int64) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-self.ctx.Done():
			level.Info(self.logger).Log("msg", "submit canceled")
			return
		default:
		}

		canSubmit, err := self.contractInstance.IsReporter(&bind.CallOpts{}, self.account.Address)
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
			case <-self.ctx.Done():
				level.Info(self.logger).Log("msg", "submit canceled")
				return
			default:
			}

			// val, err := self.aggregator.GetValueForIDWithDefaultGranularity(reqID, time.Now())
			// if err != nil {
			// 	level.Error(self.logger).Log("msg", "getting the value from the aggregator", "err", err)
			// 	return
			// }
			val := int64(999)

			level.Info(self.logger).Log(
				"msg", "sending solution to the chain",
				"ID", reqID,
				"val", val,
			)

			f := func(auth *bind.TransactOpts) (*types.Transaction, error) {
				_reqID := big.NewInt(reqID)
				_val := big.NewInt(val)
				return self.contractInstance.SubmitValue(auth, _reqID, _val)
			}
			tx, recieipt, err := self.transactor.Transact(self.ctx, f)
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

			return
		}
	}
}
