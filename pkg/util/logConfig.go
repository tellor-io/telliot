// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package util

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

// Entry holds specific component log level.
type Entry struct {
	Level     string `json:"level"`
	Component string `json:"component"`
}

// LogConfig holds individual log level settings.
type LogConfig struct {
	levels map[string]LogLevel
}

var (
	sharedConfig *LogConfig
)

// ParseLoggingConfig parses the given JSON log level config file for use in log configuration.
func ParseLoggingConfig(file string) error {

	if len(file) > 0 {
		info, err := os.Stat(file)
		if os.IsNotExist(err) {
			return errors.Wrapf(err, "loggingConfigPath references an invalid file at: %s", file)
		}
		if info.IsDir() {
			return errors.Wrapf(err, "logging config file %s is a directory", file)
		}

		configFile, err := os.Open(file)
		defer func() {
			err := configFile.Close()
			if err != nil {
				fmt.Print("error closing the file", err)
			}
		}()
		if err != nil {
			return err
		}
		var entries []Entry

		dec := json.NewDecoder(configFile)
		err = dec.Decode(&entries)
		if err != nil {
			return err
		}
		cfg := &LogConfig{make(map[string]LogLevel)}
		for _, e := range entries {
			lvl, err := StringToLevel(e.Level)
			if err != nil {
				return err
			}
			cfg.levels[e.Component] = lvl
		}
		sharedConfig = cfg
	} else {
		sharedConfig = &LogConfig{make(map[string]LogLevel)}
	}
	// Initialize all the loggers that have already been declared as global vars.
	initLoggers(sharedConfig)
	return nil
}

// GetLoggingConfig retrieves a shared logging config.
func GetLoggingConfig() *LogConfig {
	return sharedConfig
}

// GetLevel the log level.
func (cfg *LogConfig) GetLevel(pkg string, component string) LogLevel {
	key := pkg + "." + component
	lvl := cfg.levels[key]
	if lvl == 0 {
		return InfoLogLevel
	}
	return lvl
}
