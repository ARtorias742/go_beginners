package crypto

import (
	"bytes"
	"testing"
)

func TestAESEncryptionDecryption(t *testing.T) {
	key := []byte("thisis32byteslongsecretkey123456")
	plaintext := []byte("Test message for encryption")

	ciphertext, err := EncryptAES(plaintext, key)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	decrypted, err := DecryptAES(ciphertext, key)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("Decrypted text doesn't match original. Got %s, want %s", decrypted, plaintext)
	}
}

func TestInvalidKey(t *testing.T) {
	invalidKey := []byte("shortkey")
	plaintext := []byte("test")

	_, err := EncryptAES(plaintext, invalidKey)
	if err == nil {
		t.Error("Expected error with invalid key length, got none")
	}
}

func TestInvalidCiphertext(t *testing.T) {
	key := []byte("thisis32byteslongsecretkey123456")
	invalidCiphertext := []byte("short")

	_, err := DecryptAES(invalidCiphertext, key)
	if err == nil {
		t.Error("Expected error with invalid ciphertext, got none")
	}
}
