// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/phayes/freeport"
)

var mainConfig = `
{
    "PublicAddress": "0x92f91500e105e3051f3cf94616831b58f6bce1e8",
    "DbFile": "/tellorDB",
    "EnvFile": "` + filepath.Join("..", "..", "configs", ".env.example") + `",
    "ApiFile": "` + filepath.Join("..", "..", "configs", "api.json") + `",
    "ManualDataFile": "` + filepath.Join("..", "..", "configs", "manualData.json") + `"
}`

func OpenTestConfig(t *testing.T) *Config {
	mainConfigFile, err := ioutil.TempFile(os.TempDir(), "testing")
	if err != nil {
		t.Fatal("Cannot create temporary file", err)
	}
	defer os.Remove(mainConfigFile.Name())

	if _, err = mainConfigFile.Write([]byte(mainConfig)); err != nil {
		t.Fatal("write the main config file", err)
	}
	if err := mainConfigFile.Close(); err != nil {
		t.Fatal(err)
	}
	cfg, err := ParseConfig(mainConfigFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	port, err := freeport.GetFreePort()
	if err != nil {
		t.Fatal(err)
	}
	cfg.Mine.ListenPort = uint(port)
	// Don't need any trackers for the tests.
	cfg.Trackers.Names = make(map[string]bool)

	return cfg
}
