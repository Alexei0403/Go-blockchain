package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

// Define Block type
type Block struct {
	nonce        int
	previousHash [32]byte
	transactions []*Transaction
	timeStamp    int64
}

// Create new block
func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.nonce = nonce
	b.previousHash = previousHash
	b.timeStamp = time.Now().UnixNano()
	return b
}

// Print the block
func (b *Block) Print() {
	fmt.Printf("nonce           %d\n", b.nonce)
	fmt.Printf("previous_hash   %x\n", b.previousHash)
	fmt.Printf("transactions    %s\n", b.transactions)
	fmt.Printf("time_stamp      %d\n", b.timeStamp)

	for _, t := range b.transactions {
		t.Print()
	}
}

// Define Blockchain type
type Blockchain struct {
	transactionPool []*Transaction
	chain           []*Block
}

// Create blockchain (including genesis block)
func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

// Create new block
func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

// Get last block
func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

// Print blockchain
func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain: %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

// Generate sha256 hash from a block, for a block
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)

	fmt.Println(string(m))

	return sha256.Sum256([]byte(m))
}

// Marshal block to JSON (translate struct to JSON)
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
		TimeStamp    int64          `json:"time_stamp"`
	}{
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
		TimeStamp:    b.timeStamp,
	})
}

// Add transaction to the transaction pool
func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

// Create transaction type
type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

// Create new transaction
func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

// Print transaction
func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address    %s\n", t.senderBlockchainAddress)
	fmt.Printf(" recipient_blockchain_address %s\n", t.recipientBlockchainAddress)
	fmt.Printf(" value                        %.1f\n", t.value)
}

// Marshal transaction to JSON (translate struct to JSON)
func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct { // Creating struct on the fly (only for MarshalJSON)
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}

// Function to initialize the logger
func init() { //? TODO: Why we need this?
	log.SetPrefix("Blockchain: ")
}

// Main function
func main() {
	// Initialize a new blockchain ()
	blockChain := NewBlockchain()
	blockChain.Print()

	// Create a new block
	blockChain.AddTransaction("Alice", "Bob", 1.0)
	previousHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(5, previousHash)
	blockChain.Print()

	// Create a new block
	previousHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(2, previousHash)
	blockChain.Print()
}
