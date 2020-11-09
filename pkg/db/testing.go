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
	}
	db, err := Open(tmpdir)
	if err != nil {
		log.Fatal(err)
	}

	cleanup := func() {
		if err := db.Close(); err != nil {
			if err != leveldb.ErrClosed {
				t.Fatal("closing the DB", err)
			}
		}
		if err := os.RemoveAll(tmpdir); err != nil {
			t.Fatal("removing temp DB dir", err)

		}
	}
	return db, cleanup
}
