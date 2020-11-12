// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package util

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"
	"github.com/tellor-io/TellorMiner/pkg/testutil"
)

func TestLogConfig(t *testing.T) {
	path := filepath.Join("..", "..", "configs", "loggingConfig.json")
	err := ParseLoggingConfig(path)
	if err != nil {
		testutil.Ok(t, err)
		t.Fatal(err)
	}
	cfg := GetLoggingConfig()
	if cfg.levels["config.Config"] == 0 {
		testutil.Ok(t, errors.New(fmt.Sprintf("Config did not parse correctly: %v", cfg.levels)))
	} else {
		t.Logf("Parsed log level: %d", cfg.levels["config.Config"])
	}
}
