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
		return v.ComputeHash()
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

	oldhash := testvotes.ComputeHash()

	// change the hash of one vote
	testvotes[0].For = "Person 1 modified"

	newhash := testvotes.ComputeHash()

	if bytes.Compare(oldhash, newhash) == 0 {
		t.Errorf("Votes hash still the same!")
	}
}

func TestVotesHashOnlyLimitedSize(t *testing.T) {
	f := func(vs *Votes) bool {
		return len(vs.ComputeHash()) == sha256.Size
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
