// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package db

import (
	"io/ioutil"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/logging"
)

func OpenTestDB(cfg *config.Config) (DB, func() error, error) {
	tmpdir, err := ioutil.TempDir("", "test")
	if err != nil {
		return nil, nil, err
	}

	db, err := Open(logging.NewLogger(), cfg, tmpdir)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() error {
		if err := db.Close(); err != nil {
			if err != leveldb.ErrClosed {
				if err != nil {
					return err
				}
			}
		}
		if err := os.RemoveAll(tmpdir); err != nil {
			return err
		}
		return nil

	}
	return db, cleanup, nil
}
