package pow

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"sync"
	"time"
	"math/rand"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/contracts"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/util"
	"golang.org/x/crypto/ripemd160"
)

type miningRequest struct {
	challenge  []byte
	difficulty *big.Int
	requestID  *big.Int
	nonce      string
	value      *big.Int
}

//Worker state for mining operation
type Worker struct {
	id             int
	mining         bool
	canMine        bool
	checkInterval  time.Duration
	miningWait     sync.WaitGroup
	exitCh         chan os.Signal
	log            *util.Logger
	submitter      tellorCommon.TransactionSubmitter
	currentRequest *miningRequest
	pendingRequest *miningRequest
}

//CreateWorker creates a new worker instance
func CreateWorker(exitCh chan os.Signal, id int, submitter tellorCommon.TransactionSubmitter, checkIntervalSeconds time.Duration) *Worker {
	if checkIntervalSeconds == 0 {
		checkIntervalSeconds = 15
	}
	return &Worker{canMine: true, id: id, mining: false, submitter: submitter, checkInterval: checkIntervalSeconds, exitCh: exitCh, log: util.NewLogger("pow", "MiningWorker-"+strconv.Itoa(id))}
}

//Start kicks of go routines to check for challenge changes, mine, and submit solutions
func (w *Worker) Start(ctx context.Context) {
	w.log.Info("Starting mining worker", w.id)
	ticker := time.NewTicker(w.checkInterval * time.Second)
	//run immediately
	w.checkForChallengeChanges(ctx)
	go func() {
		for {
			select {
			case _ = <-w.exitCh:
				{
					w.log.Info("Shutting down miner", w.id)
					w.stopMining()
					ticker.Stop()
					return
				}

			case _ = <-ticker.C:
				{
					w.checkForChallengeChanges(ctx)
				}
			}
		}
	}()
}

//CanMine checks whether this mining worker is allowed to mine right now
func (w *Worker) CanMine() bool {
	return w.canMine
}

func (w *Worker) checkForChallengeChanges(ctx context.Context) {
	//check whether we should start mining a new request
	w.log.Info("Checking for mining challenge changes")
	mined, err := w.alreadyMined(ctx)
	if err != nil {
		w.log.Error("Problem reading if already mined: %v\n", err)
	} else if !mined {
		changed, err := w.challengeChanged(ctx)

		if err != nil {
			w.log.Error("Problem determining if challenge changed: %v\n", err)
		} else if changed {
			if w.mining {
				w.log.Info("New challenge found for mining. Interrupting current mining cycle and working on current challenge instead")
				w.stopMining()
			} else {
				w.log.Info("Found new challenge to mine")
			}

			//reset mining flag since it was used to stop the computation loop
			w.canMine = true
			w.currentRequest = w.pendingRequest
			w.pendingRequest = nil
			w.miningWait.Add(1)
			go w.solveChallenge(ctx)
		} else {
			w.log.Info("Challenge has not changed. Will wait for next cycle to mine again")
		}
	} else if mined {
		if w.mining {
			w.log.Warn("Another thread mined solution, stopping miner and will wait for next challenge")
			w.stopMining()
		} else {
			w.log.Info("Miner already mined the current challenge. Will wait for next cycle to mine again")
		}
	}
}

func (w *Worker) stopMining() {
	w.canMine = false
	if w.mining {
		w.log.Info("Stopping mining cycle")
		w.miningWait.Wait() //wait for current mining cycle to complete
	}
}

func (w *Worker) alreadyMined(ctx context.Context) (bool, error) {
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)
	//if there is a pending challenge and it matches the current one, we already worked it.
	currentChallenge, err := DB.Get(db.CurrentChallengeKey)
	if err != nil {
		w.log.Error("Problem reading current challenge from DB", err)
		return false, err
	}

	requestID, err := DB.Get(db.RequestIdKey)
	if err != nil {
		w.log.Error("Problem reading request id from DB: %v\n", err)
		return false, err
	}
	if len(requestID) == 0 {
		w.log.Debug("No requestID stored, nothing to mine")
		return true, nil
	}

	asInt, err := hexutil.DecodeBig(string(requestID))
	if err != nil {
		w.log.Error("Problem decoding request id as big int: %v\n", err)
		return false, err
	}

	if asInt.Cmp(big.NewInt(0)) == 0 {
		w.log.Info("No current challenge, marking as already mined")
		return true, nil
	}

	pendingChallenge, err := DB.Get(db.PendingChallengeKey)
	if err != nil {
		w.log.Error("Problem reading pending challenge from DB", err)
		return false, err
	}
	if bytes.Compare(pendingChallenge, currentChallenge) == 0 {
		//we've already done it
		w.log.Info("Miner has pending txn for current challenge, already mined")
		return true, nil
	} else {
		//if it's not the same, we don't care if it's pending anymore...we just care
		//that we new have work to do
		DB.Delete(db.PendingChallengeKey)
	}

	cfg, err := config.GetConfig()
	if err != nil {
		w.log.Error("Could not get configuration while checking mining status", err)
		return false, err
	}

	//get address from config
	_fromAddress := cfg.PublicAddress

	//convert to address
	fromAddress := common.HexToAddress(_fromAddress)

	instance := ctx.Value(tellorCommon.MasterContractContextKey).(*contracts.TellorMaster)
	var asBytes32 [32]byte
	copy(asBytes32[:], currentChallenge)
	didIMine, err := instance.DidMine(nil, asBytes32, fromAddress)
	if err != nil {
		w.log.Error("Problem reading whether this miner mined from on-chain: %v\n", err)
		return false, err
	}
	if didIMine {
		w.log.Info("Already mined according to on-chain contract status")
		return true, nil
	}
	return false, nil
}

func (w *Worker) challengeChanged(ctx context.Context) (bool, error) {
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)
	currentChallenge, err := DB.Get(db.CurrentChallengeKey)
	if err != nil {
		w.log.Error("Problem reading challenge in miner run loop: %v\n", err)
		return false, err
	}
	if len(currentChallenge) == 0 {
		w.log.Debug("No current challenge stored, nothing to mine")
		return false, nil
	}

	diff, err := DB.Get(db.DifficultyKey)
	if err != nil {
		w.log.Error("Problem reading difficult from DB: %v\n", err)
		return false, err
	}
	if len(diff) == 0 {
		w.log.Debug("No difficulty stored, nothing to mine")
		return false, nil
	}

	difficulty, err := hexutil.DecodeBig(string(diff))
	if err != nil {
		w.log.Error("Problem decoding difficulty: %v\n", err)
		return false, err
	}

	requestID, err := DB.Get(db.RequestIdKey)
	if err != nil {
		w.log.Error("Problem reading request id from DB: %v\n", err)
		return false, err
	}
	if len(requestID) == 0 {
		w.log.Debug("No requestID stored, nothing to mine")
		return false, nil
	}

	asInt, err := hexutil.DecodeBig(string(requestID))
	if err != nil {
		w.log.Error("Problem decoding request id as big int: %v\n", err)
		return false, err
	}

	if asInt.Cmp(big.NewInt(0)) == 0 {
		w.log.Info("No current challenge, will not mine this cycle")
		return false, nil
	}

	value, err := w.readAndDecodeLatestValue(ctx, asInt)
	if err != nil {
		w.log.Error("Problem reading latest price value for request: %v. %v\n", requestID, err)
		return false, err
	}
	if value == nil {
		w.log.Warn("Tracker has not fetched price data for request %v yet", asInt)
		return false, nil
	}

	if w.currentRequest == nil || bytes.Compare(w.currentRequest.challenge, currentChallenge) != 0 {
		w.pendingRequest = &miningRequest{challenge: currentChallenge, difficulty: difficulty, requestID: asInt, nonce: "", value: value}
		return true, nil
	}
	return false, nil
}

//SolveChallenge performs PoW
func (w *Worker) solveChallenge(ctx context.Context) {
	defer func() {
		w.mining = false
		w.miningWait.Done()
		w.log.Info("Finished mining. Cleaning up this cycle.")
	}()

	challenge := w.currentRequest.challenge
	_difficulty := w.currentRequest.difficulty
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	if !w.canMine {
		w.log.Warn("Miner will not solve challenge since it is flagged to not mine right now")
		w.miningWait.Done()
		return
	}
	w.mining = true

	w.log.Info("Mining on challenge: %x", challenge)
	w.log.Info("Solving for difficulty: %d", _difficulty)


	//Generaete random start for worker
	rand.Seed(time.Now().UnixNano())
	i := rand.Int()

	//i := 0
	startTime := time.Now()
	
	// Constructors for loop objects
	numHash := new(big.Int)
	x := new(big.Int)
	compare_zero := big.NewInt(0)
    
	for {

		i++
		if i%100000000 == 0 {
			w.log.Info("Still Mining")
		}
		if !w.canMine {
			w.log.Info("Stopping computation loop")
			return
		}

		//nn := randInt() //do we need to use big number?
		nn := strconv.Itoa(i)
		
		nonce := fmt.Sprintf("%x", nn)
		
		_string := fmt.Sprintf("%x", challenge) + cfg.PublicAddress + nonce
		
		hash := solsha3.SoliditySHA3(
			solsha3.Bytes32(decodeHex(_string)),
		)

		hasher := ripemd160.New()
		//Consider moving hasher constructor outside loop and replacing with hasher.Reset()
		
		hasher.Write([]byte(hash))
		hash1 := hasher.Sum(nil)
		n := sha256.Sum256(hash1)
		q := fmt.Sprintf("%x", n)
		
		
		numHash, ok := numHash.SetString(q, 16)
		if !ok {
			w.log.Error("!!!!!SetString: error")
			return
		}

		x.Mod(numHash, _difficulty)
		
		if x.Cmp(compare_zero) == 0 {
			diff := time.Now().Sub(startTime)
			w.log.Info("Solution Found: %s in %f secs", nn, diff.Seconds())
			w.currentRequest.nonce = nn
			w.submitSolution(ctx)
			return
		}
	}
}

func (w *Worker) readAndDecodeLatestValue(ctx context.Context, requestID *big.Int) (*big.Int, error) {
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)
	val, err := DB.Get(fmt.Sprintf("%s%d", db.QueriedValuePrefix, requestID.Uint64()))
	if requestID.Uint64() > 5 && requestID.Uint64() < 51 {
		val, err = DB.Get(fmt.Sprintf("%s%d", db.QueriedValuePrefix, 1))
	}
	if err != nil {
		w.log.Error("Problem reading price data from DB: %v\n", err)
		return nil, err
	}
	if len(val) == 0 {
		w.log.Warn("Have not retrieved requestID value. Will wait for next tracker cycle to complete before mining")
		return nil, nil
	}
	value, err := hexutil.DecodeBig(string(val))
	if err != nil {
		w.log.Error("Problem decoding price value: %v\n", err)
		return nil, err
	}
	return value, nil
}

func (w *Worker) txnGenerator(ctx context.Context, contract tellorCommon.ContractInterface) (*types.Transaction, error) {
	if w.currentRequest == nil || w.currentRequest.nonce == "" {
		w.log.Warn("Somehow attempting to submit incomplete solution")
		return nil, nil
	}

	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)
	pending, err := DB.Get(db.PendingChallengeKey)
	if err != nil {
		w.log.Error("Problem reading pending challenge from DB: %v", err)
		return nil, err
	}

	if bytes.Compare(pending, w.currentRequest.challenge) == 0 {
		//we've already submitted a solution!
		w.log.Error("Already submitted solution for challenge. Pending txn should have been checked already")
		return nil, nil
	}

	cChallenge := [32]byte{}
	copy(cChallenge[:], w.currentRequest.challenge)
	mined, err := contract.DidMine(cChallenge)
	if err != nil {
		w.log.Info("Problem reading whether this miner has already mined challenge: %v", err)
		return nil, err
	}
	if mined {
		w.log.Info("Miner has already mined current challenge, will not submit txn again")
		return nil, nil
	}

	txn, err := contract.SubmitSolution(w.currentRequest.nonce, w.currentRequest.requestID, w.currentRequest.value)
	e2 := DB.Put(db.PendingChallengeKey, w.currentRequest.challenge)
	if e2 != nil {
		w.log.Error("Problem storing pending challenge: %v", e2)
	}
	return txn, err
}

func (w *Worker) submitSolution(ctx context.Context) {
	value, err := w.readAndDecodeLatestValue(ctx, w.currentRequest.requestID)
	if err != nil {
		w.log.Info("Could not read price data for current request %v, using last known value: %v", w.currentRequest.requestID, w.currentRequest.value)
		value = w.currentRequest.value
	}
	if value == nil {
		w.log.Info("Latest prices data is missing from DB, using last known price for request %v: %v", w.currentRequest.requestID, w.currentRequest.value)
		value = w.currentRequest.value
	}
	w.currentRequest.value = value
	err = w.submitter.PrepareTransaction(ctx, "SubmitSolution", w.txnGenerator)
	if err != nil {
		w.log.Error("Problem submitting solution to network: %v", err)
	}

}
