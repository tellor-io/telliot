package db

import (
	"fmt"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

const (
	// degradationWarnInterval specifies how often warning should be printed if the
	// leveldb database cannot keep up with requested writes.
	degradationWarnInterval = time.Minute

	// minCache is the minimum amount of memory in megabytes to allocate to leveldb
	// read and write caching, split half and half.
	minCache = 4

	// minHandles is the minimum number of files handles to allocate to the open
	// database files.
	minHandles = 8

	// metricsGatheringInterval specifies the interval to retrieve leveldb database
	// compaction, io and pause stats to report to the user.
	metricsGatheringInterval = 3 * time.Second
)

//DB is the primary interface to an underlying datastore
type DB interface {
	Has(key string) (bool, error)
	Put(key string, value []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
	Close() error
}

type impl struct {
	db *leveldb.DB
}

//Open the database using the given DB file as its data store
func Open(file string) (DB, error) {
	// Open the db and recover any potential corruptions
	db, err := leveldb.OpenFile(file, &opt.Options{
		OpenFilesCacheCapacity: minHandles,
		BlockCacheCapacity:     minCache / 2 * opt.MiB,
		WriteBuffer:            minCache / 4 * opt.MiB, // Two of these are used internally
		Filter:                 filter.NewBloomFilter(10),
	})
	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		db, err = leveldb.RecoverFile(file, nil)
	}
	if err != nil {
		return nil, err
	}
	return &impl{db}, nil
}

func (i *impl) Close() error {
	fmt.Println("Closing DB...")
	return i.db.Close()
}

func (i *impl) Has(key string) (bool, error) {
	return i.db.Has([]byte(key), nil)
}

func (i *impl) Put(key string, value []byte) error {
	return i.db.Put([]byte(key), value, nil)
}

func (i *impl) Get(key string) ([]byte, error) {
	return i.db.Get([]byte(key), nil)
}

func (i *impl) Delete(key string) error {
	return i.db.Delete([]byte(key), nil)
}
