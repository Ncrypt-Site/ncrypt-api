package key_manager

import (
	"ncrypt-api/helpers"
)

const (
	StorageDirectory = "./storage/keys"
	KeySize          = 4096
)

func GenerateNewKeyPair() error {
	pkey, err := generatePrivateKey()
	if err != nil {
		return err
	}

	err = helpers.WriteToFile(convertPrivateKeyToByte(pkey), StorageDirectory+getKeyId()+".rsa", 0600)
	if err != nil {
		return err
	}

	return nil
}
