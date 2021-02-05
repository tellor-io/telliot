// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package db

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/logging"
)

const ComponentName = "db"

const (

	// minCache is the minimum amount of memory in megabytes to allocate to leveldb
	// read and write caching, split half and half.
	minCache = 4

	// minHandles is the minimum number of files handles to allocate to the open
	// database files.
	minHandles = 8
)

// DB is the primary interface to an underlying datastore.
type DB interface {
	Has(key string) (bool, error)
	Put(key string, value []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
	Close() error
}

type impl struct {
	db     *leveldb.DB
	logger log.Logger
}

// Open the database using the given DB file as its data store.
func Open(logger log.Logger, cfg *config.Config, file string) (DB, error) {
	// Open the db and recover any potential corruptions.
	db, err := leveldb.OpenFile(file, &opt.Options{
		OpenFilesCacheCapacity: minHandles,
		BlockCacheCapacity:     minCache / 2 * opt.MiB,
		WriteBuffer:            minCache / 4 * opt.MiB, // Two of these are used internally.
		Filter:                 filter.NewBloomFilter(10),
	})
	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		db, err = leveldb.RecoverFile(file, nil)
	}
	if err != nil {
		return nil, err
	}

	logger, err = logging.ApplyFilter(*cfg, ComponentName, logger)
	if err != nil {
		return nil, err
	}

	i := &impl{db: db, logger: log.With(logger, "component", ComponentName)}
	level.Info(i.logger).Log("msg", "created DB", "at", file)
	return i, nil
}

func (i *impl) Close() error {
	level.Info(i.logger).Log("msg", "closing db")
	return i.db.Close()
}

func (i *impl) Has(key string) (bool, error) {
	return i.db.Has([]byte(key), nil)
}

func (i *impl) Put(key string, value []byte) error {
	level.Debug(i.logger).Log(
		"msg", "adding DB entry",
		"key", key,
		"bytes", len(value),
	)
	return i.db.Put([]byte(key), value, nil)
}

func (i *impl) Get(key string) ([]byte, error) {
	b, e := i.db.Get([]byte(key), nil)
	if e == errors.ErrNotFound {
		level.Debug(i.logger).Log(
			"msg", "did not find value",
			"key", key,
		)
		return nil, nil
	}
	return b, e
}

func (i *impl) Delete(key string) error {
	level.Debug(i.logger).Log("msg", "deleting key", "key", key)
	return i.db.Delete([]byte(key), nil)
}
