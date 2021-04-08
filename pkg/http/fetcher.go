// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package http

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func Fetch(ctx context.Context, logger log.Logger, url string, retryDelay time.Duration) ([]byte, error) {
	client := http.Client{}
	ticker := time.NewTicker(retryDelay)

	logger = log.With(logger, "url", url)

	for {
		r, err := client.Get(url)
		if err != nil {
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
}
