package main

type Election struct {
	Type   ElectionType
	Result *Block
}

type ElectionType func(*Block) map[string]int
