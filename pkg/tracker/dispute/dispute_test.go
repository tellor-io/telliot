// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package dispute

import (
	"context"
	stdlog "log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/testutil"
	"github.com/tellor-io/telliot/pkg/tracker/index"
)

func TestMain(m *testing.M) {

	cfg, err := config.OpenTestConfig("../../../")
	if err != nil {
		stdlog.Fatal(err)
	}

	logger := logging.NewLogger()
	DB, cleanup, err := db.OpenTestDB(cfg)
	if err != nil {
		stdlog.Fatal(err)
	}
	client := rpc.NewMockClient()
	defer func() {
		if err := cleanup(); err != nil {
			stdlog.Fatal(err)
		}
	}()
	proxy, err := db.OpenLocal(logging.NewLogger(), cfg, DB)
	if err != nil {
		stdlog.Fatal(err)
	}

	if _, err := index.BuildIndexTrackers(logger, cfg, proxy, client); err != nil {
		stdlog.Fatal(err)
	}
	os.Exit(m.Run())

}

func TestDisputeCheckerInRange(t *testing.T) {
	client := rpc.NewMockClient()
	contract, err := contracts.NewITellor(client)
	testutil.Ok(t, err)
	ctx := context.Background()
	ethUSDPairs := index.GetIndexes()["ETH/USD"]
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	time.Sleep(2 * time.Second)
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	cfg, err := config.OpenTestConfig("../../../")
	if err != nil {
		stdlog.Fatal(err)
	}
	disputeChecker := &disputeChecker{lastCheckedBlock: 500, config: cfg, logger: log.NewNopLogger(), client: client, contract: contract}
	testutil.Ok(t, disputeChecker.Exec(ctx))
}

func TestDisputeCheckerOutOfRange(t *testing.T) {
	cfg, err := config.OpenTestConfig("../../../")
	testutil.Ok(t, err)

	logger := logging.NewLogger()
	client := rpc.NewMockClient()
	contract, err := contracts.NewITellor(client)
	testutil.Ok(t, err)

	disputeChecker := NewDisputeChecker(logger, cfg, client, contract, 500)

	ethUSDPairs := index.GetIndexes()["ETH/USD"]
	ctx := context.Background()
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	time.Sleep(2 * time.Second)
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	testutil.Ok(t, disputeChecker.Exec(ctx))

	files, err := filepath.Glob("possible-dispute-*.txt")
	if err != nil {
		panic(err)
	}
	testutil.Assert(t, len(files) >= 1, "expected a possible-dispute file")

	for _, f := range files {
		testutil.Ok(t, os.Remove(f))
	}
}

func execEthUsdPsrs(ctx context.Context, t *testing.T, psrs []*index.IndexTracker) {
	for _, psr := range psrs {
		err := psr.Exec(ctx)

		testutil.Ok(t, err)

	}
}
