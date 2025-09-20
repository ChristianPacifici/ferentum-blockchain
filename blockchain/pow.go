package blockchain

// ProofOfWork represents a proof-of-work
type ProofOfWork struct {
	Block  *Block
	Target string
}

// NewProofOfWork creates a new proof-of-work
func NewProofOfWork(b *Block) *ProofOfWork {
	target := "0000"
	return &ProofOfWork{b, target}
}

// Run performs the proof-of-work
func (pow *ProofOfWork) Run() (int, string) {
	nonce := 0
	var hash string
	for {
		hash = pow.Block.CalculateHash()
		if hash[:len(pow.Target)] == pow.Target {
			break
		}
		nonce++
		pow.Block.Nonce = nonce
	}
	return nonce, hash
}