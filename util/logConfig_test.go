package util

import (
	"testing"
)

func TestLogConfig(t *testing.T) {
	path := "./testConfig.json"
	err := ParseLoggingConfig(path)
	if err != nil {
		t.Fatal(err)
	}
	cfg,_ := GetLoggingConfig()
	if cfg.levels["config.Config"] == 0 {
		t.Fatalf("Config did not parse correctly: %v", cfg.levels)
	} else {
		t.Logf("Parsed log level: %d", cfg.levels["config.Config"])
	}
}
