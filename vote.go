package main

import (
	"bytes"
)

type Vote struct {
	For string
}

func (v *Vote) ComputeHash() []byte {
	forb := []byte(v.For)
	return DoubleSha256(forb)
}

type Votes []Vote

func (vs *Votes) ComputeHash() []byte {
	if vs == nil {
		return nil
	}

	votesarray := []Vote(*vs)
	var totalshasum bytes.Buffer
	for _, vote := range votesarray {
		votesha := vote.ComputeHash()
		totalshasum.Write(votesha)
	}

	return DoubleSha256(totalshasum.Bytes())
}
