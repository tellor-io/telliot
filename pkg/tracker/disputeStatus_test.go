// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/rpc"
	"github.com/tellor-io/TellorMiner/pkg/testutil"
	"github.com/tellor-io/TellorMiner/pkg/util"
)

func TestDisputeString(t *testing.T) {
	logSetup := util.SetupLogger()
	logger := logSetup("debug")
	tracker := NewDisputeTracker(logger)
	res := tracker.String()
	if res != DisputeTrackerName {
		testutil.Ok(t, errors.New(fmt.Sprintf("didn't return expected string", DisputeTrackerName)))
	}
}

func TestDisputeStatus(t *testing.T) {
	startBal := big.NewInt(356000)
	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}, DisputeStatus: big.NewInt(1)}
	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_disputeStatus"))
	if err != nil {
		testutil.Ok(t, err)
	}
	logSetup := util.SetupLogger()
	logger := logSetup("debug")
	tracker := NewDisputeTracker(logger)
	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)
	err = tracker.Exec(ctx)
	if err != nil {
		testutil.Ok(t, err)
	}
	v, err := DB.Get(db.DisputeStatusKey)
	if err != nil {
		testutil.Ok(t, err)
	}
	b, err := hexutil.DecodeBig(string(v))
	if err != nil {
		testutil.Ok(t, err)
	}
	t.Logf("Dispute Status stored: %v\n", string(v))
	if b.Cmp(big.NewInt(1)) != 0 {
		testutil.Ok(t, errors.New(fmt.Sprintf("Dispute Status from client did not match what should have been stored in DB. %s != %s", b, "one")))
	}
	DB.Close()
}

func TestDisputeStatusNegativeBalance(t *testing.T) {
	startBal := big.NewInt(356000)
	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}, DisputeStatus: big.NewInt(0)}
	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_disputeStatus"))
	if err != nil {
		testutil.Ok(t, err)
	}
	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	context.WithValue(ctx, common.DBContextKey, DB)

	v, err := DB.Get(db.DisputeStatusKey)
	if err != nil {
		testutil.Ok(t, err)
	}
	b, err := hexutil.DecodeBig(string(v))
	if err != nil {
		testutil.Ok(t, err)
	}
	t.Logf("Dispute Status stored: %v\n", string(v))
	if b.Cmp(big.NewInt(1)) != 0 {
		testutil.Ok(t, errors.New(fmt.Sprintf("Dispute Status from client did not match what should have been stored in DB. %s != %s", b, "one")))
	}
}
