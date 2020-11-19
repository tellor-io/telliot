// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package util

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

var httpFetchLog = NewLogger("util", "HTTPFetchWithRetries")

const (
	// GET is a GET request type.
	GET = iota + 1
	// POST is a POST request type.
	POST
)

// HTTPFetchRequest holds info for a request.
type HTTPFetchRequest struct {
	Method   int
	QueryURL string
	Payload  []byte
	Timeout  time.Duration
}

// HTTPWithRetries will keep trying the given request until non-error result or timeout.
func HTTPWithRetries(req *HTTPFetchRequest) ([]byte, error) {
	return _recReq(req, time.Now().Add(req.Timeout))
}

func _recReq(req *HTTPFetchRequest, expiration time.Time) ([]byte, error) {
	httpFetchLog.Debug("Fetch request will expire at: %v (timeout: %v)", expiration, req.Timeout)
	var r *http.Response
	var err error
	if req.Method == GET {
		r, err = http.Get(req.QueryURL)
	} else {
		r, err = http.Post(req.QueryURL, "application/json", bytes.NewBuffer(req.Payload))
	}
	if err != nil {
		// Log local non-timeout errors for now.
		httpFetchLog.Warn("Problem fetching data from: %s. %v", req.QueryURL, err)
		now := time.Now()
		if now.After(expiration) {
			httpFetchLog.Error("Timeout expired, not retrying query and passing error up")
			return nil, err
		}
		// FIXME: should this be configured as fetch error sleep duration?
		time.Sleep(500 * time.Millisecond)

		// Try again.
		httpFetchLog.Warn("Trying fetch again...")
		return _recReq(req, expiration)
	}

	data, _ := ioutil.ReadAll(r.Body)

	if r.StatusCode < 200 || r.StatusCode > 299 {
		httpFetchLog.Warn("Response from fetching  %s. Response code: %d, payload: %s", req.QueryURL, r.StatusCode, data)
		// Log local non-timeout errors for now.
		now := time.Now()
		if now.After(expiration) {
			return nil, errors.Wrapf(err, "Giving up fetch request after request timeout: %d", r.StatusCode)
		}
		// FIXME: should this be configured as fetch error sleep duration?
		time.Sleep(500 * time.Millisecond)

		// Try again.
		return _recReq(req, expiration)
	}
	return data, nil
}
