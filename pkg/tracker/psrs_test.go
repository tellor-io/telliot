// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"testing"

	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/testutil"
)

func TestMeanAt(t *testing.T) {
	cfg := config.OpenTestConfig(t)
	DB, cleanup := db.OpenTestDB(t)
	testClient := rpc.NewMockClient()
	defer t.Cleanup(cleanup)
	if _, err := BuildIndexTrackers(cfg, DB, testClient); err != nil {
		testutil.Ok(t, err)
	}
	ethIndexes := indexes["ETH/USD"]
	execEthUsdPsrs(context.Background(), t, ethIndexes)

	MeanAt(ethIndexes, clck.Now())
}
