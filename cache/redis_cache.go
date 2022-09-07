package cache

import (
	"time"

	"github.com/go-redis/redis/v7"
)

type RedisCache struct {
	Host    string
	DB      int
	Expires time.Duration
}

func NewRedisCache(host string, db int, expires time.Duration) *RedisCache {
	return &RedisCache{
		Host:    host,
		DB:      db,
		Expires: expires,
	}
}

func (cache *RedisCache) GetClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.Host,
		Password: "",
		DB:       cache.DB,
	})
}
