package keyManager

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func generatePrivateKey() (*rsa.PrivateKey, error) {
	key, err := rsa.GenerateKey(rand.Reader, KeySize)
	if err != nil {
		return nil, err
	}

	err = key.Validate()
	if err != nil {
		return nil, err
	}

	return key, nil
}

func convertPrivateKeyToByte(k *rsa.PrivateKey) []byte {
	key := x509.MarshalPKCS1PrivateKey(k)

	keyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: key,
	}
	b := pem.EncodeToMemory(keyBlock)
	return b
}
