package db

import (
	"fmt"

	"github.com/tellor-io/TellorMiner/util"
)

type localProxy struct {
	localDB DB
	log     *util.Logger
}

//OpenLocalProxy creates a local data proxy so that the miner operations are seamless regardless
//whether accessing data remotely or locally
func OpenLocalProxy(localDB DB) (DataServerProxy, error) {
	log := util.NewLogger("db", "LocalDataProxy")
	log.Info("Using local data proxy to pull data from local DB")
	return &localProxy{localDB: localDB, log: log}, nil
}

func (l *localProxy) Get(key string) ([]byte, error) {
	keys := []string{key}
	r, err := l.BatchGet(keys)
	if err != nil {
		return nil, err
	}
	return r[key], nil
}

func (l *localProxy) BatchGet(keys []string) (map[string][]byte, error) {
	outMap := map[string][]byte{}
	for _, k := range keys {

		bts, err := l.localDB.Get(k)

		if err != nil {
			return nil, err
		}
		if bts != nil {
			outMap[k] = bts
		}
	}
	l.log.Debug("Requested keys: %v, resulting output: %v", keys, outMap)
	return outMap, nil
}

func (l *localProxy) Put(key string, value []byte) (map[string][]byte, error) {
	values := make([][]byte, 1)
	values[0] = value
	return l.BatchPut([]string{key}, values)
}
func (l *localProxy) BatchPut(keys []string, values [][]byte) (map[string][]byte, error) {

	if len(values) > 0 && len(keys) != len(values) {
		return nil, fmt.Errorf("Keys and values must have same array dimensions")
	}
	for idx, k := range keys {
		err := l.localDB.Put(k, values[idx])
		if err != nil {
			return nil, err
		}
	}

	return l.BatchGet(keys)
}

func (l *localProxy) IncomingRequest(data []byte) ([]byte, error) {
	return nil, fmt.Errorf("Local proxy should never be called with incoming requests")
}
