package pow

import (
	"crypto/sha256"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"golang.org/x/crypto/ripemd160"
	"math/big"
	"strconv"
)



type CpuMiner uint64

func NewCpuMiner(step uint64) *CpuMiner {
	x := CpuMiner(step)
	return &x
}

func (c *CpuMiner)StepSize() uint64 {
	return uint64(*c)
}

func (c *CpuMiner)CheckRange(base []byte, difficulty *big.Int,  start uint64, n uint64) (string, error) {
	baseLen := len(base)

	numHash := new(big.Int)
	x := new(big.Int)
	compareZero := big.NewInt(0)

	for i := start; i < (start + n); i++ {
		nn := strconv.FormatUint(i, 10)
		base = base[:baseLen]
		base = append(base, []byte(nn)...)
		hash(base, numHash)
		x.Mod(numHash, difficulty)
		if x.Cmp(compareZero) == 0 {
			return nn, nil
		}
	}
	return "", nil
}

func hash(data []byte, result *big.Int) {

	hash := solsha3.SoliditySHA3(data)

	//Consider moving hasher constructor outside loop and replacing with hasher.Reset()
	hasher := ripemd160.New()

	hasher.Write(hash)
	hash1 := hasher.Sum(nil)
	n := sha256.Sum256(hash1)
	result.SetBytes(n[:])
}

