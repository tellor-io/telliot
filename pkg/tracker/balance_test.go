// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/rpc"
	"github.com/tellor-io/TellorMiner/pkg/util"
)

func TestStringId(t *testing.T) {
	tracker := NewBalanceTracker(util.SetupLogger("debug"))
	res := tracker.String()
	if res != BalanceTrackerName {
		t.Fatal("didn't return expected string", BalanceTrackerName)
	}
}
func TestPositiveBalance(t *testing.T) {
	startBal := big.NewInt(356000)
	dbBalanceTest(startBal, t)
}

func TestZeroBalance(t *testing.T) {
	startBal := big.NewInt(0)
	dbBalanceTest(startBal, t)
}

func TestNegativeBalance(t *testing.T) {
	startBal := big.NewInt(-753)
	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}}
	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_balance"))
	if err != nil {
		t.Fatal(err)
	}
	tracker := NewBalanceTracker(util.SetupLogger("debug"))
	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)
	err = tracker.Exec(ctx)
	if err == nil {
		t.Fatal(err)
	}
}

func dbBalanceTest(startBal *big.Int, t *testing.T) {
	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}}
	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_balance"))
	if err != nil {
		t.Fatal(err)
	}
	tracker := NewBalanceTracker(util.SetupLogger("debug"))
	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)
	err = tracker.Exec(ctx)
	if err != nil {
		t.Fatal(err)
	}
	v, err := DB.Get(db.BalanceKey)
	if err != nil {
		t.Fatal(err)
	}
	b, err := hexutil.DecodeBig(string(v))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Balance stored: %v\n", string(v))
	if b.Cmp(startBal) != 0 {
		t.Fatalf("Balance from client did not match what should have been stored in DB. %s != %s", b, startBal)
	}
	DB.Close()
}
