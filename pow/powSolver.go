package pow

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/tellor-io/TellorMiner/config"
	tellor1 "github.com/tellor-io/TellorMiner/contracts1"
	"golang.org/x/crypto/ripemd160"
)

func randInt() string {
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))

	//Generate cryptographically strong pseudo-random between 0 - max
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		//error handling
	}
	//n := big.NewInt(9000)
	//String representation of n in base 16
	nonce := fmt.Sprintf("%x", n.String()) //n.Text(16)
	return nonce
}

func decodeHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return b
}

func SolveChallenge(challenge []byte, _difficulty *big.Int) string {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 100000000; i++ {
		nonce := randInt() //do we need to use big number?
		fmt.Println(nonce)
		_string := fmt.Sprintf("%x", challenge) + cfg.PublicAddress + nonce
		fmt.Println("String created", _string)
		hash := solsha3.SoliditySHA3(
			solsha3.Bytes32(decodeHex(_string)),
		)
		hasher := ripemd160.New()
		hasher.Write([]byte(hash))
		hash1 := hasher.Sum(nil)
		n := sha256.Sum256([]byte(hash1))
		q := fmt.Sprintf("%x", n)
		fmt.Println("Sha256 found", q)
		p := new(big.Int)
		p, ok := p.SetString(q, 16)
		if !ok {
			fmt.Println("SetString: error")
			return ""
		}
		x := new(big.Int)
		x.Mod(p, _difficulty)
		fmt.Println(x)
		if x.Cmp(big.NewInt(0)) == 0 {
			fmt.Println("Solution Found", p)
			return nonce
		}
	}
	return ""
}

func SubmitTransaction(solution string, value, requestId *big.Int) error {

	//get the single config instance
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
		return err
	}

	client, err := ethclient.Dial(cfg.NodeURL)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	cost := new(big.Int)
	cost.Mul(gasPrice, big.NewInt(700000))
	if balance.Cmp(cost) >= 0 {
		auth := bind.NewKeyedTransactor(privateKey)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)     // in wei
		auth.GasLimit = uint64(300000) // in units
		auth.GasPrice = gasPrice

		contractAddress := common.HexToAddress(cfg.ContractAddress)
		instance, err := tellor1.NewTellorTransactor(contractAddress, client)
		if err != nil {
			log.Fatal(err)
		}
		solution := string(solution)
		fmt.Println(auth, solution, requestId, value)
		tx, err := instance.SubmitMiningSolution(auth, solution, requestId, value)
		if err != nil {
			log.Fatal(err)
			return err
		}

		fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	}
	return nil
}
