package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"tech.pacifici/blockchain/blockchain"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ferentum-blockchain",
	Short: "Ferentum Blockchain CLI",
}

var transferCmd = &cobra.Command{
	Use:   "transfer [from] [to] [amount]",
	Short: "Create and add a transaction to blockchain",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		amount, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			fmt.Println("Invalid amount:", err)
			return
		}

		bc := blockchain.LoadBlockchain("ferentum-blockchain.dat")
		tx := blockchain.NewTransaction(args[0], args[1], amount)
		bc.AddBlock([]*blockchain.Transaction{tx})
		bc.SaveBlockchain("ferentum-blockchain.dat")
		fmt.Printf("Transaction from %s to %s: %.2f FRN\n", args[0], args[1], amount)
		fmt.Printf("Block #%d mined with hash: %s\n", len(*bc)-1, (*bc)[len(*bc)-1].Hash)
	},
}

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print the blockchain",
	Run: func(cmd *cobra.Command, args []string) {
		bc := blockchain.LoadBlockchain("ferentum-blockchain.dat")
		for _, block := range *bc {
			fmt.Printf("Index: %d\n", block.Index)
			fmt.Printf("Timestamp: %s\n", block.Timestamp)
			fmt.Printf("Difficulty: %d\n", block.Difficulty)
			fmt.Printf("Transactions: %d\n", len(block.Transactions))
			for j, tx := range block.Transactions {
				fmt.Printf("  [%d] %s -> %s: %.2f FRN\n", j, tx.From, tx.To, tx.Amount)
			}
			fmt.Printf("MerkleRoot: %s\n", block.MerkleRoot)
			fmt.Printf("PrevHash: %s\n", block.PrevHash)
			fmt.Printf("Hash: %s\n", block.Hash)
			fmt.Printf("Nonce: %d\n", block.Nonce)
			fmt.Println("---")
		}
	},
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate the blockchain",
	Run: func(cmd *cobra.Command, args []string) {
		bc := blockchain.LoadBlockchain("ferentum-blockchain.dat")
		fmt.Println("Is blockchain valid?", bc.IsValid())
	},
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the blockchain to genesis",
	Run: func(cmd *cobra.Command, args []string) {
		bc := blockchain.Blockchain{*blockchain.GenesisBlock()}
		bc.SaveBlockchain("ferentum-blockchain.dat")
		fmt.Println("Blockchain reset to genesis.")
	},
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show blockchain info",
	Run: func(cmd *cobra.Command, args []string) {
		bc := blockchain.LoadBlockchain("ferentum-blockchain.dat")
		totalTxs := 0
		for _, block := range *bc {
			totalTxs += len(block.Transactions)
		}
		fmt.Printf("Number of blocks: %d\n", len(*bc))
		fmt.Printf("Total transactions: %d\n", totalTxs)
		lastBlock := (*bc)[len(*bc)-1]
		fmt.Printf("Last block hash: %s\n", lastBlock.Hash)
		fmt.Printf("Current difficulty: %d\n", lastBlock.Difficulty)
		fmt.Printf("Is valid: %t\n", bc.IsValid())
	},
}

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search block by index or hash",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bc := blockchain.LoadBlockchain("ferentum-blockchain.dat")
		query := args[0]

		// Try to parse as index
		if index, err := strconv.Atoi(query); err == nil {
			block := bc.SearchBlockByIndex(index)
			if block != nil {
				printBlock(block)
				return
			}
		}

		// Try as hash
		block := bc.SearchBlockByHash(query)
		if block != nil {
			printBlock(block)
			return
		}

		fmt.Println("Block not found")
	},
}

var addressCmd = &cobra.Command{
	Use:   "address [address]",
	Short: "Get all transactions for an address",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bc := blockchain.LoadBlockchain("ferentum-blockchain.dat")
		blocks := bc.GetBlocksByAddress(args[0])

		if len(blocks) == 0 {
			fmt.Printf("No transactions found for %s\n", args[0])
			return
		}

		fmt.Printf("Found %d blocks with transactions for %s:\n", len(blocks), args[0])
		for _, block := range blocks {
			fmt.Printf("\nBlock #%d (Hash: %s):\n", block.Index, block.Hash[:16]+"...")
			for _, tx := range block.Transactions {
				if tx.From == args[0] || tx.To == args[0] {
					fmt.Printf("  %s -> %s: %.2f FRN\n", tx.From, tx.To, tx.Amount)
				}
			}
		}
	},
}

func printBlock(block *blockchain.Block) {
	fmt.Printf("Index: %d\n", block.Index)
	fmt.Printf("Timestamp: %s\n", block.Timestamp)
	fmt.Printf("Difficulty: %d\n", block.Difficulty)
	fmt.Printf("Transactions: %d\n", len(block.Transactions))
	for j, tx := range block.Transactions {
		fmt.Printf("  [%d] %s -> %s: %.2f FRN\n", j, tx.From, tx.To, tx.Amount)
	}
	fmt.Printf("MerkleRoot: %s\n", block.MerkleRoot)
	fmt.Printf("PrevHash: %s\n", block.PrevHash)
	fmt.Printf("Hash: %s\n", block.Hash)
	fmt.Printf("Nonce: %d\n", block.Nonce)
}

func init() {
	rootCmd.AddCommand(transferCmd, printCmd, validateCmd, resetCmd, infoCmd, searchCmd, addressCmd)
}

func main() {
	if _, err := os.Stat("ferentum-blockchain.dat"); os.IsNotExist(err) {
		bc := blockchain.Blockchain{*blockchain.GenesisBlock()}
		bc.SaveBlockchain("ferentum-blockchain.dat")
	}
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
