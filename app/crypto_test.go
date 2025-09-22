package app

import (
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	plaintext := "8c8d5882648d47ef994443fdeeb38ecf"
	encrypted, err := Crypto().Encrypt(plaintext)
	if err != nil {
		t.Errorf("Error occurred [%v]", err)
	}
	decrypted, err := Crypto().Decrypt(encrypted)
	if err != nil {
		t.Errorf("Error occurred [%v]", err)
	}
	if decrypted != plaintext {
		t.Errorf("Expected decrypted [%v], got [%v]", plaintext, decrypted)
	}
}
