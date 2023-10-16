package utils

import (
	"crypto/sha256"
)

// SHA256 calculates the hash for the bytes.
func SHA256(b []byte) []byte {
	h := sha256.Sum256(b)
	return h[:]
}
