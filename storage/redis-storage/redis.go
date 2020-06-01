package redisStorage

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
	"ncrypt-api/models"
	"time"
)

//RedisStorage holds client connection data
type RedisStorage struct {
	Client *redis.Client
}

//BuildConfiguration connect to redis and return the interface model
func (r RedisStorage) BuildConfiguration(c models.Config) (models.StorageInterface, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     c.RedisConfig.Addr,
		Password: c.RedisConfig.Password, // no password set
		DB:       c.RedisConfig.Database, // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	r.Client = client

	return r, nil
}

//Store stores data to redis
func (r RedisStorage) Store(id uuid.UUID, data []byte, duration time.Duration) error {
	err := r.Client.Set(id.String(), data, duration).Err()
	if err != nil {
		return err
	}

	return nil
}

//Exists checks whether a note exists in redis
func (r RedisStorage) Exists(id uuid.UUID) bool {
	result, err := r.Client.Exists(id.String()).Result()
	if err != nil || result == 0 {
		return false
	}

	return true
}

//Retrieve fetch a note from redis and return it as a SecureMessage model
func (r RedisStorage) Retrieve(id uuid.UUID) (models.SecureMessage, error) {
	secureMessage := models.SecureMessage{}

	note, err := r.Client.Get(id.String()).Result()
	if err != nil {
		return secureMessage, err
	}

	if err := json.Unmarshal([]byte(note), &secureMessage); err != nil {
		return secureMessage, err
	}

	return secureMessage, nil
}

//Delete destroy a note on the redis server
func (r RedisStorage) Delete(id uuid.UUID) error {
	result, err := r.Client.Del(id.String()).Result()
	if err != nil {
		return err
	}
	if result != 1 {
		return errors.New("the note was not deleted due to an error on redis")
	}
	return nil
}
