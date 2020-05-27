package models

type Config struct {
	StorageDriver string
	RedisConfig
	ApplicationConfig
}

type ApplicationConfig struct {
	ApplicationUrlConfig
}

type RedisConfig struct {
	Addr     string
	Password string
	Database int
}

type ApplicationUrlConfig struct {
	ApiBaseUrl string
	AppBaseUrl string
}
