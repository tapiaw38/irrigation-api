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

func NewRedisCache(config *RedisCache) *RedisCache {
	return &RedisCache{
		Host:    config.Host,
		DB:      config.DB,
		Expires: config.Expires,
	}
}

func (cache *RedisCache) GetClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.Host,
		Password: "",
		DB:       cache.DB,
	})
}
