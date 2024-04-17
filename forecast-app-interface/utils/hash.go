package utils

import "crypto/sha512"

func HashValue(value []byte) []byte {
	h := sha512.New()
	h.Write(value)
	return h.Sum(nil)
}
