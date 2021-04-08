// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package index

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/tellor-io/telliot/pkg/apiOracle"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/testutil"
)

type TestCase struct {
	URL      string
	Param    string
	Payload  string
	Expected []float64
}

func TestIndexTracker(t *testing.T) {
	// Load the testdata from test_api.json file.
	// The testdata is genertaed using indextracker_testdata script.
	var testdata map[string][]TestCase
	rawJSON, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "tracker", "testdata", "test_api.json"))
	testutil.Ok(t, err)
	err = json.Unmarshal(rawJSON, &testdata)
	testutil.Ok(t, err)

	// Test jsonpath parsing per test cases.
	for _, indexers := range testdata {
		for _, testCase := range indexers {
			actual, err := (&IndexTracker{Param: testCase.Param}).ParsePayload([]byte(testCase.Payload))
			if err != nil {
				testutil.Ok(t, fmt.Errorf("parse payload(URL: %v): %v", testCase.URL, err))
			}
			testutil.Equals(t, testCase.Expected, actual)
		}
	}
}

var configJSON = `{
	"Trackers": {"names":{}},
    "DbFile": "/tellorDB",
	"Logger": {"db.Db":"DEBUG"},
    "EnvFile": "` + filepath.Join("..", "..", "configs", ".env.example") + `",
    "ApiFile": "` + filepath.Join("..", "..", "configs", "api.json") + `",
    "ManualDataFile": "` + filepath.Join("..", "..", "configs", "manualData.json") + `"
}`

func TestMain(m *testing.M) {
	mainConfigFile, err := ioutil.TempFile(os.TempDir(), "testing")
	if err != nil {
		log.Fatal("Cannot create temporary file", err)
	}
	defer os.Remove(mainConfigFile.Name())

	if _, err = mainConfigFile.Write([]byte(configJSON)); err != nil {
		log.Fatal("write the main config file", err)
	}
	if err := mainConfigFile.Close(); err != nil {
		log.Fatal(err)
	}
	cfg, err := config.ParseConfig(mainConfigFile.Name())
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "parse config %v\n", err)
		os.Exit(-1)
	}
	if err := apiOracle.EnsureValueOracle(logging.NewLogger(), cfg); err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}
