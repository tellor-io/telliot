// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package util

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/tellor-io/TellorMiner/pkg/testutil"
)

const API = "json(https://api.gdax.com/products/ETH-USD/ticker).price"

func TestJSONParser(t *testing.T) {
	res, err := testFetch(API)
	testutil.Ok(t, err)

	for _, r := range res {
		if r > 0 {
			t.Logf("Parsed json properly: %v", res)

		} else {
			testutil.Ok(t, errors.Errorf("Json not parsed properly: %v", res))
		}
	}

}

func testFetch(queryString string) ([]float64, error) {

	url, args := ParseQueryString(queryString)
	resp, _ := http.Get(url)

	input, _ := ioutil.ReadAll(resp.Body)
	return ParsePayload(input, args)
}
