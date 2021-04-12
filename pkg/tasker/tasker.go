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
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/mining"
)

const ComponentName = "taskerNewChallenge"

// SubmitCanceler will be used to cancel current submits when new event arrives.
type SubmitCanceler interface {
	CancelPendingSubmit()
}

type Tasker struct {
	ctx              context.Context
	close            context.CancelFunc
	logger           log.Logger
	accounts         []*config.Account
	contractInstance *contracts.ITellor
	client           contracts.ETHClient
	cfg              *config.Config
	workSinks        map[string]chan *mining.Work
	SubmitCancelers  []SubmitCanceler
	txPending        context.CancelFunc
}

func NewTasker(
	ctx context.Context,
	logger log.Logger,
	cfg *config.Config,
	client contracts.ETHClient,
	contract *contracts.ITellor,
	accounts []*config.Account,
) (*Tasker, map[string]chan *mining.Work, error) {
	ctx, close := context.WithCancel(ctx)
	workSinks := make(map[string]chan *mining.Work)
	for _, acc := range accounts {
		workSinks[acc.Address.String()] = make(chan *mining.Work)
	}
	logger, err := logging.ApplyFilter(*cfg, ComponentName, logger)
	if err != nil {
		close()
		return nil, nil, errors.Wrap(err, "apply filter logger")
	}
	tasker := &Tasker{
		ctx:              ctx,
		close:            close,
		accounts:         accounts,
		contractInstance: contract,
		workSinks:        workSinks,
		logger:           log.With(logger, "component", ComponentName),
		cfg:              cfg,
		client:           client,
		SubmitCancelers:  make([]SubmitCanceler, 0),
	}
	return tasker, tasker.workSinks, nil
}

func (self *Tasker) AddSubmitCanceler(SubmitCanceler SubmitCanceler) {
	self.SubmitCancelers = append(self.SubmitCancelers, SubmitCanceler)
}

func (self *Tasker) newSub(output chan *tellor.ITellorNewChallenge) (event.Subscription, error) {
	var tellorFilterer *tellor.ITellorFilterer
	tellorFilterer, err := tellor.NewITellorFilterer(self.contractInstance.Address, self.client)
	if err != nil {
		return nil, errors.Wrap(err, "getting filter instance")
	}
	sub, err := tellorFilterer.WatchNewChallenge(&bind.WatchOpts{Context: self.ctx}, output, nil)
	if err != nil {
		return nil, errors.Wrap(err, "getting subscription channel")
	}
	return sub, nil
}

func (self *Tasker) sendWork(challenge *tellor.ITellorNewChallenge) {
	newChallenge := &mining.MiningChallenge{
		Challenge:  challenge.CurrentChallenge[:],
		Difficulty: challenge.Difficulty,
		RequestIDs: challenge.CurrentRequestId,
	}
	for _, acc := range self.accounts {
		level.Info(self.logger).Log("msg", "new event",
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
	level.Info(self.logger).Log("msg", "starting")

	// Getting current challenge from the contract.
	newVariables, err := self.contractInstance.GetNewCurrentVariables(nil)
	if err != nil {
		level.Warn(self.logger).Log("msg", "getting new current variables", "err", err)
		return errors.Wrap(err, "getting GetNewCurrentVariables")
	}

	currentChallenge := &tellor.ITellorNewChallenge{
		CurrentChallenge: newVariables.Challenge,
		Difficulty:       newVariables.Difficutly,
		CurrentRequestId: newVariables.RequestIds,
		TotalTips:        newVariables.Tip,
	}

	level.Info(self.logger).Log("msg", "sending the initial event")
	self.sendWork(currentChallenge)

	// Subscribe and wait until the context cancellation event.
	events := make(chan *tellor.ITellorNewChallenge)
	var sub event.Subscription

	// Initial subscription.
	for {
		select {
		case <-self.ctx.Done():
			return nil
		default:
		}
		sub, err = self.newSub(events)
		if err != nil {
			level.Error(self.logger).Log("msg", "initial subscription to events failed")
			<-ticker.C
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
					"subscription error",
					"err", err)
			}

			// Trying to resubscribe until it succeeds.
			for {
				select {
				case <-self.ctx.Done():
					return nil
				default:
				}
				sub, err = self.newSub(events)
				if err != nil {
					level.Error(self.logger).Log("msg", "re-subscribing to events failed")
					<-ticker.C
					continue
				}
				break
			}
			level.Info(self.logger).Log("msg", "re-subscribed to events")
		case event := <-events:
			level.Debug(self.logger).Log("msg", "new event", "reorg", event.Raw.Removed)
			if self.txPending != nil {
				self.txPending()
				self.txPending = nil
			}

			if !event.Raw.Removed { // For reorg events just cancel the old TXs without sending this one.
				ctxPending, ctxPendingCncl := context.WithCancel(self.ctx)
				self.txPending = ctxPendingCncl
				go self.sendWhenConfirmed(ctxPending, event)
			}
		}
	}
}

func (self *Tasker) sendWhenConfirmed(ctx context.Context, vLog *tellor.ITellorNewChallenge) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		// Send the event only when the tx that emitted the event has been confirmed.
		receipt, err := self.client.TransactionReceipt(ctx, vLog.Raw.TxHash)
		if receipt != nil {
			// Send it only when the TX ReceiptStatusSuccessful or otherwise ignore.
			if receipt.Status == types.ReceiptStatusSuccessful {
				for _, canceler := range self.SubmitCancelers {
					canceler.CancelPendingSubmit()
				}
				self.sendWork(vLog)
			}
			return
		}
		if err != nil {
			level.Error(self.logger).Log("msg", "getting TX receipt", "err", err)
		} else {
			level.Debug(self.logger).Log("msg", "transaction not yet mined", "tx", vLog.Raw.TxHash)
		}

		<-ticker.C
		continue
	}
}

func (self *Tasker) Stop() {
	self.close()
}
