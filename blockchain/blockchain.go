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
	Index        int
	Timestamp    string
	Transactions []*Transaction
	PrevHash     string
	Hash         string
	Nonce        int
	Difficulty   int
	MerkleRoot   string
}

// Blockchain is a slice of blocks
type Blockchain []Block


// NewBlock creates a new block with proof-of-work
func NewBlock(index int, transactions []*Transaction, prevHash string, difficulty int) *Block {
	merkleTree := NewMerkleTree(transactions)
	block := &Block{
		Index:        index,
		Timestamp:    time.Now().String(),
		Transactions: transactions,
		PrevHash:     prevHash,
		Difficulty:   difficulty,
		MerkleRoot:   merkleTree.GetRoot(),
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return block
}

// GenesisBlock creates the first block in the blockchain
func GenesisBlock() *Block {
	return NewBlock(0, []*Transaction{}, "", 1)
}

// CalculateHash computes the SHA-256 hash of a block
func (b *Block) CalculateHash() string {
	record := string(b.Index) + b.Timestamp + b.MerkleRoot + b.PrevHash + string(b.Nonce) + string(b.Difficulty)
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

// AddBlock adds a new block to the blockchain with transactions
func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	prevBlock := (*bc)[len(*bc)-1]
	newDifficulty := bc.CalculateDifficulty(prevBlock)
	newBlock := NewBlock(prevBlock.Index+1, transactions, prevBlock.Hash, newDifficulty)
	*bc = append(*bc, *newBlock)
}

// CalculateDifficulty adjusts difficulty based on block time
// Targets ~10 second block time
func (bc *Blockchain) CalculateDifficulty(prevBlock Block) int {
	if prevBlock.Index == 0 {
		return 1
	}

	currentTime := time.Now().Unix()
	blockTime, _ := time.Parse(time.UnixDate, prevBlock.Timestamp)
	timeDiff := currentTime - blockTime.Unix()

	targetTime := int64(10)
	difficulty := prevBlock.Difficulty

	if timeDiff < targetTime {
		difficulty++
	} else if timeDiff > targetTime*2 && difficulty > 1 {
		difficulty--
	}

	return difficulty
}

// SearchBlockByIndex finds a block by its index
func (bc Blockchain) SearchBlockByIndex(index int) *Block {
	if index >= 0 && index < len(bc) {
		return &bc[index]
	}
	return nil
}

// SearchBlockByHash finds a block by its hash
func (bc Blockchain) SearchBlockByHash(hash string) *Block {
	for i := range bc {
		if bc[i].Hash == hash {
			return &bc[i]
		}
	}
	return nil
}

// GetBlocksByAddress returns all blocks containing transactions for an address
func (bc Blockchain) GetBlocksByAddress(address string) []*Block {
	var blocks []*Block
	for i := range bc {
		for _, tx := range bc[i].Transactions {
			if tx.From == address || tx.To == address {
				blocks = append(blocks, &bc[i])
				break
			}
		}
	}
	return blocks
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

