package main

import (
	"bytes"
	"context"
	"encoding/binary"
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
	fmt.Println("Starting Tellor Miner")
	mineTributes(ds)
}
func mineTributes(ds *dataServer.DataServer) {
	DB := ds.DB
	var nonce string
	var prevCurrentChallenge []byte
	x := 0
	fmt.Println("starting miner")
	for {
		currentChallenge, _ := DB.Get(db.CurrentChallengeKey)
		if bytes.Compare(prevCurrentChallenge, currentChallenge) != 0 {
			prevCurrentChallenge = currentChallenge
			difficulty, _ := DB.Get(db.DifficultyKey)
			data := binary.BigEndian.Uint64(difficulty)
			ndata := big.NewInt(int64(data))
			nonce = pow.SolveChallenge(currentChallenge, ndata)
			if nonce != "" {
				requestID, _ := DB.Get(db.RequestIdKey)
				data = binary.BigEndian.Uint64(requestID)
				rdata := big.NewInt(int64(data))
				value, _ := DB.Get(fmt.Sprint(requestID))
				data = binary.BigEndian.Uint64(value)
				ndata = big.NewInt(int64(data))
				pow.SubmitTransaction(nonce, ndata, rdata)
				nonce = ""
			}
			x++
		} else {
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
