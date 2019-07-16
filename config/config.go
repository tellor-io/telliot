package config

import (
	"encoding/json"
	"os"

	"github.com/tellor-io/TellorMiner/cli"
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
func ParseConfig(path string) (*Config, error) {
	if len(path) == 0 {
		path = cli.GetFlags().ConfigPath
		if len(path) == 0 {
			panic("Invalid config path. Not provided and not a command line option")
		}
	}
	configFile, err := os.Open(path)
	defer configFile.Close()
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(configFile)
	err = dec.Decode(&config)
	config.logger = util.NewLogger("config", "Config", util.InfoLogLevel)
	config.logger.Info("config: %+v", config)
	return config, nil
}

//GetConfig returns a shared instance of config
func GetConfig() (*Config, error) {
	if config == nil {
		_, err := ParseConfig("")
		if err != nil {
			return nil, err
		}
		return config, nil
	}
	return config, nil
}
