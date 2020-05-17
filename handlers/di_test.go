package handlers

import (
	"ncrypt-api/config"
	"ncrypt-api/models"
	"ncrypt-api/storage"
	redisStorage "ncrypt-api/storage/redis-storage"
	"testing"
)

type storageShadowInterface struct {
	redisStorage.RedisStorage
}

func (s storageShadowInterface) BuildConfiguration(c models.Config) (models.StorageInterface, error) {
	return s, nil
}

func TestFindStorage(t *testing.T) {
	_, err := findStorageDriver("redis")
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuildDI(t *testing.T) {
	storage.Storage["shadow"] = storageShadowInterface{}

	c := config.BuildConfig()
	c.StorageDriver = "shadow"

	_, err := BuildDI(c)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuildDIWithWrongStorageDriver(t *testing.T) {
	c := config.BuildConfig()
	c.StorageDriver = "shadows"

	_, err := BuildDI(c)
	if err == nil || err.Error() != "unsupported storage" {
		t.Fail()
	}
}

func TestBuildDiWithStorageBuildConfigError(t *testing.T) {
	c := config.BuildConfig()

	/*
		currently the default storage is redis and will fail
		We expect it to fails.
	*/
	_, err := BuildDI(c)
	if err == nil {
		t.Fail()
	}
}
