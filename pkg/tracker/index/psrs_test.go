// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package index

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
	cfg, err := config.OpenTestConfig("../../../")
	testutil.Ok(t, err)
	logger := logging.NewLogger()
	DB, cleanup, err := db.OpenTestDB(cfg)
	testutil.Ok(t, err)
	defer func() {
		testutil.Ok(t, cleanup())
	}()
	proxy, err := db.OpenLocal(logger, cfg, DB)
	testutil.Ok(t, err)
	testClient := rpc.NewMockClient()
	if _, err := BuildIndexTrackers(logger, cfg, proxy, testClient); err != nil {
		testutil.Ok(t, err)
	}
	ethIndexes := indexes["ETH/USD"]

	for _, psr := range ethIndexes {
		err := psr.Exec(context.Background())
		testutil.Ok(t, err)

	}

	_, _, err = MeanAt(ethIndexes, clck.Now(), cfg.Trackers.SleepCycle.Seconds())
	testutil.Ok(t, err)
}
