// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
)

// Client utilized for all HTTP requests.
var client http.Client

func init() {
	client = http.Client{}
}

// FetchRequest holds info for a request.
// TODO: add mock fetch.
type FetchRequest struct {
	queryURL string
	timeout  time.Duration
}

func fetchWithRetries(logger log.Logger, req *FetchRequest) ([]byte, error) {
	return _recFetch(logger, req, clck.Now().Add(req.timeout))
}

func _recFetch(logger log.Logger, req *FetchRequest, expiration time.Time) ([]byte, error) {
	level.Debug(logger).Log(
		"msg", "fetch request will expire",
		"at", expiration,
		"timeout", req.timeout,
	)

	now := clck.Now()
	client.Timeout = expiration.Sub(now)

	r, err := client.Get(req.queryURL)
	if err != nil {
		//log local non-timeout errors for now
		level.Warn(logger).Log(
			"msg", "fetching data",
			"from", req.queryURL,
			"err", err,
		)
		now := clck.Now()
		if now.After(expiration) {
			return nil, errors.Wrap(err, "retry timeout expired, last error is wrapped")
		}
		//FIXME: should this be configured as fetch error sleep duration?
		time.Sleep(1000 * time.Millisecond)

		//try again
		level.Warn(logger).Log("msg", "trying fetch again")
		return _recFetch(logger, req, expiration)
	}

	data, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		return nil, errors.Wrap(err, "read response body")
	}

	if r.StatusCode < 200 || r.StatusCode > 299 {
		level.Warn(logger).Log(
			"msg", "response from fetching",
			"queryURL", req.queryURL,
			"statusCode", r.StatusCode,
			"payload", data,
		)
		//log local non-timeout errors for now
		// this is a duplicated error that is unlikely to be triggered since expiration is updated above
		now := clck.Now()
		if now.After(expiration) {
			return nil, errors.Errorf("giving up fetch request after request timeout:%v", r.StatusCode)
		}
		//FIXME: should this be configured as fetch error sleep duration?
		time.Sleep(500 * time.Millisecond)

		//try again
		return _recFetch(logger, req, expiration)
	}
	return data, nil
}
