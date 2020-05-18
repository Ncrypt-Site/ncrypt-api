package processors

import (
	"errors"
	"ncrypt-api/models"
)

func RetrieveSecureNote(storage models.StorageInterface, payload models.RetrieveNoteRequest) ([]byte, error) {
	if !storage.Exists(payload.Id) {
		return nil, errors.New("note does not exist")
	}

	note, err := storage.Retrieve(payload.Id)
	if err != nil {
		return nil, err
	}

	// todo: need improvement and better handling
	if note.DestructAfterOpening {
		err := storage.Delete(payload.Id)
		if err != nil {
			return nil, err
		}
	}

	return note.Note, nil
}
