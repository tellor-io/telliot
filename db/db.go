package db

import (
	"github.com/syndtr/goleveldb/leveldb"
)

//DB is the primary interface to an underlying datastore
type DB interface {
	Store(key string, value []byte) error
}

type impl struct {
	db *leveldb.DB
}

func Open(path string) (DB, error) {

}
