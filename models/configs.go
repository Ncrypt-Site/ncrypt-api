package models

// Config holds the application configuration data
type Config struct {
	StorageDriver string
	RedisConfig
	ApplicationConfig
}

// ApplicationConfig holds application configuration data
type ApplicationConfig struct {
	ApplicationUrlConfig
}

// RedisConfig holds redis connection configuration data
type RedisConfig struct {
	Addr     string
	Password string
	Database int
}

// ApplicationUrlConfig holds different URLs used inside the app
type ApplicationUrlConfig struct {
	ApiBaseUrl string
	AppBaseUrl string
}
