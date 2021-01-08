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
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/testutil"
	"github.com/tellor-io/telliot/pkg/util"
)

func TestDisputeString(t *testing.T) {
	cfg := config.OpenTestConfig(t)
	DB, cleanup := db.OpenTestDB(t)
	defer t.Cleanup(cleanup)
	logger := util.SetupLogger("debug")
	tracker := NewDisputeTracker(logger, cfg, DB, nil, nil)
	res := tracker.String()
	testutil.Assert(t, res == DisputeTrackerName, "didn't return expected string", DisputeTrackerName)
}

func TestDisputeStatus(t *testing.T) {
	cfg := config.OpenTestConfig(t)
	startBal := big.NewInt(356000)
	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}, DisputeStatus: big.NewInt(1)}
	client := rpc.NewMockClientWithValues(opts)

	DB, cleanup := db.OpenTestDB(t)
	defer t.Cleanup(cleanup)
	logger := util.SetupLogger("debug")
	contract, err := contracts.NewTellor(cfg, client)
	testutil.Ok(t, err)
	account, err := rpc.NewAccount(cfg)
	testutil.Ok(t, err)
	tracker := NewDisputeTracker(logger, cfg, DB, &contract, &account)
	testutil.Ok(t, tracker.Exec(context.Background()))
	v, err := DB.Get(db.DisputeStatusKey)
	testutil.Ok(t, err)
	b, err := hexutil.DecodeBig(string(v))
	testutil.Ok(t, err)
	t.Logf("Dispute Status stored: %v\n", string(v))
	testutil.Equals(t, b.Cmp(big.NewInt(1)), 0, "dispute status from client did not match what should have been stored in DB. %s != %s", b, "one")
	DB.Close()
}

func TestDisputeStatusNegativeBalance(t *testing.T) {
	cfg := config.OpenTestConfig(t)
	startBal := big.NewInt(356000)
	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}, DisputeStatus: big.NewInt(1)}
	client := rpc.NewMockClientWithValues(opts)

	DB, cleanup := db.OpenTestDB(t)
	defer t.Cleanup(cleanup)
	logger := util.SetupLogger("debug")
	contract, err := contracts.NewTellor(cfg, client)
	testutil.Ok(t, err)
	account, err := rpc.NewAccount(cfg)
	testutil.Ok(t, err)
	tracker := NewDisputeTracker(logger, cfg, DB, &contract, &account)
	testutil.Ok(t, tracker.Exec(context.Background()))
	v, err := DB.Get(db.DisputeStatusKey)
	testutil.Ok(t, err)
	b, err := hexutil.DecodeBig(string(v))
	testutil.Ok(t, err)
	t.Logf("Dispute Status stored: %v\n", string(v))
	if b.Cmp(big.NewInt(1)) != 0 {
		testutil.Ok(t, errors.Errorf("Dispute Status from client did not match what should have been stored in DB. %s != %s", b, "one"))
	}
}
