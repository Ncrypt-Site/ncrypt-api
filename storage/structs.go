package storage

import redisStorage "ncrypt-api/storage/redis-storage"

//Storage holds the list of all supported storage interfaces
var Storage = map[string]interface{}{
	"redis": &redisStorage.RedisStorage{},
}
