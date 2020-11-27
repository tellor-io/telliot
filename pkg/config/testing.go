// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/phayes/freeport"
	"github.com/tellor-io/telliot/pkg/util"
)

var mainConfig = `
{
    "contractAddress": "0x0Ba45A8b5d5575935B8158a88C631E9F9C95a2e5",
    "nodeURL": "https://mainnet.infura.io/v3/7f11ed6dfxxxxxxxxxxxxxxxxx",
    "databaseURL": "http://localhost7545",
    "publicAddress": "3808C441402D0418C52237B697bBC5D18c84839A",
    "privateKey": "A000000000000000000000000000000000000000000000000000000000000000",
    "serverHost": "localhost",
    "serverPort": 5000,
    "trackerCycle": 1,
    "trackers": {},
    "dbFile": "/tellorDB",
    "requestTips": 1,
    "configFolder": "` + filepath.Join("..", "..", "configs") + `"
}`

const loggingConfig = `
[
    {
        "component": "config.Config",
        "level": "DEBUG"
    },
    {
        "component":"db.DB",
        "level": "WARN"
    },
    {
        "component": "rpc.client",
        "level": "INFO"
    },
    {
        "component": "rpc.ABICodec",
        "level": "INFO"
    },
    {
        "component": "rpc.mockClient",
        "level": "INFO"
    },
    {
        "component": "tracker.Top50Tracker",
        "level": "INFO"
    },
    {
        "component": "tracker.FetchDataTracker",
        "level": "ERROR"
    },
    {
        "component": "pow.MiningWorker-0",
        "level": "ERROR"
    },    {
        "component": "pow.MiningWorker-1",
        "level": "ERROR"
    },    {
        "component": "pow.MiningTasker-0",
        "level": "ERROR"
    },    {
        "component": "pow.MiningTasker-1",
        "level": "ERROR"
    },
    {
        "component":"tracker.PSRTracker",
        "level":"INFO"
    }
]
`

func OpenTestConfig(t *testing.T) *Config {
	mainConfigFile, err := ioutil.TempFile(os.TempDir(), "testing")
	if err != nil {
		t.Fatal("Cannot create temporary file", err)
	}
	defer os.Remove(mainConfigFile.Name())

	if _, err = mainConfigFile.Write([]byte(mainConfig)); err != nil {
		t.Fatal("Failed to write the main config file", err)
	}
	if err := mainConfigFile.Close(); err != nil {
		t.Fatal(err)
	}
	err = ParseConfig(mainConfigFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	loggingConfigFile, err := ioutil.TempFile(os.TempDir(), "testing")
	if err != nil {
		t.Fatal("Cannot create temporary file", err)
	}
	defer os.Remove(loggingConfigFile.Name())

	if _, err = loggingConfigFile.Write([]byte(loggingConfig)); err != nil {
		t.Fatal("Failed to write the main config file", err)
	}
	if err := loggingConfigFile.Close(); err != nil {
		t.Fatal(err)
	}
	err = util.ParseLoggingConfig(loggingConfigFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	cfg := GetConfig()
	port, err := freeport.GetFreePort()
	if err != nil {
		t.Fatal(err)
	}
	cfg.ServerPort = uint(port)

	return cfg
}
