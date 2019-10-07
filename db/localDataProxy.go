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

func (l *localProxy) IncomingRequest(data []byte) ([]byte, error) {
	return nil, fmt.Errorf("Local proxy should never be called with incoming requests")
}
