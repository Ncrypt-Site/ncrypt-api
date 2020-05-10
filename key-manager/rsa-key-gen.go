package keyManager

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"ncrypt-api/helpers"
)

func GenerateNewKeyPair() error {
	pkey, err := generatePrivateKey()
	if err != nil {
		return err
	}

	filename, err := getKeyPath()
	if err != nil {
		return err
	}

	filename += "/" + GetKeyId() + ".rsa"
	err = helpers.WriteToFile(filename, convertPrivateKeyToByte(pkey), 0600)
	if err != nil {
		return err
	}

	return nil
}

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
