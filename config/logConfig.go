package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tellor-io/TellorMiner/util"
)

//Entry holds specific component log level
type Entry struct {
	Level     string `json:"level"`
	Component string `json:"component"`
}

//LogConfig holds individual log level settings
type LogConfig struct {
	Levels map[string]util.LogLevel
}

var (
	sharedConfig *LogConfig
)

//ParseLoggingConfig parses the given JSON log level config file for use in log configuration
func ParseLoggingConfig(file string) (*LogConfig, error) {
	if len(file) == 0 {
		return nil, fmt.Errorf("Invalid log config file")
	}
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return nil, err
	}
	var entries []Entry
	sharedConfig = &LogConfig{make(map[string]util.LogLevel)}

	dec := json.NewDecoder(configFile)
	err = dec.Decode(&entries)
	if err != nil {
		return nil, err
	}
	for _, e := range entries {
		lvl, err := util.StringToLevel(e.Level)
		if err != nil {
			return nil, err
		}
		sharedConfig.Levels[e.Component] = lvl
	}
	return sharedConfig, nil
}

//GetLoggingConfig retrieves a shared logging config
func GetLoggingConfig() (*LogConfig, error) {
	if sharedConfig == nil {
		return nil, fmt.Errorf("Logging config was not initialized ")
	}
	return sharedConfig, nil
}
