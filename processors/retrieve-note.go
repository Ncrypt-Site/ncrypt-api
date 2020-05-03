package processors

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
	"ncrypt-api/cypher"
	keyManager "ncrypt-api/key-manager"
	"ncrypt-api/models"
)

func RetrieveSecureNote(client *redis.Client, currentKey models.Key, payload models.RetrieveNoteRequest) ([]byte, error) {
	if !checkIfNoteExists(client, payload.Id) {
		return nil, errors.New("note does not exist")
	}

	note, err := retrieveNote(client, payload.Id)
	if err != nil {
		return nil, err
	}

	if len(note.Password) != 0 {
		if !cypher.ValidatePassword([]byte(payload.Password), []byte(note.Password)) {
			return nil, errors.New("invalid password provided")
		}
	}

	noteDecoded, err := decryptNote(note.Message, note.KeyId, currentKey)
	if err != nil {
		return nil, err
	}

	// todo: need improvement and better handling
	if note.DestructAfterOpening {
		err := deleteNote(client, payload.Id)
		if err != nil {
			return nil, err
		}
	}

	return noteDecoded, nil
}

func checkIfNoteExists(client *redis.Client, id uuid.UUID) bool {
	result, err := client.Exists(id.String()).Result()
	if err != nil || result == 0 {
		return false
	}

	return true
}

func retrieveNote(client *redis.Client, id uuid.UUID) (models.SecureMessage, error) {
	secureMessage := models.SecureMessage{}

	note, err := client.Get(id.String()).Result()
	if err != nil {
		return secureMessage, err
	}

	if err := json.Unmarshal([]byte(note), &secureMessage); err != nil {
		return secureMessage, err
	}

	return secureMessage, nil
}

// todo: this probably can be improved, think about it.
func decryptNote(note []byte, keyId string, k models.Key) ([]byte, error) {
	var message []byte
	var err error

	if keyId != keyManager.GetKeyId() {
		privateKey, err := keyManager.LoadKey(keyId)
		if err != nil {
			return nil, err
		}

		message, err = cypher.DecryptMessage(note, privateKey)
	} else {
		message, err = cypher.DecryptMessage(note, k.PrivateKey)
	}

	if err != nil {
		return nil, err
	}

	return message, nil
}

func deleteNote(client *redis.Client, id uuid.UUID) error {
	result, err := client.Del(id.String()).Result()
	if err != nil {
		return err
	}
	if result != 1 {
		return errors.New("the note was not deleted due to an error on redis")
	}
	return nil
}
