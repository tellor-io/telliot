// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package db

// import (
// 	"crypto/ecdsa"
// 	"io"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"sync"
// 	"time"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/go-kit/kit/log"
// 	"github.com/go-kit/kit/log/level"
// 	lru "github.com/hashicorp/golang-lru"
// 	"github.com/pkg/errors"
// 	"github.com/tellor-io/telliot/pkg/config"
// 	"github.com/tellor-io/telliot/pkg/logging"
// 	"github.com/tellor-io/telliot/pkg/util"
// )

// // DataServerProxy interface for local interaction/abstraction/testing.
// type DataServerProxy interface {
// 	// RequestSigner
// 	// RequestValidator
// 	// local call to get a data server value by its key.
// 	Get(key string) ([]byte, error)

// 	// local call to put data into the data server's store. All keys should
// 	// be prefixed with the calling miner's public ETH key to avoid conflicts.
// 	// Implementation must ensure thread safety from multiple miners attempting
// 	// to write at the same time.
// 	Put(key string, value []byte) error

// 	// put multiple keys and values on remote data server.
// 	BatchPut(keys []string, values [][]byte) error

// 	// local call to get several data server values by their keys.
// 	BatchGet(keys []string) (map[string][]byte, error)

// 	// notification that a remote miner has requested data.
// 	IncomingRequest(data []byte) ([]byte, error)

// 	io.Closer
// }

// // how long a signed request is good for before reject it. Semi-protection against replays.
// const _validityThreshold = 2 //seconds

// /***************************************************************************************
// ** NOTE: This component is used to proxy data requests from approved miner processes. Miner
// ** public addresses are whitelisted and a small history of requests is retained to mitigate
// ** replay attacks. All incoming requests must be signed by the miner making the request so
// ** that the miner's public address can be verified. Best practice is to batch data lookups
// ** into single requests to improve performance and security.
// **
// ** This component does NOT, repeat NOT, prevent DDoS attacks. Users must
// ** use their own solution to prevent such attacks if operating this code in a publicly
// ** accessible environment. USE AT YOUR OWN RISK
// ***************************************************************************************/

// type remoteImpl struct {
// 	privateKey    *ecdsa.PrivateKey
// 	publicAddress string
// 	localDB       DB
// 	whitelist     map[string]bool
// 	postURL       string
// 	logger        log.Logger
// 	wlHistory     map[string]*lru.ARCCache
// 	isRemote      bool
// 	rwLock        sync.RWMutex
// }

// func OpenRemote(logger log.Logger, cfg map[string]string, localDB DB) (DataServerProxy, error) {
// 	return open(logger, cfg, localDB, true)
// }

// func OpenLocal(logger log.Logger, cfg map[string]string, localDB DB) (DataServerProxy, error) {
// 	return open(logger, cfg, localDB, false)
// }

// // OpenRemoteDB establishes a proxy to a remote data server.
// func open(logger log.Logger, cfg map[string]string, localDB DB, isRemote bool) (DataServerProxy, error) {
// 	logger, err := logging.ApplyFilter(cfg, ComponentName, logger)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "apply filter logger")
// 	}

// 	// Using the public key from the first private key.
// 	_privateKeys := os.Getenv(config.PrivateKeysEnvName)
// 	privateKeys := strings.Split(_privateKeys, ",")
// 	privateKey, err := crypto.HexToECDSA(strings.TrimSpace(privateKeys[0]))
// 	if err != nil {
// 		return nil, errors.Wrap(err, "getting private key to ECDSA")
// 	}
// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		return nil, errors.New("casting public key to ECDSA")
// 	}
// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

// 	whitelist := cfg.ServerWhitelist
// 	wlMap := make(map[string]bool)
// 	wlLRU := make(map[string]*lru.ARCCache)
// 	for _, a := range whitelist {
// 		addr := common.HexToAddress(a)
// 		hist, err := lru.NewARC(50)
// 		if err != nil {
// 			return nil, err
// 		}
// 		wlLRU[addr.Hex()] = hist
// 		wlMap[addr.Hex()] = true
// 	}

// 	url := "http://" + cfg.Mine.RemoteDBHost + ":" + strconv.Itoa(int(cfg.Mine.RemoteDBPort))
// 	i := &remoteImpl{
// 		privateKey:    privateKey,
// 		publicAddress: fromAddress.Hex(),
// 		localDB:       localDB,
// 		postURL:       url,
// 		whitelist:     wlMap,
// 		wlHistory:     wlLRU,
// 		logger:        log.With(logger, "component", ComponentName),
// 		isRemote:      isRemote,
// 	}

// 	if isRemote {
// 		level.Info(i.logger).Log(
// 			"msg", "created remote data connector",
// 			"host", cfg.Mine.RemoteDBHost,
// 			"port", cfg.Mine.RemoteDBPort,
// 		)
// 	} else {
// 		level.Info(i.logger).Log(
// 			"msg", "created local data connector",
// 		)
// 	}
// 	return i, nil
// }

// // Check whether an incoming storage request key is prefixed by one of our known miner
// // keys. Otherwise, we have to reject the request as invalid or not coming from this
// // codebase.
// func (i *remoteImpl) hasAddressPrefix(key string) bool {
// 	for k := range i.whitelist {
// 		if strings.HasPrefix(key, k) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (i *remoteImpl) IncomingRequest(data []byte) ([]byte, error) {
// 	req, err := decodeRequest(i.logger, data, i)
// 	if err != nil {
// 		return errorResponse("decoding incoming request")
// 	}

// 	if req == nil {
// 		return errorResponse("decode request!")
// 	}

// 	if req.dbKeys == nil {
// 		return errorResponse("No keys found in request!")
// 	}

// 	if i.localDB == nil {
// 		return errorResponse("Missing localDB instance!")
// 	}

// 	if req.dbValues != nil && len(req.dbValues) > 0 {

// 		//lock out other threads from reading/writing until the write is done
// 		i.rwLock.Lock()

// 		//make sure we unlock when we leave so next thread can get in
// 		defer i.rwLock.Unlock()

// 		//if we're writing values to the DB
// 		if len(req.dbKeys) != len(req.dbValues) {
// 			return errorResponse("Keys and values must have the same array dimensions")
// 		}

// 		//request to write data locally
// 		for idx, k := range req.dbKeys {
// 			//make sure key is prefixed with address
// 			if !i.hasAddressPrefix(k) {
// 				return errorResponse("All remote data storage request keys must be prefixed with miner public Ethereum address")
// 			}
// 			v := req.dbValues[idx]
// 			if err := i.localDB.Put(k, v); err != nil {
// 				return errorResponse(err.Error())
// 			}
// 		}

// 	} else {
// 		//we're not writing, so we just need a read lock
// 		i.rwLock.RLock()
// 		defer i.rwLock.RUnlock()
// 	}

// 	level.Info(i.logger).Log("msg", "getting remote request for keys", req.dbKeys)

// 	outMap := map[string][]byte{}
// 	for _, k := range req.dbKeys {
// 		if req.dbValues == nil && !isKnownKey(k) {
// 			return errorResponse("Invalid lookup key: " + k)
// 		}
// 		level.Debug(i.logger).Log("looking up for local DB key", "key", k)
// 		bts, err := i.localDB.Get(k)

// 		if err != nil {
// 			return errorResponse(err.Error())
// 		}
// 		if bts != nil {
// 			level.Debug(i.logger).Log("msg", "get bytes of result", "bytes", len(bts))
// 			outMap[k] = bts
// 		}
// 	}

// 	resp := &responsePayload{dbVals: outMap, errorMsg: ""}
// 	return encodeResponse(resp)
// }

// func (i *remoteImpl) Get(key string) ([]byte, error) {
// 	keys := []string{key}
// 	res, err := i.BatchGet(keys)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res[key], nil
// }

// func (i *remoteImpl) Put(key string, value []byte) error {
// 	//every key must be prefixed with miner's address

// 	keys := []string{key}
// 	vals := [][]byte{value}
// 	return i.BatchPut(keys, vals)
// }

// func (i *remoteImpl) BatchGet(keys []string) (map[string][]byte, error) {
// 	if !i.isRemote {
// 		outMap := map[string][]byte{}
// 		for _, k := range keys {

// 			bts, err := i.localDB.Get(k)

// 			if err != nil {
// 				return nil, err
// 			}
// 			if bts != nil {
// 				outMap[k] = bts
// 			}
// 		}
// 		level.Debug(i.logger).Log(
// 			"msg", "requested keys result in output",
// 			"keys", keys,
// 			"outMap", outMap,
// 		)
// 		return outMap, nil
// 	}
// 	req, err := createRequest(i.logger, keys, nil, i)
// 	if err != nil {
// 		return nil, err
// 	}
// 	data, err := encodeRequest(i.logger, req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	httpReq := &util.HTTPFetchRequest{Method: util.POST, QueryURL: i.postURL, Payload: data, Timeout: time.Duration(10 * time.Second)}

// 	respData, err := util.HTTPWithRetries(i.logger, httpReq)
// 	if err != nil {
// 		return nil, errors.Wrapf(err, "retrieving data after retries")
// 	}
// 	remResp, err := decodeResponse(respData)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(remResp.errorMsg) > 0 {
// 		return nil, errors.New(remResp.errorMsg)
// 	}
// 	return remResp.dbVals, nil
// }

// func (i *remoteImpl) BatchPut(keys []string, values [][]byte) error {
// 	if !i.isRemote {
// 		if len(values) > 0 && len(keys) != len(values) {
// 			return errors.Errorf("keys and values must have same array dimensions")
// 		}
// 		for idx, k := range keys {
// 			err := i.localDB.Put(k, values[idx])
// 			if err != nil {
// 				return err
// 			}
// 		}

// 		return nil
// 	}

// 	//must prefix all keys with public address
// 	dbKeys := make([]string, len(keys))
// 	for idx, k := range keys {
// 		if !strings.HasPrefix(k, i.publicAddress) {
// 			dbKeys[idx] = i.publicAddress + "-" + k
// 		} else {
// 			dbKeys[idx] = k
// 		}
// 	}
// 	req, err := createRequest(i.logger, dbKeys, values, i)
// 	if err != nil {
// 		return err
// 	}
// 	data, err := encodeRequest(i.logger, req)
// 	if err != nil {
// 		return err
// 	}
// 	httpReq := &util.HTTPFetchRequest{
// 		Method:   util.POST,
// 		QueryURL: i.postURL,
// 		Payload:  data,
// 		Timeout:  time.Duration(10 * time.Second),
// 	}
// 	respData, err := util.HTTPWithRetries(i.logger, httpReq)
// 	if err != nil {
// 		//return nil, err
// 		return errors.Wrap(err, "put data after retries")
// 	}
// 	remResp, err := decodeResponse(respData)
// 	if err != nil {
// 		return err
// 	}
// 	if len(remResp.errorMsg) > 0 {
// 		return errors.New(remResp.errorMsg)
// 	}
// 	return nil
// }

// func (i *remoteImpl) Sign(hash []byte) ([]byte, error) {
// 	return crypto.Sign(hash, i.privateKey)
// }

// func (i *remoteImpl) Verify(hash []byte, timestamp int64, sig []byte) error {
// 	pubKey, err := crypto.SigToPub(hash, sig)
// 	if err != nil {
// 		return err
// 	}
// 	addr := crypto.PubkeyToAddress(*pubKey)
// 	ashex := addr.Hex()
// 	level.Debug(i.logger).Log(
// 		"msg", "verifying signature against whitelist",
// 		"address", ashex,
// 		"whitlisted", i.whitelist[ashex],
// 	)
// 	if !i.whitelist[ashex] {
// 		level.Warn(i.logger).Log("msg", "unauthorized miner detected", "address", ashex)
// 		return errors.Errorf("Unauthorized")
// 	}

// 	cache := i.wlHistory[ashex]
// 	if cache == nil {
// 		return errors.Errorf("No history found for address")
// 	}
// 	if cache.Contains(timestamp) {
// 		level.Debug(i.logger).Log(
// 			"msg", "miner already made request",
// 			"miner", ashex,
// 			"timestamp", timestamp,
// 		)
// 		expr := time.Unix(timestamp+_validityThreshold, 0)
// 		now := time.Now()
// 		if now.After(expr) {
// 			level.Warn(i.logger).Log(
// 				"msg", "request time expired",
// 				"timestamp", time.Unix(timestamp, 0),
// 				"now", now,
// 			)
// 			return errors.Errorf("Request expired")
// 		}
// 		level.Debug(i.logger).Log(
// 			"msg", "time of last request",
// 			"comparing", expr,
// 			"to", now,
// 		)

// 	} else {
// 		level.Debug(i.logger).Log(
// 			"msg", "never seen miner before",
// 			"address", ashex,
// 			"timestamp", timestamp,
// 		)
// 	}
// 	cache.Add(timestamp, true)
// 	return nil
// }

// func (l *remoteImpl) Close() error {
// 	return l.localDB.Close()
// }
