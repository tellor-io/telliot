// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package apiOracle

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/logging"
)

const ComponentName = "apiOracle"

// maps symbol to a time window of values.
var valueHistory map[string]*Window
var valueHistoryMutex sync.RWMutex

func GetNearestTwoRequestValue(id string, at time.Time) (before, after *PriceStamp) {
	valueHistoryMutex.RLock()
	defer valueHistoryMutex.RUnlock()
	w, ok := valueHistory[id]
	if !ok {
		return nil, nil
	}
	return w.ClosestTwo(at)
}

func GetRequestValuesForTime(id string, at time.Time, delta time.Duration) []*PriceStamp {
	valueHistoryMutex.RLock()
	defer valueHistoryMutex.RUnlock()
	w, ok := valueHistory[id]
	if !ok {
		return nil
	}
	return w.WithinRange(at, delta)
}

func SetRequestValue(id string, at time.Time, info PriceInfo) {

	valueHistoryMutex.Lock()
	_, ok := valueHistory[id]
	if !ok {
		valueHistory[id] = NewWindow(7 * 24 * time.Hour)
	}
	valueHistory[id].Insert(&PriceStamp{
		Created:   at,
		PriceInfo: info,
	})
	valueHistoryMutex.Unlock()
}

func writeOutHistory(logger log.Logger, cfg *config.Config) {
	valueHistoryMutex.Lock()
	for _, v := range valueHistory {
		v.Trim()
	}
	data, err := json.MarshalIndent(valueHistory, "", "\t")

	// In order to not hold up the rest of the program, we release the mutex while we write out the file
	// this function is single threaded, but we need mutex to access multithreaded history.
	valueHistoryMutex.Unlock()
	if err != nil {
		level.Error(logger).Log("msg", "marshal PSR values", "err", err)
		return
	}

	psrSavedDataTmp := cfg.HistoryFile + ".tmp"
	err = ioutil.WriteFile(psrSavedDataTmp, data, 0644)
	if err != nil {
		level.Error(logger).Log(
			"msg", "write out PSR values",
			"psrSavedDataTmp", psrSavedDataTmp,
			"err", err,
		)
		return
	}
	// Rename tmp file to old file (should be atomic on most modern OS)
	err = os.Rename(psrSavedDataTmp, cfg.HistoryFile)
	if err != nil {
		level.Error(logger).Log("msg", "move new PSR save onto old", "err", err)
		return
	}
}

func EnsureValueOracle(logger log.Logger, cfg *config.Config) error {
	if valueHistory != nil {
		return nil
	}

	valueHistoryMutex.Lock()
	defer valueHistoryMutex.Unlock()

	// Check again after we grabbed mutex
	if valueHistory != nil {
		return nil
	}

	logger, err := logging.ApplyFilter(*cfg, ComponentName, logger)
	if err != nil {
		return errors.Wrap(err, "creating filter logger")
	}
	logger = log.With(logger, "component", ComponentName)

	_, err = os.Stat(cfg.HistoryFile)
	exists := true
	if err != nil {
		if os.IsNotExist(err) {
			exists = false
		} else {
			return errors.Wrapf(err, "stat error file: %v", cfg.HistoryFile)
		}
	}

	if exists {
		byteValue, err := ioutil.ReadFile(cfg.HistoryFile)
		if err != nil {
			return errors.Wrapf(err, "read psr file:%v", cfg.HistoryFile)
		}
		err = json.Unmarshal(byteValue, &valueHistory)
		if err != nil {
			return errors.Errorf("unmarshal saved values")
		}
	} else {
		valueHistory = make(map[string]*Window)
	}
	// Periodically flush the value history to disk to create a record for disputes
	go func() {
		for {
			time.Sleep(2 * time.Minute)
			writeOutHistory(logger, cfg)
		}
	}()
	return nil
}
