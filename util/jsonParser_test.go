package util

import (
	"testing"
)

const API = "json(https://api.gdax.com/products/ETH-USD/ticker).price"

func TestJSONParser(t *testing.T) {
	res := fetchAPI(1000, API)
	if res > 0 {
		t.Logf("Parsed json properly: %d", res)

	} else {
		t.Fatalf("Json not parsed properly: %v", res)
	}
}
