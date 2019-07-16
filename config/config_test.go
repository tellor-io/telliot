package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	cfg, err := GetConfig()
	if err != nil {
		t.Fatal(err)
	}
	if len(cfg.ContractAddress) == 0 {
		t.Fatal("Config did not parse correctly")
	}
}
