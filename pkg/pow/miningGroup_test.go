// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package pow

import (
	"github.com/pkg/errors"

	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/testutil"

	"github.com/ethereum/go-ethereum/common/math"
)

func createChallenge(id int, difficulty int64) *MiningChallenge {
	hash := math.PaddedBigBytes(big.NewInt(int64(id)), 32)
	var b32 [32]byte
	for i, v := range hash {
		b32[i] = v
	}

	return &MiningChallenge{
		Challenge:  b32[:],
		Difficulty: big.NewInt(difficulty),
		RequestIDs: [5]*big.Int{big.NewInt(1)},
	}
}

func CheckSolution(t *testing.T, challenge *MiningChallenge, nonce string) {
	cfg := config.OpenTestConfig(t)
	_string := fmt.Sprintf("%x", challenge.Challenge) + cfg.PublicAddress
	hashIn := decodeHex(_string)
	hashIn = append(hashIn, []byte(nonce)...)

	a, err := hashFn(hashIn)
	testutil.Ok(t, err)

	a.Mod(a, challenge.Difficulty)
	if !a.IsUint64() || a.Uint64() != 0 {
		testutil.Ok(t, errors.Errorf("nonce: %s remainder: %s\n", string(hashIn[52:]), a.Text(10)))
	}
}

func DoCompleteMiningLoop(t *testing.T, impl Hasher, diff int64) {
	cfg := config.OpenTestConfig(t)
	exitCh := make(chan os.Signal)
	group := NewMiningGroup([]Hasher{impl}, exitCh)

	timeout := time.Millisecond * 200

	input := make(chan *Work)
	output := make(chan *Result)

	go group.Mine(input, output)

	testVectors := []int{19, 133, 8, 442, 1231}
	for _, v := range testVectors {
		challenge := createChallenge(v, diff)
		input <- &Work{Challenge: challenge, Start: 0, PublicAddr: cfg.PublicAddress, N: math.MaxInt64}

		// Wait for a solution to be found.
		select {
		case result := <-output:
			if result == nil {
				testutil.Ok(t, errors.Errorf("nil result for challenge %v", v))
			} else {
				//Fixing a possible nil pointer deference... not sure if that's the appropriate way to do it
				CheckSolution(t, challenge, result.Nonce)
			}
		case <-time.After(timeout):
			testutil.Ok(t, errors.Errorf("Expected result for challenge in less than %s", timeout.String()))
		}
	}
}

func TestCpuMiner(t *testing.T) {
	impl := NewCpuMiner(0)
	DoCompleteMiningLoop(t, impl, 100)
}

func TestGpuMiner(t *testing.T) {
	cfg := config.OpenTestConfig(t)
	gpus, err := GetOpenCLGPUs()
	testutil.Ok(t, err)
	if len(gpus) == 0 {
		t.Skip("no mining gpus")
	}

	impl, err := NewGpuMiner(gpus[0], cfg.GPUConfig[gpus[0].Name()], false)
	testutil.Ok(t, err)
	DoCompleteMiningLoop(t, impl, 1000)
}

func TestMulti(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	cfg := config.OpenTestConfig(t)

	var hashers []Hasher
	for i := 0; i < 4; i++ {
		hashers = append(hashers, NewCpuMiner(int64(i)))
	}
	gpus, err := GetOpenCLGPUs()
	if err != nil {
		fmt.Println(gpus)
		testutil.Ok(t, err)
	}
	for _, gpu := range gpus {
		impl, err := NewGpuMiner(gpu, cfg.GPUConfig[gpu.Name()], false)
		testutil.Ok(t, err)
		hashers = append(hashers, impl)
	}
	fmt.Printf("Using %d hashers\n", len(hashers))
	exitCh := make(chan os.Signal)
	group := NewMiningGroup(hashers, exitCh)
	input := make(chan *Work)
	output := make(chan *Result)
	go group.Mine(input, output)

	challenge := createChallenge(0, math.MaxInt64)
	input <- &Work{Challenge: challenge, Start: 0, PublicAddr: cfg.PublicAddress, N: math.MaxInt64}
	time.Sleep(1 * time.Second)

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

	for k, v := range testVectors {
		nonce := fmt.Sprintf("%x", fmt.Sprintf("%d", k))
		_string := fmt.Sprintf("%x", challenge.Challenge) + "abcd0123" + nonce
		bytes := decodeHex(_string)
		result, err := hashFn(bytes)
		testutil.Ok(t, err)
		if result.Text(16) != v {
			testutil.Ok(t, errors.Errorf("wrong hash:\nexpected:\n%s\ngot:\n%s\n", v, result.Text(16)))
		}
	}
}

func BenchmarkHashFunction(b *testing.B) {
	challenge := createChallenge(0, 500)
	nonce := fmt.Sprintf("%x", fmt.Sprintf("%d", 10))
	_string := fmt.Sprintf("%x", challenge.Challenge) + "abcd0123" + nonce
	bytes := decodeHex(_string)

	for i := 0; i < b.N; i++ {
		_, err := hashFn(bytes)
		testutil.Ok(b, err)
	}
}

var configJSON = `{
    "contractAddress": "0x0Ba45A8b5d5575935B8158a88C631E9F9C95a2e5",
    "databaseURL": "http://localhost7545",
    "publicAddress": "92f91500e105e3051f3cf94616831b58f6bce1e8",
    "serverHost": "localhost",
    "serverPort": 5000,
    "trackerCycle": 1,
    "trackers": {},
    "dbFile": "/tellorDB",
    "requestTips": 1,
    "configFolder": "` + filepath.Join("..", "..", "configs") + `",
    "envFile": "` + filepath.Join("..", "..", "configs", ".env.example") + `"
}`

func TestMain(m *testing.M) {
	err := config.ParseConfigBytes([]byte(configJSON))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse mock config: %v\n", err)
		os.Exit(-1)
	}
	os.Exit(m.Run())
}
