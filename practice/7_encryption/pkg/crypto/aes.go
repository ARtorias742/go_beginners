package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// EncryptAES encrypts data using AES-256 in CBC mode
func EncryptAES(plaintext, key []byte) ([]byte, error) {
	// Create cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Pad plaintext to block size
	padded := pad(plaintext, aes.BlockSize)

	// Create IV
	ciphertext := make([]byte, aes.BlockSize+len(padded))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// Encrypt
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], padded)

	return ciphertext, nil
}

// DecryptAES decrypts data using AES-256 in CBC mode
func DecryptAES(ciphertext, key []byte) ([]byte, error) {
	// Create cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Check minimum length
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	// Extract IV
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// Check if ciphertext length is valid
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of block size")
	}

	// Decrypt
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// Unpad
	plaintext, err = unpad(plaintext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// pad adds PKCS7 padding to the input
func pad(buf []byte, blockSize int) []byte {
	padding := blockSize - len(buf)%blockSize
	padtext := make([]byte, len(buf)+padding)
	copy(padtext, buf)
	for i := len(buf); i < len(padtext); i++ {
		padtext[i] = byte(padding)
	}
	return padtext
}

// unpad removes PKCS7 padding from the input
func unpad(buf []byte) ([]byte, error) {
	if len(buf) == 0 {
		return nil, errors.New("empty buffer")
	}
	padding := int(buf[len(buf)-1])
	if padding > len(buf) || padding == 0 {
		return nil, errors.New("invalid padding")
	}
	return buf[:len(buf)-padding], nil
}
