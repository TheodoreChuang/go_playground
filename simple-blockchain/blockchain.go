package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Block - a block
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	PrevHash  string
	Hash      string
}

// Message - data for a block
type Message struct {
	BPM int
}

// Blockchain - global state
var Blockchain []Block

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(prevBlock Block, BPM int) (Block, error) {
	var nextBlock Block
	t := time.Now()

	nextBlock.Index = prevBlock.Index + 1
	nextBlock.Timestamp = t.String()
	nextBlock.BPM = BPM
	nextBlock.PrevHash = prevBlock.Hash
	nextBlock.Hash = calculateHash(nextBlock)

	return nextBlock, nil
}

func isBlockValid(nextBlock, prevBlock Block) bool {
	if prevBlock.Index+1 != nextBlock.Index {
		return false
	}

	if prevBlock.Hash != nextBlock.PrevHash {
		return false
	}

	if calculateHash(nextBlock) != nextBlock.Hash {
		return false
	}

	return true
}

func replaceChain(newBlocks []Block) {
	// longest chain rules
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
