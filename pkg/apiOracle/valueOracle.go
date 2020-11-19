// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package apiOracle

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/util"
)

var logger = util.NewLogger("apiOracle", "valueOracle")

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

func writeOutHistory() {
	valueHistoryMutex.Lock()
	for _, v := range valueHistory {
		v.Trim()
	}
	data, err := json.MarshalIndent(valueHistory, "", "\t")

	// In order to not hold up the rest of the program, we release the mutex while we write out the file
	// this function is single threaded, but we need mutex to access multithreaded history.
	valueHistoryMutex.Unlock()
	if err != nil {
		logger.Error("failed to marshal PSR values: %s", err.Error())
		return
	}

	cfg := config.GetConfig()
	psrSavedData := filepath.Join(cfg.ConfigFolder, "saved.json")
	psrSavedDataTmp := psrSavedData + ".tmp"
	err = ioutil.WriteFile(psrSavedDataTmp, data, 0644)
	if err != nil {
		logger.Error("failed to write out PSR values to %s: %s", psrSavedDataTmp, err.Error())
		return
	}
	// Rename tmp file to old file (should be atomic on most modern OS)
	err = os.Rename(psrSavedDataTmp, psrSavedData)
	if err != nil {
		logger.Error("failed move new PSR save onto old: %s", err.Error())
		return
	}
}

func EnsureValueOracle() error {
	if valueHistory != nil {
		return nil
	}

	valueHistoryMutex.Lock()
	defer valueHistoryMutex.Unlock()

	// Check again after we grabbed mutex
	if valueHistory != nil {
		return nil
	}

	cfg := config.GetConfig()

	historyPath := filepath.Join(cfg.ConfigFolder, "saved.json")

	_, err := os.Stat(historyPath)
	exists := true
	if err != nil {
		if os.IsNotExist(err) {
			exists = false
		} else {
			return errors.Wrapf(err, "stats for file: %v", historyPath)
		}
	}

	if exists {
		byteValue, err := ioutil.ReadFile(historyPath)
		if err != nil {
			return errors.Wrapf(err, "failed to read psr file @: %v", historyPath)
		}
		err = json.Unmarshal(byteValue, &valueHistory)
		if err != nil {
			return errors.Errorf("failed to unmarshal saved")
		}
	} else {
		valueHistory = make(map[string]*Window)
	}
	// Periodically flush the value history to disk to create a record for disputes
	go func() {
		for {
			time.Sleep(2 * time.Minute)
			writeOutHistory()
		}
	}()
	return nil
}
