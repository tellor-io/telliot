// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package math

import (
	"math/big"
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

func TestFloatToBigInt18(t *testing.T) {
	type testcase struct {
		input    float64
		expected string
	}

	cases := []testcase{
		{
			1,
			"1000000000000000000",
		},
		{
			10,
			"10000000000000000000",
		},
		{
			0.1,
			"100000000000000000",
		},
		{
			0.01,
			"10000000000000000",
		},
	}

	for i, tc := range cases {

		expected, ok := big.NewInt(0).SetString(tc.expected, 10)
		testutil.Assert(t, ok)

		ii, err := FloatToBigInt18e(tc.input)
		testutil.Ok(t, err)

		testutil.Equals(t, expected, ii, "Case:"+strconv.Itoa(i))
	}
}

func TestBigInt18eToFloat(t *testing.T) {
	type testcase struct {
		input    string
		expected float64
	}

	cases := []testcase{
		{
			"1000000000000000000",
			1,
		},
		{
			"10000000000000000000",
			10,
		},
		{
			"100000000000000000",
			0.1,
		},
		{
			"10000000000000000",
			0.01,
		},
	}

	for i, tc := range cases {

		input, ok := big.NewInt(0).SetString(tc.input, 10)
		testutil.Assert(t, ok)

		act := BigInt18eToFloat(input)

		testutil.Equals(t, tc.expected, act, "Case:"+strconv.Itoa(i))
	}
}
