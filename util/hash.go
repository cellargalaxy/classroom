package util

import (
	"crypto/sha256"
)

func Sha256String(data string) string {
	hashed := Sha256([]byte(data))
	return HexEncode(hashed)
}

func Sha256(data []byte) []byte {
	hashed := sha256.Sum256(data)
	return hashed[:]
}
