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
	"github.com/tellor-io/telliot/pkg/rpc"
)

const ComponentName = "tasker"

// SubmissionCanceler will be used to cancel current submits when new challenge has been arrived.
type SubmissionCanceler interface {
	CancelPendingSubmit()
}

type Tasker struct {
	ctx                 context.Context
	close               context.CancelFunc
	logger              log.Logger
	proxy               db.DataServerProxy
	accounts            []*rpc.Account
	contractInstance    *contracts.ITellor
	client              contracts.ETHClient
	cfg                 *config.Config
	workSinks           map[string]chan *mining.Work
	SubmissionCancelers []SubmissionCanceler
}

func NewTasker(ctx context.Context, logger log.Logger, cfg *config.Config, proxy db.DataServerProxy, client contracts.ETHClient, contract *contracts.ITellor, accounts []*rpc.Account) (*Tasker, map[string]chan *mining.Work, error) {
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
		ctx:                 ctx,
		close:               close,
		proxy:               proxy,
		accounts:            accounts,
		contractInstance:    contract,
		workSinks:           workSinks,
		logger:              log.With(filterLog, "component", ComponentName),
		cfg:                 cfg,
		client:              client,
		SubmissionCancelers: make([]SubmissionCanceler, 0),
	}
	return tasker, tasker.workSinks, nil
}

func (mt *Tasker) AddSubmitCanceler(submissionCanceler SubmissionCanceler) {
	mt.SubmissionCancelers = append(mt.SubmissionCancelers, submissionCanceler)
}

func (mt *Tasker) getNewChallengeChannel() (chan *tellor.ITellorNewChallenge, event.Subscription, error) {
	sink := make(chan *tellor.ITellorNewChallenge)
	var tellorFilterer *tellor.ITellorFilterer
	tellorFilterer, err := tellor.NewITellorFilterer(mt.contractInstance.Address, mt.client)
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting ITellorFilterer instance")
	}
	sub, err := tellorFilterer.WatchNewChallenge(&bind.WatchOpts{}, sink, nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting NewChallenge channel")
	}
	return sink, sub, nil
}

func (mt *Tasker) sendWork(challenge *tellor.ITellorNewChallenge, delay time.Duration) {
	if challenge.CurrentRequestId[0].Int64() > int64(100) || challenge.CurrentRequestId[0].Int64() == 0 {
		level.Warn(mt.logger).Log("msg", "new current variables request ID not correct - contract about to be upgraded")
		return
	}
	newChallenge := &mining.MiningChallenge{
		Challenge:  challenge.CurrentChallenge[:],
		Difficulty: challenge.Difficulty,
		RequestIDs: challenge.CurrentRequestId,
	}
	for _, acc := range mt.accounts {
		level.Info(mt.logger).Log("msg", "new challenge",
			"addr", acc.Address.String(),
			"challenge", fmt.Sprintf("%x", newChallenge.Challenge),
			"difficulty", newChallenge.Difficulty,
			"requestIDs", fmt.Sprintf("%+v", newChallenge.RequestIDs),
		)
		mt.workSinks[acc.Address.String()] <- &mining.Work{Challenge: newChallenge, PublicAddr: acc.Address.String(), Start: uint64(rand.Int63()), N: math.MaxInt64}
		select {
		case <-mt.ctx.Done():
			return
		case <-time.After(delay):
			continue
		}
	}

}

func (mt *Tasker) Start() error {
	var err error
	level.Info(mt.logger).Log("msg", "tasker has been started")

	// Getting current challenge from the contract.
	newVariables, err := mt.contractInstance.GetNewCurrentVariables(nil)
	if err != nil {
		level.Warn(mt.logger).Log("msg", "new current variables retrieval - contract might not be upgraded", "err", err)
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

	level.Info(mt.logger).Log("msg", "sending the initial challenge to the miner")
	mt.sendWork(currentChallenge, 5*time.Second)

	// Subscribe and wait until the context cancellation event.
	var sink chan *tellor.ITellorNewChallenge
	var sub event.Subscription

	// Initial subscription.
	for {
		sink, sub, err = mt.getNewChallengeChannel()
		if err != nil {
			level.Error(mt.logger).Log("msg", "initial subscribing to NewChallenge events failed")
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}

	for {
		select {
		case <-mt.ctx.Done():
			level.Info(mt.logger).Log("msg", "unsubscribing from NewChallenge events")
			if sub != nil {
				sub.Unsubscribe()
			}
			level.Info(mt.logger).Log("msg", "tasker shutdown complete")
			return mt.ctx.Err()
		case err := <-sub.Err():
			if err != nil {
				level.Error(mt.logger).Log(
					"msg",
					"new challenge subscription error",
					"err", err)
			}

			// Trying to resubscribe until it succeeds.
			for {
				sink, sub, err = mt.getNewChallengeChannel()
				if err != nil {
					level.Error(mt.logger).Log("msg", "re-subscribing to NewChallenge events failed")
					time.Sleep(1 * time.Second)
					continue
				}
				break
			}
			level.Info(mt.logger).Log("msg", "re-subscribed to NewChallenge events")
		case vLog := <-sink:
			for _, canceler := range mt.SubmissionCancelers {
				canceler.CancelPendingSubmit()
			}
			mt.sendWork(vLog, 0)
		}
	}
}

func (mt *Tasker) Stop() {
	mt.close()
}
