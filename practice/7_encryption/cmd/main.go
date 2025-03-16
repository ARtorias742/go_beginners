package main

import (
	"fmt"

	"github.com/ARtorias742/pkg/crypto" // Correct import path
)

func main() {
	// Example key (32 bytes for AES-256)
	key := []byte("thisis32byteslongsecretkey123456")
	plaintext := "Hello, this is a secret message!"

	// Encrypt
	ciphertext, err := crypto.EncryptAES([]byte(plaintext), key)
	if err != nil {
		fmt.Printf("Encryption error: %v", err) // Removed \n as Printf already adds it
		return
	}
	fmt.Printf("Encrypted: %x", ciphertext) // Removed \n

	// Decrypt
	decrypted, err := crypto.DecryptAES(ciphertext, key)
	if err != nil {
		fmt.Printf("Decryption error: %v", err)
		return
	}
	fmt.Printf("Decrypted: %s", decrypted) // Removed \n
}
