package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time" // Import time for timestamp functionality
)

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "The blockchain"
	MINING_REWARD     = 1.0
)

type Transaction struct {
	// Define your Transaction struct here
}

type Block struct {
	timestamp     int64
	nonce         int
	previousHash  [32]byte
	transactions  []*Transaction
}

// NewBlock creates a new block
func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions
	return b
}

// Print prints block details
func (b *Block) Print() {
	fmt.Printf("Timestamp:            %d\n", b.timestamp)
	fmt.Printf("Nonce:                %d\n", b.nonce)
	fmt.Printf("Previous Hash:        %x\n", b.previousHash) // Use %x for byte arrays

	for _, t := range b.transactions {
		// Assuming Transaction has a Print method
		t.Print()
	}
}

// Hash calculates the hash of the block
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256(m) // Use m directly without converting to byte slice again
}

// MarshalJSON marshals the block to JSON
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64       `json:"timestamp"`
		Nonce        int         `json:"nonce"`
		PreviousHash [32]byte    `json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

type Blockchain struct {
	transactionPool []*Transaction
	chain []*Block
}

func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreatedBlock(0, b.Hash())
	return bc 

}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte ) *Block {
	b = NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain =append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}
func (bc * Blockchain ) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i,
		strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}
// Main function for testing
func main() {
	// Example usage
	previousHash := sha256.Sum256([]byte("previous block"))
	transactions := []*Transaction{} // Populate with actual transactions
	block := NewBlock(0, previousHash, transactions)

	block.Print()
	fmt.Printf("Hash: %x\n", block.Hash()) // Print the block hash
}
