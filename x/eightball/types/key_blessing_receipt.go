package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// BlessingReceiptKeyPrefix is the prefix to retrieve all BlessingReceipt
	BlessingReceiptKeyPrefix = "BlessingReceipt/value/"
)

// BlessingReceiptKey returns the store key to retrieve a BlessingReceipt from the index fields
func BlessingReceiptKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
