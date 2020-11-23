// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/tellor-io/telliot/pkg/common"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/testutil"
	"github.com/tellor-io/telliot/pkg/util"
)

func TestETHGasStation(t *testing.T) {
	logSetup := util.SetupLogger()
	logger := logSetup("debug")
	tracker := NewGasTracker(logger)
	opts := &rpc.MockOptions{ETHBalance: big.NewInt(300000), Nonce: 1, GasPrice: big.NewInt(7000000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}}
	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "ethGas_test"))
	testutil.Ok(t, err)
	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)
	err = tracker.Exec(ctx)
	testutil.Ok(t, err)
	v, err := DB.Get(db.GasKey)
	testutil.Ok(t, err)

	t.Logf("Gas Price stored: %v\n", string(v))

}

// func TestGas(t *testing.T) {
// 	opts := &rpc.MockOptions{ETHBalance: big.NewInt(300000), Nonce: 1, GasPrice: big.NewInt(7000000000),
// 		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}}
// 	client := rpc.NewMockClientWithValues(opts)

// 	DB, err := db.Open(filepath.Join(os.TempDir(), "test_gas"))
// 	if err != nil {
// 		testutil.Ok(t, err)
// 	}
// 	tracker := &GasTracker{}
// 	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
// 	ctx = context.WithValue(ctx, common.DBContextKey, DB)
// 	err = tracker.Exec(ctx)
// 	if err != nil {
// 		testutil.Ok(t, err)
// 	}
// 	v, err := DB.Get(db.GasKey)
// 	if err != nil {
// 		testutil.Ok(t, err)
// 	}
// 	b, err := hexutil.DecodeBig(string(v))
// 	if err != nil {
// 		testutil.Ok(t, err)
// 	}
// 	t.Logf("Gas PriceStamp stored: %v\n", string(v))
// 	if b.Cmp(big.NewInt(7000000000)) != 0 {
// 		testutil.Ok(t, errors.New(fmt.Sprintf("Balance from client did not match what should have been stored in DB. %s != %s", b, "Should be 1")))
// 	}
// }
