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
	// test that the hash changes when the miner changes
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

	// test that the hash changes when the nonce changes
	f = func(b *Block) bool {
		oldhash := b.Hash()

		// change the nonce
		b.Nonce++

		newhash := b.Hash()

		// hashes should be different
		return bytes.Compare(oldhash, newhash) != 0
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMinedBlockHasCorrectHash(t *testing.T) {
	minerID, err := uuid.FromString("edc39bcc-166d-442b-8dbe-7b47facd1158")
	if err != nil {
		panic(err)
	}

	block := Block{
		Difficulty: 1,
		Miner:      &Device{Id: minerID},
		Votes: &Votes{
			Vote{For: "Trapo1"},
			Vote{For: "Trapo2"},
		},
	}

	err = block.Mine()
	if err != nil {
		t.Error("Couldn't mine an easy block!")
	}

	if !block.Valid() {
		t.Error("Mined block isn't valid!")
	}
}
