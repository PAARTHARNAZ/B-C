package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Block represents a single block in the blockchain.
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	CurrentHash  string
}

// Blockchain is a list of validated blocks.
var Blockchain []Block

// NewBlock creates a new block and returns it.
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
		CurrentHash:  CalculateHash(transaction + string(nonce) + previousHash),
	}
	Blockchain = append(Blockchain, *block)
	return block
}

// CalculateHash returns the hash of a string using SHA256.
func CalculateHash(stringToHash string) string {
	hash := sha256.New()
	hash.Write([]byte(stringToHash))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}

// ListBlocks prints the details of all blocks in the blockchain.
func ListBlocks() {
	for i, block := range Blockchain {
		fmt.Printf("Block %d\n", i+1)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Current Hash: %s\n", block.CurrentHash)
		fmt.Println("----------")
	}
}

// ChangeBlock modifies the transaction of a given block and recalculates the hash.
func ChangeBlock(blockIndex int, newTransaction string) {
    if blockIndex < 0 || blockIndex >= len(Blockchain) {
        fmt.Println("Invalid block index")
        return
    }
    
    // Modify the specified block
    Blockchain[blockIndex].Transaction = newTransaction
    Blockchain[blockIndex].CurrentHash = CalculateHash(newTransaction + string(Blockchain[blockIndex].Nonce) + Blockchain[blockIndex].PreviousHash)
    
    // Recalculate hashes for all subsequent blocks
    for i := blockIndex + 1; i < len(Blockchain); i++ {
        Blockchain[i].PreviousHash = Blockchain[i-1].CurrentHash
        Blockchain[i].CurrentHash = CalculateHash(Blockchain[i].Transaction + string(Blockchain[i].Nonce) + Blockchain[i].PreviousHash)
    }
    
    fmt.Println("Block modified and hashes recalculated")
}


// VerifyChain checks if the blockchain is valid by ensuring each block’s hash matches the previous block’s hash.
func VerifyChain() bool {
	for i := 1; i < len(Blockchain); i++ {
		if Blockchain[i].PreviousHash != Blockchain[i-1].CurrentHash {
			fmt.Printf("Blockchain is invalid at block %d\n", i)
			return false
		}
	}
	fmt.Println("Blockchain is valid")
	return true
}
