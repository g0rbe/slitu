package slitu

import (
	"crypto/sha256"
	"encoding/hex"
)

// SHA256Str calculates the SHA256 checksum of v and returns the hex encoded string result.
func SHA256Str(v string) string {

	r := sha256.Sum256([]byte(v))

	return hex.EncodeToString(r[:])
}
