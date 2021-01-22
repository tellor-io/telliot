// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"testing"

	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/testutil"
	"github.com/tellor-io/telliot/pkg/util"
)

func TestTimeOutString(t *testing.T) {
	cfg := config.OpenTestConfig(t)
	DB, cleanup := db.OpenTestDB(t)
	defer t.Cleanup(cleanup)
	proxy, err := db.OpenLocal(cfg, DB)
	testutil.Ok(t, err)
	logger := util.SetupLogger("debug")

	tracker := NewTimeOutTracker(logger, cfg, proxy, nil, nil)
	res := tracker.String()

	testutil.Assert(t, res == "TimeOutTracker", "should return 'TimeOutTracker' string")

}

//Can't make the test compile with current setup. Need a contract.Getter connected to the client, but didn't managed to compile this
// func TestTimeOutTracker(t *testing.T) {

// 	DB, cleanup := db.OpenTestDB(t)
// 	defer t.Cleanup(cleanup)

// 	startBal := big.NewInt(456000)
// 	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
// 		TokenBalance: startBal, Top50Requests: []*big.Int{}}
// 	client := rpc.NewMockClientWithValues(opts)

// 	add := commom.Address
// 	contractTellorInstance, err := tellor.NewTellor(common.Address, client)
// 	testutil.Ok(t, err)

// 	contractGetterInstance, err := getter.NewTellorGetters(contractAddress, client)

// 	ctx := context.Background()
// 	// ctx = context.WithValue(ctx, common.ClientContextKey, client)
// 	// ctx = context.WithValue(ctx, common.DBContextKey, db)

// 	contract := contracts.Tellor{}
// 	account := rpc.Account{}
// 	logSetup := util.SetupLogger()
// 	logger := logSetup("debug")
// 	tracker := NewTimeOutTracker(logger, DB, contract, account)
// 	if err := tracker.Exec(ctx); err != nil {
// 		testutil.Ok(t, err)
// 	}
// }
