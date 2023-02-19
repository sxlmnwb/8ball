package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SignedScriptureKeyPrefix is the prefix to retrieve all SignedScripture
	SignedScriptureKeyPrefix = "SignedScripture/value/"
)

// SignedScriptureKey returns the store key to retrieve a SignedScripture from the index fields
func SignedScriptureKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
