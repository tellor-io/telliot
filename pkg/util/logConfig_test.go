// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package util

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/testutil"
)

func TestLogConfig(t *testing.T) {
	var defaultEntries = map[string]string{
		"config.Config":            "INFO",
		"db.DB":                    "INFO",
		"rpc.client":               "INFO",
		"rpc.ABICodec":             "INFO",
		"rpc.mockClient":           "INFO",
		"tracker.Top50Tracker":     "INFO",
		"tracker.FetchDataTracker": "INFO",
		"pow.MiningWorker-0:":      "INFO",
		"pow.MiningWorker-1:":      "INFO",
		"pow.MiningTasker-0:":      "INFO",
		"pow.MiningTasker-1:":      "INFO",
		"tracker.PSRTracker":       "INFO",
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
