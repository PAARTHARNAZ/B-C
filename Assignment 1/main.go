package main

import (
	"assignment01bca/assignment01bca" // Use the correct module path
	"fmt"
)

func main() {
	fmt.Println("Creating Blockchain...")

	// Create the first block (Genesis Block)
	genesisBlock := assignment01bca.NewBlock("Genesis Block", 0, "0")
	fmt.Printf("Created Block: %+v\n", genesisBlock)

	// Add more blocks
	block2 := assignment01bca.NewBlock("Alice to Bob", 1, genesisBlock.CurrentHash)
	fmt.Printf("Created Block: %+v\n", block2)

	block3 := assignment01bca.NewBlock("Bob to Carol", 2, block2.CurrentHash)
	fmt.Printf("Created Block: %+v\n", block3)

	// List all blocks
	assignment01bca.ListBlocks()

	// Modify a block
	assignment01bca.ChangeBlock(1, "Alice to Charlie")
	assignment01bca.ListBlocks()

	// Verify the chain
	valid := assignment01bca.VerifyChain()
	if valid {
		fmt.Println("Blockchain verified!")
	}
}
