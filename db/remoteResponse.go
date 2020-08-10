package db

import (
	"bytes"
)

//
// --- Response from remote request contains either an error message
// --- or a map of requested keys and their values
//
type responsePayload struct {
	errorMsg string
	dbVals   map[string][]byte
}

/**
 * Create an outgoing response
 */
func createResponse(dbVals map[string][]byte, errorMsg string) (*responsePayload, error) {
	return &responsePayload{errorMsg: errorMsg, dbVals: dbVals}, nil
}

/**
 * Encode the given request for transport over the wire
 */
func encodeResponse(r *responsePayload) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := encodeString(buf, r.errorMsg); err != nil {
		return nil, err
	}

	if r.dbVals != nil {
		if err := encode(buf, uint32(len(r.dbVals))); err != nil {
			return nil, err
		}
		for k, v := range r.dbVals {
			if err := encodeString(buf, k); err != nil {
				return nil, err
			}
			if err := encodeBytes(buf, v); err != nil {
				return nil, err
			}
		}
	} else {
		if err := encode(buf, uint32(0)); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

/**
 * Decode a request from the given bytes. The signer is used to validate keys
 * and whitelisted miners
 */
func decodeResponse(data []byte, validator RequestValidator) (*responsePayload, error) {
	buf := bytes.NewReader(data)
	errMsg, err := decodeString(buf)
	if err != nil {
		return nil, err
	}
	mapLen := uint32(0)
	if err = decode(buf, &mapLen); err != nil {
		return nil, err
	}
	dbVals := make(map[string][]byte)
	for i := uint32(0); i < mapLen; i++ {
		k, err := decodeString(buf)
		if err != nil {
			return nil, err
		}
		bts, err := decodeBytes(buf)
		if err != nil {
			return nil, err
		}
		dbVals[k] = bts
	}

	return &responsePayload{errorMsg: errMsg, dbVals: dbVals}, nil
}

/**
 * convenience function to encode an error message for remote caller
 */
func errorResponse(msg string) ([]byte, error) {
	r := &responsePayload{errorMsg: msg}
	return encodeResponse(r)
}
