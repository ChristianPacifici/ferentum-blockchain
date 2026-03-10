package blockchain

import "strings"

// ProofOfWork represents a proof-of-work
type ProofOfWork struct {
	Block *Block
}

// NewProofOfWork creates a new proof-of-work
func NewProofOfWork(b *Block) *ProofOfWork {
	return &ProofOfWork{b}
}

// Run performs the proof-of-work
func (pow *ProofOfWork) Run() (int, string) {
	nonce := 0
	target := strings.Repeat("0", pow.Block.Difficulty)

	for {
		hash := pow.Block.CalculateHash()
		if strings.HasPrefix(hash, target) {
			return nonce, hash
		}
		nonce++
		pow.Block.Nonce = nonce
	}
}