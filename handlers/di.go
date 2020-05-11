package handlers

import (
	"github.com/go-redis/redis/v7"
	"ncrypt-api/models"
)

type DI struct {
	RedisClient *redis.Client
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

func buildRedisConnection(config models.RedisConfig) (*redis.Client, error) {
	c := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password, // no password set
		DB:       config.Database, // use default DB
	})

	_, err := c.Ping().Result()
	if err != nil {
		return nil, err
	}

	return c, nil
}
