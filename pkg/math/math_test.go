// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package math

import (
	"strconv"
	"testing"

	"github.com/tellor-io/telliot/pkg/testutil"
)

func TestPercentageDiff(t *testing.T) {

	type testcase struct {
		old      float64
		new      float64
		expected float64
	}

	cases := []testcase{
		{
			1,
			10,
			90,
		},
		{
			10,
			1,
			-90,
		},
		{
			0.01,
			0.1,
			90,
		},
		{
			0.1,
			0.01,
			-90,
		},
		{
			0,
			1,
			100,
		},
		{
			1,
			0,
			-100,
		},
		{
			1,
			-1,
			-200,
		},
	}

	for i, tc := range cases {

		testutil.Equals(t, tc.expected, PercentageDiff(tc.old, tc.new), "Case:"+strconv.Itoa(i))
	}
}
