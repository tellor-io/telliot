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

	req, err := http.NewRequest("GET", ExpandTimeVars(url), nil)
	if err != nil {
		return nil, err
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

func ExpandTimeVars(url string) string {
	yesterday := time.Now().UTC().AddDate(0, 0, -1)
	yesterdayBod := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, yesterday.Location())
	yesterdayEod := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 23, 59, 59, 999, yesterday.Location())
	yesterdayBodMilliseconds := strconv.Itoa(int(yesterdayBod.Unix() * 1000))
	yesterdayEodMilliseconds := strconv.Itoa(int(yesterdayEod.Unix() * 1000))
	yesterdayEodSeconds := strconv.Itoa(int(yesterdayEod.Unix()))
	url = strings.Replace(url, "$BOD_MILLISECONDS", yesterdayBodMilliseconds, -1)
	url = strings.Replace(url, "$EOD_MILLISECONDS", yesterdayEodMilliseconds, -1)
	url = strings.Replace(url, "$EOD_SECONDS", yesterdayEodSeconds, -1)
	return url
}