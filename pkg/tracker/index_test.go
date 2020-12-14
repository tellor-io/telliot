// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"fmt"
	"testing"

	"github.com/tellor-io/telliot/pkg/testutil"
)

type TestDataSource struct {
	Payload string
}

func (i TestDataSource) Get() ([]byte, error) {
	return []byte(i.Payload), nil
}

type TestCase struct {
	IndexTracker *IndexTracker
	Expected     []float64
}

func TestIndexTracker(t *testing.T) {
	// Create some test cases.
	testCases := []TestCase{
		{
			IndexTracker: &IndexTracker{
				Name:       "test1",
				Identifier: "id1",
				JSONPath:   "$[0][4]",
				Source: TestDataSource{
					Payload: `[[324.34,53453.534,4443.3,45.53453,53.63653]]`,
				},
			},
			Expected: []float64{53.63653},
		},
		{
			IndexTracker: &IndexTracker{
				Name:       "test2",
				Identifier: "id2",
				JSONPath:   `$["test"]["a","b"]`,
				Source: TestDataSource{
					Payload: `{"test":{"a":879.54,"b":876.5}}`,
				},
			},
			Expected: []float64{879.54, 876.5},
		},
		{
			IndexTracker: &IndexTracker{
				Name:       "test3",
				Identifier: "id3",
				JSONPath:   `$[dummy][test]`,
				Source: TestDataSource{
					Payload: `{"dummy":{"test": "1321,67.3"}}`,
				},
			},
			Expected: []float64{132167.3},
		},
		{
			IndexTracker: &IndexTracker{
				Name:       "test4",
				Identifier: "id4",
				JSONPath:   "$[0][7,8]",
				Source: TestDataSource{
					Payload: "[[768,68,324.34,53453.534,4443.3,45.53453,53.63653,454.534,454.837]]",
				},
			},
			Expected: []float64{454.534, 454.837},
		},
	}
	// Test jsonpath parsing per test cases.
	for _, testCase := range testCases {
		payload, _ := testCase.IndexTracker.Source.Get()
		actual, err := testCase.IndexTracker.parsePayload(payload)
		if err != nil {
			testutil.Ok(t, fmt.Errorf("failed to parse payload: %v", err))
		}
		testutil.Equals(t, testCase.Expected, actual)

	}
}
