package db

import (
	"reflect"
	"testing"
	"time"
)

func setup() (DataServerProxy, DB, error) {
	DB, err := Open("/tmp/test_remoteRequest")
	if err != nil {
		DB.Close()
		return nil, DB, err
	}
	remote, err := OpenRemoteDB(DB)
	return remote, DB, err
}

func TestRemoteRequestCodec(t *testing.T) {
	remote, DB, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	defer DB.Close()

	keys := []string{RequestIdKey, DifficultyKey}
	req, err := createRequest(keys, remote)
	if err != nil {
		t.Fatal(err)
	}
	if req.timestamp == 0 {
		t.Fatal("Expected a timestamp to get applied to request")
	}
	if req.sig == nil {
		t.Fatal("Expected a signature to be attached to request")
	}
	if req.dbKeys == nil || len(req.dbKeys) == 0 {
		t.Fatal("Expected request to have dbKeys")
	}

	encoded, err := encodeRequest(req)
	if err != nil {
		t.Fatal(err)
	}
	decReq, err := decodeRequest(encoded, remote)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(decReq.sig, req.sig) {
		t.Fatal("Signatures did not match after codec")
	}
	if !reflect.DeepEqual(decReq.dbKeys, req.dbKeys) {
		t.Fatal("DBKeys did not match after codec")
	}
	if decReq.timestamp != req.timestamp {
		t.Fatal("Timestamps do not match after codec")
	}
}

func TestRequestReplayAttack(t *testing.T) {

	remote, DB, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	defer DB.Close()

	keys := []string{RequestIdKey, DifficultyKey}
	req, err := createRequest(keys, remote)
	if err != nil {
		t.Fatal(err)
	}
	if req.timestamp == 0 {
		t.Fatal("Expected a timestamp to get applied to request")
	}
	if req.sig == nil {
		t.Fatal("Expected a signature to be attached to request")
	}
	if req.dbKeys == nil || len(req.dbKeys) == 0 {
		t.Fatal("Expected request to have dbKeys")
	}

	encoded, err := encodeRequest(req)
	if err != nil {
		t.Fatal(err)
	}
	_, err = decodeRequest(encoded, remote)
	if err != nil {
		t.Fatal(err)
	}

	//that simulated a call that was decoded. Now we'll wait for timeout on request
	time.Sleep((_validityThreshold * 1500) * time.Millisecond)

	_, err = decodeRequest(encoded, remote)
	if err == nil {
		t.Fatal("Expected failure when decoding request as a replay after expiration period")
	}
}

func TestRequestForData(t *testing.T) {
	remote, DB, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	defer DB.Close()
	DB.Delete(RequestIdKey)
	DB.Delete(DifficultyKey)
	err = DB.Put(RequestIdKey, []byte("1"))
	if err != nil {
		t.Fatal(err)
	}
	err = DB.Put(DifficultyKey, []byte("2"))
	if err != nil {
		t.Fatal(err)
	}
	keys := []string{RequestIdKey, DifficultyKey}
	req, err := createRequest(keys, remote)
	if err != nil {
		t.Fatal(err)
	}
	encoded, err := encodeRequest(req)
	if err != nil {
		t.Fatal(err)
	}
	data, err := remote.IncomingRequest(encoded)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := decodeResponse(data, remote)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.errorMsg) != 0 {
		t.Fatal(resp.errorMsg)
	}
	reqID := string(resp.dbVals[RequestIdKey])
	diff := string(resp.dbVals[DifficultyKey])
	if reqID != "1" {
		t.Fatalf("Expected result map to map request id to '1': %v", resp.dbVals)
	}
	if diff != "2" {
		t.Fatalf("Expected difficulty to be mapped to '2': %v", resp.dbVals)
	}

}
