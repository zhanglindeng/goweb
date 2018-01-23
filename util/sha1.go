package util

import (
	"encoding/hex"
	"crypto/sha1"
)

func Sha1(s string) string {
	b := sha1.Sum([]byte(s))
	return hex.EncodeToString(b[:])
}
