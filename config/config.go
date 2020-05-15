package config

import (
	"ncrypt-api/models"
	"os"
	"strconv"
)

func BuildConfig() models.Config {
	config := models.Config{}

	config.RedisConfig = BuildRedisConfig()
	config.StorageDriver = BuildStorageDriver()

	return config
}

func BuildStorageDriver() string {
	if driver, ok := os.LookupEnv("NCRYPT_API_STORAGE_DRIVER"); ok {
		return driver
	}

	return "redis"
}

func BuildRedisConfig() models.RedisConfig {
	c := models.RedisConfig{}

	if addr, ok := os.LookupEnv("NCRYPT_API_REDIS_HOST"); ok {
		if port, ok := os.LookupEnv("NCRYPT_API_REDIS_PORT"); ok {
			c.Addr = addr + ":" + port
		} else {
			c.Addr = addr + ":6379"
		}
	} else {
		c.Addr = "127.0.0.1:6379"
	}

	if db, ok := os.LookupEnv("NCRYPT_API_REDIS_DB"); ok {
		if v, err := strconv.Atoi(db); err == nil {
			c.Database = v
		}
	}

	if password, ok := os.LookupEnv("NCRYPT_API_REDIS_PASSWORD"); ok {
		c.Password = password
	}

	return c
}
