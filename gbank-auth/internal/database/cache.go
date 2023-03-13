package database

import (
	"context"

	"github.com/guilhermealvess/gbank/gbank-auth/internal/repository"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	repository.Cache
	//repository.SqlDB
	client *redis.Client
}

func (c *Cache) Set(key string, value string) error {
	var ctx = context.Background()

	return c.client.Set(ctx, key, value, 0).Err()
}

func (c *Cache) Get(key string) (string, error) {
	var ctx = context.Background()
	return c.client.Get(ctx, "key").Result()
}

func NewCache(client *redis.Client) *Cache {
	return &Cache{
		client: client,
	}
}
