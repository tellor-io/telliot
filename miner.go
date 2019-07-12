/*
Tellor Miner

This is the workhorse of the Miner system as it takes on solving the PoW challenge.
As for the feel of this piece, we are going to follow the Geth model to keep a familiar interface for users.

The main goals for this build are:
	Solve PoW
	Submit Transaction to the Ethereum network

Other features that it should also have to properly do these main tasks:
	Grab challenge details from database
	Grab request data from database
	Optimize transaction details (gas price)
	Pull current gas price from database
	Handle transaction failure when submitting to Ethereum
	Show low gas warning (more ETH needed for gas)
	IPC interface for interacting with Miner (Geth model)

Resources:
https://goethereumbook.org/client-setup/

To create go file:
	cd contracts
	abigen -sol TellorMaster.sol -pkg contracts -out tellorMaster.go


*/

package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	tellor "./contracts"
	tellor1 "./contracts1"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"golang.org/x/crypto/ripemd160"
)

type Config struct {
	Contract_Address string `json:"contract_Address"`
	NodeURL          string `json:"nodeURL"`
	Private_Key      string `json:"private_Key"`
	DatabaseURL      string `json:"databaseURL"`
	Public_address   string `json:"public_address"`
}

func getCurrentChallenge(Config Config) ([]byte, *big.Int, *big.Int, string, *big.Int, *big.Int) {

	client, err := ethclient.Dial(Config.NodeURL)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(Config.Contract_Address)
	instance, err := tellor.NewTellorMaster(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	currentChallenge, requestId, difficulty, queryString, granularity, totalTip, err := instance.GetCurrentVariables(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("%x", currentChallenge), requestId, difficulty, queryString, granularity, totalTip) // 60806...10029
	return currentChallenge[:], requestId, difficulty, queryString, granularity, totalTip
}

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

func solveChallenge(Config Config, challenge []byte, _difficulty *big.Int) string {
	for i := 0; i < 100000000; i++ {
		nonce := randInt() //do we need to use big number?
		fmt.Println(nonce)
		_string := fmt.Sprintf("%x", challenge) + Config.Public_address + nonce
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

func getRequestedValues(_requestId, _granularity *big.Int) (bool, *big.Int) {
	value := new(big.Int)
	value.Mul(big.NewInt(1000), _granularity)
	fmt.Println("Getting Request Id", _requestId)
	return true, value

}

func submitTransaction(Config Config, solution string, value, requestId *big.Int) {
	client, err := ethclient.Dial(Config.NodeURL)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(Config.Private_Key)
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

		contractAddress := common.HexToAddress(Config.Contract_Address)
		instance, err := tellor1.NewTellorTransactor(contractAddress, client)
		if err != nil {
			log.Fatal(err)
		}
		solution := string(solution)
		fmt.Println(auth, solution, requestId, value)
		tx, err := instance.SubmitMiningSolution(auth, solution, requestId, value)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	}
}

func main() {

	var config Config
	configFile, err := os.Open("config.json")
	fmt.Println(configFile)
	defer configFile.Close()
	if err != nil {
		return
	}
	dec := json.NewDecoder(configFile)
	err = dec.Decode(&config)
	fmt.Println("config", config)
	var nonce string
	var prevCurrentChallenge []byte
	x := 0
	fmt.Println("starting miner")
	for x < 10 {
		currentChallenge, requestId, difficulty, _, granularity, _ := getCurrentChallenge(config)
		fmt.Println("challenge retrieved", requestId.String())
		if x == 0 || bytes.Compare(prevCurrentChallenge, currentChallenge) != 0 {
			prevCurrentChallenge = currentChallenge
			fmt.Println("Going to get Challenge")
			nonce = solveChallenge(config, currentChallenge, difficulty)
			fmt.Println("Challenge Solved", nonce)
			if nonce != "" {
				goodValue, value := getRequestedValues(requestId, granularity)
				fmt.Println("Value Retrieved", value.String())
				if goodValue {
					submitTransaction(config, nonce, value, requestId)
					fmt.Println("Transaction submitted!")
				} else {
					fmt.Println("Value Error")
				}
				nonce = ""
			}
			x++
		} else {
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
