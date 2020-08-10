package db

import (
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/tellor-io/TellorMiner/util"
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
	db  *leveldb.DB
	log *util.Logger
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

	i := &impl{db: db, log: util.NewLogger("db", "DB")}
	i.log.Info("Created DB at path: %s\n", file)
	return i, nil
}

func (i *impl) Close() error {
	i.log.Info("Closing DB...")
	return i.db.Close()
}

func (i *impl) Has(key string) (bool, error) {
	return i.db.Has([]byte(key), nil)
}

func (i *impl) Put(key string, value []byte) error {
	i.log.Debug("Adding DB entry: %s with %d bytes of data", key, len(value))
	return i.db.Put([]byte(key), value, nil)
}

func (i *impl) Get(key string) ([]byte, error) {
	b, e := i.db.Get([]byte(key), nil)
	if e == errors.ErrNotFound {
		i.log.Debug("Did not find value for key: %s", key)
		return nil, nil
	}
	return b, e
}

func (i *impl) Delete(key string) error {
	i.log.Debug("Deleting key: %s", key)
	return i.db.Delete([]byte(key), nil)
}
