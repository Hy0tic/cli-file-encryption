package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// decryptFile decrypts the content of the input file and writes the decrypted data to the output file.
func decryptFile(inputFile, outputFile, keyString string) error {
	// Convert the hex string key to bytes
	key, err := hex.DecodeString(keyString)
	if err != nil {
		return fmt.Errorf("failed to decode hex key: %v", err)
	}

	// Read the encrypted file content
	ciphertext, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file: %v", err)
	}

	// Generate AES cipher using the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("failed to create cipher: %v", err)
	}

	// Create a GCM (Galois/Counter Mode) cipher
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("failed to create GCM: %v", err)
	}

	// Get the nonce size and separate it from the ciphertext
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return fmt.Errorf("ciphertext too short")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt the ciphertext using the nonce
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return fmt.Errorf("failed to decrypt file: %v", err)
	}

	// Write the decrypted data to the output file
	if err := os.WriteFile(outputFile, plaintext, 0644); err != nil {
		return fmt.Errorf("failed to write output file: %v", err)
	}

	fmt.Printf("File decrypted successfully and saved as: %s\n", outputFile)
	return nil
}

func createDecryptCmd() *cobra.Command {
	var inputFile string
	var outputFile string
	var key string

	var decryptCmd = &cobra.Command{
		Use:   "decrypt",
		Short: "Decrypts an encrypted file using AES-256",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Ensure required flags are provided
			if inputFile == "" || outputFile == "" || key == "" {
				return fmt.Errorf("input file, output file, and key are required")
			}

			// Decrypt the file
			err := decryptFile(inputFile, outputFile, key)
			if err != nil {
				return fmt.Errorf("error decrypting file: %v", err)
			}
			return nil
		},
	}

	// Add flags for input, output files, and key
	decryptCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Path to the encrypted file")
	decryptCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Path to save the decrypted file")
	decryptCmd.Flags().StringVarP(&key, "key", "k", "", "Decryption key (in hex)")

	return decryptCmd
}
