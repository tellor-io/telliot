package tracker

import (
	"encoding/json"
	"fmt"
	"github.com/tellor-io/TellorMiner/config"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//maps requestID to a time window of values
var valueHistory map[uint64]*Window
var valueHistoryMutex sync.RWMutex

//last time PSR windows written to disk
var lastHistoryWriteAttempt time.Time

func GetLatestRequestValue(id uint64) *TimedInt {
	valueHistoryMutex.RLock()
	defer valueHistoryMutex.RUnlock()
	w, ok := valueHistory[id]
	if !ok {
		return nil
	}
	return w.Latest()
}

func GetRequestValuesForTime(id uint64, at time.Time, delta time.Duration) []*TimedInt {
	valueHistoryMutex.RLock()
	defer valueHistoryMutex.RUnlock()
	w, ok := valueHistory[id]
	if !ok {
		return nil
	}
	return w.WithinRange(at, delta)
}

func setRequestValue(id uint64, at time.Time, val *big.Int) {

	valueHistoryMutex.Lock()
	_, ok := valueHistory[id]
	if !ok {
		valueHistory[id] = NewWindow(7 * 24 * time.Hour)
	}
	valueHistory[id].Insert(&TimedInt{
		Created: at,
		Val:val.Uint64(),
	})
	valueHistoryMutex.Unlock()
}

func writeOutHistory() {
	valueHistoryMutex.Lock()
	for _,v := range valueHistory {
		v.Trim()
	}
	data, err := json.MarshalIndent(valueHistory, "", "\t")

	//in order to not hold up the rest of the program, we release the mutex while we write out the file
	//this function is single threaded, but we need mutex to access multithreaded history
	valueHistoryMutex.Unlock()
	if err != nil {
		psrLog.Error("failed to marshal PSR values: %s", err.Error())
		return
	}

	cfg := config.GetConfig()
	psrSavedData := filepath.Join(cfg.PSRFolder, "saved.json")
	psrSavedDataTmp := psrSavedData + ".tmp"
	err = ioutil.WriteFile(psrSavedDataTmp, data, 0644)
	if err != nil {
		psrLog.Error("failed to write out PSR values to %s: %s", psrSavedDataTmp, err.Error())
		return
	}
	//rename tmp file to old file (should be atomic on most modern OS)
	err = os.Rename(psrSavedDataTmp, psrSavedData)
	if err != nil {
		psrLog.Error("failed move new PSR save onto old: %s", err.Error())
		return
	}
}

func EnsureValueOracle() error {
	if valueHistory != nil {
		return nil
	}

	valueHistoryMutex.Lock()
	defer valueHistoryMutex.Unlock()

	//check again after we grabbed mutex
	if valueHistory != nil {
		return nil
	}

	cfg := config.GetConfig()

	historyPath := filepath.Join(cfg.PSRFolder, "saved.json")

	_, err := os.Stat(historyPath)
	exists := true
	if err != nil {
		if os.IsNotExist(err) {
			exists = false
		} else {
			return fmt.Errorf("file %s stat error: %v", historyPath, err)
		}
	}

	if exists {
		byteValue, err := ioutil.ReadFile(historyPath)
		if err != nil {
			return fmt.Errorf("failed to read psr file @ %s: %v", historyPath, err)
		}
		err = json.Unmarshal(byteValue, &valueHistory)
		if err != nil {
			return fmt.Errorf("failed to unmarshal saved")
		}
	} else {
		valueHistory = make(map[uint64]*Window)
	}
	//periodically flush the value history to disk to create a record for disputes
	go func() {
		for {
			time.Sleep(2 * time.Minute)
			writeOutHistory()
		}
	}()
	return nil
}

