// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package db

import (
	"bytes"

	"github.com/pkg/errors"

	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/testutil"
)

func TestRemoteRequestCodec(t *testing.T) {
	cfg := config.OpenTestConfig(t)
	cfg.ServerWhitelist = []string{"0x053b09e98ede40997546e8bb812cd838f18bb146"}

	DB, cleanup := OpenTestDB(t)
	defer t.Cleanup(cleanup)
	remote, err := OpenRemoteDB(DB)
	testutil.Ok(t, err)

	keys := []string{RequestIdKey, DifficultyKey}
	req, err := createRequest(keys, nil, remote.(*remoteImpl))
	testutil.Ok(t, err)

	testutil.Assert(t, req.timestamp > 0, "Expected a timestamp to get applied to request")
	testutil.Assert(t, req.sig != nil, "Expected a signature to be attached to request")
	testutil.Assert(t, req.dbKeys != nil && len(req.dbKeys) > 0, "Expected request to have dbKeys")

	encoded, err := encodeRequest(req)
	testutil.Ok(t, err)
	decReq, err := decodeRequest(encoded, remote.(*remoteImpl))
	testutil.Ok(t, err)

	testutil.Assert(t, reflect.DeepEqual(decReq.sig, req.sig), "Signatures did not match after codec")
	testutil.Assert(t, reflect.DeepEqual(decReq.dbKeys, req.dbKeys), "DBKeys did not match after codec")
	testutil.Assert(t, decReq.timestamp == req.timestamp, "Timestamps do not match after codec")

}

func TestRequestReplayAttack(t *testing.T) {
	cfg := config.OpenTestConfig(t)
	cfg.ServerWhitelist = []string{"0x053b09e98ede40997546e8bb812cd838f18bb146"}

	DB, cleanup := OpenTestDB(t)
	defer t.Cleanup(cleanup)
	remote, err := OpenRemoteDB(DB)
	testutil.Ok(t, err)

	keys := []string{RequestIdKey, DifficultyKey}
	req, err := createRequest(keys, nil, remote.(*remoteImpl))
	testutil.Ok(t, err)

	testutil.Assert(t, req.timestamp > 0, "Expected a timestamp to get applied to request")
	testutil.Assert(t, req.sig != nil, "Expected a signature to be attached to request")
	testutil.Assert(t, req.dbKeys != nil && len(req.dbKeys) > 0, "Expected request to have dbKeys")

	encoded, err := encodeRequest(req)
	testutil.Ok(t, err)

	_, err = decodeRequest(encoded, remote.(*remoteImpl))
	testutil.Ok(t, err)

	// That simulated a call that was decoded. Now we'll wait for timeout on request.
	time.Sleep((_validityThreshold * 1500) * time.Millisecond)

	_, err = decodeRequest(encoded, remote.(*remoteImpl))
	testutil.Ok(t, err)
}

func TestRequestForData(t *testing.T) {
	cfg := config.OpenTestConfig(t)
	cfg.ServerWhitelist = []string{"0x053b09e98ede40997546e8bb812cd838f18bb146"}

	DB, cleanup := OpenTestDB(t)
	defer t.Cleanup(cleanup)
	remote, err := OpenRemoteDB(DB)
	testutil.Ok(t, err)

	if err := DB.Delete(RequestIdKey); err != nil {
		testutil.Ok(t, err)
	}
	if err := DB.Delete(DifficultyKey); err != nil {
		testutil.Ok(t, err)
	}
	if err := DB.Put(RequestIdKey, []byte("1")); err != nil {
		testutil.Ok(t, err)
	}
	err = DB.Put(DifficultyKey, []byte("2"))
	testutil.Ok(t, err)

	keys := []string{RequestIdKey, DifficultyKey}
	req, err := createRequest(keys, nil, remote.(*remoteImpl))
	testutil.Ok(t, err)

	encoded, err := encodeRequest(req)

	testutil.Ok(t, err)
	data, err := remote.IncomingRequest(encoded)
	testutil.Ok(t, err)

	resp, err := decodeResponse(data)
	testutil.Ok(t, err)

	if len(resp.errorMsg) != 0 {
		testutil.Ok(t, errors.New(resp.errorMsg))
	}
	reqID := string(resp.dbVals[RequestIdKey])
	diff := string(resp.dbVals[DifficultyKey])
	testutil.Assert(t, reqID == "1", "Expected result map to map request id to '1': %v", resp.dbVals)
	testutil.Assert(t, diff == "2", "Expected difficulty to be mapped to '2': %v", resp.dbVals)

}

func TestRequestPut(t *testing.T) {
	cfg := config.OpenTestConfig(t)
	cfg.ServerWhitelist = []string{"0x053b09e98ede40997546e8bb812cd838f18bb146"}

	DB, cleanup := OpenTestDB(t)
	defer t.Cleanup(cleanup)
	remote, err := OpenRemoteDB(DB)
	testutil.Ok(t, err)

	_fromAddress := cfg.PublicAddress

	// Convert to address
	fromAddress := common.HexToAddress(_fromAddress)
	pubKey := strings.ToLower(fromAddress.Hex())
	dbKey := pubKey + "-" + CurrentChallengeKey
	vals := make([][]byte, 1)
	vals[0] = []byte("TEST_CHALLENGE")
	req, err := createRequest([]string{dbKey}, vals, remote.(*remoteImpl))
	testutil.Ok(t, err)

	err = DB.Put(dbKey, vals[0])
	testutil.Ok(t, err)

	bts, err := encodeRequest(req)

	_, err = remote.IncomingRequest(bts)
	testutil.Ok(t, err)

	data, err := DB.Get(dbKey)
	testutil.Ok(t, err)

	testutil.Assert(t, bytes.Equal(data, vals[0]), "DB bytes did not match expected put request data")

}
