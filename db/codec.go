package db

import (
	"bytes"
	"encoding/binary"
	"io"
)

func encode(buf *bytes.Buffer, any interface{}) error {
	return binary.Write(buf, binary.BigEndian, any)
}

func decode(buf io.Reader, any interface{}) error {
	return binary.Read(buf, binary.BigEndian, any)
}

/**
 * encodeBytes to the given bufer by encoding array length and
 * content
 */
func encodeBytes(buf *bytes.Buffer, bts []byte) error {
	if err := encode(buf, uint32(len(bts))); err != nil {
		return err
	}
	return encode(buf, bts)
}

/**
 * decodeBytes from the given buffer by reading a length and then
* the bytes
*/
func decodeBytes(buf io.Reader) ([]byte, error) {

	len := uint32(0)
	if err := decode(buf, &len); err != nil {
		return nil, err
	}
	bts := make([]byte, len, len)
	if err := decode(buf, &bts); err != nil {
		return nil, err
	}
	return bts, nil
}

/**
 * encodeString encodes the length/bytes of the given string to the buffer
 */
func encodeString(buf *bytes.Buffer, str string) error {
	return encodeBytes(buf, []byte(str))
}

/**
 * decodeString decodes the length/bytes of a string from the given stream
 */
func decodeString(buf io.Reader) (string, error) {
	bts, err := decodeBytes(buf)
	if err != nil {
		return "", err
	}
	return string(bts), nil
}
