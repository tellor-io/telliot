package util

import (
	"net/http"
)

// RoundTripFunc The mock interface to implement. Can be configured to return
// a hardcoded response or a response that depends on the request
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip The method called by transport to get a response
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//CreateTestClient returns *http.Client with Transport replaced to avoid making real calls
func CreateTestClient(client *http.Client, fn RoundTripFunc) {
	client.Transport = RoundTripFunc(fn)
}
