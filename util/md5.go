package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(s string) string {
	b := md5.Sum([]byte(s))
	return hex.EncodeToString(b[:])
}
