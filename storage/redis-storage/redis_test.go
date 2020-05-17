package redis_storage

import (
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
	"ncrypt-api/models"
	"testing"
	"time"
)

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
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	storage := RedisStorage{
		Client: client,
	}

	err = storage.Store(uuid.New(), []byte(""), time.Minute*10)
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
