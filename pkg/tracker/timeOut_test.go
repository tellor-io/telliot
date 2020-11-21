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

func TestTimeOutString(t *testing.T) {
	logSetup := util.SetupLogger()
	logger := logSetup("debug")
	tracker := NewTimeOutTracker(logger)
	res := tracker.String()

	testutil.Assert(t, res == "TimeOutTracker", "should return 'TimeOutTracker' string")

}

func TestTimeOutTracker(t *testing.T) {

	db, err := db.Open(filepath.Join(os.TempDir(), "test_timeOut"))
	testutil.Ok(t, err)

	startBal := big.NewInt(456000)
	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: startBal, Top50Requests: []*big.Int{}}
	client := rpc.NewMockClientWithValues(opts)
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, db)

	logSetup := util.SetupLogger()
	logger := logSetup("debug")
	tracker := NewTimeOutTracker(logger)
	if err := tracker.Exec(ctx); err != nil {
		testutil.Ok(t, err)
	}
}
