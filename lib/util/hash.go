package util

import (
	"crypto/sha256"
	"encoding/base64"
)

// Hash retorna um checksum SHA224 para o input
func Hash(input string) string {
	sum := sha256.Sum224([]byte(input))
	return base64.
		RawURLEncoding.
		EncodeToString(sum[:])
}
