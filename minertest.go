package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"strings"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"golang.org/x/crypto/ripemd160"
	// "crypto/sha256"
	// "golang.org/x/crypto/ripemd160"
)

const public_address string = "0xe037ec8ec9ec423826750853899394de7f024fee"
const challenge string = "a9fd1780508babd96cc82ee6f0fce53cff9127119129fb4b2a33053458dc4495"
const difficulty int64 = 1

func main() {
	mine()
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
func hexaNumberToInteger(hexaString string) string {
	// replace 0x or 0X with empty String
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return numberStr
}
func mine() uint {
	var i uint = 0 //
	for i < 1 {
		//nonce := randInt(0, 100000000000)
		nonce := fmt.Sprintf("%02x", fmt.Sprint(9000))
		_string := challenge + public_address[2:] + nonce
		hash := solsha3.SoliditySHA3(
			solsha3.Bytes32(decodeHex(_string)),
		)
		hasher := ripemd160.New()
		hasher.Write([]byte(hash))
		hash1 := hasher.Sum(nil)
		n := sha256.Sum256([]byte(hash1))
		q := fmt.Sprintf("%x", n)
		p := new(big.Int)
		p, ok := p.SetString(q, 16)
		if !ok {
			fmt.Println("SetString: error")
			return 1
		}
		fmt.Println(p)
		v := big.NewInt(difficulty)
		x := new(big.Int)
		fmt.Println(v)
		x.Mod(p, v)
		fmt.Println(x)
		if x.Cmp(big.NewInt(0)) == 0 {
			fmt.Println("Solution Found", p)
			return i
		}
		fmt.Println("Solution Not Found", p)
		i++
	}
	return i
}
func decodeHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return b
}
