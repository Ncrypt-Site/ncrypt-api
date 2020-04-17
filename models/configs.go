package models

type Config struct {
	RedisConfig
}

type RedisConfig struct {
	Addr     string
	Password string
	Database int
}
