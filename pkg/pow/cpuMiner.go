// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package pow

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/tellor-io/TellorMiner/pkg/rpc"
	// nolint:staticcheck
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

func (c *CpuMiner) CheckRange(hash *HashSettings, start uint64, n uint64) (string, uint64, error) {
	baseLen := len(hash.prefix)
	hashInput := make([]byte, len(hash.prefix))
	copy(hashInput, hash.prefix)

	x := new(big.Int)
	compareZero := big.NewInt(0)

	for i := start; i < (start + n); i++ {
		nn := strconv.FormatUint(i, 10)
		hashInput = hashInput[:baseLen]
		hashInput = append(hashInput, []byte(nn)...)
		numHash, err := rpc.HashFn(hashInput)
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
