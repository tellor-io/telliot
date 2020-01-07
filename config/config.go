package config

import (
	"fmt"
	"encoding/json"
	"log"
	"os"
	"strings"
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
	RequestTips					 int64          `json: "requestTips"`
	MiningInterruptCheckInterval time.Duration `json:"miningInterruptCheckInterval"` //in seconds
	GasMultiplier                float32       `json:"gasMultiplier"`
	GasMax                       uint          `json:"gasMax"`
	NumProcessors                int           `json:"numProcessors"`
	Heartbeat                    time.Duration `json:"heartbeat"`
	ServerWhitelist              []string      `json:"serverWhitelist"`
	UseGPU					     bool 	       `json:"useGPU"`
	logger                       *util.Logger
	mux                          sync.Mutex
}

const defaultTimeout = 30 //30 second fetch timeout

const defaultRequestInterval = 30 //30 seconds between data requests (0-value tipping)
const defaultMiningInterrupt = 15 //every 15 seconds, check for new challenges that could interrupt current mining
const defaultCores = 2

const defaultHeartbeat = 10000000 //check miner speed every 10 ^ 8 cycles

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
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Fatalf("Invalid ConfigPath setting: %s", path)
	}
	if info.IsDir() {
		log.Fatalf("ConfigPath is a directory: %s", path)
	}

	configFile, err := os.Open(path)
	defer configFile.Close()
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(configFile)
	err = dec.Decode(&config)
	config.logger = util.NewLogger("config", "Config")
	if config.UseGPU == false {
		fmt.Println("Not using GPU's, check config file")
	}
	if config.FetchTimeout == 0 {
		config.FetchTimeout = defaultTimeout
	}
	if config.RequestDataInterval == 0 {
		config.RequestDataInterval = defaultRequestInterval
	}
	if config.MiningInterruptCheckInterval == 0 {
		config.MiningInterruptCheckInterval = defaultMiningInterrupt
	}
	if config.NumProcessors == 0 {
		config.NumProcessors = defaultCores
	}

	if config.Heartbeat == 0 {
		config.Heartbeat = defaultHeartbeat
	}

	if len(config.ServerWhitelist) == 0{
		if strings.Contains(config.PublicAddress, "0x") {
			config.ServerWhitelist = append(config.ServerWhitelist,config.PublicAddress)
		}else{
			config.ServerWhitelist = append(config.ServerWhitelist,"0x" + config.PublicAddress)
		}
	}

	config.PrivateKey = strings.ReplaceAll(config.PrivateKey, "0x", "")
	config.PublicAddress = strings.ReplaceAll(config.PublicAddress, "0x", "")

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
