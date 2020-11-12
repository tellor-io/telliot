// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package config

import (
	"errors"
	"os"
	"testing"

	"github.com/tellor-io/TellorMiner/pkg/testutil"
)

func createEnvFile(t *testing.T) func() {
	f, err := os.Create(".env")
	if err != nil {
		testutil.Ok(t, err)
	}
	_, err = f.WriteString("ETH_PRIVATE_KEY=\"0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef\"")
	if err != nil {
		f.Close()
		testutil.Ok(t, err)
	}

	return func() {
		os.Remove(".env")
	}
}

func TestConfig(t *testing.T) {
	//Creating a mock .ENV file to go around this issue with godotenv:
	//https://github.com/joho/godotenv/issues/43
	cleanup := createEnvFile(t)
	defer t.Cleanup(cleanup)

	cfg := OpenTestConfig(t)

	//Asserting Default Values
	if cfg.GasMax == 0 {
		testutil.Ok(t, errors.New("GasMax should have value"))
	}
	if cfg.GasMultiplier == 0 {
		testutil.Ok(t, errors.New("GasMultiplier should have value"))
	}
	if cfg.MinConfidence == 0 {
		testutil.Ok(t, errors.New("MinConfidence should have value"))
	}
	if cfg.DisputeThreshold == 0 {
		testutil.Ok(t, errors.New("DisputeThreshold should have value"))
	}
}
