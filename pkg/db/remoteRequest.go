// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package db

import (
	"bytes"
	"io"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
)

// RequestSigner handles signing an outgoing request. It's just an abstraction
// so we can test, etc.
type RequestSigner interface {
	// Sign the given payload hash with a private key and return the signature bytes.
	Sign(payload []byte) ([]byte, error)
}

// RequestValidator validates that a miner's signature is valid, that its address
// is whitelisted, and minizes chances that the requested hash isn't being replayed.
type RequestValidator interface {
	// Verify the given signature was signed by a valid/whitelisted miner address.
	Verify(hash []byte, timestamp int64, sig []byte) error
}

// Request payload is encoded and comes from a remote client (miner) that is
// asking for specific data. Every request has a signature to verify it's
// coming from a whitelisted client and that it has not been already requested
// based on its timestamp.
type requestPayload struct {

	// dbKeys to access the DB.
	dbKeys []string

	// dbValues to store in the DB.
	dbValues [][]byte

	// timestamp when the request was sent. Aids in avoiding replay attacks.
	timestamp int64

	// signature of op, dbKey, dbVal, and timestamp.
	sig []byte
}

// Create an outgoing request for the given keys.
func createRequest(logger log.Logger, dbKeys []string, values [][]byte, signer RequestSigner) (*requestPayload, error) {

	t := time.Now().Unix()
	buf := new(bytes.Buffer)
	level.Debug(logger).Log("msg", "encoding initial keys and timestamp")
	err := encodeKeysValuesAndTime(logger, buf, dbKeys, values, t)
	if err != nil {
		return nil, err
	}

	level.Debug(logger).Log("msg", "generating request hash")
	hash := crypto.Keccak256(buf.Bytes())
	level.Debug(logger).Log("msg", "signing hash")
	sig, err := signer.Sign(hash)

	if err != nil {
		return nil, errors.Wrap(err, "signing hash")
	}
	if sig == nil {
		return nil, errors.Errorf("generate a signature for hash: %v", hash)
	}
	return &requestPayload{dbKeys: dbKeys, dbValues: values, timestamp: t, sig: sig}, nil
}

// Since we use keys and time for sig hashing, we have a specific function for
// encoding just those parts.
func encodeKeysValuesAndTime(logger log.Logger, buf *bytes.Buffer, dbKeys []string, values [][]byte, timestamp int64) error {

	level.Debug(logger).Log("msg", "encoding timestamp")
	if err := encode(buf, timestamp); err != nil {
		return err
	}
	if dbKeys == nil {
		return errors.Errorf("no keys to encode")
	}

	level.Debug(logger).Log("msg", "encoding dbKeys")
	if err := encode(buf, uint32(len(dbKeys))); err != nil {
		return errors.Wrap(err, "encoding dbKeys")
	}
	for _, k := range dbKeys {
		level.Debug(logger).Log("msg", "encoding key", k)
		if err := encodeString(buf, k); err != nil {
			return errors.Wrap(err, "encoding key")
		}
	}

	if values != nil {
		if err := encode(buf, uint32(len(values))); err != nil {
			return errors.Wrap(err, "encoding values length")
		}
		for _, v := range values {
			if err := encodeBytes(buf, v); err != nil {
				return errors.Wrap(err, "encoding value bytes")
			}
		}

	} else {
		if err := encode(buf, uint32(0)); err != nil {
			return errors.Wrap(err, "encode zero value")
		}
	}

	return nil
}

// Decodes just the key and timestamp portions of a buffer.
func decodeKeysValuesAndTime(buf io.Reader) ([]string, [][]byte, int64, error) {
	var time int64
	if err := decode(buf, &time); err != nil {
		return nil, nil, 0, err
	}
	len := uint32(0)
	if err := decode(buf, &len); err != nil {
		return nil, nil, 0, err
	}
	dbKeys := make([]string, len)
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
	values := make([][]byte, len)
	for i := uint32(0); i < len; i++ {
		bts, err := decodeBytes(buf)
		if err != nil {
			return nil, nil, 0, err
		}
		values[i] = bts
	}
	return dbKeys, values, time, nil
}

// Encode the given request for transport over the wire.
func encodeRequest(logger log.Logger, r *requestPayload) ([]byte, error) {
	buf := new(bytes.Buffer)

	if r.dbKeys == nil || len(r.dbKeys) == 0 {
		return nil, errors.Errorf("No keys in request. No point in making a request if there are no keys")
	}
	if r.sig == nil {
		return nil, errors.Errorf("Cannot encode a request without a signature attached")
	}

	// Capture keys and timestamp.
	level.Debug(logger).Log("msg", "encoding keys and time")
	if err := encodeKeysValuesAndTime(logger, buf, r.dbKeys, r.dbValues, r.timestamp); err != nil {
		return nil, errors.Wrap(err, "encoding keys and time")
	}

	// Then if there is a sig, encode it.
	level.Debug(logger).Log("msg", "encoding signature")
	if err := encodeBytes(buf, r.sig); err != nil {
		return nil, errors.Wrap(err, "encoding signature")
	}

	return buf.Bytes(), nil
}

// Decode a request from the given bytes. The signer is used to validate keys
// and whitelisted miners.
func decodeRequest(logger log.Logger, data []byte, validator RequestValidator) (*requestPayload, error) {
	buf := bytes.NewReader(data)
	keys, vals, time, err := decodeKeysValuesAndTime(buf)
	if err != nil {
		return nil, err
	}
	if len(keys) == 0 {
		return nil, errors.Errorf("No dbKeys in incoming request")
	}
	sig, err := decodeBytes(buf)
	if err != nil {
		return nil, err
	}
	hBuf := new(bytes.Buffer)
	if err := encodeKeysValuesAndTime(logger, hBuf, keys, vals, time); err != nil {
		return nil, err
	}
	hash := crypto.Keccak256(hBuf.Bytes())
	if err := validator.Verify(hash, time, sig); err != nil {
		return nil, err
	}
	return &requestPayload{dbKeys: keys, dbValues: vals, timestamp: time, sig: sig}, nil
}
