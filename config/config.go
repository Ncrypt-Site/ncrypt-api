package config

import (
	"ncrypt-api/models"
	"os"
	"strconv"
)

//BuildConfig build config struct
func BuildConfig() models.Config {
	config := models.Config{}

	config.RedisConfig = BuildRedisConfig()
	config.StorageDriver = BuildStorageDriver()
	config.ApplicationConfig = BuildApplicationConfig()

	return config
}

//BuildApplicationConfig build application config struct
func BuildApplicationConfig() models.ApplicationConfig {
	c := models.ApplicationConfig{}

	apiBaseUrl, ok := os.LookupEnv("NCRYPT_API_API_BASE_URL")
	if ok {
		c.ApplicationUrlConfig.ApiBaseUrl = apiBaseUrl
	}

	appBaseUrl, ok := os.LookupEnv("NCRYPT_API_APP_BASE_URL")
	if ok {
		c.ApplicationUrlConfig.AppBaseUrl = appBaseUrl
	}

	return c
}

//BuildStorageDriver return storage driver
func BuildStorageDriver() string {
	if driver, ok := os.LookupEnv("NCRYPT_API_STORAGE_DRIVER"); ok {
		return driver
	}

	return "redis"
}

//BuildRedisConfig build redis config struct
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
