// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/aggregator"
	"github.com/tellor-io/telliot/pkg/dataServer"
	"github.com/tellor-io/telliot/pkg/submitter"
	"github.com/tellor-io/telliot/pkg/tracker/index"
	"github.com/tellor-io/telliot/pkg/util"
)

const (
	TellorAddress      = "0x88dF592F8eb5D7Bd38bFeF7dEb0fBc02cf3778a0"
	LensAddressMainnet = "0x577417CFaF319a1fAD90aA135E3848D2C00e68CF"
	LensAddressRinkeby = "0xebEF7ceB7C43850898e258be0a1ea5ffcdBc3205"
)

// Config holds global config info derived from config.json.
type Config struct {
	Submitter        submitter.Config
	DataServer       dataServer.Config
	IndexTracker     index.Config
	Aggregator       aggregator.Config
	EthClientTimeout uint
	DBFile           string
	GasMultiplier    float32
	GasMax           uint
	Logger           map[string]string
	// EnvFile location that include all private details like private key etc.
	EnvFile string `json:"envFile"`
}

var defaultConfig = Config{
	GasMax:        10,
	GasMultiplier: 1,
	Submitter: submitter.Config{
		ListenHost: "localhost",
		ListenPort: 9090,
		// MinSubmitPeriod is the time limit between each submit for a staked miner.
		// We added a 1 second delay here as a workaround to prevent failed transactions.
		MinSubmitPeriod: util.Duration{15*time.Minute + 1*time.Second},
	},
	DataServer: dataServer.Config{
		ListenHost: "localhost",
		ListenPort: 5000,
	},
	Aggregator: aggregator.Config{
		MinConfidence: 0.2,
		Interval:      util.Duration{30 * time.Second},
	},
	DBFile:           "db",
	EthClientTimeout: 3000,
	IndexTracker: index.Config{
		Interval:       util.Duration{30 * time.Second},
		FetchTimeout:   util.Duration{30 * time.Second},
		ApiFile:        "configs/api.json",
		ManualDataFile: "configs/manualData.json",
	},
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

func Populate(cfg *Config) (*Config, error) {
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
			cfg.DataServer.ServerWhitelist = append(cfg.DataServer.ServerWhitelist, acc.Address.String())
		}
	}
	return cfg, nil
}
