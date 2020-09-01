package tracker

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/math"

	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rpc"
)

func TestRunner(t *testing.T) {
	exitCh := make(chan int)
	
	startBal := big.NewInt(356000)

	hash := math.PaddedBigBytes(big.NewInt(256), 32)
	var b32 [32]byte
	for i, v := range hash {
		b32[i] = v
	}

	top50 := make([]*big.Int, 51)
	mockQueryParams := &rpc.MockQueryMeta{QueryString: "json(https://api.binance.com/api/v1/klines?symbol=ETHBTC&interval=1d&limit=1).0.4", Granularity: 1000}
	paramsMap := make(map[uint]*rpc.MockQueryMeta)
	for i := range top50 {
		top50[i] = big.NewInt(int64(i + 51))
		paramsMap[uint(i+51)] = mockQueryParams
	}

	queryStr := "json(https://coinbase.com)"
	chal := &rpc.CurrentChallenge{ChallengeHash: b32, RequestID: big.NewInt(1),
		Difficulty: big.NewInt(500), QueryString: queryStr,
		Granularity: big.NewInt(1000), Tip: big.NewInt(0)}
	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: big.NewInt(0),MiningStatus:true, Top50Requests:top50, CurrentChallenge: chal,DisputeStatus: big.NewInt(1),QueryMetadata: paramsMap}
	client := rpc.NewMockClientWithValues(opts)
	
	db, err := db.Open(filepath.Join(os.TempDir(), "test_leveldb"))
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	runner, _ := NewRunner(client, db)

	ctx := context.Background()
	runner.Ready()
	runner.Start(ctx, exitCh)
	fmt.Println("runner done")
	time.Sleep(2 * time.Second)
	close(exitCh)
	time.Sleep(1 * time.Second)
}
