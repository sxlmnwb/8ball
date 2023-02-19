package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ScriptureKeyPrefix is the prefix to retrieve all Scripture
	ScriptureKeyPrefix = "Scripture/value/"
)

// ScriptureKey returns the store key to retrieve a Scripture from the index fields
func ScriptureKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
