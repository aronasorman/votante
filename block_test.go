package main

import (
	"bytes"
	"crypto/sha256"
	"github.com/satori/go.uuid"
	"testing"
	"testing/quick"
)

func TestBlockIsHashable(t *testing.T) {
	// test whether the return value is the right length
	f := func(b *Block) int {
		return len(b.Hash())
	}

	g := func(b *Block) int {
		return sha256.Size
	}

	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func TestBlockHashChanges(t *testing.T) {
	f := func(b *Block) bool {
		oldhash := b.Hash()

		// change the miner id
		b.Miner.Id = uuid.NewV4()

		newhash := b.Hash()

		// hashes should be different
		return bytes.Compare(oldhash, newhash) != 0
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
