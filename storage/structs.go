package storage

import redisStorage "ncrypt-api/storage/redis-storage"

var Storage = map[string]interface{}{
	"redis": &redisStorage.RedisStorage{},
}
