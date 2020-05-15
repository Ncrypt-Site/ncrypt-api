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

	storageInterface, err := findStorageDriver(config.StorageDriver)
	if err != nil {
		log.Fatal(err)
	}

	di.StorageDriver, err = storageInterface.(models.StorageInterface).BuildConfiguration(config)
	if err != nil {
		log.Fatal(err)
	}

	return di, nil
}

func findStorageDriver(d string) (interface{}, error) {
	s, ok := storage.Storage[d]
	if !ok {
		return nil, errors.New("unsupported storage")
	}

	return s, nil
}
