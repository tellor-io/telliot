package pow

import (
	"fmt"
	"github.com/tellor-io/TellorMiner/config"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/math"
)

func createChallenge(id int, difficulty int64) *MiningChallenge {
	hash := math.PaddedBigBytes(big.NewInt(int64(id)), 32)
	var b32 [32]byte
	for i, v := range hash {
		b32[i] = v
	}

	return &MiningChallenge{
		challenge:  b32[:],
		difficulty: big.NewInt(difficulty),
		requestID:  big.NewInt(1),
	}
}


func createRandomChallenge(difficulty int64) *MiningChallenge {
	rand.Seed(time.Now().UnixNano())
	i := rand.Int()
	return createChallenge(i, difficulty)
}

//func TestBasicMiningLoop(t *testing.T) {
//	miner, err := createMiningLoop(1, NewCpuMiner(50e3))
//	if err != nil {
//		t.Fatal(err)
//	}
//	ctx := context.Background()
//	miner.Start(ctx)
//	time.Sleep(300 * time.Millisecond)
//	if miner.mining {
//		t.Fatal("Should not be mining after start without challenge issued")
//	}
//	challenge := createRandomChallenge(1000000)
//
//	var receivedSolution *miningSolution
//	go func() {
//		receivedSolution = <-miner.solutionCh
//	}()
//	miner.taskCh <- challenge
//	time.Sleep(10 * time.Millisecond)
//	if !miner.mining {
//		t.Fatal("Expected miner to start mining once issued a challenge")
//	}
//	miner.taskCh <- challenge
//	if !miner.mining && receivedSolution == nil {
//		t.Fatal("Duplicate task should interrupt mining")
//	}
//	miner.cancelCh <- true
//	time.Sleep(10 * time.Millisecond)
//	if miner.mining {
//		t.Fatal("Expected miner to stop mining on cancel")
//	}
//}

func CheckSolution(t *testing.T, challenge *MiningChallenge, nonce string) {
	cfg, err := config.GetConfig()
	if err != nil {
		t.Fatal(err)
	}
	_string := fmt.Sprintf("%x", challenge.challenge) + cfg.PublicAddress
	hashIn := decodeHex(_string)
	hashIn = append(hashIn, []byte(nonce)...)
	a := new(big.Int)
	hashFn(hashIn, a)

	a.Mod(a, challenge.difficulty)
	if !a.IsUint64() || a.Uint64() != 0 {
		t.Fatalf("nonce: %s remainder: %s\n", string(hashIn[52:]), a.Text(10))
	}
}



func DoCompleteMiningLoop(t *testing.T, impl Hasher, diff int64) {

	cfg, err := config.GetConfig()
	if err != nil {
		t.Fatal(err)
	}

	group := NewMiningGroup([]Hasher{impl})


	timeout := time.Millisecond * 100

	testVectors := []int{19, 133, 8, 442, 1231}
	for _,v := range testVectors {
		challenge := createChallenge(v, diff)
		settings := NewHashSettings(challenge, cfg.PublicAddress)
		sol, err := group.Mine(settings, 0, 1e6, timeout)
		if err != nil {
			t.Fatal(err)
		}
		if sol == "" {
			t.Fatal("Expected a mining solution in the end")
		}
		CheckSolution(t, challenge, sol)
	}
}

func TestCpuMiner(t *testing.T) {
	impl := NewCpuMiner(0)
	DoCompleteMiningLoop(t, impl, 100)
}

func TestGpuMiner(t *testing.T) {
	gpus, err := GetOpenCLGPUs()
	if err != nil {
		fmt.Println(gpus)
		t.Fatal(err)
	}
	impl, err := NewGpuMiner(gpus[0])
	if err != nil {
		t.Fatal(err)
	}
	DoCompleteMiningLoop(t, impl, 1000)
}

func TestMulti(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	var hashers []Hasher
	for i := 0; i < 4; i++ {
		hashers = append(hashers, NewCpuMiner(int64(i)))
	}
	gpus, err := GetOpenCLGPUs()
	if err != nil {
		fmt.Println(gpus)
		t.Fatal(err)
	}
	for _,gpu := range gpus {
		impl, err := NewGpuMiner(gpu)
		if err != nil {
			t.Fatal(err)
		}
		hashers = append(hashers, impl)
	}

	fmt.Printf("Using %d hashers\n", len(hashers))

	cfg, err := config.GetConfig()
	if err != nil {
		t.Fatal(err)
	}

	challenge := createChallenge(0, math.MaxInt64)
	settings := NewHashSettings(challenge, cfg.PublicAddress)
	group := NewMiningGroup(hashers)
	_, err = group.Mine(settings, 0, 1e12, 1000*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	group.PrintHashRateSummary()
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
		hashFn(bytes, result)
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
		hashFn(bytes, result)
	}
}


