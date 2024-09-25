package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func createEncryptCmd() *cobra.Command {
	var inputFile string
	var outputFile string
	var key string

	var encryptCmd = &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypts a file using AES-256",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Make sure all required flags are provided
			if inputFile == "" || outputFile == "" || key == "" {
				return fmt.Errorf("input file, output file, and key are required")
			}

			// Encrypt the file
			err := encryptFile(inputFile, outputFile, key)
			if err != nil {
				return fmt.Errorf("error encrypting file: %v", err)
			}
			return nil
		},
	}

	// Add flags for input, output files, and key
	encryptCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Path to the input file")
	encryptCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Path to the output encrypted file")
	encryptCmd.Flags().StringVarP(&key, "key", "k", "", "Encryption key (in hex)")

	return encryptCmd
}

func main() {
	rootCmd := createRootCmd()
	rootCmd.AddCommand(createGreetCmd())
	rootCmd.AddCommand(createEncryptCmd()) // Add the encrypt command
	rootCmd.AddCommand(createDecryptCmd())

	rootCmd.Execute()
}
