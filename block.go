package main

import (
	"bytes"
	"errors"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"testing/quick"
)

type Block struct {
	Miner      *Device
	Votes      *Votes
	Counter    int32
	Nonce      int64
	Difficulty int32
	Previous   *Block
}

func (b *Block) Generate(r *rand.Rand, size int) reflect.Value {
	miner, _ := quick.Value(reflect.TypeOf(&Device{}), r)
	votes, _ := quick.Value(reflect.TypeOf(&Votes{}), r)
	counter, _ := quick.Value(reflect.TypeOf(int32(0)), r)
	nonce, _ := quick.Value(reflect.TypeOf(int64(0)), r)
	difficulty, _ := quick.Value(reflect.TypeOf(int32(0)), r)

	b = &Block{}
	reflect.ValueOf(&b.Miner).Elem().Set(miner)
	reflect.ValueOf(&b.Votes).Elem().Set(votes)
	reflect.ValueOf(&b.Counter).Elem().Set(counter)
	reflect.ValueOf(&b.Nonce).Elem().Set(nonce)
	reflect.ValueOf(&b.Difficulty).Elem().Set(difficulty)

	// we leave the Previous block as nil

	return reflect.ValueOf(b)

}

func (b *Block) Hash() []byte {
	var totalshasum bytes.Buffer
	totalshasum.Write(b.Miner.Hash())
	totalshasum.Write(b.Votes.Hash())
	totalshasum.Write([]byte(strconv.Itoa(int(b.Nonce))))
	totalshasum.Write([]byte(strconv.Itoa(int(b.Counter))))

	return DoubleSha256(totalshasum.Bytes())
}

func (b *Block) Mine() error {
	var i int64
	for ; i <= math.MaxInt32; i++ {
		b.Nonce = i
		if b.Valid() {
			return nil
		}
	}

	return errors.New("Couldn't find valid block")
}

func (b *Block) Valid() bool {
	for i, byte := range b.Hash() {
		if int32(i) > b.Difficulty {
			return true
		} else if byte != 0 {
			return false
		}
	}
	return false
}
