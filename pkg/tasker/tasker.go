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


type MiningTasker struct {
	ctx              context.Context
	close            context.CancelFunc
	logger           log.Logger
	proxy            db.DataServerProxy
	accounts         []*rpc.Account
	currChallenge    *mining.MiningChallenge
	contractInstance *contracts.ITellor
	client           contracts.ETHClient
	cfg              *config.Config
	workSink         chan *mining.Work
	Running          bool
	done             chan bool
	resubscribe      chan bool
}

func CreateTasker(ctx context.Context, logger log.Logger, cfg *config.Config, proxy db.DataServerProxy, client contracts.ETHClient, contract *contracts.ITellor, accounts []*rpc.Account) (*MiningTasker, chan *mining.Work) {
	ctx, close := context.WithCancel(ctx)
	tasker := &MiningTasker{
		ctx:              ctx,
		close:            close,
		proxy:            proxy,
		accounts:         accounts,
		contractInstance: contract,
		workSink:         make(chan *mining.Work, 1),
		done:             make(chan bool),
		resubscribe:      make(chan bool),
		logger:           log.With(logger, "component", ComponentName),
		cfg:              cfg,
		client:           client,
	}
	return tasker, tasker.workSink
}

func (mt *MiningTasker) getCurrentChallenge() (*tellor.ITellorNewChallenge, error) {
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

func (mt *MiningTasker) getNewChallengeChannel() (chan *tellor.ITellorNewChallenge, event.Subscription, error) {
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

func (mt *MiningTasker) subscribeToNewChallenge() error {
	sink, sub, err := mt.getNewChallengeChannel()
	if err != nil {
		return err
	}
	level.Info(mt.logger).Log("msg", "subscribed to NewChallenge events")

	go func() {
		for {
			select {
			case <-mt.done:
				sub.Unsubscribe()
				level.Info(mt.logger).Log("msg", "unsubscribed to NewChallenge events")
				return
			case err := <-sub.Err():
				if err != nil {
					level.Error(mt.logger).Log(
						"msg",
						"new challenge subscription error",
						"err", err)
				}
				mt.resubscribe <- true
			case vLog := <-sink:
				mt.sendWork(vLog)
			}
		}
	}()
	return nil
}

func (mt *MiningTasker) sendWork(challenge *tellor.ITellorNewChallenge) {
	if challenge.CurrentRequestId[0].Int64() > int64(100) || challenge.CurrentRequestId[0].Int64() == 0 {
		level.Warn(mt.logger).Log("msg", "new current variables request ID not correct - contract about to be upgraded")
		return
	}
	work := mt.CreateWork(challenge)
	// Send new work to the sink.
	if work != nil {
		mt.workSink <- work
	}
}

func (mt *MiningTasker) Start() error {
	level.Info(mt.logger).Log("msg", "tasker has been started")
	currentChallenge, err := mt.getCurrentChallenge()
	if err != nil {
		return errors.Wrap(err, "tasker getting the current challenge")
	}
	level.Info(mt.logger).Log("msg", "tasker is sending the initial challenge to the miner")
	mt.sendWork(currentChallenge)
	err = mt.subscribeToNewChallenge()
	if err != nil {
		return errors.Wrap(err, "tasker subscribing to new challenges")
	}
	for {
		select {
		case <-mt.resubscribe:
			err = mt.subscribeToNewChallenge()
			if err != nil {
				return errors.Wrap(err, "tasker resubscribing to new challenges")
			}
		case <-mt.ctx.Done():
			mt.done <- true
		}
	}
}

func (mt *MiningTasker) Stop() {
	mt.close()
	level.Info(mt.logger).Log("msg", "tasker shutdown complete")
}

func (mt *MiningTasker) CreateWork(challenge *tellor.ITellorNewChallenge) *mining.Work {
	newChallenge := &mining.MiningChallenge{
		Challenge:  challenge.CurrentChallenge[:],
		Difficulty: challenge.Difficulty,
		RequestIDs: challenge.CurrentRequestId,
	}

	level.Debug(mt.logger).Log("msg", "new challenge for mining",
		"hex", fmt.Sprintf("%x", newChallenge.Challenge),
		"difficulty", newChallenge.Difficulty,
		"requestIDs", fmt.Sprintf("%+v", newChallenge.RequestIDs),
	)

	mt.currChallenge = newChallenge
	return &mining.Work{Challenge: newChallenge, PublicAddr: mt.accounts[0].Address.String(), Start: uint64(rand.Int63()), N: math.MaxInt64}
}
