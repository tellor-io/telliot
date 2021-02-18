// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tasker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strconv"

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

/* const (
	statusWaitNext = iota + 1
	statusFailure
	statusSuccess
) */

/**
 * Tasker role is to pull challenge and other information from the data server
 * and push either new challenges to an output channel or boolean values to a
 * cancel channel. It doesn't know anything about the mining loop so that we
 * can operate in isolation and only evaluate incoming information based on a
 * set of simple rules. Here are the rules:
 *
 * - If the new challenge is zero, issue cancel
 * - If the miner address is in dispute, end program entirely
 * - If there is a pending txn for the miner address, issue cancel
 * - If there is no price data available for the current request, issue cancel
 * - Otherwise, push new challenge to output channel
 */

type MiningTasker struct {
	ctx                           context.Context
	close                         context.CancelFunc
	logger                        log.Logger
	proxy                         db.DataServerProxy
	accounts                      []*rpc.Account
	currChallenge                 *mining.MiningChallenge
	contractInstance              *contracts.ITellor
	instantSubmitSentForThisBlock bool
	client                        contracts.ETHClient
	cfg                           *config.Config
	workSink                      chan *mining.Work
	Running                       bool
	done                          chan bool
	resubscribe                   chan bool
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
	sub, err := mt.contractInstance.ITellorFilterer.WatchNewChallenge(&bind.WatchOpts{}, sink, nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error in getting NewChallenge channel")
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
	/* TODO: Do we need these anymore?
	dispKeys := []string{}
	for _, account := range mt.accounts {
		dispKeys = append(dispKeys, db.DisputeStatusKeyFor(account.Address))

	}
	keys := []string{
		db.RequestIdKey,
		db.RequestIdKey0,
		db.RequestIdKey1,
		db.RequestIdKey2,
		db.RequestIdKey3,
		db.RequestIdKey4,
		db.LastNewValueKey,
	}
	keys = append(keys, dispKeys...)

	m, err := mt.proxy.BatchGet(keys)
	if err != nil {
		level.Error(mt.logger).Log("msg", "get data from data proxy, cannot continue at all")
		return nil
	}

	level.Debug(mt.logger).Log("msg", "received data", "data", m)

	for _, dispKey := range dispKeys {
		if mt.checkDispute(m[dispKey]) == statusWaitNext {
			return nil
		}
	}*/

	diff := challenge.Difficulty
	reqIDs := challenge.CurrentRequestId
	for i := 0; i < 5; i++ {
		valKey := fmt.Sprintf("%s%d", db.QueriedValuePrefix, reqIDs[i].Uint64())
		m2, err := mt.proxy.BatchGet([]string{valKey})
		if err != nil {
			level.Info(mt.logger).Log(
				"msg", "retrieve pricing data for current request id",
				"err ", err,
			)
			//return nil, false
		}
		val := m2[valKey]
		if len(val) == 0 {
			if err != nil {
				level.Error(mt.logger).Log(
					"msg", "parsing config",
					"err ", err,
				)
				return nil
			}
			jsonFile, err := os.Open(mt.cfg.ManualDataFile)
			if err != nil {
				return nil
			}
			defer jsonFile.Close()
			byteValue, _ := ioutil.ReadAll(jsonFile)
			var result map[string]map[string]uint
			_ = json.Unmarshal([]byte(byteValue), &result)
			_id := strconv.FormatUint(reqIDs[i].Uint64(), 10)
			val := result[_id]["VALUE"]
			if val == 0 {
				level.Info(mt.logger).Log(
					"msg", "pricing data not available for request",
					"request", reqIDs[i].Uint64(),
				)
				return nil
			}
			level.Info(mt.logger).Log("msg", "USING MANUALLY ENTERED VALUE!!!! USE CAUTION")
		}
	}

	newChallenge := &mining.MiningChallenge{
		Challenge:  challenge.CurrentChallenge[:],
		Difficulty: diff,
		RequestIDs: reqIDs,
	}

	// If this challenge is already sent out, don't do it again.
	if mt.currChallenge != nil &&
		mt.instantSubmitSentForThisBlock && // Not instanst submit or instant submit but already has been sent out.
		bytes.Equal(newChallenge.Challenge, mt.currChallenge.Challenge) { // This a new oracle block so a new challenge.
		return nil
	}

	// When it is a new challenge reset the instant submit status.
	if mt.currChallenge != nil && !bytes.Equal(newChallenge.Challenge, mt.currChallenge.Challenge) {
		mt.instantSubmitSentForThisBlock = false
	}

	level.Debug(mt.logger).Log("msg", "new challenge for mining",
		"hex", fmt.Sprintf("%x", newChallenge.Challenge),
		"difficulty", diff,
		"requestIDs", fmt.Sprintf("%+v", reqIDs),
	)

	mt.currChallenge = newChallenge
	return &mining.Work{Challenge: newChallenge, PublicAddr: mt.accounts[0].Address.String(), Start: uint64(rand.Int63()), N: math.MaxInt64}
}

/* func (mt *MiningTasker) checkDispute(disp []byte) int {
	disputed, stat := mt.getInt(disp)
	if stat == statusWaitNext || stat == statusFailure {
		return stat
	}

	if disputed.Cmp(big.NewInt(1)) != 0 {
		level.Error(mt.logger).Log("msg", "miner is in dispute, cannot continue")
		return statusFailure
	}
	level.Debug(mt.logger).Log("msg", "miner is not in dispute, continuing")
	return statusSuccess
}

func (mt *MiningTasker) getInt(data []byte) (*big.Int, int) {
	if len(data) == 0 {
		return nil, statusWaitNext
	}

	val, err := hexutil.DecodeBig(string(data))
	if err != nil {
		level.Error(mt.logger).Log("msg", "decoding int", "err", err)
		return nil, statusFailure
	}
	return val, statusSuccess
} */
