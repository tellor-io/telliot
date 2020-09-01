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
	"time"

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

func (mt *MiningTasker) GetWork(input chan *Work) (*Work,bool) {
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

	if stat := mt.checkDispute(m[dispKey]); stat == statusWaitNext {
		return nil,false
	}
	diff, stat := mt.getInt(m[db.DifficultyKey])
	if stat == statusWaitNext || stat == statusFailure {
		return nil,false
	}
	var reqIDs [5] *big.Int

	l, stat := mt.getInt(m[db.LastNewValueKey]) 
	instantSubmit := false
	looper := 1
	if l != nil{
		looper = 5
		today := time.Now() 
		tm := time.Unix(l.Int64(), 0)
		fmt.Println("This long since last value:  ",today.Sub(tm) )
		if today.Sub(tm) >= time.Duration(15) * time.Minute {
			instantSubmit = true
		}
		r, stat := mt.getInt(m[db.RequestIdKey0])
		if stat == statusWaitNext || stat == statusFailure {
			return nil,false
		}
		reqIDs[0] = r
		r, stat = mt.getInt(m[db.RequestIdKey1])
		if stat == statusWaitNext || stat == statusFailure {
			return nil,false
		}	
		reqIDs[1] = r
		r, stat = mt.getInt(m[db.RequestIdKey2])
		if stat == statusWaitNext || stat == statusFailure {
			return nil,false
		}
		reqIDs[2] = r
		r, stat = mt.getInt(m[db.RequestIdKey3])
		if stat == statusWaitNext || stat == statusFailure {
			return nil,false
		}
		reqIDs[3] = r
		r, stat = mt.getInt(m[db.RequestIdKey4])
		if stat == statusWaitNext || stat == statusFailure {
			return nil,false
		}
		reqIDs[4] = r
	}else {
		r, stat := mt.getInt(m[db.RequestIdKey])
		if stat == statusWaitNext || stat == statusFailure {
			return nil,false
		}
		reqIDs[0] = r
	
		if reqIDs[0].Uint64() == 0 {
			mt.log.Info("Request ID is zero")
			return nil,false
		}
	}
	for i := 0;i<looper;i++ {
		valKey := fmt.Sprintf("%s%d", db.QueriedValuePrefix, reqIDs[i].Uint64())
		m2, err := mt.proxy.BatchGet([]string{valKey})
		if err != nil {
			mt.log.Info("Could not retrieve pricing data for current request id: %v", err)
			return nil,false
		}
		val := m2[valKey]
		if val == nil || len(val) == 0 {
				jsonFile, err := os.Open("manualData.json")
				if err != nil {
					fmt.Println("manualData read error",err)
					return nil,false
				}
				defer jsonFile.Close()
				byteValue, _ := ioutil.ReadAll(jsonFile)
				var result map[string]map[string]uint
				json.Unmarshal([]byte(byteValue), &result)
				_id := strconv.FormatUint(reqIDs[i].Uint64(), 10)
				val := result[_id]["VALUE"]
			if val == 0{
				mt.log.Info("Pricing data not available for request %d", reqIDs[i].Uint64())
				return nil,false
			}else{
				fmt.Println("Using Manually entered value: ",val)
			}
		}
	}


	newChallenge := &MiningChallenge{
		Challenge:  m[db.CurrentChallengeKey],
		Difficulty: diff,
		RequestID:  reqIDs[0],
		RequestIDs: reqIDs,
	}

	//if we already sent this challenge out, don't do it again
	if mt.currChallenge != nil {
		if bytes.Compare(newChallenge.Challenge, mt.currChallenge.Challenge) == 0 {
			return nil,false
		}
	}
	mt.currChallenge = newChallenge
	return &Work{Challenge: newChallenge, PublicAddr: mt.pubKey[2:], Start: uint64(rand.Int63()), N: math.MaxInt64},instantSubmit
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
	if challenge.RequestID.Cmp(big.NewInt(0)) == 0 && challenge.RequestIDs[0].Cmp(big.NewInt(0)) == 0 {
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
