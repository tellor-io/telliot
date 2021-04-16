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
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/mining"
	"github.com/tellor-io/telliot/pkg/submitter"
	"github.com/tellor-io/telliot/pkg/tasker"
	"github.com/tellor-io/telliot/pkg/tracker/index"
	"github.com/tellor-io/telliot/pkg/tracker/profit"
	"github.com/tellor-io/telliot/pkg/transactor"
	"github.com/tellor-io/telliot/pkg/util"
)

// Config is the top-level configuration that holds configs for all components.
type Config struct {
	Mining        mining.Config
	Submitter     submitter.Config
	ProfitTracker profit.Config
	Tasker        tasker.Config
	Transactor    transactor.Config
	DataServer    dataServer.Config
	IndexTracker  index.Config
	Ethereum      ethereum.Config
	Aggregator    aggregator.Config
	Db            db.Config
	// EnvFile location that include all private details like private key etc.
	EnvFile string `json:"envFile"`

	// Exposes metrics on this host and port.
	ListenHost string
	ListenPort uint
}

var defaultConfig = Config{
	ListenHost: "localhost",
	ListenPort: 9090,

	Db: db.Config{
		LogLevel: "info",
		Path:     "db",
	},
	Tasker: tasker.Config{
		LogLevel: "info",
	},
	ProfitTracker: profit.Config{
		LogLevel: "info",
	},
	Ethereum: ethereum.Config{
		LogLevel: "info",
		Timeout:  3000,
	},
	Transactor: transactor.Config{
		GasMax:        10,
		LogLevel:      "info",
		GasMultiplier: 1,
	},
	Submitter: submitter.Config{
		LogLevel: "info",
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

	IndexTracker: index.Config{
		Interval:       util.Duration{30 * time.Second},
		FetchTimeout:   util.Duration{30 * time.Second},
		ApiFile:        "configs/api.json",
		ManualDataFile: "configs/manualData.json",
	},
	EnvFile: "configs/.env",
}

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

	if err := godotenv.Load(cfg.EnvFile); err != nil && !os.IsNotExist(err) {
		return nil, errors.Wrap(err, "loading env vars from env file")
	}

	return cfg, nil
}
