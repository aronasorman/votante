package main

import (
	"bytes"
)

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
