package pow

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"math/big"
	"math/rand"
	"os"
	"io/ioutil"
	"encoding/json"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/util"
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

func (mt *MiningTasker) GetWork() *Work {
	mt.log.Info("Pulling current data from data server...")
	dispKey := mt.pubKey + "-" + db.DisputeStatusKey
	keys := []string{
		db.DifficultyKey,
		db.CurrentChallengeKey,
		db.RequestIdKey,
		dispKey,
	}

	m, err := mt.proxy.BatchGet(keys)
	if err != nil {
		mt.log.Error("Could not get data from data proxy, cannot continue at all")
		log.Fatal(err)
	}

	mt.log.Debug("Received data: %v", m)

	if stat := mt.checkDispute(m[dispKey]); stat == statusWaitNext {
		return nil
	}

	diff, stat := mt.getInt(m[db.DifficultyKey])
	if stat == statusWaitNext || stat == statusFailure {
		return nil
	}

	reqID, stat := mt.getInt(m[db.RequestIdKey])
	if stat == statusWaitNext || stat == statusFailure {
		return nil
	}

	if reqID.Uint64() == 0 {
		mt.log.Info("Request ID is zero")
		return nil
	}

	valKey := fmt.Sprintf("%s%d", db.QueriedValuePrefix, reqID.Uint64())
	m2, err := mt.proxy.BatchGet([]string{valKey})
	if err != nil {
		mt.log.Info("Could not retrieve pricing data for current request id: %v", err)
		return nil
	}
	val := m2[valKey]
	if val == nil || len(val) == 0 {
			jsonFile, err := os.Open("manualData.json")
			if err != nil {
				fmt.Println(err)
			}
			defer jsonFile.Close()
			byteValue, _ := ioutil.ReadAll(jsonFile)
			var result map[string]map[string]uint
			json.Unmarshal([]byte(byteValue), &result)
			_id := strconv.FormatUint(reqID.Uint64(), 10)
			val := result[_id]["VALUE"]
		if val == 0{
			mt.log.Info("Pricing data not available for request %d", reqID.Uint64())
			return nil
		}else{
			fmt.Println("Using Manually entered value: ",val)
		}
	}

	newChallenge := &MiningChallenge{
		Challenge:  m[db.CurrentChallengeKey],
		Difficulty: diff,
		RequestID:  reqID,
	}

	//if we already sent this challenge out, don't do it again
	if mt.currChallenge != nil {
		if bytes.Compare(newChallenge.Challenge, mt.currChallenge.Challenge) == 0 {
			return nil
		}
	}
	mt.currChallenge = newChallenge
	return &Work{Challenge: newChallenge, PublicAddr: mt.pubKey[2:], Start: uint64(rand.Int63()), N: math.MaxInt64}
}

func (mt *MiningTasker) checkDispute(disp []byte) int {
	disputed, stat := mt.getInt(disp)
	if stat == statusWaitNext || stat == statusFailure {
		if stat == statusWaitNext {
			mt.log.Info("No dispute results from data server, waiting for next cycle")
		}
		return stat
	}

	if disputed.Cmp(big.NewInt(1)) != 0 {
		mt.log.Error("Miner is in dispute, cannot continue")
		log.Fatal("Miner in dispute")
		return statusFailure //never gets here but just for completeness
	}
	mt.log.Info("Miner is not in dispute, continuing")
	return statusSuccess
}

func (mt *MiningTasker) isEmptyChallenge(challenge *MiningChallenge) bool {
	mt.log.Info("Checking whether current challenge is empty")
	if challenge.RequestID.Cmp(big.NewInt(0)) == 0 {
		mt.log.Info("Current challenge has 0-value request ID, Cancelling any ongoing mining since previous challenge is complete")
		return true
	}
	if challenge.Challenge == nil || len(challenge.Challenge) == 0 {
		mt.log.Info("Current challenge has empty nonce. Cancelling any ongoing mining since previous challenge is complete")
		return true
	}

	mt.log.Info("Current challenge looks good")
	return false
}

func (mt *MiningTasker) getInt(data []byte) (*big.Int, int) {
	if data == nil || len(data) == 0 {
		return nil, statusWaitNext
	}

	val, err := hexutil.DecodeBig(string(data))
	if err != nil {
		mt.log.Error("Problem decoding int: %v", err)
		return nil, statusFailure
	}
	return val, statusSuccess
}
