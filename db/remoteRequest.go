package db

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/tellor-io/TellorMiner/util"
)

type requestType int

//RequestSigner handles signing an outgoing request. It's just an abstraction
//so we can test, etc.
type RequestSigner interface {
	//Sign the given payload hash with a private key and return the
	//signature bytes
	Sign(payload []byte) ([]byte, error)
}

//RequestValidator validates that a miner's signature is valid, that its address
//is whitelisted, and minizes chances that the requested hash isn't being replayed
type RequestValidator interface {
	//Verify the given signature was signed by a valid/whitelisted miner address
	Verify(hash []byte, timestamp int64, sig []byte) error
}

//
// --- Request payload is encoded and comes from a remote client (miner) that is
// --- asking for specific data. Every request has a signature to verify it's
// --- coming from a whitelisted client and that it has not been already requested
// --- based on its timestamp
//
type requestPayload struct {

	//key to access the DB
	dbKeys []string

	//values to store in the DB
	dbValues [][]byte

	//time when the request was sent. Aids in avoiding replay attacks
	timestamp int64

	//signature of op, dbKey, dbVal, and timestamp
	sig []byte
}

var rrlog *util.Logger = util.NewLogger("db", "RemoteRequest")

/**
 * Create an outgoing request for the given keys
 */
func createRequest(dbKeys []string, values [][]byte, signer RequestSigner) (*requestPayload, error) {

	//rrlog = util.NewLogger("db", "RemoteRequest")
	t := time.Now().Unix()
	buf := new(bytes.Buffer)
	rrlog.Debug("Encoding initial keys and timestamp")
	err := encodeKeysValuesAndTime(buf, dbKeys, values, t)
	if err != nil {
		return nil, err
	}

	log.Debug("Generating request hash...")
	hash := crypto.Keccak256(buf.Bytes())
	log.Debug("Signing hash")
	sig, err := signer.Sign(hash)

	if err != nil {
		log.Error("Signature failed", err.Error())
		return nil, err
	}
	if sig == nil {
		log.Error("Signature was not generated")
		return nil, fmt.Errorf("Could not generate a signature for  hash: %v", hash)
	}
	return &requestPayload{dbKeys: dbKeys, dbValues: values, timestamp: t, sig: sig}, nil
}

/**
 * Since we use keys and time for sig hashing, we have a specific function for
 * encoding just those parts
 **/
func encodeKeysValuesAndTime(buf *bytes.Buffer, dbKeys []string, values [][]byte, timestamp int64) error {

	rrlog.Debug("Encoding timestamp")
	if err := encode(buf, timestamp); err != nil {
		return err
	}
	if dbKeys == nil {
		rrlog.Error("No keys to encode")
		return fmt.Errorf("No keys to encode")
	}

	rrlog.Debug("Encoding dbKeys")
	if err := encode(buf, uint32(len(dbKeys))); err != nil {
		rrlog.Error("Problem encoding dbKeys", err.Error())
		return err
	}
	for _, k := range dbKeys {
		rrlog.Debug("Encoding key", k)
		if err := encodeString(buf, k); err != nil {
			rrlog.Error("Problem encoding key", err.Error())
			return err
		}
	}

	if values != nil {
		if err := encode(buf, uint32(len(values))); err != nil {
			rrlog.Error("Problem encoding values length", err.Error())
			return err
		}
		for _, v := range values {
			if err := encodeBytes(buf, v); err != nil {
				rrlog.Error("Problem encoding value bytes", err.Error())
				return err
			}
		}

	} else {
		if err := encode(buf, uint32(0)); err != nil {
			rrlog.Error("Could not encode zero value", err.Error())
			return err
		}
	}

	return nil
}

/**
 * Decodes just the key and timestamp portions of a buffer
 */
func decodeKeysValuesAndTime(buf io.Reader) ([]string, [][]byte, int64, error) {
	var time int64
	if err := decode(buf, &time); err != nil {
		return nil, nil, 0, err
	}
	len := uint32(0)
	if err := decode(buf, &len); err != nil {
		return nil, nil, 0, err
	}
	dbKeys := make([]string, len, len)
	for i := uint32(0); i < len; i++ {
		s, err := decodeString(buf)
		if err != nil {
			return nil, nil, 0, err
		}
		dbKeys[i] = s
	}
	len = uint32(0)
	if err := decode(buf, &len); err != nil {
		return nil, nil, 0, err
	}
	values := make([][]byte, len, len)
	for i := uint32(0); i < len; i++ {
		bts, err := decodeBytes(buf)
		if err != nil {
			return nil, nil, 0, err
		}
		values[i] = bts
	}
	return dbKeys, values, time, nil
}

/**
 * Encode the given request for transport over the wire
 */
func encodeRequest(r *requestPayload) ([]byte, error) {
	buf := new(bytes.Buffer)

	if r.dbKeys == nil || len(r.dbKeys) == 0 {
		return nil, fmt.Errorf("No keys in request. No point in making a request if there are no keys")
	}
	if r.sig == nil {
		return nil, fmt.Errorf("Cannot encode a request without a signature attached")
	}

	//capture keys and timestamp
	rrlog.Debug("Encoding keys and time...")
	if err := encodeKeysValuesAndTime(buf, r.dbKeys, r.dbValues, r.timestamp); err != nil {
		rrlog.Error("Problem encoding keys and time", err)
		return nil, err
	}

	//then if there is a sig, encode it
	rrlog.Debug("Encoding signature...")
	if err := encodeBytes(buf, r.sig); err != nil {
		rrlog.Error("Problem encoding signature", err)
		return nil, err
	}

	return buf.Bytes(), nil
}

/**
 * Decode a request from the given bytes. The signer is used to validate keys
 * and whitelisted miners
 */
func decodeRequest(data []byte, validator RequestValidator) (*requestPayload, error) {
	buf := bytes.NewReader(data)
	keys, vals, time, err := decodeKeysValuesAndTime(buf)
	if err != nil {
		return nil, err
	}
	if len(keys) == 0 {
		return nil, fmt.Errorf("No dbKeys in incoming request")
	}
	sig, err := decodeBytes(buf)
	if err != nil {
		return nil, err
	}
	hBuf := new(bytes.Buffer)
	if err := encodeKeysValuesAndTime(hBuf, keys, vals, time); err != nil {
		return nil, err
	}
	hash := crypto.Keccak256(hBuf.Bytes())
	if err := validator.Verify(hash, time, sig); err != nil {
		return nil, err
	}
	return &requestPayload{dbKeys: keys, dbValues: vals, timestamp: time, sig: sig}, nil
}
