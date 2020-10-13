// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package util

import (
	"path/filepath"
	"testing"
)

func TestLogConfig(t *testing.T) {
	path := filepath.Join("..", "configs", "loggingConfig.json")
	err := ParseLoggingConfig(path)
	if err != nil {
		t.Fatal(err)
	}
	cfg := GetLoggingConfig()
	if cfg.levels["config.Config"] == 0 {
		t.Fatalf("Config did not parse correctly: %v", cfg.levels)
	} else {
		t.Logf("Parsed log level: %d", cfg.levels["config.Config"])
	}
}
