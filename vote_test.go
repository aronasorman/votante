package main

import (
	"bytes"
	"crypto/sha256"
	"testing"
	"testing/quick"
)

func TestVoteIsHashable(t *testing.T) {
	f := func(v *Vote) []byte {
		forb := []byte(v.For)
		digest := DoubleSha256(forb)
		return digest
	}

	g := func(v *Vote) []byte {
		return v.Hash()
	}

	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}

}

func TestVotesHashChangeWhenSingleVoteChanges(t *testing.T) {
	testvotes := Votes{
		Vote{For: "person1"},
		Vote{For: "person2"},
	}

	oldhash := testvotes.Hash()

	// change the hash of one vote
	testvotes[0].For = "Person 1 modified"

	newhash := testvotes.Hash()

	if bytes.Compare(oldhash, newhash) == 0 {
		t.Errorf("Votes hash still the same!")
	}
}

func TestVotesOnlyLimitedSize(t *testing.T) {
	f := func(vs *Votes) bool {
		return len(vs.Hash()) == sha256.Size
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
