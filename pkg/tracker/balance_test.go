// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"math/big"
	"testing"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/testutil"
	"github.com/tellor-io/telliot/pkg/util"
)

func TestStringId(t *testing.T) {
	client := rpc.NewMockClient()
	config.OpenTestConfig(t)
	DB, cleanup := db.OpenTestDB(t)
	defer t.Cleanup(cleanup)
	cfg := config.OpenTestConfig(t)
	proxy, err := db.OpenLocal(cfg, DB)
	testutil.Ok(t, err)
	logger := util.SetupLogger("debug")
	tracker := NewBalanceTracker(logger, proxy, client, nil)
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
	cfg := config.OpenTestConfig(t)
	startBal := big.NewInt(-753)
	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}}
	client := rpc.NewMockClientWithValues(opts)

	DB, cleanup := db.OpenTestDB(t)
	defer t.Cleanup(cleanup)
	proxy, err := db.OpenLocal(cfg, DB)
	testutil.Ok(t, err)

	logger := util.SetupLogger("debug")
	account, err := rpc.NewAccount(cfg)
	testutil.Ok(t, err)
	tracker := NewBalanceTracker(logger, proxy, client, &account)
	err = tracker.Exec(context.Background())
	testutil.NotOk(t, err, "should have error")
}

func dbBalanceTest(startBal *big.Int, t *testing.T) {
	cfg := config.OpenTestConfig(t)
	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}}
	client := rpc.NewMockClientWithValues(opts)

	DB, cleanup := db.OpenTestDB(t)
	defer t.Cleanup(cleanup)
	proxy, err := db.OpenLocal(cfg, DB)
	testutil.Ok(t, err)
	logger := util.SetupLogger("debug")
	account, err := rpc.NewAccount(cfg)
	testutil.Ok(t, err)
	tracker := NewBalanceTracker(logger, proxy, client, &account)
	err = tracker.Exec(context.Background())
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
