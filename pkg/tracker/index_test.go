// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/tellor-io/telliot/pkg/testutil"
)

type TestCase struct {
	URL      string
	Param    string
	Payload  string
	Expected []float64
}

func TestIndexTracker(t *testing.T) {
	// Load the testdata from test_indexes.json file.
	// The testdata is genertaed using indextracker_testdata script.
	var testdata map[string][]TestCase
	rawJSON, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "tracker", "testdata", "test_indexes.json"))
	testutil.Ok(t, err)
	err = json.Unmarshal(rawJSON, &testdata)
	testutil.Ok(t, err)

	// Test jsonpath parsing per test cases.
	for _, indexers := range testdata {
		for _, testCase := range indexers {
			actual, err := (&IndexTracker{Param: testCase.Param}).ParsePayload([]byte(testCase.Payload))
			if err != nil {
				testutil.Ok(t, fmt.Errorf("failed to parse payload(URL: %v): %v", testCase.URL, err))
			}
			testutil.Equals(t, testCase.Expected, actual)
		}
	}
}
