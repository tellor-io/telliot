package pow

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/math"
)

func createChallenge(id int, difficulty int64) *miningChallenge {
	hash := math.PaddedBigBytes(big.NewInt(int64(id)), 32)
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


func createRandomChallenge(difficulty int64) *miningChallenge {
	rand.Seed(time.Now().UnixNano())
	i := rand.Int()
	return createChallenge(i, difficulty)
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
	challenge := createRandomChallenge(1000000)

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
	miner, err := createMiningLoopWithStart(1, 0)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	miner.Start(ctx)

	testVectors := make(map[int]string)
	testVectors[0] = "19"
	testVectors[1] = "133"
	testVectors[2] = "8"
	testVectors[3] = "442"
	testVectors[4] = "528"
	for k,v := range testVectors {
		challenge := createChallenge(k, 500)
		miner.taskCh <- challenge
		timeout := time.Millisecond * 50
		select {
		case sol := <-miner.solutionCh:
			if sol == nil {
				t.Fatal("Expected a mining solution in the end")
			}
			if sol.nonce != v {
				t.Fatal("expected a different number of hashes before finding solution")
			}
		case <-time.After(timeout):
			t.Fatalf("Expected mining to finish within %s", timeout.String())
		}

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
	challenge := createRandomChallenge(1000000)
	challenge2 := createRandomChallenge(500)
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

	challenge := createRandomChallenge(1000000)
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

	challenge := createRandomChallenge(1000000)
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

func TestHashFunction(t *testing.T) {

	challenge := createChallenge(734561, 500)

	testVectors := make(map[int]string)
	testVectors[46] = "7a29a4ea30744b40ff70d9a3ef8e6cc1ec8aa0a80a8a914ad4c0e9c9ea781b7"
	testVectors[3751] = "94c7bbe18751463f8e84a433c3414602b3d569b840e403c92bae8e5b81726c6d"
	testVectors[982879] = "866db7221f0bfcd36efd3e00da593a081c8519995659f8abcf97f189ecba6c64"
	testVectors[5] = "acb584c01027480cc06a039f9dba9b1d834efa9b34fa41da95245956bcf353a1"
	testVectors[0] = "b864b47407a9f328a3d5eee5c1996ea048ac35e2f3a96396c34555aa7ea4ff4a"
	testVectors[1] = "6ad05c010b7ec871d7d72a7e8d12ad69f00f73ada2553ad517185fbfc1e3da82"

	result := new(big.Int)
	for k,v := range testVectors {
		nonce := fmt.Sprintf("%x", fmt.Sprintf("%d", k))
		_string := fmt.Sprintf("%x", challenge.challenge) + "abcd0123" + nonce
		bytes := decodeHex(_string)
		hash(bytes, result)
		if result.Text(16) != v {
			t.Fatalf("wrong hash:\nexpected:\n%s\ngot:\n%s\n", v, result.Text(16))
		}
	}
}

func BenchmarkHashFunction(b *testing.B) {
	challenge := createChallenge(0, 500)
	result := new(big.Int)
	nonce := fmt.Sprintf("%x", fmt.Sprintf("%d", 10))
	_string := fmt.Sprintf("%x", challenge.challenge) + "abcd0123" + nonce
	bytes := decodeHex(_string)

	for i := 0; i < b.N; i++ {
		hash(bytes, result)
	}
}


