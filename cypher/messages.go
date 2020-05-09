package cypher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

func EncryptNote(note []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	encryptedText := gcm.Seal(nonce, nonce, note, nil)
	return encryptedText, nil
}

func DecryptNote(encryptedNote []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	nonce := gcm.NonceSize()
	if len(encryptedNote) < nonce {
		return nil, errors.New("invalid message")
	}

	messageNonce, message := encryptedNote[:nonce], encryptedNote[nonce:]

	decryptedNote, err := gcm.Open(nil, messageNonce, message, nil)
	if err != nil {
		return nil, err
	}

	return decryptedNote, nil
}
