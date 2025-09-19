package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

// Blockchain is a slice of blocks
type Blockchain []Block

// CalculateHash computes the SHA-256 hash of a block
func (b *Block) CalculateHash() string {
	record := string(b.Index) + b.Timestamp + b.Data + b.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}


// NewBlock creates a new block
func NewBlock(index int, data, prevHash string) Block {
	block := Block{
		Index:     index,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevHash,
	}
	block.Hash = block.CalculateHash()
	return block
}


// GenesisBlock creates the first block in the blockchain
func GenesisBlock() Block {
	return NewBlock(0, "Genesis Block", "")
}


// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := (*bc)[len(*bc)-1]
	newBlock := NewBlock(prevBlock.Index+1, data, prevBlock.Hash)
	*bc = append(*bc, newBlock)
}


// NewBlockchain creates a new blockchain with the genesis block
func NewBlockchain() Blockchain {
	genesisBlock := GenesisBlock()
	return Blockchain{genesisBlock}
}

func main() {
	// Create a new blockchain
	blockchain := NewBlockchain()

	// Add blocks
	blockchain.AddBlock("First Block")
	blockchain.AddBlock("Second Block")

	// Print all blocks
	for _, block := range blockchain {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println("---")
	}
}
