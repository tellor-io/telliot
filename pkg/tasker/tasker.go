// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tasker

import (
	"context"
	"fmt"
	"math"
	"math/rand"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/contracts/tellor"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/mining"
	"github.com/tellor-io/telliot/pkg/rpc"
)

const ComponentName = "tasker"

type Tasker struct {
	ctx              context.Context
	close            context.CancelFunc
	logger           log.Logger
	proxy            db.DataServerProxy
	accounts         []*rpc.Account
	contractInstance *contracts.ITellor
	client           contracts.ETHClient
	cfg              *config.Config
	workSink         chan *mining.Work
	Running          bool
	done             chan bool
	resubscribe      chan bool
}

func CreateTasker(ctx context.Context, logger log.Logger, cfg *config.Config, proxy db.DataServerProxy, client contracts.ETHClient, contract *contracts.ITellor, accounts []*rpc.Account) (*Tasker, chan *mining.Work) {
	ctx, close := context.WithCancel(ctx)
	tasker := &Tasker{
		ctx:              ctx,
		close:            close,
		proxy:            proxy,
		accounts:         accounts,
		contractInstance: contract,
		workSink:         make(chan *mining.Work, 1),
		done:             make(chan bool),
		resubscribe:      make(chan bool, 1),
		logger:           log.With(logger, "component", ComponentName),
		cfg:              cfg,
		client:           client,
	}
	return tasker, tasker.workSink
}

func (mt *Tasker) getCurrentChallenge() (*tellor.ITellorNewChallenge, error) {
	newVariables, err := mt.contractInstance.GetNewCurrentVariables(nil)
	if err != nil {
		level.Warn(mt.logger).Log("msg", "new current variables retrieval - contract might not be upgraded", "err", err)
		return nil, err
	}
	return &tellor.ITellorNewChallenge{
		CurrentChallenge: newVariables.Challenge,
		Difficulty:       newVariables.Difficutly,
		CurrentRequestId: newVariables.RequestIds,
		TotalTips:        newVariables.Tip,
	}, err
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

func (mt *Tasker) subscribeToNewChallenge() error {
	var sink chan *tellor.ITellorNewChallenge
	var err error
	var sub event.Subscription
	var failedSubscriptionCount int

	// Initial subscription.
	sink, sub, err = mt.getNewChallengeChannel()
	if err != nil {
		level.Error(mt.logger).Log("msg", "initial subscribing to NewChallenge events failed")
		mt.resubscribe <- true
	}
	for {
		select {
		case <-mt.done:
			if sub != nil {
				sub.Unsubscribe()
			}
			level.Info(mt.logger).Log("msg", "unsubscribed to NewChallenge events")
			return nil
		case err := <-sub.Err():
			if err != nil {
				level.Error(mt.logger).Log(
					"msg",
					"new challenge subscription error",
					"err", err)
			}
			mt.resubscribe <- true
		case <-mt.resubscribe:
			if failedSubscriptionCount == 10 {
				return errors.New("failed to subscribe to NewChallenge events after 10 retries")
			}
			sink, sub, err = mt.getNewChallengeChannel()
			if err != nil {
				failedSubscriptionCount++
				level.Error(mt.logger).Log("msg", "subscribing to NewChallenge events failed", "retry", failedSubscriptionCount)
				mt.resubscribe <- true
				continue
			}
			failedSubscriptionCount = 0
			level.Info(mt.logger).Log("msg", "subscribed to NewChallenge events")
		case vLog := <-sink:
			mt.sendWork(vLog)
		}
	}
}

func (mt *Tasker) sendWork(challenge *tellor.ITellorNewChallenge) {
	if challenge.CurrentRequestId[0].Int64() > int64(100) || challenge.CurrentRequestId[0].Int64() == 0 {
		level.Warn(mt.logger).Log("msg", "new current variables request ID not correct - contract about to be upgraded")
		return
	}
	newChallenge := &mining.MiningChallenge{
		Challenge:  challenge.CurrentChallenge[:],
		Difficulty: challenge.Difficulty,
		RequestIDs: challenge.CurrentRequestId,
	}

	level.Info(mt.logger).Log("msg", "sending new challenge to the miner manager",
		"challenge", fmt.Sprintf("%x", newChallenge.Challenge),
		"difficulty", newChallenge.Difficulty,
		"requestIDs", fmt.Sprintf("%+v", newChallenge.RequestIDs),
	)
	mt.workSink <- &mining.Work{Challenge: newChallenge, PublicAddr: mt.accounts[0].Address.String(), Start: uint64(rand.Int63()), N: math.MaxInt64}
}

func (mt *Tasker) Start() error {
	level.Info(mt.logger).Log("msg", "tasker has been started")
	currentChallenge, err := mt.getCurrentChallenge()
	if err != nil {
		return errors.Wrap(err, "tasker getting the current challenge")
	}
	level.Info(mt.logger).Log("msg", "tasker is sending the initial challenge to the miner")
	mt.sendWork(currentChallenge)
	// Subscribe and wait until the context cancellation event.
	return mt.subscribeToNewChallenge()
}

func (mt *Tasker) Stop() {
	mt.close()
	level.Info(mt.logger).Log("msg", "tasker shutdown complete")
}
