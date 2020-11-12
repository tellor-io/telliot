// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package db

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

func OpenTestDB(t *testing.T) (DB, func()) {
	tmpdir, err := ioutil.TempDir("", "test")
	if err != nil {
		log.Fatal(err)
		// Ok(t, err)
	}
	db, err := Open(tmpdir)
	if err != nil {
		log.Fatal(err)
		// Ok(t, err)
	}

	cleanup := func() {
		if err := db.Close(); err != nil {
			if err != leveldb.ErrClosed {
				log.Fatal(err)
				// Ok(t, errors.Wrap(err, "closing the DB"))
			}
		}
		if err := os.RemoveAll(tmpdir); err != nil {
			log.Fatal(err)
			// testutil.Ok(t, errors.Wrap(err, "removing temp DB dir"))

		}
	}
	return db, cleanup
}
