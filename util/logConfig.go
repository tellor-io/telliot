package util

import (
	"encoding/json"
	"log"
	"os"

	"github.com/tellor-io/TellorMiner/cli"
)

//Entry holds specific component log level
type Entry struct {
	Level     string `json:"level"`
	Component string `json:"component"`
}

//LogConfig holds individual log level settings
type LogConfig struct {
	levels map[string]LogLevel
}

var (
	sharedConfig *LogConfig
)

//ParseLoggingConfig parses the given JSON log level config file for use in log configuration
func ParseLoggingConfig(file string) (*LogConfig, error) {
	if len(file) == 0 {
		lCfg := cli.GetFlags()
		file = lCfg.LoggingConfigPath
	}
	if sharedConfig != nil {
		return sharedConfig, nil
	}

	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		log.Fatalf("LoggingConfigPath references an invalid file at: %s", file)
	}
	if info.IsDir() {
		log.Fatalf("Logging config file %s is a directory", file)
	}

	if len(file) > 0 {
		configFile, err := os.Open(file)
		defer configFile.Close()
		if err != nil {
			return nil, err
		}
		var entries []Entry

		dec := json.NewDecoder(configFile)
		err = dec.Decode(&entries)
		if err != nil {
			return nil, err
		}
		cfg := &LogConfig{make(map[string]LogLevel)}
		for _, e := range entries {
			lvl, err := StringToLevel(e.Level)
			if err != nil {
				return nil, err
			}
			cfg.levels[e.Component] = lvl
		}
		sharedConfig = cfg
	} else {
		sharedConfig = &LogConfig{make(map[string]LogLevel)}
	}
	return sharedConfig, nil
}

//GetLoggingConfig retrieves a shared logging config
func GetLoggingConfig() (*LogConfig, error) {
	if sharedConfig == nil {
		_, err := ParseLoggingConfig("")
		if err != nil {
			return nil, err
		}
	}
	return sharedConfig, nil
}

//GetLevel the log level
func (cfg *LogConfig) GetLevel(pkg string, component string) LogLevel {
	key := pkg + "." + component
	lvl := cfg.levels[key]
	if lvl == 0 {
		return InfoLogLevel
	}
	return lvl
}
