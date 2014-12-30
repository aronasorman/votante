package main

import (
	"crypto/sha256"
)

type Hasher interface {
	Hash() []byte
}

func DoubleSha256(data []byte) []byte {
	sha := sha256.New()
	sha.Write(data)
	hash := sha.Sum(nil)
	sha.Reset()
	sha.Write(hash)
	return sha.Sum(nil)
}
