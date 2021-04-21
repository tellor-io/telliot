// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
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
	"github.com/tellor-io/telliot/pkg/web"
)

// Config is the top-level configuration that holds configs for all components.
type Config struct {
	Web           web.Config
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
}

var defaultConfig = Config{
	Mining: mining.Config{
		LogLevel:  "info",
		Heartbeat: time.Minute,
	},
	Web: web.Config{
		LogLevel:   "info",
		ListenHost: "", // Listen on all addresses.
		ListenPort: 9090,
	},
	Db: db.Config{
		LogLevel:      "info",
		Path:          "db",
		RemoteTimeout: util.Duration{Duration: 5 * time.Second},
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
		LogLevel:      "info",
		GasMax:        10,
		GasMultiplier: 1,
	},
	Submitter: submitter.Config{
		LogLevel: "info",
		// MinSubmitPeriod is the time limit between each submit for a staked miner.
		// We added a 1 second delay here as a workaround to prevent failed transactions.
		MinSubmitPeriod: util.Duration{Duration: 15*time.Minute + 1*time.Second},
	},
	Aggregator: aggregator.Config{
		LogLevel:      "info",
		MinConfidence: 0.2,
	},

	IndexTracker: index.Config{
		LogLevel:       "info",
		Interval:       util.Duration{Duration: 30 * time.Second},
		FetchTimeout:   util.Duration{Duration: 30 * time.Second},
		ApiFile:        "configs/api.json",
		ManualDataFile: "configs/manualData.json",
	},
	EnvFile: ".local/.env",
}

const NodeURLEnvName = "NODE_WEBSOCKET_URL"

func ParseConfig(logger log.Logger, path string) (*Config, error) {
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
	var noConfigFile bool
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, errors.Wrap(err, "open config file")
		}
		noConfigFile = true
		level.Warn(logger).Log("msg", "no config file on disk so using defaults", "path", path)
	}

	if !noConfigFile {
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
	}

	if err := godotenv.Load(cfg.EnvFile); err != nil && !os.IsNotExist(err) {
		return nil, errors.Wrap(err, "loading env vars from env file")
	}

	return cfg, nil
}
