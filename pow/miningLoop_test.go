package pow

import (
	"bytes"
	"context"
	"math/big"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/math"
)

func createChallenge(difficulty int64) *miningChallenge {
	rand.Seed(time.Now().UnixNano())
	i := rand.Int()
	hash := math.PaddedBigBytes(big.NewInt(int64(i)), 32)
	var b32 [32]byte
	for i, v := range hash {
		b32[i] = v
	}

	return &miningChallenge{
		challenge:  b32[:],
		difficulty: big.NewInt(difficulty),
		requestID:  big.NewInt(1),
	}
}

func TestBasicMiningLoop(t *testing.T) {
	miner, err := createMiningLoop(1)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	miner.Start(ctx)
	time.Sleep(300 * time.Millisecond)
	if miner.mining {
		t.Fatal("Should not be mining after start without challenge issued")
	}
	challenge := createChallenge(1000000)

	var receivedSolution *miningSolution
	go func() {
		receivedSolution = <-miner.solutionCh
	}()
	miner.taskCh <- challenge
	time.Sleep(10 * time.Millisecond)
	if !miner.mining {
		t.Fatal("Expected miner to start mining once issued a challenge")
	}
	miner.taskCh <- challenge
	if !miner.mining && receivedSolution == nil {
		t.Fatal("Duplicate task should interrupt mining")
	}
	miner.cancelCh <- true
	time.Sleep(10 * time.Millisecond)
	if miner.mining {
		t.Fatal("Expected miner to stop mining on cancel")
	}
}

func TestCompletedMiningLoop(t *testing.T) {
	miner, err := createMiningLoop(1)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	miner.Start(ctx)

	var receivedSolution *miningSolution
	go func() {
		receivedSolution = <-miner.solutionCh
	}()
	challenge := createChallenge(500)
	miner.taskCh <- challenge
	start := time.Now()
	expiration := time.Now().Add(5 * time.Second)
	for {
		if start.After(expiration) {
			t.Fatal("Expected mining to finish within 5 seconds")
		}
		time.Sleep(1 * time.Second)
		if !miner.mining {
			break
		}
	}

	if receivedSolution == nil {
		t.Fatal("Expected a mining solution in the end")
	}
}

func TestMiningLoopNewChallengeInterrupt(t *testing.T) {
	miner, err := createMiningLoop(1)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	miner.Start(ctx)

	var receivedSolution *miningSolution
	go func() {
		receivedSolution = <-miner.solutionCh
	}()
	challenge := createChallenge(1000000)
	challenge2 := createChallenge(500)
	miner.taskCh <- challenge
	time.Sleep(10 * time.Millisecond)
	if receivedSolution == nil {
		miner.taskCh <- challenge2
		time.Sleep(1 * time.Second)
		if receivedSolution == nil {
			t.Fatal("Expected a solution to be mined quickly on second challenge")
		}
		if bytes.Compare(receivedSolution.challenge.challenge, challenge2.challenge) != 0 {
			t.Fatal("Expected second challenge to have replaced original challenge for mining")
		}
	}
}

func TestMiningLoopStop(t *testing.T) {
	miner, err := createMiningLoop(1)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	miner.Start(ctx)

	challenge := createChallenge(1000000)
	miner.taskCh <- challenge
	time.Sleep(10 * time.Millisecond)
	if miner.mining {
		go miner.stopMining()
	}
	start := time.Now()
	exp := start.Add(5 * time.Second)
	for {
		if time.Now().After(exp) {
			t.Fatal("Expected to stop mining by now")
		}
		if !miner.mining {
			break
		}
	}
}

func TestMiningLoopOSInterrupt(t *testing.T) {
	miner, err := createMiningLoop(1)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	miner.Start(ctx)

	challenge := createChallenge(1000000)
	miner.taskCh <- challenge
	time.Sleep(10 * time.Millisecond)
	if miner.mining {
		miner.exitCh <- os.Interrupt
	}
	start := time.Now()
	exp := start.Add(5 * time.Second)
	for {
		if time.Now().After(exp) {
			t.Fatal("Expected to stop mining by now")
		}
		if !miner.mining {
			break
		}
	}
}
