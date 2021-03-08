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
	workSinks        map[string]chan *mining.Work
}

func CreateTasker(ctx context.Context, logger log.Logger, cfg *config.Config, proxy db.DataServerProxy, client contracts.ETHClient, contract *contracts.ITellor, accounts []*rpc.Account) (*Tasker, map[string]chan *mining.Work) {
	ctx, close := context.WithCancel(ctx)
	workSinks := make(map[string]chan *mining.Work)
	for _, acc := range accounts {
		workSinks[acc.Address.String()] = make(chan *mining.Work)
	}
	tasker := &Tasker{
		ctx:              ctx,
		close:            close,
		proxy:            proxy,
		accounts:         accounts,
		contractInstance: contract,
		workSinks:        workSinks,
		logger:           log.With(logger, "component", ComponentName),
		cfg:              cfg,
		client:           client,
	}
	return tasker, tasker.workSinks
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
	for _, acc := range mt.accounts {
		level.Info(mt.logger).Log("msg", "new challenge",
			"addr", acc.Address.String(),
			"challenge", fmt.Sprintf("%x", newChallenge.Challenge),
			"difficulty", newChallenge.Difficulty,
			"requestIDs", fmt.Sprintf("%+v", newChallenge.RequestIDs),
		)
		go func(acc *rpc.Account) {
			mt.workSinks[acc.Address.String()] <- &mining.Work{Challenge: newChallenge, PublicAddr: acc.Address.String(), Start: uint64(rand.Int63()), N: math.MaxInt64}
		}(acc)
	}

}

func (mt *Tasker) Start() error {
	var err error
	level.Info(mt.logger).Log("msg", "tasker has been started")
	currentChallenge, err := mt.getCurrentChallenge()

	if err != nil {
		return errors.Wrap(err, "tasker getting the current challenge")
	}

	level.Info(mt.logger).Log("msg", "tasker is sending the initial challenge to the miner")
	mt.sendWork(currentChallenge)

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
			level.Info(mt.logger).Log("msg", "closing the tasker channels")
			for _, ch := range mt.workSinks {
				close(ch)
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
			mt.sendWork(vLog)
		}
	}
}

func (mt *Tasker) Stop() {
	mt.close()
}
