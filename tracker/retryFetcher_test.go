package tracker

import (
	"testing"
	"time"
)

func TestFetchRetry(t *testing.T) {
	req := &FetchRequest{queryURL: "https://api.gdax.com/products/ETH-USD/ticker", timeout: time.Duration(5 * time.Second)}

	res, err := fetchWithRetries(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Result from query: %s\n", string(res))

}
