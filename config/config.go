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

//unfortunate hack to enable json parsing of human readable time strings
//see https://github.com/golang/go/issues/10275
//code from https://stackoverflow.com/questions/48050945/how-to-unmarshal-json-into-durations
type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value * float64(time.Second))
		return nil
	case string:
		dur, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		d.Duration = dur
		return nil
	default:
		return fmt.Errorf("invalid duration")
	}
}


//Config holds global config info derived from config.json
type Config struct {
	ContractAddress              string        `json:"contractAddress"`
	NodeURL                      string        `json:"nodeURL"`
	PrivateKey                   string        `json:"privateKey"`
	DatabaseURL                  string   `json:"databaseURL"`
	PublicAddress                string   `json:"publicAddress"`
	EthClientTimeout             uint     `json:"ethClientTimeout"`
	TrackerSleepCycle            Duration     `json:"trackerCycle"` //in seconds
	Trackers                     []string `json:"trackers"`
	DBFile                       string   `json:"dbFile"`
	ServerHost                   string   `json:"serverHost"`
	ServerPort                   uint     `json:"serverPort"`
	FetchTimeout                 Duration `json:"fetchTimeout"`
	RequestData                  uint     `json:"requestData"`
	RequestDataInterval          Duration `json:"requestDataInterval"` //in seconds
	RequestTips                  int64    `json: "requestTips"`
	MiningInterruptCheckInterval Duration `json:"miningInterruptCheckInterval"` //in seconds
	GasMultiplier                float32  `json:"gasMultiplier"`
	GasMax                       uint     `json:"gasMax"`
	NumProcessors                int      `json:"numProcessors"`
	Heartbeat                    Duration `json:"heartbeat"`
	ServerWhitelist              []string `json:"serverWhitelist"`
	UseGPU                       bool     `json:"useGPU"`
	logger                       *util.Logger
}

const defaultTimeout = 30 * time.Second //30 second fetch timeout

const defaultRequestInterval = 30 * time.Second //30 seconds between data requests (0-value tipping)
const defaultMiningInterrupt = 15 * time.Second //every 15 seconds, check for new challenges that could interrupt current mining
const defaultCores = 2

const defaultHeartbeat = 15 * time.Second //check miner speed every 10 ^ 8 cycles

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
	if config.FetchTimeout.Seconds() == 0 {
		config.FetchTimeout.Duration = defaultTimeout
	}
	if config.RequestDataInterval.Seconds() == 0 {
		config.RequestDataInterval.Duration = defaultRequestInterval
	}
	if config.MiningInterruptCheckInterval.Seconds() == 0 {
		config.MiningInterruptCheckInterval.Duration = defaultMiningInterrupt
	}
	if config.NumProcessors == 0 {
		config.NumProcessors = defaultCores
	}

	if config.Heartbeat.Seconds() == 0 {
		config.Heartbeat.Duration = defaultHeartbeat
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
