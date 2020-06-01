package redisStorage

import (
	"encoding/json"
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
	"ncrypt-api/models"
	"testing"
	"time"
)

func getMiniRedisClient() (RedisStorage, *miniredis.Miniredis) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	storage := RedisStorage{
		Client: client,
	}

	return storage, s
}

func TestRedisStorage_BuildConfiguration(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	c := models.Config{
		StorageDriver: "redis",
		RedisConfig: models.RedisConfig{
			Addr:     s.Addr(),
			Password: "",
			Database: 0,
		},
	}

	redisStorage := RedisStorage{}
	_, err = redisStorage.BuildConfiguration(c)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRedisStorage_Store(t *testing.T) {
	storage, ms := getMiniRedisClient()
	defer ms.Close()

	err := storage.Store(uuid.New(), []byte(""), time.Minute*10)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRedisStorage_StoreFailure(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "fake-add:1990",
	})

	storage := RedisStorage{
		Client: client,
	}

	err := storage.Store(uuid.New(), []byte(""), time.Minute*10)
	if err == nil {
		t.Fail()
	}
}

func TestRedisStorage_RetrieveWithValidData(t *testing.T) {
	storage, ms := getMiniRedisClient()
	defer ms.Close()

	secureMessage := models.SecureMessage{
		Note:                 []byte("perhaps those who are best suited to power are those who have never sought it."),
		DestructAfterOpening: false,
	}
	jsonMessage, _ := json.Marshal(secureMessage)
	id := uuid.New()

	err := storage.Store(id, jsonMessage, time.Minute*2)
	if err != nil {
		t.Fatal(err)
	}

	messageModel, err := storage.Retrieve(id)
	if err != nil ||
		string(messageModel.Note) != string(secureMessage.Note) {
		t.Fail()
	}
}

func TestRedisStorage_RetrieveWithRedisFailure(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "fake-add:1990",
	})

	storage := RedisStorage{
		Client: client,
	}

	_, err := storage.Retrieve(uuid.New())
	if err == nil {
		t.Fail()
	}
}

func TestRedisStorage_RetrieveWithJsonUnmarshalFailure(t *testing.T) {
	storage, ms := getMiniRedisClient()
	defer ms.Close()

	message := []byte("perhaps those who are best suited to power are those who have never sought it.")
	id := uuid.New()

	err := storage.Store(id, message, time.Minute*2)
	if err != nil {
		t.Fatal(err)
	}

	_, err = storage.Retrieve(id)
	if err == nil {
		t.Fail()
	}
}

func TestRedisStorage_Delete(t *testing.T) {
	storage, ms := getMiniRedisClient()
	defer ms.Close()

	message := []byte("perhaps those who are best suited to power are those who have never sought it.")
	id := uuid.New()

	err := storage.Store(id, message, time.Minute*2)
	if err != nil {
		t.Fatal(err)
	}

	err = storage.Delete(id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRedisStorage_DeleteWithInvalidNoteId(t *testing.T) {
	storage, ms := getMiniRedisClient()
	defer ms.Close()

	err := storage.Delete(uuid.New())
	if err == nil {
		t.Fail()
	}
}

func TestRedisStorage_DeleteWithRedisFailure(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "fake-add:1990",
	})

	storage := RedisStorage{
		Client: client,
	}

	err := storage.Delete(uuid.New())
	if err == nil {
		t.Fail()
	}
}

func TestRedisStorage_Exists(t *testing.T) {
	storage, ms := getMiniRedisClient()
	defer ms.Close()

	message := []byte("perhaps those who are best suited to power are those who have never sought it.")
	id := uuid.New()

	err := storage.Store(id, message, time.Minute*2)
	if err != nil {
		t.Fatal(err)
	}

	exists := storage.Exists(id)
	if !exists {
		t.Fail()
	}
}

func TestRedisStorage_ExistsWithInvalidNoteId(t *testing.T) {
	storage, ms := getMiniRedisClient()
	defer ms.Close()

	exists := storage.Exists(uuid.New())
	if exists {
		t.Fail()
	}
}

func TestRedisStorage_ExistsWithRedisFailure(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "fake-add:1990",
	})

	storage := RedisStorage{
		Client: client,
	}

	exists := storage.Exists(uuid.New())
	if exists {
		t.Fail()
	}
}
