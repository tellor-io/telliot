package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	cfg := GetConfig()
	if len(cfg.ContractAddress) == 0 {
		t.Fatal("Config did not parse correctly")
	}
}
