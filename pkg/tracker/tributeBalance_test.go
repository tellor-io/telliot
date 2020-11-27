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

func TestTributeBalance(t *testing.T) {
	startBal := big.NewInt(456000)
	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: startBal, Top50Requests: []*big.Int{}}
	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_Tributebalance"))
	testutil.Ok(t, err)
	logSetup := util.SetupLogger()
	logger := logSetup("debug")
	tracker := NewTributeTracker(logger)
	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)
	err = tracker.Exec(ctx)
	testutil.Ok(t, err)
	v, err := DB.Get(db.TributeBalanceKey)
	testutil.Ok(t, err)
	b, err := hexutil.DecodeBig(string(v))
	testutil.Ok(t, err)
	t.Logf("Tribute Balance stored: %v\n", b)
	if b.Cmp(startBal) != 0 {
		testutil.Ok(t, errors.Errorf("Balance from client did not match what should have been stored in DB. %s != %s", b, startBal))
	}
}
