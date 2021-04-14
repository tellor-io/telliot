// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

const (
	TellorAddress      = "0x88dF592F8eb5D7Bd38bFeF7dEb0fBc02cf3778a0"
	LensAddressMainnet = "0x577417CFaF319a1fAD90aA135E3848D2C00e68CF"
	LensAddressRinkeby = "0xebEF7ceB7C43850898e258be0a1ea5ffcdBc3205"
)

// Unfortunate hack to enable json parsing of human readable time strings
// see https://github.com/golang/go/issues/10275
// code from https://stackoverflow.com/questions/48050945/how-to-unmarshal-json-into-durations.
type Duration struct {
	time.Duration
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return []byte("\"" + d.String() + "\""), nil
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
	// Minimum percent of profit when submitting a solution.
	// For example if the tx cost is 0.01 ETH and current reward is 0.02 ETH
	// a ProfitThreshold of 200% or more will wait until the reward is increased or
	// the gas cost is lowered.
	// a ProfitThreshold of 199% or less will submit
	ProfitThreshold              uint64
	Heartbeat                    Duration
	MiningInterruptCheckInterval Duration
	MinSubmitPeriod              Duration
}

type Trackers struct {
	SleepCycle       Duration
	FetchTimeout     Duration
	MinConfidence    float64
	DisputeTimeDelta Duration // Ignore data further than this away from the value we are checking.
	DisputeThreshold float64  // Maximum allowed relative difference between observed and submitted value.
	Names            map[string]bool
}

// Config holds global config info derived from config.json.
type Config struct {
	Mine             Mine
	DataServer       DataServer
	Trackers         Trackers
	EthClientTimeout uint
	DBFile           string
	GasMultiplier    float32
	GasMax           uint
	ServerWhitelist  []string
	ApiFile          string
	ManualDataFile   string
	HistoryFile      string
	Logger           map[string]string
	// EnvFile location that include all private details like private key etc.
	EnvFile string `json:"envFile"`
}

var defaultConfig = Config{
	GasMax:        10,
	GasMultiplier: 1,
	Mine: Mine{
		ListenHost:                   "localhost",
		ListenPort:                   9090,
		Heartbeat:                    Duration{15 * time.Second},
		MiningInterruptCheckInterval: Duration{15 * time.Second},
		// MinSubmitPeriod is the time limit between each submit for a staked miner.
		// We added a 1 second delay here as a workaround to prevent failed transactions.
		MinSubmitPeriod: Duration{15*time.Minute + 1*time.Second},
	},
	DataServer: DataServer{
		ListenHost: "localhost",
		ListenPort: 5000,
	},
	DBFile:           "db",
	EthClientTimeout: 3000,
	Trackers: Trackers{
		SleepCycle:       Duration{30 * time.Second},
		FetchTimeout:     Duration{30 * time.Second},
		MinConfidence:    0.2,
		DisputeTimeDelta: Duration{5 * time.Minute},
		DisputeThreshold: 0.01,
		Names: map[string]bool{
			"gas":            true,
			"indexers":       true,
			"disputeChecker": false,
		},
	},
	ApiFile:        "configs/api.json",
	ManualDataFile: "configs/manualData.json",
	HistoryFile:    "configs/saved.json",
	Logger: map[string]string{
		"db":         "info",
		"rpc":        "info",
		"dataServer": "info",
		"tracker":    "info",
		"pow:":       "info",
		"ops":        "info",
		"rest":       "info",
		"apiOracle":  "info",
	},
	EnvFile: "configs/.env",
}

const PrivateKeysEnvName = "ETH_PRIVATE_KEYS"
const NodeURLEnvName = "NODE_WEBSOCKET_URL"

func ParseConfig(path string) (*Config, error) {
	if path == "" {
		path = filepath.Join("configs", "config.json")
	}

	cfg := &Config{}
	// DeepCopy the default config into the final.
	{
		b, err := json.Marshal(defaultConfig)
		if err != nil {
			return nil, errors.Wrap(err, "marshal default config")
		}

		if err := json.Unmarshal(b, cfg); err != nil {
			return nil, errors.Wrap(err, "copy default config")
		}
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open config file")
	}
	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields()
	for {
		// Override defaults with the custom configs.
		if err := dec.Decode(cfg); err == io.EOF {
			break
		} else if err != nil {
			return nil, errors.Wrap(err, "parse config")
		}

	}

	return Populate(cfg)
}

func validate(cfg *Config) error {
	return nil
}

func Populate(cfg *Config) (*Config, error) {
	if err := validate(cfg); err != nil {
		return nil, errors.Wrap(err, "validate config")
	}

	err := godotenv.Load(cfg.EnvFile)
	if err != nil && !os.IsNotExist(err) {
		return nil, errors.Wrap(err, "loading env vars from env file")
	}

	// Parsing private keys and add their public keys to cfg.ServerWhitelist if any.
	accounts, err := GetAccounts()
	if err != nil {
		return nil, errors.Wrap(err, "parsing private keys")
	}
	if len(accounts) != 0 {
		for _, acc := range accounts {
			cfg.ServerWhitelist = append(cfg.ServerWhitelist, acc.Address.String())
		}
	}
	return cfg, nil
}

func ValidateDataServerConfig(cfg *Config) error {
	if len(cfg.ServerWhitelist) == 0 {
		return errors.New("ServerWhitelist shouldn't be empty while running as dataserver")
	}
	return nil
}

func ValidateMinerConfig(cfg *Config) error {
	_privateKeys := os.Getenv(PrivateKeysEnvName)
	privateKeys := strings.Split(_privateKeys, ",")

	if len(privateKeys) == 0 {
		return errors.New("PrivateKeysEnvName env shouldn't be empty while running as miner")
	}
	return nil
}
