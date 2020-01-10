package pow

import (
	"fmt"
	"github.com/tellor-io/TellorMiner/config"
	"log"
	"math"
	"math/big"
	"time"
)

type HashSettings struct {
	prefix []byte
	difficulty *big.Int
}

// interface for all mining implementations
type Hasher interface {
	//base is a 52 byte slice containing the challenge and public address
	// the guessed nonce is appended to this slice and used as input to the first hash fn
	// returns a valid nonce, or empty string if none was found
	CheckRange(hash *HashSettings,  start uint64, n uint64) (string, error)

	//number of hashes this backend checks at a time
	StepSize() uint64

	Name() string
}


type Backend struct {
	Hasher
	TotalHashes uint64
	HashRateEstimate float64
}

//miningChallenge holds information about a PoW challenge
type miningChallenge struct {
	challenge  []byte
	difficulty *big.Int
	requestID  *big.Int
}

func NewHashSettings(challenge *miningChallenge, publicAddr string) *HashSettings {
	_string := fmt.Sprintf("%x", challenge.challenge) + publicAddr
	hashPrefix := decodeHex(_string)
	return &HashSettings{
		prefix:     hashPrefix,
		difficulty: challenge.difficulty,
	}
}

// the mining group will attempt to size the chunk it gives each hasher so that it takes roughly this long to complete
// if you make it too low, overall mining efficiency will drop due to exessive overhead
// if you make it too high, the miner won't respond quickly to commands (stop, change challenge, etc)
// right now 50ms seems like a good default. This could perhaps be made configurable, but I don't see much benefit
const targetChunkTime = 50 * time.Millisecond

type MiningGroup struct {
	Backends []*Backend
}

func NewMiningGroup(hashers []Hasher) *MiningGroup {
	group := &MiningGroup{
		Backends: make([]*Backend, len(hashers), len(hashers)),
	}
	for i,hasher := range hashers {
		//start with a small estimate for hash rate, much faster to increase the gusses rather than decrease
		group.Backends[i] = &Backend{Hasher:hasher, HashRateEstimate:100e3}
	}

	return group
}


type result struct {
	nonce string
	err error
	started time.Time
	finished time.Time
	n uint64
	backend *Backend
}

// do some work and write the result back to a channel
func (b *Backend)doWork(hash *HashSettings, start uint64, n uint64, resultCh chan *result) {
	timeStarted := time.Now()
	sol, err := b.CheckRange(hash, start, n)
	if err != nil {
		resultCh <- &result{err: err}
		return
	}
	resultCh <- &result{
		nonce: sol,
		started: timeStarted,
		finished: time.Now(),
		n: n,
		backend: b,
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

func (g *MiningGroup)PrintHashRateSummary() {
	totalHashrate := 0.0
	for _,b := range g.Backends {
		totalHashrate += b.HashRateEstimate
	}
	fmt.Printf("Total hashrate %s\n", formatHashRate(totalHashrate))
	for _,b := range g.Backends {
		fmt.Printf("\t%8s (%4.1f%%): %s \n", formatHashRate(b.HashRateEstimate), (b.HashRateEstimate/totalHashrate)*100,b.Name())
	}
}

//dispatches a chunk and returns the number of hashes chosen
func (b *Backend)dispatchWork(hash *HashSettings, start uint64, resultCh chan *result) uint64 {
	target := b.HashRateEstimate * targetChunkTime.Seconds()
	step := b.StepSize()
	nsteps := uint64(math.Round(target /float64(step)))
	if nsteps == 0 {
		nsteps = 1
	}
	n := nsteps * step
	go b.doWork(hash, start, n, resultCh)
	return n
}


func (g *MiningGroup)Mine(hash *HashSettings, start uint64, n uint64, timeout time.Duration) (string, error) {
	sent := uint64(0)
	recv := uint64(0)
	timeStarted := time.Now()

	resultChannel := make(chan *result, len(g.Backends)*2)
	elapsed := time.Duration(0)

	//keep track of how many miner threads there are
	nrunning := 0

	//send 1 chunk to each miner
	for _,b := range g.Backends {
		//dispatch work
		sent += b.dispatchWork(hash, start + sent, resultChannel)
		nrunning++
	}
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	nextHeartbeat := cfg.Heartbeat.Duration

	// now loop until we either run out of time, finish the requested number of hashes, or find solution
	// each time a hasher finishes a chunk, give it a new one to work on.
	// always waits for all miners to finish their chunks before returning (thread safety with OpenCL)
	// EXCEPT in the case of an error, but then the app is almost certainly just quitting anyways!
	sol := ""
	for (recv < n && elapsed < timeout && sol == "") || (nrunning > 0) {
		if elapsed > nextHeartbeat {
			g.PrintHashRateSummary()
			nextHeartbeat += cfg.Heartbeat.Duration
		}
		result := <-resultChannel
		nrunning--
		if result.err != nil {
			return "", fmt.Errorf("hasher failed: %s", result.err.Error())
		} else {
			result.backend.TotalHashes += result.n
			recv += result.n
			if result.nonce != "" {
				sol = result.nonce
			}
			//only update the hashRateEstimate if we didn't find a solution - otherwise the rate could be wrong
			result.backend.HashRateEstimate = float64(result.n)/(result.finished.Sub(result.started).Seconds())

			if sent < n && elapsed < timeout {
				sent += result.backend.dispatchWork(hash, start + sent, resultChannel)
				nrunning++
			}
		}
		elapsed = time.Now().Sub(timeStarted)
	}

	return sol, nil
}




