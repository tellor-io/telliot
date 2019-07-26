package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

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
	time.Sleep(1000 * time.Millisecond)
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
			var data uint64
			for i, x := range difficulty {
				data |= uint64(x) << uint64(i*8)
			}
			fmt.Println("Difficulty", difficulty)
			ndata := big.NewInt(int64(data))
			nonce = pow.SolveChallenge(currentChallenge, ndata)
			fmt.Println("nonce", nonce)
			if nonce != "" {
				requestID, _ := DB.Get(db.RequestIdKey)
				for i, x := range requestID {
					data |= uint64(x) << uint64(i*8)
				}
				rdata := big.NewInt(int64(data))
				value, _ := DB.Get(fmt.Sprint(requestID))
				for i, x := range value {
					data |= uint64(x) << uint64(i*8)
				}
				ndata = big.NewInt(int64(data))
				fmt.Println("Submitting Solution: nonce,ndata,rdata")
				pow.SubmitTransaction(nonce, ndata, rdata)
				nonce = ""
			}
			x++
		} else {
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
