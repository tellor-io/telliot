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
	rawJSON, err := ioutil.ReadFile(filepath.Join("..", "..", "..", "test", "tracker", "testdata", "test_api.json"))
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

func TestMain(m *testing.M) {
	cfg, err := config.OpenTestConfig("../../../")
	if err != nil {
		log.Fatal(err)
	}

	if err := apiOracle.EnsureValueOracle(logging.NewLogger(), cfg); err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}
