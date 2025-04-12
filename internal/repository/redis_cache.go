package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

func (r *RedisCache) Get(key string) (string, error) {
	ctx := context.Background()
	return r.client.Get(ctx, key).Result()
}

func (r *RedisCache) Set(key string, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return r.client.Set(ctx, key, value, 0).Err()
}
