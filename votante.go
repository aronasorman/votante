package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"reflect"
)

type Hasher interface {
	Hash() []byte
}

type Device struct {
	Id uuid.UUID
}

func (d *Device) Generate(rand *rand.Rand, size int) reflect.Value {
	d = &Device{Id: uuid.NewV4()}

	return reflect.ValueOf(d)
}

func (d *Device) Hash() []byte {
	return DoubleSha256([]byte(d.Id.String()))
}

type Block struct {
	Miner   *Device
	Votes   *Votes
	Counter int32
	Nonce   int64
}

func (b *Block) Hash() []byte {
	var totalshasum bytes.Buffer
	totalshasum.Write(b.Miner.Hash())
	totalshasum.Write(b.Votes.Hash())

	return DoubleSha256(totalshasum.Bytes())
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
