package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tellor-io/TellorMiner/dataServer"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/pow"
)

func main() {
	//Start Database, wait until it is synced
	fmt.Println("Starting Tellor Miner Database")
	exitCh := make(chan int)
	ds, err := dataServer.CreateServer()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Initiated")
	ctx := context.Background()
	ds.Start(ctx, exitCh)
	fmt.Println("Database Ready")

	//Once synced, start miner
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("Starting Tellor Miner")
	mineTributes(ds)
}
func mineTributes(ds *dataServer.DataServer) {
	DB := ds.DB
	var nonce string
	var prevCurrentChallenge []byte
	x := 0
	fmt.Println("starting miner")
	for x < 1 {
		currentChallenge, _ := DB.Get(db.CurrentChallengeKey)
		if bytes.Compare(prevCurrentChallenge, currentChallenge) != 0 {
			prevCurrentChallenge = currentChallenge
			difficulty, _ := DB.Get(db.DifficultyKey)
			ndata, _ := hexutil.DecodeBig(string(difficulty))
			nonce = pow.SolveChallenge(currentChallenge, ndata)
			fmt.Println("nonce", nonce)
			if nonce != "" {
				requestID, _ := DB.Get(db.RequestIdKey)
				rdata, _ := hexutil.DecodeBig(string(requestID))
				value, _ := DB.Get(fmt.Sprint(rdata))
				ndata, _ := hexutil.DecodeBig(string(value))
				fmt.Println("Submitting Solution:", nonce, ndata, rdata)
				//pow.SubmitTransaction(nonce, ndata, rdata)
				nonce = ""
			}
			x++
		} else {
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
