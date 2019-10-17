package pow

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/util"
	"golang.org/x/crypto/ripemd160"
)

/**
 * This module is responsible for the primary execution of the mining loop. Its
 * primary purpose is to find a solution to a PoW challenge and pushes solutions
 * through a channel. It has three source channels: one for exit signals, one for
 * mining interruptions/cancel, and one for challenge for PoW nonces. Its output
 * channel is for nonce solutions. Something downstream handles post-nonce verification
 * and submission on-chain
 */

//miningChallenge holds information about a PoW challenge
type miningChallenge struct {
	challenge  []byte
	difficulty *big.Int
	requestID  *big.Int
}

//miningSolution associates the nonce solution for its requestID and challenge
type miningSolution struct {
	challenge *miningChallenge
	nonce     string
}

//taskChannel is for sending new challenges to the miner
type taskChannel chan *miningChallenge

//cancelChannel is for interrupting and stopping the mining loop
type cancelChannel chan bool

//solution channel
type solutionChannel chan *miningSolution

type exitChannel chan os.Signal

//miningLoop represents the main mining loop to response to external channel signals
type miningLoop struct {
	id               int
	mining           bool
	canMine          bool
	exitCh           chan os.Signal
	taskCh           taskChannel
	cancelCh         cancelChannel
	solutionCh       solutionChannel
	currentChallenge *miningChallenge
	pendingChallenge *miningChallenge
	miningCycle      sync.WaitGroup
	log              *util.Logger
}

func createMiningLoop(id int) (*miningLoop, error) {
	return &miningLoop{
		id:         id,
		log:        util.NewLogger("pow", "MiningWorker-"+strconv.Itoa(id)),
		exitCh:     make(exitChannel),
		taskCh:     make(taskChannel),
		cancelCh:   make(cancelChannel),
		solutionCh: make(solutionChannel),
	}, nil
}

func (ml *miningLoop) Start(ctx context.Context) {
	ml.log.Info("Starting mining loop %d", ml.id)
	go func() {
		for {
			select {
			//listen for system exists
			case _ = <-ml.exitCh:
				{
					ml.log.Info("Shutting down mining loop %d on OS interrupt", ml.id)
					ml.stopMining()
					return
				}
				//wait for new tasks
			case chal := <-ml.taskCh:
				{
					ml.startMining(ctx, chal)
				}
				//wait for cancellations
			case _ = <-ml.cancelCh:
				{
					ml.stopMining()
				}
			}
		}
	}()
}

func (ml *miningLoop) stopMining() {
	ml.canMine = false
	if ml.mining {
		ml.log.Info("Stopping mining loop %d", ml.id)
		ml.miningCycle.Wait() //wait for current mining cycle to complete
	}
}

func (ml *miningLoop) startMining(ctx context.Context, chal *miningChallenge) {
	ml.log.Info("Receiving new challege request...")
	if ml.currentChallenge != nil {
		if bytes.Compare(ml.currentChallenge.challenge, chal.challenge) == 0 {
			ml.log.Info("Ignoring duplicate challenge request")
			//same challenge ignore
			return
		}
	}
	if ml.mining {
		ml.log.Info("Interrupting current mining since challenge changed")
		ml.stopMining()
	}
	ml.canMine = true
	ml.currentChallenge = chal
	ml.pendingChallenge = nil
	ml.miningCycle.Add(1)
	go ml.solveChallenge()
}

func (ml *miningLoop) solveChallenge() {
	defer func() {
		ml.mining = false
		ml.miningCycle.Done()
		ml.log.Info("Finished mining. Cleaning up this cycle.")
	}()

	challenge := ml.currentChallenge.challenge
	_difficulty := ml.currentChallenge.difficulty

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	if !ml.canMine {
		ml.log.Warn("Miner will not solve challenge since it is flagged to not mine right now")
		return
	}

	ml.mining = true

	ml.log.Info("Mining on challenge: %x", challenge)
	ml.log.Info("Solving for difficulty: %d", _difficulty)

	//Generaete random start for worker
	rand.Seed(time.Now().UnixNano())
	i := rand.Int()

	//i := 0
	startTime := time.Now()

	// Constructors for loop objects
	numHash := new(big.Int)
	x := new(big.Int)
	compareZero := big.NewInt(0)

	for {

		i++
		if i%100000000 == 0 {
			ml.log.Info("Still Mining")
		}
		if !ml.canMine {
			ml.log.Info("Stopping computation loop since asked to stop mining")
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
			ml.log.Error("Could not set string on numHash")
			return
		}

		x.Mod(numHash, _difficulty)

		if x.Cmp(compareZero) == 0 {
			diff := time.Now().Sub(startTime)
			ml.log.Info("Solution Found: %s in %f secs", nn, diff.Seconds())
			sol := &miningSolution{
				challenge: ml.currentChallenge,
				nonce:     nn,
			}
			ml.currentChallenge = nil
			ml.solutionCh <- sol
			ml.log.Info("Solution sent to output channel")
			return
		}
	}
}
