package pow

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/common"
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

type miningTasker struct {
	exitCh           exitChannel
	running          bool
	id               int
	taskCh           taskChannel
	cancelCh         cancelChannel
	log              *util.Logger
	checkInterval    time.Duration
	proxy            db.DataServerProxy
	currentChallenge *miningChallenge
	pubKey           string
}

func createTasker(id int,
	taskCh taskChannel,
	cancelCh cancelChannel,
	checkInterval time.Duration,
	proxy db.DataServerProxy) (*miningTasker, error) {

	if checkInterval == 0 {
		checkInterval = 5
	}

	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Printf("Problem getting config, can't proceed at all: %v\n", err)
		log.Fatal(err)
	}

	//get address from config
	_fromAddress := cfg.PublicAddress

	//convert to address
	fromAddress := common.HexToAddress(_fromAddress)
	pubKey := strings.ToLower(fromAddress.Hex())

	return &miningTasker{
		id:            id,
		taskCh:        taskCh,
		cancelCh:      cancelCh,
		exitCh:        make(exitChannel),
		checkInterval: checkInterval,
		proxy:         proxy,
		pubKey:        pubKey,
		log:           util.NewLogger("pow", "MiningTasker-"+strconv.Itoa(id)),
	}, nil
}

func (mt *miningTasker) Start(ctx context.Context) {
	mt.log.Info("Starting mining tasker")
	mt.running = true
	ticker := time.NewTicker(mt.checkInterval * time.Second)

	//then start run loop
	go func() {
		//pull latest updates
		mt.pullUpdates()
		mt.log.Info("Starting mining tasker loop")
		for {
			select {
			case _ = <-mt.exitCh:
				{
					mt.log.Info("Stopping mining tasker on OS interrupt")
					ticker.Stop()
					mt.running = false
					return
				}
			case _ = <-ticker.C:
				{
					mt.pullUpdates()
				}
			}
		}
	}()
}

func (mt *miningTasker) pullUpdates() {
	mt.log.Info("Pulling current data from data server...")
	dispKey := mt.pubKey + "-" + db.DisputeStatusKey
	pendingKey := mt.pubKey + "-" + db.PendingChallengeKey
	keys := []string{
		db.DifficultyKey,
		db.CurrentChallengeKey,
		db.RequestIdKey,
		dispKey,
		pendingKey,
	}

	m, err := mt.proxy.BatchGet(keys)
	if err != nil {
		mt.log.Error("Could not get data from data proxy, cannot continue at all")
		log.Fatal(err)
	}

	mt.log.Debug("Received data: %v", m)

	if stat := mt.checkDispute(m[dispKey]); stat == statusWaitNext {
		return
	}

	if mt.hasPendingTxn(m[pendingKey], m[db.CurrentChallengeKey]) {
		mt.log.Info("Already have a pending solution for the challenge, stopping any ongoing mining")
		mt.cancelCh <- true
		return
	}

	diff, stat := mt.getInt(m[db.DifficultyKey])
	if stat == statusWaitNext || stat == statusFailure {
		return
	}

	reqID, stat := mt.getInt(m[db.RequestIdKey])
	if stat == statusWaitNext || stat == statusFailure {
		return
	}

	valKey := fmt.Sprintf("%s%d", db.QueriedValuePrefix, reqID.Uint64())
	m2, err := mt.proxy.BatchGet([]string{valKey})
	if err != nil {
		mt.log.Info("Could not retrieve pricing data for current request id: %v", err)
		return
	}
	val := m2[valKey]
	if val == nil || len(val) == 0 {
		mt.log.Info("Pricing data not available for request %d, cannot mine yet", reqID.Uint64())
		return
	}

	newChallenge := &miningChallenge{
		challenge:  m[db.CurrentChallengeKey],
		difficulty: diff,
		requestID:  reqID,
	}

	if mt.isEmptyChallenge(newChallenge) {
		mt.log.Info("Current challenge is empty, cancelling any ongoing mining threads")
		mt.cancelCh <- true
		return
	}

	//means we have a valid challenge
	mt.log.Info("Issuing presumably good challenge")

	//remember it for next cycle
	mt.currentChallenge = newChallenge

	//and send to output
	mt.taskCh <- newChallenge

	mt.log.Info("Challenge submitted for mining")
}

func (mt *miningTasker) checkDispute(disp []byte) int {
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

func (mt *miningTasker) hasPendingTxn(pendingChallenge []byte, currentChallenge []byte) bool {
	mt.log.Info("Checking whether miner has a pending txn for current challenge...")

	//if the pending challenge is the same as the current challenge, it means we've
	//already sent a txn for the challenge, and thus we have a pending (or confirmed) txn
	//for the current challenge and we shouldn't do anything
	if pendingChallenge != nil && len(pendingChallenge) > 0 {
		if bytes.Compare(pendingChallenge, currentChallenge) == 0 {
			return true
		}
	}

	mt.log.Info("No pending challenge matches current challenge, continuing")
	//otherwise, we either have a new challenge that doesn't match our old one, no pending
	//txn, or an empty new challenge that indicates that we should stop anyway
	return false
}

func (mt *miningTasker) isEmptyChallenge(challenge *miningChallenge) bool {
	mt.log.Info("Checking whether current challenge is empty")
	if challenge.requestID.Cmp(big.NewInt(0)) == 0 {
		mt.log.Info("Current challenge has 0-value request ID, Cancelling any ongoing mining since previous challenge is complete")
		return true
	}
	if challenge.challenge == nil || len(challenge.challenge) == 0 {
		mt.log.Info("Current challenge has empty nonce. Cancelling any ongoing mining since previous challenge is complete")
		return true
	}

	mt.log.Info("Current challenge looks good")
	return false
}

func (mt *miningTasker) getInt(data []byte) (*big.Int, int) {
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
