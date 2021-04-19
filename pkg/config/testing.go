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
	cfg.Web.ListenPort = uint(port)

	cfg.IndexTracker.ApiFile = filepath.Join(rootDir, cfg.IndexTracker.ApiFile)
	cfg.EnvFile = filepath.Join(rootDir, cfg.EnvFile+".example")
	cfg.IndexTracker.ManualDataFile = filepath.Join(rootDir, cfg.IndexTracker.ManualDataFile)

	return &cfg, nil

}
