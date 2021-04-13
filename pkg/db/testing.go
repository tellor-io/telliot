// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package db

import (
	"io/ioutil"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/tsdb"
	"github.com/tellor-io/telliot/pkg/config"
)

func OpenTestDB(cfg *config.Config) (*tsdb.DB, func() error, error) {
	tmpdir, err := ioutil.TempDir("", "test")
	if err != nil {
		return nil, nil, err
	}

	db, err := tsdb.Open(tmpdir, log.NewNopLogger(), nil, tsdb.DefaultOptions())
	if err != nil {
		return nil, nil, errors.Wrapf(err, "creating tsdb DB")
	}

	cleanup := func() error {
		if err := db.Close(); err != nil {
			return err
		}
		if err := os.RemoveAll(tmpdir); err != nil {
			return err
		}
		return nil

	}
	return db, cleanup, nil
}
