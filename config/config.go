package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tellor-io/TellorMiner/util"
)

//Config holds global config info derived from config.json
type Config struct {
	ContractAddress   string   `json:"contractAddress"`
	NodeURL           string   `json:"nodeURL"`
	PrivateKey        string   `json:"privateKey"`
	DatabaseURL       string   `json:"databaseURL"`
	PublicAddress     string   `json:"publicAddress"`
	EthClientTimeout  uint     `json:"ethClientTimeout"`
	TrackerSleepCycle uint     `json:"trackerCycle"` //in seconds
	Trackers          []string `json:"trackers"`
	DBFile            string   `json:"dbFile"`
	ServerHost        string   `json:"serverHost"`
	ServerPort        uint     `json:"serverPort"`
	logger            *util.Logger
}

var config *Config

//ParseConfig and set a shared config entry
func ParseConfig(path string) {
	if len(path) == 0 {
		panic("Invalid config path")
	}
	configFile, err := os.Open(path)
	defer configFile.Close()
	if err != nil {
		panic(err.Error())
	}
	dec := json.NewDecoder(configFile)
	err = dec.Decode(&config)
	config.logger = util.NewLogger("config", "Config", util.InfoLogLevel)
	config.logger.Info("config: %+v", config)
}

//GetConfig returns a shared instance of config
func GetConfig() (*Config, error) {
	if config == nil {
		return nil, fmt.Errorf("Config was not initialized")
	}
	return config, nil
}
