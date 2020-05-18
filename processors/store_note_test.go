package processors

import (
	"errors"
	"github.com/google/uuid"
	"ncrypt-api/models"
	redisStorage "ncrypt-api/storage/redis-storage"
	"testing"
	"time"
)

type storage struct {
	redisStorage.RedisStorage
}

func (m storage) Store(id uuid.UUID, data []byte, duration time.Duration) error {
	return nil
}

type brokenStorage struct {
	redisStorage.RedisStorage
}

func (m brokenStorage) Store(id uuid.UUID, data []byte, duration time.Duration) error {
	return errors.New("oops, we failed")
}

func TestStoreMessage(t *testing.T) {
	storage := storage{}
	m := models.SecureMessageRequest{
		Note:                 "No post on Sundays.",
		SelfDestruct:         0,
		DestructAfterOpening: false,
	}

	_, err := StoreMessage(storage, m)
	if err != nil {
		t.Fatal(err)
	}

	brokenStorage := brokenStorage{}

	_, err = StoreMessage(brokenStorage, m)
	if err == nil {
		t.Fail()
	}

}
