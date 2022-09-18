package model

import (
	"crypto/md5"
	"encoding/hex"
)

type ID string

func genID(value string) string {
	b := []byte(value)
	sum := md5.Sum(b)
	return hex.EncodeToString(sum[:])
}
