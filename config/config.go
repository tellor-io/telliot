package config

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
	"strings"
	"time"
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
	Groups int `json:"groups"`
	//number of iterations within a thread
	Count uint32 `json:"count"`

	Disabled bool `json:"disabled"`
}

//Config holds global config info derived from config.json
type Config struct {
	ContractAddress              string                `json:"contractAddress"`
	NodeURL                      string                `json:"nodeURL"`
	DatabaseURL                  string                `json:"databaseURL"`
	PublicAddress                string                `json:"publicAddress"`
	EthClientTimeout             uint                  `json:"ethClientTimeout"`
	TrackerSleepCycle            Duration              `json:"trackerCycle"`
	Trackers                     []string              `json:"trackers"`
	DBFile                       string                `json:"dbFile"`
	ServerHost                   string                `json:"serverHost"`
	ServerPort                   uint                  `json:"serverPort"`
	FetchTimeout                 Duration              `json:"fetchTimeout"`
	RequestData                  uint                  `json:"requestData"`
	MinConfidence                float64                 `json:"minConfidence"`
	RequestDataInterval          Duration              `json:"requestDataInterval"`
	RequestTips                  int64                 `json:"requestTips"`
	MiningInterruptCheckInterval Duration              `json:"miningInterruptCheckInterval"`
	GasMultiplier                float32               `json:"gasMultiplier"`
	GasMax                       uint                  `json:"gasMax"`
	NumProcessors                int                   `json:"numProcessors"`
	Heartbeat                    Duration              `json:"heartbeat"`
	ServerWhitelist              []string              `json:"serverWhitelist"`
	GPUConfig                    map[string]*GPUConfig `json:"gpuConfig"`
	EnablePoolWorker             bool                  `json:"enablePoolWorker"`
	Worker                       string                `json:"worker"`
	Password                     string                `json:"password"`
	PoolURL                      string                `json:"poolURL"`
	IndexFolder                  string               `json:"indexFolder"`
	DisputeTimeDelta             Duration              `json:"disputeTimeDelta"` //ignore data further than this away from the value we are checking
	DisputeThreshold             float64               `json:"disputeThreshold"` //maximum allowed relative difference between observed and submitted value

	//config parameters excluded from the json config file
	PrivateKey                   string 			   `json:"privateKey"`
}

const defaultTimeout = 30 * time.Second //30 second fetch timeout

const defaultRequestInterval = 30 * time.Second //30 seconds between data requests (0-value tipping)
const defaultMiningInterrupt = 15 * time.Second //every 15 seconds, check for new challenges that could interrupt current mining
const defaultCores = 2

const defaultHeartbeat = 15 * time.Second       //check miner speed every 10 ^ 8 cycles
var (
	config *Config
)

const defaultMaxParallelPSR = 4

const defaultTrackerInterval = 30 * time.Second

const DefaultMaxCheckTimeDelta = 5 * time.Minute

//threshold, a percentage of the expected value
const DefaultDisputeThreshold = 0.01

const PrivateKeyEnvName = "ETH_PRIVATE_KEY"

//ParseConfig and set a shared config entry
func ParseConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to open config file %s: %v", path, err)
	}
	return ParseConfigBytes(data)
}

func ParseConfigBytes(data []byte) error {
	//parse the json
	err := json.Unmarshal(data, &config)
	if err != nil {
		return fmt.Errorf("failed to parse json: %s", err.Error())
	}

	//check if the env is already set, only try loading .env if its not there
	if config.PrivateKey == "" {
		//load the env
		err = godotenv.Load()
		if err != nil {
			return fmt.Errorf("error reading .env file: %v", err)
		}

		config.PrivateKey = os.Getenv(PrivateKeyEnvName)
		if config.PrivateKey == "" {
			return fmt.Errorf("missing ethereum wallet private key environment variable '%s'", PrivateKeyEnvName)
		}
	}

	if config.MinConfidence == 0 {
		config.MinConfidence = 0.2
	}

	if config.FetchTimeout.Seconds() == 0 {
		config.FetchTimeout.Duration = defaultTimeout
	}
	if config.RequestDataInterval.Duration == 0 {
		config.RequestDataInterval.Duration = defaultRequestInterval
	}
	if config.MiningInterruptCheckInterval.Seconds() == 0 {
		config.MiningInterruptCheckInterval.Duration = defaultMiningInterrupt
	}
	if config.NumProcessors == 0 {
		config.NumProcessors = defaultCores
	}
	if config.TrackerSleepCycle.Duration == 0 {
		config.TrackerSleepCycle.Duration = defaultTrackerInterval
	}

	if config.Heartbeat.Seconds() == 0 {
		config.Heartbeat.Duration = defaultHeartbeat
	}

	if len(config.ServerWhitelist) == 0 {
		if strings.Contains(config.PublicAddress, "0x") {
			config.ServerWhitelist = append(config.ServerWhitelist, config.PublicAddress)
		} else {
			config.ServerWhitelist = append(config.ServerWhitelist, "0x"+config.PublicAddress)
		}
	}

	config.PrivateKey = strings.ToLower(strings.ReplaceAll(config.PrivateKey, "0x", ""))
	config.PublicAddress = strings.ToLower(strings.ReplaceAll(config.PublicAddress, "0x", ""))

	if config.DisputeThreshold == 0 {
		config.DisputeThreshold = DefaultDisputeThreshold
	}
	if config.DisputeTimeDelta.Duration == 0 {
		config.DisputeTimeDelta.Duration = DefaultMaxCheckTimeDelta
	}

	err = validateConfig(config)
	if err != nil {
		return fmt.Errorf("validation failed: %s", err)
	}
	return nil
}

func validateConfig(cfg *Config) error {
	b, err := hex.DecodeString(cfg.PublicAddress)
	if err != nil || len(b) != 20 {
		return fmt.Errorf("expecting 40 hex character public address, got \"%s\"", cfg.PublicAddress)
	}
	if cfg.EnablePoolWorker  {
		if len(cfg.Worker) == 0 {
			return fmt.Errorf("worker name required for pool")
		}
		if len(cfg.Password) == 0 {
			return fmt.Errorf("password name required for pool")
		}
	} else {
		b, err = hex.DecodeString(cfg.PrivateKey)
		if err != nil || len(b) != 32 {
			return fmt.Errorf("expecting 64 hex character private key, got \"%s\"", cfg.PrivateKey)
		}
		if len(cfg.ContractAddress) != 42 {
			return fmt.Errorf("expecting 40 hex character contract address, got \"%s\"", cfg.ContractAddress)
		}
		b, err = hex.DecodeString(cfg.ContractAddress[2:])
		if err != nil || len(b) != 20 {
			return fmt.Errorf("expecting 40 hex character contract address, got \"%s\"", cfg.ContractAddress)
		}

		if cfg.GasMultiplier < 0 || cfg.GasMultiplier > 20 {
			return fmt.Errorf("gas multiplier out of range [0, 20] %f", cfg.GasMultiplier)
		}
	}

	for name, gpuConfig := range cfg.GPUConfig {
		if gpuConfig.Disabled {
			continue
		}
		if gpuConfig.Count == 0 {
			return fmt.Errorf("gpu '%s' requires 'count' > 0", name)
		}
		if gpuConfig.GroupSize == 0 {
			return fmt.Errorf("gpu '%s' requires 'groupSize' > 0", name)
		}
		if gpuConfig.Groups == 0 {
			return fmt.Errorf("gpu '%s' requires 'groups' > 0", name)
		}
	}

	return nil
}

//GetConfig returns a shared instance of config
func GetConfig() *Config {
	return config
}
