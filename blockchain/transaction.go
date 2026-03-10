package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Transaction represents a transaction in the blockchain
type Transaction struct {
	ID        string
	From      string
	To        string
	Amount    float64
	Timestamp string
}

// NewTransaction creates a new transaction
func NewTransaction(from, to string, amount float64) *Transaction {
	tx := &Transaction{
		From:      from,
		To:        to,
		Amount:    amount,
		Timestamp: time.Now().String(),
	}
	tx.ID = tx.CalculateHash()
	return tx
}

// CalculateHash computes the SHA-256 hash of a transaction
func (t *Transaction) CalculateHash() string {
	record := t.From + t.To + fmt.Sprintf("%.2f", t.Amount) + t.Timestamp
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

