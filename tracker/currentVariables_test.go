package tracker

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	tellor "github.com/tellor-io/TellorMiner/contracts"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rpc"
)

func TestCurrentVarableString(t *testing.T) {
	tracker := &CurrentVariablesTracker{}
	res := tracker.String()
	if res != "CurrentVariablesTracker" {
		t.Fatalf("should return 'CurrentVariablesTracker' string")
	}
}

func TestCurrentVariables(t *testing.T) {

	startBal := big.NewInt(356000)

	hash := math.PaddedBigBytes(big.NewInt(256), 32)
	var b32 [32]byte
	for i, v := range hash {
		b32[i] = v
	}
	queryStr := "json(https://coinbase.com)"
	chal := &rpc.CurrentChallenge{ChallengeHash: b32, RequestID: big.NewInt(1),
		Difficulty: big.NewInt(500), QueryString: queryStr,
		Granularity: big.NewInt(1000), Tip: big.NewInt(0)}
	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: big.NewInt(0), MiningStatus:true,Top50Requests: []*big.Int{}, CurrentChallenge: chal}
	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_currentVariables"))
	if err != nil {
		t.Fatal(err)
	}
	cfg := config.GetConfig()
	tracker := &CurrentVariablesTracker{}
	ctx := context.WithValue(context.Background(), tellorCommon.ClientContextKey, client)
	ctx = context.WithValue(ctx, tellorCommon.DBContextKey, DB)

	masterInstance := ctx.Value(tellorCommon.MasterContractContextKey)
	if masterInstance == nil {
		contractAddress := common.HexToAddress(cfg.ContractAddress)
		// TODO create error state flag for mock client
		masterInstance, err = tellor.NewTellorMaster(contractAddress,client)
		if err != nil {
			runnerLog.Error("Problem creating tellor master instance: %v\n", err)
			return
		}
		ctx = context.WithValue(ctx, tellorCommon.MasterContractContextKey, masterInstance)
	}
	
	fmt.Println("Working to Line 41")
	err = tracker.Exec(ctx)
	if err != nil {
		t.Fatal(err)
	}
	v, err := DB.Get(db.RequestIdKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Working to Line 51", v)
	b, err := hexutil.DecodeBig(string(v))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("RequestID stored: %v\n", string(v))
	if b.Cmp(big.NewInt(1)) != 0 {
		t.Fatalf("Current Request ID from client did not match what should have been stored in DB. %s != %s", b, string(1))
	}

	v, err = DB.Get(db.QueryStringKey)
	if err != nil {
		t.Fatal(err)
	}
	if string(v) != queryStr {
		t.Fatalf("Expected query string to match test input: %s != %s\n", string(v), queryStr)
	}

}
