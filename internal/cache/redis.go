package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Set(ctx context.Context, key string, value string, ttl time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}

type RedisCache struct {
    client *redis.Client
}

var _ Cache = (*RedisCache)(nil)

func NewRedisCache(addr string) *RedisCache {
    rdb := redis.NewClient(&redis.Options{
        Addr: addr,
        DB:   0,
    })

    return &RedisCache{client: rdb}
}

func (r *RedisCache) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
    return r.client.Set(ctx, key, value, ttl).Err()
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
    return r.client.Get(ctx, key).Result()
}

func (r *RedisCache) Delete(ctx context.Context, key string) error {
    return r.client.Del(ctx, key).Err()
}