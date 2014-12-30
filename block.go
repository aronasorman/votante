package main

import (
	"bytes"
	"errors"
	"math"
	"strconv"
)

type Block struct {
	Miner      *Device
	Votes      *Votes
	Counter    int32
	Nonce      int64
	Difficulty int32
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
