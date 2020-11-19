// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/tellor-io/TellorMiner/pkg/util"
)

// Client utilized for all HTTP requests.
var client http.Client

func init() {
	client = http.Client{}
}

var retryFetchLog = util.NewLogger("tracker", "FetchWithRetries")

// FetchRequest holds info for a request.
// TODO: add mock fetch.
type FetchRequest struct {
	queryURL string
	timeout  time.Duration
}

func fetchWithRetries(req *FetchRequest) ([]byte, error) {
	return _recFetch(req, clck.Now().Add(req.timeout))
}

func _recFetch(req *FetchRequest, expiration time.Time) ([]byte, error) {
	retryFetchLog.Debug("Fetch request will expire at: %v (timeout: %v)", expiration, req.timeout)

	now := clck.Now()
	client.Timeout = expiration.Sub(now)

	r, err := client.Get(req.queryURL)
	if err != nil {
		//log local non-timeout errors for now
		retryFetchLog.Warn("Problem fetching data from: %s. %v", req.queryURL, err)
		now := clck.Now()
		if now.After(expiration) {
			retryFetchLog.Error("Timeout expired, not retrying query and passing error up")
			return nil, err
		}
		//FIXME: should this be configured as fetch error sleep duration?
		time.Sleep(500 * time.Millisecond)

		//try again
		retryFetchLog.Warn("Trying fetch again...")
		return _recFetch(req, expiration)
	}

	data, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		return nil, errors.Wrap(err, "to read response body")
	}

	if r.StatusCode < 200 || r.StatusCode > 299 {
		retryFetchLog.Warn("Response from fetching  %s. Response code: %d, payload: %s", req.queryURL, r.StatusCode, data)
		//log local non-timeout errors for now
		// this is a duplicated error that is unlikely to be triggered since expiration is updated above
		now := clck.Now()
		if now.After(expiration) {
			return nil, errors.Wrapf(err, "giving up fetch request after request timeout: %d", r.StatusCode)
		}
		//FIXME: should this be configured as fetch error sleep duration?
		time.Sleep(500 * time.Millisecond)

		//try again
		return _recFetch(req, expiration)
	}
	return data, nil
}
