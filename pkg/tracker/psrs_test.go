// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"testing"

	"github.com/tellor-io/TellorMiner/pkg/tcontext"
	"github.com/tellor-io/TellorMiner/pkg/testutil"
)

func TestMeanAt(t *testing.T) {
	ctx, _, cleanup := tcontext.CreateTestContext(t)
	t.Cleanup(cleanup)
	if _, err := BuildIndexTrackers(); err != nil {
		testutil.Ok(t, err)
	}
	ethIndexes := indexes["ETH/USD"]
	execEthUsdPsrs(ctx, t, ethIndexes)

	MeanAt(ethIndexes, clck.Now())
}
