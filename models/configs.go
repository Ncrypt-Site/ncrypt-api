package models

type Config struct {
	StorageDriver string
	RedisConfig
}

type RedisConfig struct {
	Addr     string
	Password string
	Database int
}
