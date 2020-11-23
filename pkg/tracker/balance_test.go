// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tellor-io/telliot/pkg/common"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/testutil"
	"github.com/tellor-io/telliot/pkg/util"
)

func TestStringId(t *testing.T) {
	logSetup := util.SetupLogger()
	logger := logSetup("debug")
	tracker := NewBalanceTracker(logger)
	res := tracker.String()

	testutil.Equals(t, res, BalanceTrackerName, "didn't return expected string", BalanceTrackerName)
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
	testutil.Ok(t, err)
	logSetup := util.SetupLogger()
	logger := logSetup("debug")
	tracker := NewBalanceTracker(logger)
	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)
	err = tracker.Exec(ctx)
	testutil.NotOk(t, err, "should have error")
}

func dbBalanceTest(startBal *big.Int, t *testing.T) {
	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}}
	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_balance"))
	testutil.Ok(t, err)

	logSetup := util.SetupLogger()
	logger := logSetup("debug")
	tracker := NewBalanceTracker(logger)
	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)
	err = tracker.Exec(ctx)
	testutil.Ok(t, err)
	v, err := DB.Get(db.BalanceKey)
	testutil.Ok(t, err)
	b, err := hexutil.DecodeBig(string(v))
	testutil.Ok(t, err)
	t.Logf("Balance stored: %v\n", string(v))
	if b.Cmp(startBal) != 0 {
		testutil.Ok(t, errors.Errorf("Balance from client did not match what should have been stored in DB. %s != %s", b, startBal))
	}
	DB.Close()
}
