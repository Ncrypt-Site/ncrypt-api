package config

import (
	"os"
	"testing"
)

func TestBuildRedisConfigWithDefaultValue(t *testing.T) {
	config := BuildRedisConfig()

	if config.Addr != "127.0.0.1:6379" ||
		config.Database != 0 ||
		config.Password != "" {
		t.Error()
	}
}

func TestBuildRedisConfigWithCustomizedHost(t *testing.T) {
	err := os.Setenv("NCRYPT_API_REDIS_HOST", "redis-host")
	if err != nil {
		t.Fatal(err)
	}

	config := BuildRedisConfig()
	if config.Addr != "redis-host:6379" {
		t.Error()
	}

	os.Clearenv()
}

func TestBuildRedisConfigWithCustomizedHostAndPort(t *testing.T) {
	err := os.Setenv("NCRYPT_API_REDIS_HOST", "redis-host")
	err = os.Setenv("NCRYPT_API_REDIS_PORT", "1990")
	if err != nil {
		t.Fatal(err)
	}

	config := BuildRedisConfig()
	if config.Addr != "redis-host:1990" {
		t.Error()
	}

	os.Clearenv()
}

func TestBuildRedisConfigWithCustomizedDB(t *testing.T) {
	err := os.Setenv("NCRYPT_API_REDIS_DB", "5")
	if err != nil {
		t.Fatal(err)
	}

	config := BuildRedisConfig()
	if config.Database != 5 {
		t.Error()
	}

	os.Clearenv()
}

func TestBuildRedisConfigWithPassword(t *testing.T) {
	err := os.Setenv("NCRYPT_API_REDIS_PASSWORD", "Finite Incantatem")
	if err != nil {
		t.Fatal(err)
	}

	config := BuildRedisConfig()
	if config.Password != "Finite Incantatem" {
		t.Error()
	}

	os.Clearenv()
}

func TestBuildRedisConfigWithAllEnvDefined(t *testing.T) {
	err := os.Setenv("NCRYPT_API_REDIS_HOST", "redis-host")
	err = os.Setenv("NCRYPT_API_REDIS_PORT", "1990")
	err = os.Setenv("NCRYPT_API_REDIS_DB", "5")
	err = os.Setenv("NCRYPT_API_REDIS_PASSWORD", "Finite Incantatem")
	if err != nil {
		t.Fatal(err)
	}

	config := BuildRedisConfig()
	if config.Addr != "redis-host:1990" &&
		config.Database != 5 &&
		config.Password != "Finite Incantatem" {
		t.Error()
	}

	os.Clearenv()
}

func TestBuildStorageDriver(t *testing.T) {
	driver := BuildStorageDriver()
	if driver != "redis" {
		t.Fail()
	}

	err := os.Setenv("NCRYPT_API_STORAGE_DRIVER", "fake")
	if err != nil {
		t.Fatal(err)
	}

	driver = BuildStorageDriver()
	if driver != "fake" {
		t.Fail()
	}
}

func TestBuildConfig(t *testing.T) {
	c := BuildConfig()

	if c.RedisConfig.Database != 0 &&
		c.RedisConfig.Addr != "127.0.0.1:6379" &&
		c.RedisConfig.Password != "" &&
		c.StorageDriver != "redis" &&
		c.ApiBaseUrl == "" &&
		c.AppBaseUrl == "" {
		t.Fail()
	}
}

func TestBuildApplicationConfig(t *testing.T) {
	c := BuildApplicationConfig()
	if len(c.ApiBaseUrl) != 0 || len(c.AppBaseUrl) != 0 {
		t.Fail()
	}

	err := os.Setenv("NCRYPT_API_APP_BASE_URL", "https://farshad.nematdoust.com")
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("NCRYPT_API_API_BASE_URL", "https://farshad.nematdoust.com")
	if err != nil {
		t.Fatal(err)
	}

	c = BuildApplicationConfig()
	if c.ApiBaseUrl != "https://farshad.nematdoust.com" || c.AppBaseUrl != "https://farshad.nematdoust.com" {
		t.Fail()
	}
}
