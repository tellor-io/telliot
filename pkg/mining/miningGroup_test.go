// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package mining

import (
	"context"

	"github.com/pkg/errors"

	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/testutil"
	"github.com/tellor-io/telliot/pkg/util"

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
	_string := fmt.Sprintf("%x", challenge.Challenge) + "0000000000000000000000000000000000000000"
	hashIn := util.DecodeHex(_string)
	hashIn = append(hashIn, []byte(nonce)...)

	a, err := hashFn(hashIn)
	testutil.Ok(t, err)

	a.Mod(a, challenge.Difficulty)
	if !a.IsUint64() || a.Uint64() != 0 {
		testutil.Ok(t, errors.Errorf("nonce: %s remainder: %s\n", string(hashIn[52:]), a.Text(10)))
	}
}

func DoCompleteMiningLoop(t *testing.T, impl Hasher, diff int64) {
	cfg, err := config.OpenTestConfig("../..")
	testutil.Ok(t, err)

	opts := &rpc.MockOptions{
		Nonce:         1,
		GasPrice:      big.NewInt(700000000),
		TokenBalance:  big.NewInt(0),
		Top50Requests: []*big.Int{},
	}
	client := rpc.NewMockClientWithValues(opts)
	contract, err := contracts.NewITellor(client)
	if err != nil {
		testutil.Ok(t, errors.Wrap(err, "creating new contract instance"))
	}
	group, err := NewMiningGroup(logging.NewLogger(), cfg, []Hasher{impl}, contract)
	if err != nil {
		testutil.Ok(t, errors.Wrap(err, "creating new mining group"))
	}

	timeout := time.Millisecond * 200

	input := make(chan *Work)
	output := make(chan *Result)

	ctx, close := context.WithCancel(context.Background())
	go group.Mine(ctx, input, output)

	testVectors := []int{19, 133, 8, 442, 1231}
	for _, v := range testVectors {
		challenge := createChallenge(v, diff)
		input <- &Work{Challenge: challenge, Start: 0, PublicAddr: "0x0000000000000000000000000000000000000000", N: math.MaxInt64}

		// Wait for a solution to be found.
		select {
		case result := <-output:
			if result == nil {
				testutil.Ok(t, errors.Errorf("nil result for challenge %v", v))
			} else {
				// Fixing a possible nil pointer deference... not sure if that's the appropriate way to do it.
				CheckSolution(t, challenge, result.Nonce)
			}
		case <-time.After(timeout):
			testutil.Ok(t, errors.Errorf("Expected result for challenge in less than %s", timeout.String()))
		}
	}
	// Tell the mining group to close.
	close()

	// Wait for it to close.
	// select {
	// case result := <-output:
	// 	if result != nil {
	// 		testutil.Ok(t, errors.New("expected nil result when closing mining group"))
	// 	}
	// case <-time.After(timeout):
	// 	testutil.Ok(t, errors.Errorf("Expected mining group to close in less than %s", timeout.String()))
	// }
}

func TestCpuMiner(t *testing.T) {
	impl := NewCpuMiner(0)
	DoCompleteMiningLoop(t, impl, 100)
}

func TestMulti(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	cfg, err := config.OpenTestConfig("../..")
	testutil.Ok(t, err)

	opts := &rpc.MockOptions{
		Nonce:         1,
		GasPrice:      big.NewInt(700000000),
		TokenBalance:  big.NewInt(0),
		Top50Requests: []*big.Int{},
	}
	client := rpc.NewMockClientWithValues(opts)
	contract, err := contracts.NewITellor(client)
	testutil.Ok(t, err)

	var hashers []Hasher
	for i := 0; i < 4; i++ {
		hashers = append(hashers, NewCpuMiner(int64(i)))
	}

	fmt.Printf("Using %d hashers\n", len(hashers))
	group, err := NewMiningGroup(logging.NewLogger(), cfg, hashers, contract)
	if err != nil {
		testutil.NotOk(t, errors.Wrap(err, "creating new mining group"))
	}
	input := make(chan *Work)
	output := make(chan *Result)
	ctx, close := context.WithCancel(context.Background())
	go group.Mine(ctx, input, output)

	challenge := createChallenge(0, math.MaxInt64)
	input <- &Work{Challenge: challenge, Start: 0, PublicAddr: "0x0000000000000000000000000000000000000000", N: math.MaxInt64}
	time.Sleep(1 * time.Second)
	close()
	// timeout := 500 * time.Millisecond
	// select {
	// case <-output:
	// 	group.PrintHashRateSummary()
	// case <-time.After(timeout):
	// 	testutil.Ok(t, errors.Errorf("mining group didn't quit before %s", timeout.String()))
	// }
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
		bytes := util.DecodeHex(_string)
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
	bytes := util.DecodeHex(_string)

	for i := 0; i < b.N; i++ {
		_, err := hashFn(bytes)
		testutil.Ok(b, err)
	}
}
