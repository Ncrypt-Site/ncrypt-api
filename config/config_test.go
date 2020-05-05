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

func TestBuildConfig(t *testing.T) {
	config := BuildConfig()

	if config.RedisConfig.Database != 0 &&
		config.RedisConfig.Addr != "127.0.0.1:6379" &&
		config.RedisConfig.Password != "" {
		t.Fail()
	}
}
