// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package pow

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/crypto"
	// nolint:staticcheck
	"golang.org/x/crypto/ripemd160"
)

type CpuMiner int64

func NewCpuMiner(id int64) *CpuMiner {
	x := CpuMiner(id)
	return &x
}

func (c *CpuMiner) StepSize() uint64 {
	return 1
}

func (c *CpuMiner) Name() string {
	return fmt.Sprintf("CPU %d", *c)
}

func (c *CpuMiner) CheckRange(hash *HashSettings, start uint64, n uint64, ctx context.Context) (string, uint64, error) {
	baseLen := len(hash.prefix)
	hashInput := make([]byte, len(hash.prefix))
	copy(hashInput, hash.prefix)

	x := new(big.Int)
	compareZero := big.NewInt(0)

	for i := start; i < (start + n); i++ {
		select {
		case <-ctx.Done():
			return "context expired", n, nil
		default:
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
