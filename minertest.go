package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	tellor "./contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const public_address string = "0xe037ec8ec9ec423826750853899394de7f024fee"
const challenge string = "a9fd1780508babd96cc82ee6f0fce53cff9127119129fb4b2a33053458dc4495"
const difficulty int64 = 1

func main() {
	submitSolution()
}

func submitSolution() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("4bdc16637633fa4b4854670fbb83fa254756798009f52a1d3add27fb5f5a8e16")
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

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	contractAddress := common.HexToAddress("0x9f7bC259319001166e986210d846af7C15B1f644")
	instance, err := tellor.NewTellor(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	var requestId uint = 9
	var value uint = 1000
	solution := "85000"

	tx, err := instance.submitMiningSolution(auth, solution, requestId, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
}
