package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

// MerkleTree represents a Merkle tree of transactions
type MerkleTree struct {
	Tree [][][]byte
	Root []byte
}

// NewMerkleTree creates a new Merkle tree from transactions
func NewMerkleTree(transactions []*Transaction) *MerkleTree {
	if len(transactions) == 0 {
		return &MerkleTree{Root: []byte{}}
	}

	var tree [][][]byte
	var currentLevel [][]byte

	// Create leaf nodes from transactions
	for _, tx := range transactions {
		hash := sha256.Sum256([]byte(tx.ID))
		currentLevel = append(currentLevel, hash[:])
	}

	tree = append(tree, currentLevel)

	// Build tree upwards
	for len(currentLevel) > 1 {
		var nextLevel [][]byte
		for i := 0; i < len(currentLevel); i += 2 {
			if i+1 < len(currentLevel) {
				hash := sha256.Sum256(append(currentLevel[i], currentLevel[i+1]...))
				nextLevel = append(nextLevel, hash[:])
			} else {
				hash := sha256.Sum256(append(currentLevel[i], currentLevel[i]...))
				nextLevel = append(nextLevel, hash[:])
			}
		}
		tree = append(tree, nextLevel)
		currentLevel = nextLevel
	}

	root := currentLevel[0]
	return &MerkleTree{Tree: tree, Root: root}
}

// GetRoot returns the Merkle root as hex string
func (mt *MerkleTree) GetRoot() string {
	if len(mt.Root) == 0 {
		return ""
	}
	return hex.EncodeToString(mt.Root)
}

