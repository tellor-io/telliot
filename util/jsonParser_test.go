// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package util

import (
	"io/ioutil"
	"net/http"
	"testing"
)

const API = "json(https:// pi.gdax.com/products/ETH-USD/ticker).price"

func TestJSONParser(t *testing.T) {
	res, err := testFetch(API)
	if err != nil {
		t.Fatal(err)
	}

	for _, r := range res {
		if r > 0 {
			t.Logf("Parsed json properly: %v", res)

		} else {
			t.Fatalf("Json not parsed properly: %v", res)
		}
	}

}

func testFetch(queryString string) ([]float64, error) {

	url, args := ParseQueryString(queryString)
	resp, _ := http.Get(url)

	input, _ := ioutil.ReadAll(resp.Body)
	return ParsePayload(input, args)
}
