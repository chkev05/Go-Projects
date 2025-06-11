package main

import (
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	Index     int    // Index of the block in the chain
	Timestamp string // time data is written to the block
	BPM       int    // Beats per minute, the data we want to store
	Hash      string // SHA256 hash of the block
	PrevHash  string // Hash of the previous block
}

var Blockchain []Block

func calculateHash(block Block) string {
	// This function should calculate the SHA256 hash of the block
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
