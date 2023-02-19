package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// EncryptedMagicKeyShareKeyPrefix is the prefix to retrieve all EncryptedMagicKeyShare
	EncryptedMagicKeyShareKeyPrefix = "EncryptedMagicKeyShare/value/"
)

// EncryptedMagicKeyShareKey returns the store key to retrieve a EncryptedMagicKeyShare from the index fields
func EncryptedMagicKeyShareKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
