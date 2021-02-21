// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package pow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
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

type MiningTasker struct {
	logger                        log.Logger
	proxy                         db.DataServerProxy
	pubKey                        string
	currChallenge                 *MiningChallenge
	cfg                           *config.Config
	contractInstance              *contracts.ITellor
	instantSubmitSentForThisBlock bool
}

func CreateTasker(logger log.Logger, cfg *config.Config, contractInstance *contracts.ITellor, proxy db.DataServerProxy) *MiningTasker {

	return &MiningTasker{
		proxy:            proxy,
		pubKey:           cfg.PublicAddress,
		logger:           log.With(logger, "component", ComponentName),
		cfg:              cfg,
		contractInstance: contractInstance,
	}
}

func (mt *MiningTasker) GetWork() (*Work, bool) {
	dispKey := mt.pubKey + "-" + db.DisputeStatusKey
	keys := []string{
		db.DifficultyKey,
		db.CurrentChallengeKey,
		db.RequestIdKey,
		db.RequestIdKey0,
		db.RequestIdKey1,
		db.RequestIdKey2,
		db.RequestIdKey3,
		db.RequestIdKey4,
		db.LastNewValueKey,
		dispKey,
	}

	m, err := mt.proxy.BatchGet(keys)
	if err != nil {
		level.Error(mt.logger).Log("msg", "get data from data proxy, cannot continue")
		return nil, false
	}

	if mt.checkDispute(m[dispKey]) == statusWaitNext {
		level.Info(mt.logger).Log("msg", "no dispute results from data server, waiting for next cycle", "key", dispKey)
		return nil, false
	}
	diff, stat := mt.getInt(m[db.DifficultyKey])
	if stat == statusWaitNext || stat == statusFailure {
		return nil, false
	}
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
		Challenge:  m[db.CurrentChallengeKey],
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
	return &Work{Challenge: newChallenge, PublicAddr: mt.pubKey, Start: uint64(rand.Int63()), N: math.MaxInt64}, instantSubmit
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
