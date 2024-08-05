package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type cacheV1 struct {
	cache *redis.Client
}

func New(cache *redis.Client) *cacheV1 {
	return &cacheV1{cache}
}

func (c *cacheV1) ClearMultipleCacheKey(ctx context.Context, key ...string) error {
	cmd := c.cache.Del(ctx, key...)
	if err := cmd.Err(); err != nil {
		return err
	}
	return nil
}

func (c *cacheV1) GetDataFromCache(ctx context.Context, key string) ([]byte, error) {
	cmd := c.cache.Get(ctx, key)
	if err := cmd.Err(); err != nil {
		return nil, err
	}
	data, err := cmd.Bytes()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *cacheV1) SetDataToCache(ctx context.Context, key string, value []byte) error {
	cmd := c.cache.Set(ctx, key, value, time.Hour*2)
	if err := cmd.Err(); err != nil {
		return err
	}
	return nil
}
