// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tasker

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/contracts/tellor"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/mining"
)

const ComponentName = "tasker"

// SubmitCanceler will be used to cancel current submits when new challenge has been arrived.
type SubmitCanceler interface {
	CancelPendingSubmit()
}

type Tasker struct {
	ctx              context.Context
	close            context.CancelFunc
	logger           log.Logger
	proxy            db.DataServerProxy
	accounts         []*config.Account
	contractInstance *contracts.ITellor
	client           contracts.ETHClient
	cfg              *config.Config
	workSinks        map[string]chan *mining.Work
	SubmitCancelers  []SubmitCanceler
	txPending        context.CancelFunc
}

func NewTasker(ctx context.Context, logger log.Logger, cfg *config.Config, proxy db.DataServerProxy, client contracts.ETHClient, contract *contracts.ITellor, accounts []*config.Account) (*Tasker, map[string]chan *mining.Work, error) {
	ctx, close := context.WithCancel(ctx)
	workSinks := make(map[string]chan *mining.Work)
	for _, acc := range accounts {
		workSinks[acc.Address.String()] = make(chan *mining.Work)
	}
	filterLog, err := logging.ApplyFilter(*cfg, ComponentName, logger)
	if err != nil {
		close()
		return nil, nil, errors.Wrap(err, "apply filter logger")
	}
	tasker := &Tasker{
		ctx:              ctx,
		close:            close,
		proxy:            proxy,
		accounts:         accounts,
		contractInstance: contract,
		workSinks:        workSinks,
		logger:           log.With(filterLog, "component", ComponentName),
		cfg:              cfg,
		client:           client,
		SubmitCancelers:  make([]SubmitCanceler, 0),
	}
	return tasker, tasker.workSinks, nil
}

func (self *Tasker) AddSubmitCanceler(SubmitCanceler SubmitCanceler) {
	self.SubmitCancelers = append(self.SubmitCancelers, SubmitCanceler)
}

func (self *Tasker) getNewChallengeChannel() (chan *tellor.ITellorNewChallenge, event.Subscription, error) {
	sink := make(chan *tellor.ITellorNewChallenge)
	var tellorFilterer *tellor.ITellorFilterer
	tellorFilterer, err := tellor.NewITellorFilterer(self.contractInstance.Address, self.client)
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting ITellorFilterer instance")
	}
	sub, err := tellorFilterer.WatchNewChallenge(&bind.WatchOpts{}, sink, nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting NewChallenge channel")
	}
	return sink, sub, nil
}

func (self *Tasker) sendWork(challenge *tellor.ITellorNewChallenge) {
	if challenge.CurrentRequestId[0].Int64() > int64(100) || challenge.CurrentRequestId[0].Int64() == 0 {
		level.Warn(self.logger).Log("msg", "new current variables request ID not correct - contract about to be upgraded")
		return
	}
	newChallenge := &mining.MiningChallenge{
		Challenge:  challenge.CurrentChallenge[:],
		Difficulty: challenge.Difficulty,
		RequestIDs: challenge.CurrentRequestId,
	}
	for _, acc := range self.accounts {
		level.Info(self.logger).Log("msg", "new challenge",
			"addr", acc.Address.String(),
			"challenge", fmt.Sprintf("%x", newChallenge.Challenge),
			"difficulty", newChallenge.Difficulty,
			"requestIDs", fmt.Sprintf("%+v", newChallenge.RequestIDs),
		)
		self.workSinks[acc.Address.String()] <- &mining.Work{Challenge: newChallenge, PublicAddr: acc.Address.String(), Start: uint64(rand.Int63()), N: math.MaxInt64}
	}

}

func (self *Tasker) Start() error {
	var err error
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	level.Info(self.logger).Log("msg", "tasker has been started")

	// Getting current challenge from the contract.
	newVariables, err := self.contractInstance.GetNewCurrentVariables(nil)
	if err != nil {
		level.Warn(self.logger).Log("msg", "new current variables retrieval - contract might not be upgraded", "err", err)
		return err
	}

	currentChallenge := &tellor.ITellorNewChallenge{
		CurrentChallenge: newVariables.Challenge,
		Difficulty:       newVariables.Difficutly,
		CurrentRequestId: newVariables.RequestIds,
		TotalTips:        newVariables.Tip,
	}
	if err != nil {
		return errors.Wrap(err, "getting the current challenge")
	}

	level.Info(self.logger).Log("msg", "sending the initial challenge to the miner")
	self.sendWork(currentChallenge)

	// Subscribe and wait until the context cancellation event.
	var sink chan *tellor.ITellorNewChallenge
	var sub event.Subscription

	// Initial subscription.
	for {
		sink, sub, err = self.getNewChallengeChannel()
		if err != nil {
			level.Error(self.logger).Log("msg", "initial subscribing to NewChallenge events failed")
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}

	for {
		select {
		case <-self.ctx.Done():
			return nil
		case err := <-sub.Err():
			if err != nil {
				level.Error(self.logger).Log(
					"msg",
					"new challenge subscription error",
					"err", err)
			}

			// Trying to resubscribe until it succeeds.
			for {
				sink, sub, err = self.getNewChallengeChannel()
				if err != nil {
					level.Error(self.logger).Log("msg", "re-subscribing to NewChallenge events failed")
					select {
					case <-ticker.C:
						continue
					case <-self.ctx.Done():
						return nil
					}
				}
				break
			}
			level.Info(self.logger).Log("msg", "re-subscribed to NewChallenge events")
		case vLog := <-sink:
			if self.txPending != nil {
				self.txPending()
			}

			ctxPending, ctxPendingCncl := context.WithCancel(self.ctx)
			self.txPending = ctxPendingCncl

			go self.sendWhenConfirmed(ctxPending, vLog)
		}
	}
}

func (self *Tasker) sendWhenConfirmed(ctx context.Context, vLog *tellor.ITellorNewChallenge) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		// Send the event only when the tx that emitted the event has been confirmed.
		receipt, err := self.client.TransactionReceipt(ctx, vLog.Raw.TxHash)
		if receipt != nil && receipt.Status == types.ReceiptStatusSuccessful {
			// The new event arrives at the same time as the pending TX get confirmed so
			// add small delay here to avoid race conditions of canceling the pending TXs.
			time.Sleep(2 * time.Second)
			for _, canceler := range self.SubmitCancelers {
				canceler.CancelPendingSubmit()
			}
			self.sendWork(vLog)
			return
		}
		if err != nil {
			level.Error(self.logger).Log("msg", "receipt retrieval", "err", err)
		} else {
			level.Debug(self.logger).Log("msg", "transaction not yet mined", "tx", vLog.Raw.TxHash)
		}
		select {
		case <-ctx.Done():
			level.Info(self.logger).Log("msg", "tx confirmation check canceled")
			return
		case <-ticker.C:
			continue
		}
	}
}

func (self *Tasker) Stop() {
	self.close()
}
