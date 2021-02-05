// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package util

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
)

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
func HTTPWithRetries(logger log.Logger, req *HTTPFetchRequest) ([]byte, error) {
	return _recReq(logger, req, time.Now().Add(req.Timeout))
}

func _recReq(logger log.Logger, req *HTTPFetchRequest, expiration time.Time) ([]byte, error) {
	level.Debug(logger).Log(
		"msg", "fetch request will expire",
		"expiration", expiration,
		"timeout", req.Timeout,
	)
	var r *http.Response
	var err error
	if req.Method == GET {
		r, err = http.Get(req.QueryURL)
	} else {
		r, err = http.Post(req.QueryURL, "application/json", bytes.NewBuffer(req.Payload))
	}
	if err != nil {
		// Log local non-timeout errors for now.
		level.Warn(logger).Log(
			"msg", "fetching data",
			"queryURL", req.QueryURL,
			"err", err,
		)
		now := time.Now()
		if now.After(expiration) {
			return nil, errors.Wrap(err, "timeout expired, not retrying query and passing error up")
		}
		// FIXME: should this be configured as fetch error sleep duration?
		time.Sleep(500 * time.Millisecond)

		// Try again.
		level.Warn(logger).Log("msg", "trying to fetch again")
		return _recReq(logger, req, expiration)
	}

	data, _ := ioutil.ReadAll(r.Body)

	if r.StatusCode < 200 || r.StatusCode > 299 {
		level.Warn(logger).Log(
			"msg", "response from fetching",
			"queryURL", req.QueryURL,
			"statusCode", r.StatusCode,
			"payload", data,
		)
		// Log local non-timeout errors for now.
		now := time.Now()
		if now.After(expiration) {
			return nil, errors.Errorf("giving up fetch request after request timeout: %d", r.StatusCode)
		}
		// FIXME: should this be configured as fetch error sleep duration?
		time.Sleep(500 * time.Millisecond)

		// Try again.
		return _recReq(logger, req, expiration)
	}
	return data, nil
}
