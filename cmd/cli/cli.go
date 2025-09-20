package main

import (
	"fmt"
	"log"
	"os"

	"tech.pacifici/blockchain/blockchain"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ferentum-blockchain",
	Short: "Ferentum Blockchain CLI",
}

var addCmd = &cobra.Command{
	Use:   "add [data]",
	Short: "Add a block to the blockchain",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bc := blockchain.LoadBlockchain("ferentum-blockchain.dat")
		bc.AddBlock(args[0])
		bc.SaveBlockchain("ferentum-blockchain.dat")
		fmt.Println("Block added!")
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
			fmt.Printf("Data: %s\n", block.Data)
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
		fmt.Printf("Number of blocks: %d\n", len(*bc))
		fmt.Printf("Last block hash: %s\n", (*bc)[len(*bc)-1].Hash)
		fmt.Printf("Is valid: %t\n", bc.IsValid())
	},
}

var mineCmd = &cobra.Command{
	Use:   "mine [data]",
	Short: "Mine a new block with data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bc := blockchain.LoadBlockchain("ferentum-blockchain.dat")
		prevBlock := (*bc)[len(*bc)-1]
		block := blockchain.NewBlock(prevBlock.Index+1, args[0], prevBlock.Hash)
		*bc = append(*bc, *block)
		bc.SaveBlockchain("ferentum-blockchain.dat")
		fmt.Printf("Mined block #%d with hash: %s\n", block.Index, block.Hash)
	},
}

func init() {
	rootCmd.AddCommand(addCmd, printCmd, validateCmd, resetCmd, infoCmd, mineCmd)
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
