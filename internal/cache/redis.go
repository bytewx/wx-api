package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	rdb *redis.Client
}

func NewClient(addr, password string, db int) (*Client, error) {
	const operation = "wx-api.internal.cache.redis.NewClient"

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	ctx := context.Background()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("%s: error: connecting to redis: %s", operation, err)
	}

	return &Client{rdb: rdb}, nil
}

func (c *Client) Get(ctx context.Context, key string) (string, error) {
	return c.rdb.Get(ctx, key).Result()
}

func (c *Client) Set(ctx context.Context, key string, value string, ttlSeconds int) error {
	return c.rdb.Set(ctx, key, value, secondsToDuration(ttlSeconds)).Err()
}

func secondsToDuration(seconds int) time.Duration {
	return time.Duration(seconds) * time.Second
}
