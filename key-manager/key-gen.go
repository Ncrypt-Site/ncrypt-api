package keyManager

import (
	"fmt"
	"ncrypt-api/helpers"
	"path/filepath"
)

const (
	StorageDirectory = "./.storage/keys"
	KeySize          = 4096
)

func GenerateNewKeyPair() error {
	pkey, err := generatePrivateKey()
	if err != nil {
		return err
	}

	filename, err := filepath.Abs(StorageDirectory)
	if err != nil {
		return err
	}
	fmt.Println(filename)

	filename += "/" + GetKeyId() + ".rsa"
	err = helpers.WriteToFile(filename, convertPrivateKeyToByte(pkey), 0600)
	if err != nil {
		return err
	}

	return nil
}
