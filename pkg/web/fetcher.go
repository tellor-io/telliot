// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package web

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func Fetch(ctx context.Context, logger log.Logger, url string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: tr}
	ticker := time.NewTicker(1 * time.Second)

	logger = log.With(logger, "url", url)

	var errFinal error
	for i := 0; i < 5; i++ {
		r, err := client.Get(url)
		if err != nil {
			errFinal = err
			level.Error(logger).Log("msg", "fetching data", "err", err)
			select {
			case <-ticker.C:
				continue
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errFinal = err
			level.Error(logger).Log("msg", "read response body", "err", err)
			select {
			case <-ticker.C:
				continue
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}
		r.Body.Close()

		if r.StatusCode/100 != 2 {
			errFinal = errors.Errorf("response status code not ok:%v", r.StatusCode)
			level.Error(logger).Log("msg", "response status", "code", r.StatusCode, "payload", data)
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
