// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/rpc"
	"github.com/tellor-io/TellorMiner/pkg/util"
)

func TestETHGasStation(t *testing.T) {
	tracker := &GasTracker{}
	logger := util.SetupLogger("debug")
	opts := &rpc.MockOptions{ETHBalance: big.NewInt(300000), Nonce: 1, GasPrice: big.NewInt(7000000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}}
	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "ethGas_test"))
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)
	err = tracker.Exec(ctx, logger)
	if err != nil {
		t.Fatal(err)
	}
	v, err := DB.Get(db.GasKey)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Gas Price stored: %v\n", string(v))

}

// func TestGas(t *testing.T) {
// 	opts := &rpc.MockOptions{ETHBalance: big.NewInt(300000), Nonce: 1, GasPrice: big.NewInt(7000000000),
// 		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}}
// 	client := rpc.NewMockClientWithValues(opts)

// 	DB, err := db.Open(filepath.Join(os.TempDir(), "test_gas"))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	tracker := &GasTracker{}
// 	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
// 	ctx = context.WithValue(ctx, common.DBContextKey, DB)
// 	err = tracker.Exec(ctx)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	v, err := DB.Get(db.GasKey)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	b, err := hexutil.DecodeBig(string(v))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("Gas PriceStamp stored: %v\n", string(v))
// 	if b.Cmp(big.NewInt(7000000000)) != 0 {
// 		t.Fatalf("Balance from client did not match what should have been stored in DB. %s != %s", b, "Should be 1")
// 	}
// }
