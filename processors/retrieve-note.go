package processors

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
	"ncrypt-api/models"
)

func RetrieveSecureNote(client *redis.Client, payload models.RetrieveNoteRequest) ([]byte, error) {
	if !checkIfNoteExists(client, payload.Id) {
		return nil, errors.New("note does not exist")
	}

	note, err := retrieveNote(client, payload.Id)
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

	return note.Note, nil
}
