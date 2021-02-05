// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package db

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/testutil"
)

func OpenTestDB(t *testing.T) (DB, func()) {
	tmpdir, err := ioutil.TempDir("", "test")
	testutil.Ok(t, err)

	db, err := Open(logging.NewLogger(), config.GetConfig(), tmpdir)
	testutil.Ok(t, err)

	cleanup := func() {
		if err := db.Close(); err != nil {
			if err != leveldb.ErrClosed {
				testutil.Ok(t, errors.Wrap(err, "closing the DB"))
			}
		}
		if err := os.RemoveAll(tmpdir); err != nil {
			testutil.Ok(t, errors.Wrap(err, "removing temp DB dir"))

		}
	}
	return db, cleanup
}
