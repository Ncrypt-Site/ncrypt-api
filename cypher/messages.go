package cypher

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
)

func EncryptMessage(m []byte, k *rsa.PublicKey) ([]byte, error) {
	hash := sha512.New()
	ct, err := rsa.EncryptOAEP(hash, rand.Reader, k, m, nil)
	if err != nil {
		return nil, err
	}

	return ct, nil
}

func DecryptMessage(msg []byte, k *rsa.PrivateKey) ([]byte, error) {
	hash := sha512.New()
	pt, err := rsa.DecryptOAEP(hash, rand.Reader, k, msg, nil)
	if err != nil {
		return nil, err
	}

	return pt, err
}
