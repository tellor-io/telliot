// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package util

import (
	"path/filepath"
	"testing"

	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/testutil"
)

func TestLogConfig(t *testing.T) {
	path := filepath.Join("..", "..", "configs", "loggingConfig.json")
	err := ParseLoggingConfig(path)
	testutil.Ok(t, err)
	cfg := GetLoggingConfig()
	if cfg.levels["config.Config"] == 0 {
		testutil.Ok(t, errors.Errorf("Config did not parse correctly: %v", cfg.levels))
	} else {
		t.Logf("Parsed log level: %d", cfg.levels["config.Config"])
	}
}
