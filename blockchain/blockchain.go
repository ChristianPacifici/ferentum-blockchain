package blockchain

import (
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"os"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int
}

// Blockchain is a slice of blocks
type Blockchain []Block

// NewBlock creates a new block
func NewBlock(index int, data, prevHash string) *Block {
	block := &Block{
		Index:     index,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevHash,
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return block
}

// GenesisBlock creates the first block in the blockchain
func GenesisBlock() *Block {
	return NewBlock(0, "Genesis Block", "")
}

// CalculateHash computes the SHA-256 hash of a block
func (b *Block) CalculateHash() string {
	record := string(b.Index) + b.Timestamp + b.Data + b.PrevHash + string(b.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// IsValid validates the blockchain
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

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := (*bc)[len(*bc)-1]
	newBlock := NewBlock(prevBlock.Index+1, data, prevBlock.Hash)
	*bc = append(*bc, *newBlock)
}

// SaveBlockchain saves the blockchain to a file
func (bc *Blockchain) SaveBlockchain(filename string) {
	file, _ := os.Create(filename)
	encoder := gob.NewEncoder(file)
	encoder.Encode(bc)
}

// LoadBlockchain loads the blockchain from a file
func LoadBlockchain(filename string) *Blockchain {
	var bc Blockchain
	file, _ := os.Open(filename)
	defer file.Close()
	decoder := gob.NewDecoder(file)
	decoder.Decode(&bc)
	return &bc
}
