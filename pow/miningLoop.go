package pow

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/util"
)

/**
 * This module is responsible for the primary execution of the mining loop. Its
 * primary purpose is to find a solution to a PoW challenge and pushes solutions
 * through a channel. It has three source channels: one for exit signals, one for
 * mining interruptions/cancel, and one for challenge for PoW nonces. Its output
 * channel is for nonce solutions. Something downstream handles post-nonce verification
 * and submission on-chain
 */

// interface for all mining implementations
type Hasher interface {
	//base is a 56 byte slice containing the challenge and public address
	// the guessed nonce is appended to this slice and used as input to the first hash fn
	// returns a valid nonce, or empty string if none was found
	CheckRange(base []byte, difficulty *big.Int,  start uint64, n uint64) (string, error)

	//number of hashes this backend checks at a time
	StepSize() uint64
}


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
	start			 uint64
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
	backend			 Hasher
}


func createMiningLoop(id int, impl Hasher) (*miningLoop, error) {
	rand.Seed(time.Now().UnixNano())
	//this guarantees we won't overflow this variable later
	//since Int63() always returns a number between [0,2^63-1]
	//uint64 has a range [0,2^64-1]
	//the mining loop can therefore run for 2^63 iterations before overflowing
	start := uint64(rand.Int63())
	return createMiningLoopWithStart(id, start, impl)
}

func createMiningLoopWithStart(id int, start uint64, impl Hasher) (*miningLoop, error) {
	return &miningLoop{
		id:         id,
		start:		start,
		log:        util.NewLogger("pow", "MiningWorker-"+strconv.Itoa(id)),
		exitCh:     make(exitChannel),
		taskCh:     make(taskChannel),
		cancelCh:   make(cancelChannel),
		solutionCh: make(solutionChannel),
		backend:	impl,
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

	index := ml.start

	startTime := time.Now()

	// Variables for measuring hashrate
	lastIndex := index
	sampleTime := startTime
	nextHeartbeat := index + uint64(cfg.Heartbeat)

	//create the hash prefix
	_string := fmt.Sprintf("%x", challenge) + cfg.PublicAddress
	hashPrefix := decodeHex(_string)

	//check this many hashes at a time
	stepSize := ml.backend.StepSize()
	for {

		if index >= nextHeartbeat {
			now := time.Now()
			timeDelta := now.Sub(sampleTime)
			sampleTime = now
			iDelta := index - lastIndex
			lastIndex = index
			nextHeartbeat = index + uint64(cfg.Heartbeat)
			ml.log.Info("Hashrate: %v  per second", fmt.Sprintf("%.2f",float64(iDelta) / timeDelta.Seconds()))
		}
		if !ml.canMine {
			ml.log.Info("Stopping computation loop since asked to stop mining")
			return
		}

		sol, err := ml.backend.CheckRange(hashPrefix, _difficulty, index, stepSize)
		if err != nil {
			ml.log.Error("Mining backend failed: %s", err.Error())
			return
		}
		index += stepSize
		if sol != "" {
			diff := time.Now().Sub(startTime)
			ml.log.Info("Solution Found: %s in %f secs", sol, diff.Seconds())
			sol := &miningSolution{
				challenge: ml.currentChallenge,
				nonce:     sol,
			}
			ml.currentChallenge = nil
			ml.solutionCh <- sol
			ml.log.Info("Solution sent to output channel")
			return
		}
	}
}

