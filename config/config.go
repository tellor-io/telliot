package config

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

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


type GPUConfig struct {
	//number of threads in a workgroup
	GroupSize int `json:"groupSize"`
	//total number of threads
	Groups int`json:"groups"`
	//number of iterations within a thread
	Count uint32 `json:"count"`
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
	GPUConfig				     map[string]*GPUConfig  `json:"gpuConfig"`
	EnablePoolWorker             bool     `json:"enablePoolWorker"`
	PoolURL                      string   `json:"poolURL"`
	PSRPath						 string `json:"psrPath"`
	logger                       *util.Logger
}

const defaultTimeout = 30 * time.Second //30 second fetch timeout

const defaultRequestInterval = 30 * time.Second //30 seconds between data requests (0-value tipping)
const defaultMiningInterrupt = 15 * time.Second //every 15 seconds, check for new challenges that could interrupt current mining
const defaultCores = 2

const defaultHeartbeat = 15 * time.Second //check miner speed every 10 ^ 8 cycles

var (
	config *Config
)

//ParseConfig and set a shared config entry
func ParseConfig(path string) (*Config, error) {
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
	if err != nil {
		return nil, err
	}
	config.logger = util.NewLogger("config", "Config")
	if len(config.GPUConfig) == 0  {
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

	config.PrivateKey = strings.ToLower(strings.ReplaceAll(config.PrivateKey, "0x", ""))
	config.PublicAddress = strings.ToLower(strings.ReplaceAll(config.PublicAddress, "0x", ""))

	err = validateConfig(config)
	if err != nil {
		return nil, err
	}

	err = validateConfig(config)
	if err != nil {
		return nil, err
	}

	config.logger.Info("config: %+v", config)
	return config, nil
}



func validateConfig(cfg *Config) error {
	b, err := hex.DecodeString(cfg.PublicAddress)
	if err != nil || len(b) != 20 {
		return fmt.Errorf("expecting 40 hex character public address, got %s", cfg.PublicAddress)
	}
	b, err = hex.DecodeString(cfg.PrivateKey)
	if err != nil || len(b) != 32 {
		return fmt.Errorf("expecting 64 hex character private key, got %s", cfg.PublicAddress)
	}
	b, err = hex.DecodeString(cfg.ContractAddress[2:])
	if err != nil || len(b) != 20 {
		return fmt.Errorf("expecting 40 hex character contract address, got %s", cfg.ContractAddress)
	}

	if cfg.GasMultiplier < 0 || cfg.GasMultiplier > 20 {
		return fmt.Errorf("gas multiplier out of range [0, 20] %f", cfg.GasMultiplier)
	}

	for name,gpuConfig := range cfg.GPUConfig {
		if gpuConfig.Count == 0 {
			return fmt.Errorf("gpu %s requires 'count' > 0", name)
		}
		if gpuConfig.GroupSize == 0 {
			return fmt.Errorf("gpu %s requires 'groupSize' > 0", name)
		}
		if gpuConfig.Groups == 0 {
			return fmt.Errorf("gpu %s requires 'groups' > 0", name)
		}
	}

	return nil
}

//GetConfig returns a shared instance of config
func GetConfig() (*Config, error) {
	return config, nil
}
