// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package pow

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/event"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/contracts/tellor"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
)

const (
	statusWaitNext = iota + 1
	statusFailure
	statusSuccess
)

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

type WorkInfo struct {
	work          *Work
	instantSubmit bool
}

func (wi *WorkInfo) GetWork() (*Work, bool) {
	return wi.work, wi.instantSubmit
}

type MiningTasker struct {
	ctx                           context.Context
	close                         context.CancelFunc
	logger                        log.Logger
	proxy                         db.DataServerProxy
	accounts                      []*rpc.Account
	currChallenge                 *MiningChallenge
	contractInstance              *contracts.ITellor
	instantSubmitSentForThisBlock bool
	client                        contracts.ETHClient
	cfg                           *config.Config
	workSink                      chan *WorkInfo
	Running                       bool
	done                          chan bool
	resubscribe                   chan bool
}

func CreateTasker(ctx context.Context, logger log.Logger, cfg *config.Config, proxy db.DataServerProxy, client contracts.ETHClient, contract *contracts.ITellor, accounts []*rpc.Account) *MiningTasker {
	ctx, close := context.WithCancel(ctx)
	return &MiningTasker{
		ctx:              ctx,
		close:            close,
		proxy:            proxy,
		accounts:         accounts,
		contractInstance: contract,
		workSink:         make(chan *WorkInfo, 1),
		done:             make(chan bool),
		resubscribe:      make(chan bool),
		logger:           log.With(logger, "component", ComponentName),
	}
}

func (mt *MiningTasker) getCurrentChallenge() (*tellor.ITellorNewChallengeIterator, error) {
	var tellorLibraryFilterer *tellor.ITellorFilterer
	tellorLibraryFilterer, err := tellor.NewITellorFilterer(mt.contractInstance.Address, mt.client)
	if err != nil {
		return nil, err
	}
	itr, _ := tellorLibraryFilterer.FilterNewChallenge(&bind.FilterOpts{}, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error in get NewChallenge iterator")
	}
	return itr, err
}

func (mt *MiningTasker) getNewChallengeChannel() (chan *tellor.ITellorNewChallenge, event.Subscription, error) {
	sink := make(chan *tellor.ITellorNewChallenge)
	var tellorLibraryFilterer *tellor.ITellorFilterer
	tellorLibraryFilterer, err := tellor.NewITellorFilterer(mt.contractInstance.Address, mt.client)
	if err != nil {
		return nil, nil, err
	}
	sub, _ := tellorLibraryFilterer.WatchNewChallenge(&bind.WatchOpts{}, sink, nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error in get NewChallenge channel")
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
				mt.Running = false
				return
			case err := <-sub.Err():
				if err != nil {
					level.Error(mt.logger).Log(
						"msg",
						"new challenge subscription error",
						"err", err)
				}
				if mt.Running {
					mt.resubscribe <- true
				}
			case vLog := <-sink:
				mt.sendWork(vLog)
			}
		}
	}()
	return nil
}

func (mt *MiningTasker) WorkSink() chan *WorkInfo {
	return mt.workSink
}

func (mt *MiningTasker) sendWork(vLog *tellor.ITellorNewChallenge) {
	work, instantSubmit := mt.CreateWork(vLog)
	// Send new work to the sink.
	mt.workSink <- &WorkInfo{work, instantSubmit}
}

func (mt *MiningTasker) Start() error {
	mt.Running = true
	level.Info(mt.logger).Log("msg", "tasker has been started")
	currentChallenge, err := mt.getCurrentChallenge()
	if err != nil {
		return errors.Wrap(err, "tasker getting the current challenge")
	}
	level.Info(mt.logger).Log("msg", "tasker is sending the initial challenge to the miner")
	mt.sendWork(currentChallenge.Event)
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
	level.Info(mt.logger).Log("msg", "shutting down tasker...")
	mt.close()
	cnt := 0
	for {
		time.Sleep(500 * time.Millisecond)
		cnt++
		if !mt.Running {
			break
		}
		if cnt > 60 {
			level.Warn(mt.logger).Log("msg", "expected tasker to stop by now, Giving up...")
			return
		}
	}
	level.Info(mt.logger).Log("msg", "tasker shutdown complete")
}

func (mt *MiningTasker) CreateWork(challenge *tellor.ITellorNewChallenge) (*Work, bool) {
	dispKeys := []string{}
	for _, account := range mt.accounts {
		dispKeys = append(dispKeys, db.DisputeStatusPrefix+account.Address.String())

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
		level.Error(mt.logger).Log("msg", "get data from data proxy, cannot continue")
		return nil, false
	}

	level.Debug(mt.logger).Log("msg", "received data", "data", m)

	for _, dispKey := range dispKeys {
		if mt.checkDispute(m[dispKey]) == statusWaitNext {
			return nil, false
		}
	}

	diff := challenge.Difficulty
	var reqIDs [5]*big.Int

	instantSubmit := false

	timeOfLastNewValue, err := mt.contractInstance.GetUintVar(nil, rpc.Keccak256([]byte("_TIME_OF_LAST_NEW_VALUE")))
	if err != nil {
		level.Debug(mt.logger).Log("msg", "getting last submitted data in the oracle", "err", err)
		return nil, false
	}

	now := time.Now()
	tm := time.Unix(timeOfLastNewValue.Int64(), 0)
	level.Debug(mt.logger).Log("msg", "last submitted data in the oracle", "time", now.Sub(tm))
	if now.Sub(tm) >= time.Duration(15)*time.Minute {
		instantSubmit = true
	}

	r, stat := mt.getInt(m[db.RequestIdKey0])
	if stat == statusWaitNext || stat == statusFailure {
		return nil, false
	}
	reqIDs[0] = r

	r, stat = mt.getInt(m[db.RequestIdKey1])
	if stat == statusWaitNext || stat == statusFailure {
		return nil, false
	}
	reqIDs[1] = r

	r, stat = mt.getInt(m[db.RequestIdKey2])
	if stat == statusWaitNext || stat == statusFailure {
		return nil, false
	}
	reqIDs[2] = r

	r, stat = mt.getInt(m[db.RequestIdKey3])
	if stat == statusWaitNext || stat == statusFailure {
		return nil, false
	}
	reqIDs[3] = r

	r, stat = mt.getInt(m[db.RequestIdKey4])
	if stat == statusWaitNext || stat == statusFailure {
		return nil, false
	}
	reqIDs[4] = r

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
				return nil, false
			}
			jsonFile, err := os.Open(mt.cfg.ManualDataFile)
			if err != nil {
				return nil, false
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
				return nil, false
			}
			level.Info(mt.logger).Log("msg", "USING MANUALLY ENTERED VALUE!!!! USE CAUTION")
		}
	}

	newChallenge := &MiningChallenge{
		Challenge:  challenge.CurrentChallenge[:],
		Difficulty: diff,
		RequestIDs: reqIDs,
	}

	// If this challenge is already sent out, don't do it again.
	if mt.currChallenge != nil &&
		(!instantSubmit || (instantSubmit && mt.instantSubmitSentForThisBlock)) && // Not instanst submit or instant submit but already has been sent out.
		bytes.Equal(newChallenge.Challenge, mt.currChallenge.Challenge) { // This a new oracle block so a new challenge.
		return nil, false
	}

	// When it is a new challenge reset the instant submit status.
	if mt.currChallenge != nil && !bytes.Equal(newChallenge.Challenge, mt.currChallenge.Challenge) {
		mt.instantSubmitSentForThisBlock = false
	}

	if instantSubmit {
		mt.instantSubmitSentForThisBlock = true
	}

	level.Debug(mt.logger).Log("msg", "new challenge for mining",
		"hex", fmt.Sprintf("%x", m[db.CurrentChallengeKey]),
		"difficulty", diff,
		"requestIDs", fmt.Sprintf("%+v", reqIDs),
		"instantSubmit", instantSubmit,
	)

	mt.currChallenge = newChallenge
	return &Work{Challenge: newChallenge, PublicAddr: mt.accounts[0].Address.String(), Start: uint64(rand.Int63()), N: math.MaxInt64}, instantSubmit
}

func (mt *MiningTasker) checkDispute(disp []byte) int {
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
}
