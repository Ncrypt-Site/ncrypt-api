package redis_storage

import (
	"github.com/alicebob/miniredis"
	"ncrypt-api/models"
	"testing"
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
