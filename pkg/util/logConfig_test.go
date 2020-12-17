// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package util

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/testutil"
)

func TestLogConfig(t *testing.T) {
	var defaultEntries = []Entry{
		{Component: "config.Config", Level: "INFO"},
		{Component: "db.DB", Level: "INFO"},
		{Component: "rpc.client", Level: "INFO"},
		{Component: "rpc.ABICodec", Level: "INFO"},
		{Component: "rpc.mockClient", Level: "INFO"},
		{Component: "tracker.Top50Tracker", Level: "INFO"},
		{Component: "tracker.FetchDataTracker", Level: "INFO"},
		{Component: "pow.MiningWorker-0", Level: "INFO"},
		{Component: "pow.MiningWorker-1", Level: "INFO"},
		{Component: "pow.MiningTasker-0", Level: "INFO"},
		{Component: "pow.MiningTasker-1", Level: "INFO"},
		{Component: "tracker.PSRTracker", Level: "INFO"},
	}
	err := SetupLoggingConfig(defaultEntries)
	testutil.Ok(t, err)
	cfg := GetLoggingConfig()
	if cfg.levels["config.Config"] == 0 {
		testutil.Ok(t, errors.Errorf("Config did not parse correctly: %v", cfg.levels))
	} else {
		t.Logf("Parsed log level: %d", cfg.levels["config.Config"])
	}
}
