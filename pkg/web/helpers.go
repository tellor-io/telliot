// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package web

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func Get(ctx context.Context, url string, headers map[string]string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: tr}
	ticker := time.NewTicker(1 * time.Second)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if strings.Contains(url, "https://api.anyblock.tools/market/AMPL_USD_via_ALL/") {
		yesterday := time.Now().UTC().AddDate(0, 0, -1)
		yesterdayBod := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, yesterday.Location())
		yesterdayEod := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 23, 59, 59, 999, yesterday.Location())
		yesterdayBodMilliseconds := strconv.Itoa(int(yesterdayBod.Unix() * 1000))
		yesterdayEodMilliseconds := strconv.Itoa(int(yesterdayEod.Unix() * 1000))
		q := req.URL.Query()
		q.Add("start", yesterdayBodMilliseconds)
		q.Add("end", yesterdayEodMilliseconds)
		req.URL.RawQuery = q.Encode()
	} else if strings.Contains(url, "https://min-api.cryptocompare.com/data/dayAvg") {
		yesterday := time.Now().UTC().AddDate(0, 0, -1)
		yesterdayEod := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 23, 59, 59, 999, yesterday.Location())
		yesterdayEodSeconds := strconv.Itoa(int(yesterdayEod.Unix()))
		q := req.URL.Query()
		q.Add("toTs", yesterdayEodSeconds)
		req.URL.RawQuery = q.Encode()
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	var errFinal error
	for i := 0; i < 5; i++ {
		r, err := client.Do(req)
		if err != nil {
			errFinal = errors.Wrap(err, "fetching data")
			select {
			case <-ticker.C:
				continue
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errFinal = errors.Wrap(err, "read response body")
			select {
			case <-ticker.C:
				continue
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}
		r.Body.Close()

		if r.StatusCode/100 != 2 {
			errFinal = errors.Errorf("response status code not OK code:%v, payload:%v", r.StatusCode, string(data))
			select {
			case <-ticker.C:
				continue
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}
		return data, nil
	}

	return nil, errFinal

}
