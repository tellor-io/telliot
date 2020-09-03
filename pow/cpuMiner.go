package pow

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"golang.org/x/crypto/ripemd160"
)



type CpuMiner int64

func NewCpuMiner(id int64) *CpuMiner {
	x := CpuMiner(id)
	return &x
}

func (c *CpuMiner)StepSize() uint64 {
	return 1
}

func (c *CpuMiner)Name() string {
	return fmt.Sprintf("CPU %d", *c)
}

func (c *CpuMiner)CheckRange(hash *HashSettings,  start uint64, n uint64) (string, uint64, error) {
	baseLen := len(hash.prefix)
	hashInput := make([]byte, len(hash.prefix), len(hash.prefix))
	copy(hashInput, hash.prefix)

	numHash := new(big.Int)
	x := new(big.Int)
	compareZero := big.NewInt(0)

	for i := start; i < (start + n); i++ {
		nn := strconv.FormatUint(i, 10)
		hashInput = hashInput[:baseLen]
		hashInput = append(hashInput, []byte(nn)...)
		hashFn(hashInput, numHash)
		x.Mod(numHash, hash.difficulty)
		if x.Cmp(compareZero) == 0 {
			return nn, (i-start)+1, nil
		}
	}
	return "", n, nil
}

func hashFn(data []byte, result *big.Int) {

	hash := solsha3.SoliditySHA3(data)

	//Consider moving hasher constructor outside loop and replacing with hasher.Reset()
	hasher := ripemd160.New()

	hasher.Write(hash)
	hash1 := hasher.Sum(nil)
	n := sha256.Sum256(hash1)
	result.SetBytes(n[:])
}
