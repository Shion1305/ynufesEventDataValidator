package model

import (
	"crypto/md5"
	"encoding/hex"
)

type ID string

func genID(value string) ID {
	b := []byte(value)
	sum := md5.Sum(b)
	return ID(hex.EncodeToString(sum[:]))
}
