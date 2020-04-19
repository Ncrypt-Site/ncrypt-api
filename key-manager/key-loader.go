package key_manager

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

func LoadCurrentKey() (*rsa.PrivateKey, error) {
	f, err := ioutil.ReadFile("./.storage/" + getKeyId() + ".rsa")
	if err != nil {
		return nil, err
	}

	return parsePrivateKey(f)
}

func LoadKey(name string) (*rsa.PrivateKey, error) {
	f, err := ioutil.ReadFile("./.storage/" + name + ".rsa")
	if err != nil {
		return nil, err
	}
	return parsePrivateKey(f)
}

func parsePrivateKey(f []byte) (*rsa.PrivateKey, error) {
	k, _ := pem.Decode(f)
	if k.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("provided key is not RSA type")
	}

	pk, err := x509.ParsePKCS1PrivateKey(k.Bytes)
	if err != nil {
		return nil, err
	}

	return pk, nil
}
