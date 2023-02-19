package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// EncryptedPreSignKeyPrefix is the prefix to retrieve all EncryptedPreSign
	EncryptedPreSignKeyPrefix = "EncryptedPreSign/value/"
)

// EncryptedPreSignKey returns the store key to retrieve a EncryptedPreSign from the index fields
func EncryptedPreSignKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
