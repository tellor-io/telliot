// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package config

import (
	"os"
	"path/filepath"

	"github.com/phayes/freeport"
)

func OpenTestConfig(dir string) (*Config, error) {
	projectPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	rootDir := filepath.Join(projectPath, dir)
	cfg := DefaultConfig
	port, err := freeport.GetFreePort()
	if err != nil {
		return nil, err
	}
	cfg.Web.ListenPort = uint(port)

	cfg.IndexTracker.IndexFile = filepath.Join(rootDir, cfg.IndexTracker.IndexFile)
	cfg.EnvFile = filepath.Join(rootDir, cfg.EnvFile+".example")
	cfg.Aggregator.ManualDataFile = filepath.Join(rootDir, cfg.Aggregator.ManualDataFile)

	return &cfg, nil

}
