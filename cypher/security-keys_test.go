package cypher

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"testing"
)

func TestEncryptSecurityKey(t *testing.T) {
	k, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal()
	}

	message := []byte("This is us.")

	_, err = EncryptSecurityKey(message, &k.PublicKey)
	if err != nil {
		t.Fail()
	}
}

func TestDecryptSecurityKey(t *testing.T) {
	k, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal()
	}

	message := []byte("This is us.")

	encryptedSecurityKey, err := EncryptSecurityKey(message, &k.PublicKey)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	decryptedSecurityKey, err := DecryptSecurityKey(encryptedSecurityKey, k)
	if err != nil ||
		string(decryptedSecurityKey) != string(message) {
		t.Fail()
	}
}
