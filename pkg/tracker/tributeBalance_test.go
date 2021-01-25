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

func TestTributeBalance(t *testing.T) {
	cfg := config.OpenTestConfig(t)
	startBal := big.NewInt(456000)
	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: startBal, Top50Requests: []*big.Int{}}
	client := rpc.NewMockClientWithValues(opts)

	DB, cleanup := db.OpenTestDB(t)
	defer t.Cleanup(cleanup)
	proxy, err := db.OpenLocal(cfg, DB)
	testutil.Ok(t, err)
	logger := util.SetupLogger("debug")
	contract, err := contracts.NewTellor(client)
	testutil.Ok(t, err)
	account, err := rpc.NewAccount(cfg)
	testutil.Ok(t, err)
	tracker := NewTributeTracker(logger, proxy, &contract, &account)
	err = tracker.Exec(context.Background())
	testutil.Ok(t, err)
	v, err := proxy.Get(db.TributeBalanceKey)
	testutil.Ok(t, err)
	b, err := hexutil.DecodeBig(string(v))
	testutil.Ok(t, err)
	t.Logf("Tribute Balance stored: %v\n", b)
	if b.Cmp(startBal) != 0 {
		testutil.Ok(t, errors.Errorf("Balance from client did not match what should have been stored in DB. %s != %s", b, startBal))
	}
}
