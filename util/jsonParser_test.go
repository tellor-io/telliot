package util

import (
	"io/ioutil"
	"net/http"
	"testing"
)

const API = "json(https://api.gdax.com/products/ETH-USD/ticker).price"

func TestJSONParser(t *testing.T) {
	res, err := testFetch(1000, API)
	if err != nil {
		t.Fatal(err)
	}
	if res > 0 {
		t.Logf("Parsed json properly: %d", res)

	} else {
		t.Fatalf("Json not parsed properly: %v", res)
	}
}

func testFetch(_granularity uint, queryString string) (int, error) {

	url, args := ParseQueryString(queryString)
	resp, _ := http.Get(url)

	input, _ := ioutil.ReadAll(resp.Body)
	return ParsePayload(input, 1000, args)
}
