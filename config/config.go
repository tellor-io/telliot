package config

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/tellor-io/TellorMiner/cli"
	"github.com/tellor-io/TellorMiner/util"
)

//Config holds global config info derived from config.json
type Config struct {
	ContractAddress              string        `json:"contractAddress"`
	NodeURL                      string        `json:"nodeURL"`
	PrivateKey                   string        `json:"privateKey"`
	DatabaseURL                  string        `json:"databaseURL"`
	PublicAddress                string        `json:"publicAddress"`
	EthClientTimeout             uint          `json:"ethClientTimeout"`
	TrackerSleepCycle            uint          `json:"trackerCycle"` //in seconds
	Trackers                     []string      `json:"trackers"`
	DBFile                       string        `json:"dbFile"`
	ServerHost                   string        `json:"serverHost"`
	ServerPort                   uint          `json:"serverPort"`
	FetchTimeout                 uint          `json:"fetchTimeout"`
	RequestData                  uint          `json:"requestData"`
	RequestDataInterval          time.Duration `json:"requestDataInterval"`          //in seconds
	MiningInterruptCheckInterval time.Duration `json:"miningInterruptCheckInterval"` //in seconds
	GasMultiplier                uint          `json:"gasMultiplier"`
	GasMax                       uint          `json:"gasMax"`
	logger                       *util.Logger
	mux                          sync.Mutex
}

const defaultTimeout = 30 //30 second fetch timeout

const defaultRequestInterval = 30 //30 seconds between data requests (0-value tipping)
const defaultMiningInterrupt = 15 //every 15 seconds, check for new challenges that could interrupt current mining

var (
	config *Config
	mux    sync.Mutex
)

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
	config.logger = util.NewLogger("config", "Config")
	if config.FetchTimeout == 0 {
		config.FetchTimeout = defaultTimeout
	}
	if config.RequestDataInterval == 0 {
		config.RequestDataInterval = defaultRequestInterval
	}
	if config.MiningInterruptCheckInterval == 0 {
		config.MiningInterruptCheckInterval = defaultMiningInterrupt
	}
	config.logger.Info("config: %+v", config)
	return config, nil
}

//GetConfig returns a shared instance of config
func GetConfig() (*Config, error) {
	if config == nil {
		mux.Lock()
		defer mux.Unlock()
		if config == nil {
			_, err := ParseConfig("")
			if err != nil {
				return nil, err
			}
		}
	}
	return config, nil
}
