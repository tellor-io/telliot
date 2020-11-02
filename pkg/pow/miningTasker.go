// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package pow

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/util"
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
	log           *util.Logger
	proxy         db.DataServerProxy
	pubKey        string
	currChallenge *MiningChallenge
}

func CreateTasker(cfg *config.Config, proxy db.DataServerProxy) *MiningTasker {

	return &MiningTasker{
		proxy:  proxy,
		pubKey: "0x" + cfg.PublicAddress,
		log:    util.NewLogger("pow", "MiningTasker"),
	}
}

func (mt *MiningTasker) GetWork(chan *Work) (*Work, bool) {
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
		db.LastSubmissionKey,
	}

	m, err := mt.proxy.BatchGet(keys)
	if err != nil {
		mt.log.Error("Could not get data from data proxy, cannot continue at all")
		log.Fatal(err)
	}

	mt.log.Debug("Received data: %v", m)

	if mt.checkDispute(m[dispKey]) == statusWaitNext {
		return nil, false
	}
	diff, stat := mt.getInt(m[db.DifficultyKey])
	if stat == statusWaitNext || stat == statusFailure {
		return nil, false
	}
	var reqIDs [5]*big.Int

	l, _ := mt.getInt(m[db.LastNewValueKey])
	instantSubmit := false

	today := time.Now()
	tm := time.Unix(l.Int64(), 0)
	mt.log.Debug("this long since last value:%v ", today.Sub(tm))
	if today.Sub(tm) >= time.Duration(15)*time.Minute {
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
			mt.log.Info("Could not retrieve pricing data for current request id: %v", err)
			return nil, false
		}
		val := m2[valKey]
		if len(val) == 0 {
			mt.log.Info("Pricing data not available for request %d", reqIDs[i].Uint64())
			return nil, false
		}
	}

	newChallenge := &MiningChallenge{
		Challenge:  m[db.CurrentChallengeKey],
		Difficulty: diff,
		RequestIDs: reqIDs,
	}

	// If this chalange is already sent out, don't do it again.
	if mt.currChallenge != nil && !instantSubmit && bytes.Equal(newChallenge.Challenge, mt.currChallenge.Challenge) {
		return nil, false
	}
	mt.currChallenge = newChallenge
	return &Work{Challenge: newChallenge, PublicAddr: mt.pubKey[2:], Start: uint64(rand.Int63()), N: math.MaxInt64}, instantSubmit
}

func (mt *MiningTasker) checkDispute(disp []byte) int {
	disputed, stat := mt.getInt(disp)
	if stat == statusWaitNext || stat == statusFailure {
		if stat == statusWaitNext {
			mt.log.Info("no dispute results from data server, waiting for next cycle")
		}
		return stat
	}

	if disputed.Cmp(big.NewInt(1)) != 0 {
		mt.log.Error("miner is in dispute, cannot continue")
		log.Fatal("miner in dispute")
		return statusFailure // Never gets here but just for completeness.
	}
	mt.log.Debug("miner is not in dispute, continuing")
	return statusSuccess
}

func (mt *MiningTasker) getInt(data []byte) (*big.Int, int) {
	if len(data) == 0 {
		return nil, statusWaitNext
	}

	val, err := hexutil.DecodeBig(string(data))
	if err != nil {
		mt.log.Error("decoding int: %v", err)
		return nil, statusFailure
	}
	return val, statusSuccess
}
