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

// Validate Blockchain
func (bc Blockchain) IsValid() bool {
    for i := 1; i < len(bc); i++ {
        current := bc[i]
        previous := bc[i-1]
        if current.Hash != current.CalculateHash() || current.PrevHash != previous.Hash {
            return false
        }
    }
    return true
}

type ProofOfWork struct {
    Block  *Block
    Target string // e.g., "0000"
}

func NewProofOfWork(b *Block) *ProofOfWork {
    target := "0000" // Difficulty: 4 leading zeros
    return &ProofOfWork{b, target}
}

func (pow *ProofOfWork) Run() (int, string) {
    nonce := 0
    var hash string
    for {
        hash = pow.Block.CalculateHashWithNonce(nonce)
        if strings.HasPrefix(hash, pow.Target) {
            break
        }
        nonce++
    }
    return nonce, hash
}

// Add this method to the Block struct:
func (b *Block) CalculateHashWithNonce(nonce int) string {
    record := string(b.Index) + b.Timestamp + b.Data + b.PrevHash + string(nonce)
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}

func NewBlock(index int, data, prevHash string) Block {
    block := Block{
        Index:     index,
        Timestamp: time.Now().String(),
        Data:      data,
        PrevHash:  prevHash,
    }
    pow := NewProofOfWork(&block)
    nonce, hash := pow.Run()
    block.Hash = hash
    block.Nonce = nonce // Add Nonce to the Block struct
    return block
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
		fmt.Println("Is blockchain valid?", blockchain.IsValid())
		fmt.Println("---")
	}
}
