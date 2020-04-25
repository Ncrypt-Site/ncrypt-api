package keyManager

import (
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
