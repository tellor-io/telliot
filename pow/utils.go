package pow

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
)

func randInt() string {
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(126), nil).Sub(max, big.NewInt(1))

	//Generate cryptographically strong pseudo-random between 0 - max
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		//error handling
	}
	return n.String()
}

func decodeHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return b
}
