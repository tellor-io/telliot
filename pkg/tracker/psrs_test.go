// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"testing"

	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/testutil"
)

func TestMeanAt(t *testing.T) {
	cfg := config.OpenTestConfig(t)
	logger := logging.NewLogger()
	DB, cleanup := db.OpenTestDB(t)
	proxy, err := db.OpenLocal(logger, cfg, DB)
	testutil.Ok(t, err)
	testClient := rpc.NewMockClient()
	defer t.Cleanup(cleanup)
	if _, err := BuildIndexTrackers(logger, cfg, proxy, testClient); err != nil {
		testutil.Ok(t, err)
	}
	ethIndexes := indexes["ETH/USD"]
	execEthUsdPsrs(context.Background(), t, ethIndexes)

	_, _, err = MeanAt(ethIndexes, clck.Now(), cfg.Trackers.SleepCycle.Seconds())
	testutil.Ok(t, err)
}
