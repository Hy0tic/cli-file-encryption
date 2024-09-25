package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/spf13/cobra"
)

// encryptFile encrypts the content of the input file and writes the encrypted data to the output file.
func encryptFile(inputFile, outputFile, keyString string) error {
	// Convert key string to byte array
	key, err := hex.DecodeString(keyString)
	if err != nil {
		return fmt.Errorf("failed to decode hex key: %v", err)
	}

	// Read the file content
	plaintext, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file: %v", err)
	}

	// Generate a new AES cipher using the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("failed to create cipher: %v", err)
	}

	// Create a new GCM (Galois/Counter Mode) cipher mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("failed to create GCM: %v", err)
	}

	// Create a random nonce (number used once)
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return fmt.Errorf("failed to generate nonce: %v", err)
	}

	// Encrypt the plaintext and append the nonce to the ciphertext
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	// Write the encrypted data to the output file
	if err := ioutil.WriteFile(outputFile, ciphertext, 0644); err != nil {
		return fmt.Errorf("failed to write output file: %v", err)
	}

	fmt.Printf("File encrypted successfully and saved as: %s\n", outputFile)
	return nil
}

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