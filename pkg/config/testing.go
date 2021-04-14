// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package config

import (
	"os"
	"path/filepath"

	"github.com/phayes/freeport"
)

func OpenTestConfig(nestedLevel string) (*Config, error) {
	projectPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	rootDir := filepath.Join(projectPath, nestedLevel)
	cfg := defaultConfig
	port, err := freeport.GetFreePort()
	if err != nil {
		return nil, err
	}
	cfg.Mine.ListenPort = uint(port)
	// Don't need any trackers for the tests.
	cfg.Trackers.Names = make(map[string]bool)

	cfg.ApiFile = filepath.Join(rootDir, cfg.ApiFile)
	cfg.EnvFile = filepath.Join(rootDir, cfg.EnvFile+".example")
	cfg.ManualDataFile = filepath.Join(rootDir, cfg.ManualDataFile)
	cfg.HistoryFile = filepath.Join(rootDir, cfg.HistoryFile)

	return Populate(&cfg)

}
