package cypher

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
)

func EncryptSecurityKey(securityKey []byte, k *rsa.PublicKey) ([]byte, error) {
	hash := sha512.New()
	ct, err := rsa.EncryptOAEP(hash, rand.Reader, k, securityKey, nil)
	if err != nil {
		return nil, err
	}

	return ct, nil
}

func DecryptSecurityKey(encryptedSecurityKey []byte, k *rsa.PrivateKey) ([]byte, error) {
	hash := sha512.New()
	pt, err := rsa.DecryptOAEP(hash, rand.Reader, k, encryptedSecurityKey, nil)
	if err != nil {
		return nil, err
	}

	return pt, err
}
