package handlers

import (
	"errors"
	"log"
	"ncrypt-api/models"
	"ncrypt-api/storage"
)

type DI struct {
	StorageDriver models.StorageInterface
}

func BuildDI(config models.Config) (DI, error) {
	di := DI{}

	rc, err := buildRedisConnection(config.RedisConfig)
	if err != nil {
		return DI{}, err
	}
	di.RedisClient = rc

	return di, nil
}

func findStorageDriver(d string) (interface{}, error) {
	s, ok := storage.Storage[d]
	if !ok {
		return nil, errors.New("unsupported storage")
	}

	return s, nil
}
