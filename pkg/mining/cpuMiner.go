// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package mining

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/rpc"

	// nolint:staticcheck
	"golang.org/x/crypto/ripemd160"
)

type CpuMiner struct {
	id               int64
	contractInstance *contracts.ITellor
}

func NewCpuMiner(id int64, contractInstance *contracts.ITellor) *CpuMiner {
	return &CpuMiner{
		id:               id,
		contractInstance: contractInstance,
	}
}

func (c *CpuMiner) StepSize() uint64 {
	return 1
}

func (c *CpuMiner) Name() string {
	return fmt.Sprintf("CPU %d", c.id)
}

func (c *CpuMiner) CheckRange(hash *HashSettings, start uint64, n uint64) (string, uint64, error) {
	baseLen := len(hash.prefix)
	hashInput := make([]byte, len(hash.prefix))
	copy(hashInput, hash.prefix)

	x := new(big.Int)
	compareZero := big.NewInt(0)
	for i := start; i < (start + n); i++ {
		// checks the last submit value in the oracle and set a timeout of 15min - (now-lastSubmit).
		// This is because 15min after the last submit any solution will work.
		timeOfLastNewValue, err := c.contractInstance.GetUintVar(nil, rpc.Keccak256([]byte("_TIME_OF_LAST_NEW_VALUE")))
		if err != nil {
			// Return any result.
			return "", n, nil
		}
		now := time.Now()
		tm := time.Unix(timeOfLastNewValue.Int64(), 0)
		if now.Sub(tm) >= time.Duration(15)*time.Minute {
			// Return any result.
			return "", n, nil
		}
		nn := strconv.FormatUint(i, 10)
		hashInput = hashInput[:baseLen]
		hashInput = append(hashInput, []byte(nn)...)
		numHash, err := hashFn(hashInput)
		if err != nil {
			return "", 0, err
		}
		x.Mod(numHash, hash.difficulty)
		if x.Cmp(compareZero) == 0 {
			return nn, (i - start) + 1, nil
		}
	}
	return "", n, nil
}

func hashFn(input []byte) (*big.Int, error) {
	hash := crypto.Keccak256(input)
	hasher := ripemd160.New()
	if _, err := hasher.Write(hash); err != nil {
		return nil, err
	}
	hash1 := hasher.Sum(nil)
	n := sha256.Sum256(hash1)
	result := new(big.Int)
	result.SetBytes(n[:])

	return result, nil
}
