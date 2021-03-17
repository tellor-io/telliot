// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package mining

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/util"
)

const ComponentName = "miner"

type HashSettings struct {
	prefix     []byte
	difficulty *big.Int
}

// interface for all mining implementations.
type Hasher interface {
	//base is a 52 byte slice containing the challenge and public address
	// the guessed nonce is appended to this slice and used as input to the first hash fn
	// returns a valid nonce, or empty string if none was found
	CheckRange(anySolution context.Context, hash *HashSettings, start uint64, n uint64) (string, uint64, error)

	//number of hashes this backend checks at a time
	StepSize() uint64

	Name() string
}

type Backend struct {
	Hasher
	TotalHashes      uint64
	HashSincePrint   uint64
	HashRateEstimate float64
}

// MiningChallenge holds information about a PoW challenge.
type MiningChallenge struct {
	Challenge  []byte
	Difficulty *big.Int
	RequestIDs [5]*big.Int
}

func NewHashSettings(challenge *MiningChallenge, publicAddr string) *HashSettings {
	_string := fmt.Sprintf("%x", challenge.Challenge) + publicAddr[2:]
	hashPrefix := util.DecodeHex(_string)
	return &HashSettings{
		prefix:     hashPrefix,
		difficulty: challenge.Difficulty,
	}
}

// the mining group will attempt to size the chunk it gives each hasher so that it takes roughly this long to complete
// if you make it too low, overall mining efficiency will drop due to exessive overhead
// if you make it too high, the miner won't respond quickly to commands (stop, change challenge, etc)
// right now 200ms seems like a good default. This could perhaps be made configurable, but I don't see much benefit.
const targetChunkTime = 200 * time.Millisecond

const rateInitialGuess = 100e3

type MiningGroup struct {
	Backends         []*Backend
	LastPrinted      time.Time
	logger           log.Logger
	cfg              *config.Config
	contractInstance *contracts.ITellor
}

func NewMiningGroup(logger log.Logger, cfg *config.Config, hashers []Hasher, contractInstance *contracts.ITellor) (*MiningGroup, error) {
	filterLog, err := logging.ApplyFilter(*cfg, ComponentName, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	group := &MiningGroup{
		Backends:         make([]*Backend, len(hashers)),
		logger:           log.With(filterLog, "component", ComponentName),
		cfg:              cfg,
		contractInstance: contractInstance,
	}
	for i, hasher := range hashers {
		//start with a small estimate for hash rate, much faster to increase the gusses rather than decrease
		group.Backends[i] = &Backend{Hasher: hasher, HashRateEstimate: rateInitialGuess}
	}
	return group, nil
}

type backendResult struct {
	hash     *HashSettings
	nonce    string
	err      error
	started  time.Time
	finished time.Time
	n        uint64
	backend  *Backend
}

// do some work and write the result back to a channel.
func (b *Backend) doWork(anySolution context.Context, close context.CancelFunc, hash *HashSettings, start uint64, n uint64, resultCh chan *backendResult) {
	defer close()
	timeStarted := time.Now()
	sol, nchecked, err := b.CheckRange(anySolution, hash, start, n)
	if err != nil {
		resultCh <- &backendResult{err: err}
		return
	}
	resultCh <- &backendResult{
		hash:     hash,
		nonce:    sol,
		started:  timeStarted,
		finished: time.Now(),
		n:        nchecked,
		backend:  b,
	}
}

func formatHashRate(rate float64) string {
	letters := " KMGTQ"
	i := 0
	//purposely made this 10k instead of 1k. That way you won't get single digit rates
	//this function could instead just always return 3 significant digits, but then it gets hard to visually
	//see the differences between different devices when formatted into columns
	for rate > 10000 {
		rate /= 1000
		i++
	}
	return fmt.Sprintf("%.0f %cH/s", rate, letters[i])
}

func (g *MiningGroup) HashRateEstimate() float64 {
	totalHashrate := 0.0
	for _, b := range g.Backends {
		totalHashrate += b.HashRateEstimate
	}
	return totalHashrate
}

func (g *MiningGroup) PreferredWorkMultiple() uint64 {
	largest := uint64(0)
	for _, b := range g.Backends {
		if b.StepSize() > largest {
			largest = b.StepSize()
		}
	}
	return largest
}

func (g *MiningGroup) PrintHashRateSummary() {
	totalHashes := uint64(0)
	for _, b := range g.Backends {
		totalHashes += b.HashSincePrint
	}
	now := time.Now()
	delta := now.Sub(g.LastPrinted).Seconds()
	totalHashrate := float64(totalHashes) / delta
	level.Info(g.logger).Log("msg", "check total hashrate", "totalHashrate", formatHashRate(totalHashrate))
	for _, b := range g.Backends {
		hashRate := float64(b.HashSincePrint) / delta
		level.Debug(g.logger).Log(
			"msg", "print hash values",
			"hashRate", fmt.Sprintf("%8s", formatHashRate(hashRate)),
			"avgHashRate", fmt.Sprintf("%4.1f%%", (hashRate/totalHashrate)*100),
			"name", b.Name(),
		)
		b.HashSincePrint = 0
	}
	g.LastPrinted = now
}

type Work struct {
	Challenge  *MiningChallenge
	PublicAddr string
	Start      uint64
	N          uint64
}

type Result struct {
	Work  *Work
	Nonce string
}

// dispatches a chunk and returns the number of hashes chosen.
func (b *Backend) dispatchWork(parentCtx context.Context, timeOfLastNewValue *big.Int, hash *HashSettings, start uint64, resultCh chan *backendResult) uint64 {
	target := b.HashRateEstimate * targetChunkTime.Seconds()
	step := b.StepSize()
	nsteps := uint64(math.Round(target / float64(step)))
	if nsteps == 0 {
		nsteps = 1
	}
	n := nsteps * step
	tm := time.Unix(timeOfLastNewValue.Int64(), 0)
	anySolution, close := context.WithDeadline(parentCtx, tm.Add(15*time.Minute))
	go b.doWork(anySolution, close, hash, start, n, resultCh)
	return n
}

func (g *MiningGroup) getTimeOfLastNewValue() *big.Int {
	var err error
	var timeOfLastNewValue *big.Int
	for {
		// Checks the last submit value in the oracle and set a timeout of 15min - (now-lastSubmit).
		// This is because 15min after the last submit any solution will work.
		timeOfLastNewValue, err = g.contractInstance.GetUintVar(nil, rpc.Keccak256([]byte("_TIME_OF_LAST_NEW_VALUE")))
		if err == nil {
			break
		}
		level.Error(g.logger).Log("msg", "getting time of last new value from oracle, retrying in a second", "err", err)
		time.Sleep(1 * time.Second)
	}

	return timeOfLastNewValue
}

func (g *MiningGroup) Mine(ctx context.Context, input chan *Work, output chan *Result) {
	sent := uint64(0)
	recv := uint64(0)
	timeStarted := time.Now()
	g.LastPrinted = timeStarted

	resultChannel := make(chan *backendResult, len(g.Backends)*2)

	// queue of miners waiting for work.
	idleWorkers := make(chan *Backend, len(g.Backends))

	// add all available miners to the idleWorkers queue.
	for _, b := range g.Backends {
		//dispatch work
		idleWorkers <- b
	}

	nextHeartbeat := g.cfg.Mine.Heartbeat.Duration

	var currHashSettings *HashSettings
	var currWork *Work

	// Mine until context is done.
	// Each time a hasher finishes a chunk, give it a new one to work on.
	// Always waits for all miners to finish their chunks before returning.
	// EXCEPT in the case of an error, but then the app is almost certainly just quitting anyways!

	for {
		elapsed := time.Since(timeStarted)
		if elapsed > nextHeartbeat {
			g.PrintHashRateSummary()
			nextHeartbeat = elapsed + g.cfg.Mine.Heartbeat.Duration
		}
		select {
		case <-ctx.Done():
			return
		// Read in a new work block.
		case work := <-input:
			sent = 0
			recv = 0
			currWork = work
			currHashSettings = NewHashSettings(work.Challenge, work.PublicAddr)

		// Read in a result from one of the miners.
		case result := <-resultChannel:
			if result.err != nil {
				level.Error(g.logger).Log("msg", "hasher failed", "err", result.err)
				return
			}
			idleWorkers <- result.backend

			// Update the backend statistics no matter what.
			result.backend.TotalHashes += result.n
			result.backend.HashSincePrint += result.n

			// Only update the hashRateEstimate if we didn't find a solution - otherwise the rate could be wrong
			// due to returning early.
			if result.nonce == "" {
				newEst := float64(result.n) / (result.finished.Sub(result.started).Seconds())
				if result.backend.HashRateEstimate == rateInitialGuess {
					result.backend.HashRateEstimate = newEst
				} else {
					memory := 0.2
					result.backend.HashRateEstimate *= 1 - memory
					result.backend.HashRateEstimate += memory * newEst
				}
			}

			// Ignore out of date results.
			if result.hash != currHashSettings {
				break
			}

			// Did it finish the job?
			recv += result.n
			if result.nonce != "" || recv >= currWork.N {
				level.Info(g.logger).Log("msg", "found solution and sending the result",
					"challenge", fmt.Sprintf("%x", currWork.Challenge.Challenge),
					"solution", result.nonce,
					"difficulty", currWork.Challenge.Difficulty,
					"requestIDs", fmt.Sprintf("%+v", currWork.Challenge.RequestIDs),
				)

				output <- &Result{Work: currWork, Nonce: result.nonce}
				currWork = nil
				currHashSettings = nil
			}
		}
		if currWork != nil {
			for sent < currWork.N && len(idleWorkers) > 0 {
				worker := <-idleWorkers
				timeOfLastNewValue := g.getTimeOfLastNewValue()
				sent += worker.dispatchWork(ctx, timeOfLastNewValue, currHashSettings, currWork.Start+sent, resultChannel)
			}
		}
	}
}
