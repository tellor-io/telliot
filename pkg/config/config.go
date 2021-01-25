// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package config

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// Unfortunate hack to enable json parsing of human readable time strings
// see https://github.com/golang/go/issues/10275
// code from https://stackoverflow.com/questions/48050945/how-to-unmarshal-json-into-durations.
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
		return errors.Errorf("invalid duration")
	}
}

type DataServer struct {
	ListenHost string
	ListenPort uint
}

type Mine struct {
	// Connect to this remote DB.
	RemoteDBHost string
	RemoteDBPort uint
	// Exposes metrics on this host and port.
	ListenHost string
	ListenPort uint
}

// Config holds global config info derived from config.json.
type Config struct {
	Mine                         Mine
	DataServer                   DataServer
	PublicAddress                string            `json:"publicAddress"`
	EthClientTimeout             uint              `json:"ethClientTimeout"`
	MinSubmitPeriod              Duration          `json:"minSubmitPeriod"`
	TrackerSleepCycle            Duration          `json:"trackerCycle"`
	Trackers                     map[string]bool   `json:"trackers"`
	DBFile                       string            `json:"dbFile"`
	FetchTimeout                 Duration          `json:"fetchTimeout"`
	MinConfidence                float64           `json:"minConfidence"`
	MiningInterruptCheckInterval Duration          `json:"miningInterruptCheckInterval"`
	GasMultiplier                float32           `json:"gasMultiplier"`
	GasMax                       uint              `json:"gasMax"`
	NumProcessors                int               `json:"numProcessors"`
	Heartbeat                    Duration          `json:"heartbeat"`
	ServerWhitelist              []string          `json:"serverWhitelist"`
	EnablePoolWorker             bool              `json:"enablePoolWorker"`
	Worker                       string            `json:"worker"`
	Password                     string            `json:"password"`
	PoolURL                      string            `json:"poolURL"`
	ConfigFolder                 string            `json:"configFolder"`
	LogLevel                     string            `json:"logLevel"`
	Logger                       map[string]string `json:"logger"`
	DisputeTimeDelta             Duration          `json:"disputeTimeDelta"` // Ignore data further than this away from the value we are checking.
	DisputeThreshold             float64           `json:"disputeThreshold"` // Maximum allowed relative difference between observed and submitted value.
	// Minimum percent of profit when submitting a solution.
	// For example if the tx cost is 0.01 ETH and current reward is 0.02 ETH
	// a ProfitThreshold of 200% or more will wait until the reward is increased or
	// the gas cost is lowered.
	// a ProfitThreshold of 199% or less will submit
	ProfitThreshold uint64 `json:"profitThreshold"`
	// EnvFile location that include all private details like private key etc.
	EnvFile string `json:"envFile"`
}

const ConfigFolder = "configs"

// TODO remove or refactor to not be a global config instance.
var defaultConfig = Config{
	GasMax:           10,
	GasMultiplier:    1,
	MinConfidence:    0.2,
	MinSubmitPeriod:  Duration{15 * time.Minute},
	DisputeThreshold: 0.01,
	Mine: Mine{
		ListenHost: "localhost",
		ListenPort: 9090,
	},
	DataServer: DataServer{
		ListenHost: "localhost",
		ListenPort: 5000,
	},
	Heartbeat:                    Duration{15 * time.Second},
	DBFile:                       "db",
	MiningInterruptCheckInterval: Duration{15 * time.Second},
	FetchTimeout:                 Duration{30 * time.Second},
	TrackerSleepCycle:            Duration{30 * time.Second},
	DisputeTimeDelta:             Duration{5 * time.Minute},
	NumProcessors:                2,
	EthClientTimeout:             3000,
	Trackers: map[string]bool{
		"timeOut":          true,
		"balance":          true,
		"currentVariables": true,
		"disputeStatus":    true,
		"gas":              true,
		"tributeBalance":   true,
		"indexers":         true,
		"disputeChecker":   false,
	},
	ConfigFolder: ConfigFolder,
	LogLevel:     "info",
	Logger: map[string]string{
		"config.Config":            "INFO",
		"db.DB":                    "INFO",
		"rpc.client":               "INFO",
		"rpc.ABICodec":             "INFO",
		"rpc.mockClient":           "INFO",
		"tracker.Top50Tracker":     "INFO",
		"tracker.FetchDataTracker": "INFO",
		"pow.MiningWorker-0:":      "INFO",
		"pow.MiningWorker-1:":      "INFO",
		"pow.MiningTasker-0:":      "INFO",
		"pow.MiningTasker-1:":      "INFO",
		"tracker.PSRTracker":       "INFO",
	},
	EnvFile: path.Join(ConfigFolder, ".env"),
}

const PrivateKeyEnvName = "ETH_PRIVATE_KEY"
const NodeURLEnvName = "NODE_URL"

// ParseConfig and set a shared config entry.
func ParseConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Wrapf(err, "opening file:%v", path)
	}
	return ParseConfigBytes(data)
}

func ParseConfigBytes(data []byte) error {
	err := json.Unmarshal(data, &defaultConfig)
	config := &defaultConfig
	if err != nil {
		return errors.Wrap(err, "parse config json")
	}
	err = godotenv.Load(defaultConfig.EnvFile)
	// Ignore file doesn't exist errors.
	if err != nil && !os.IsNotExist(err) {
		return errors.Wrap(err, "loading .env file")
	}

	if len(config.ServerWhitelist) == 0 {
		if strings.Contains(config.PublicAddress, "0x") {
			config.ServerWhitelist = append(config.ServerWhitelist, config.PublicAddress)
		} else {
			config.ServerWhitelist = append(config.ServerWhitelist, "0x"+config.PublicAddress)
		}
	}

	os.Setenv(PrivateKeyEnvName, strings.ToLower(strings.ReplaceAll(os.Getenv(PrivateKeyEnvName), "0x", "")))
	config.PublicAddress = strings.ToLower(strings.ReplaceAll(config.PublicAddress, "0x", ""))

	err = validateConfig(config)
	if err != nil {
		return errors.Wrap(err, "config validation")
	}
	return nil
}

func validateConfig(cfg *Config) error {
	b, err := hex.DecodeString(cfg.PublicAddress)
	if err != nil || len(b) != 20 {
		return errors.Wrapf(err, "expecting 40 hex character public address, got \"%s\"", cfg.PublicAddress)
	}
	if os.Getenv(NodeURLEnvName) == "" {
		return errors.Errorf("missing nodeURL environment variable '%v'", NodeURLEnvName)
	}
	if cfg.EnablePoolWorker {
		if len(cfg.Worker) == 0 {
			return errors.Errorf("worker name required for pool")
		}
		if len(cfg.Password) == 0 {
			return errors.Errorf("password name required for pool")
		}
	} else {
		b, err = hex.DecodeString(os.Getenv(PrivateKeyEnvName))
		if err != nil || len(b) != 32 {
			return errors.Wrapf(err, "expecting 64 hex character private key, got \"%s\"", os.Getenv(PrivateKeyEnvName))
		}
		if cfg.GasMultiplier < 0 || cfg.GasMultiplier > 20 {
			return errors.Errorf("gas multiplier out of range [0, 20] %f", cfg.GasMultiplier)
		}
	}

	return nil
}

// GetConfig returns a shared instance of config.
func GetConfig() *Config {
	return &defaultConfig
}
