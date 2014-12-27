package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

type Hasher interface {
	Hash() []byte
}

type Device struct {
	Id uint64
}

func (d *Device) Hash() []byte {
	idBinary := make([]byte, 32)
	binary.PutUvarint(idBinary, d.Id)

	return DoubleSha256(idBinary)
}

type Block struct {
	Miner Device
	Votes *Votes
}

type Vote struct {
	For string
}

func (v *Vote) Hash() []byte {
	forb := []byte(v.For)
	return DoubleSha256(forb)
}

type Votes []Vote

func (vs *Votes) Hash() []byte {
	votesarray := []Vote(*vs)
	var totalshasum bytes.Buffer
	for _, vote := range votesarray {
		votesha := vote.Hash()
		totalshasum.Write(votesha)
	}

	return DoubleSha256(totalshasum.Bytes())
}

func DoubleSha256(data []byte) []byte {
	sha := sha256.New()
	sha.Write(data)
	hash := sha.Sum(nil)
	sha.Reset()
	sha.Write(hash)
	return sha.Sum(nil)
}

func main() {
	// block := newBlock()
	fmt.Println("Hello world")
}
