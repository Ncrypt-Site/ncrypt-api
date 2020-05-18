package processors

import (
	"errors"
	"github.com/google/uuid"
	"ncrypt-api/models"
	redisStorage "ncrypt-api/storage/redis-storage"
	"testing"
)

func (m storage) Exists(id uuid.UUID) bool {
	return true
}

func (m storage) Retrieve(id uuid.UUID) (models.SecureMessage, error) {
	return models.SecureMessage{}, nil
}

func (m storage) Delete(id uuid.UUID) error {
	return nil
}

func TestRetrieveSecureNote(t *testing.T) {
	storage := storage{}
	r := models.RetrieveNoteRequest{Id: uuid.New()}

	_, err := RetrieveSecureNote(storage, r)
	if err != nil {
		t.Fail()
	}
}

type storageWithRetrieveFailure struct {
	redisStorage.RedisStorage
}

func (m storageWithRetrieveFailure) Exists(id uuid.UUID) bool {
	return true
}

func (m storageWithRetrieveFailure) Retrieve(id uuid.UUID) (models.SecureMessage, error) {
	return models.SecureMessage{}, errors.New("i solemnly swear that I am up to no good")
}

func TestRetrieveSecureNoteWithExistsFailure(t *testing.T) {
	storage := storageWithRetrieveFailure{}
	r := models.RetrieveNoteRequest{Id: uuid.New()}

	_, err := RetrieveSecureNote(storage, r)
	if err == nil {
		t.Fail()
	}
}

type storageWithDestructionFailure struct {
	redisStorage.RedisStorage
}

func (m storageWithDestructionFailure) Exists(id uuid.UUID) bool {
	return true
}

func (m storageWithDestructionFailure) Retrieve(id uuid.UUID) (models.SecureMessage, error) {
	return models.SecureMessage{DestructAfterOpening: true}, nil
}

func (m storageWithDestructionFailure) Delete(id uuid.UUID) error {
	return errors.New("mischief managed")
}

func TestRetrieveSecureNoteWithDestructionFailure(t *testing.T) {
	storage := storageWithDestructionFailure{}
	r := models.RetrieveNoteRequest{Id: uuid.New()}

	_, err := RetrieveSecureNote(storage, r)
	if err == nil {
		t.Fail()
	}
}
