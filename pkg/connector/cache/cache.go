// Package cache provides ...
package cache

import (
	"context"
	"time"

	"github.com/deepzz0/appdemo/pkg/config"

	"github.com/go-redis/redis/v8"
)

// Cache cache with redis
type Cache struct {
	*redis.Client

	prefix string
}

// NewCache new cache
func NewCache(opts config.CacheRedis, appName string) (cache *Cache, err error) {
	rOpt := &redis.Options{
		Addr:     opts.Host,
		Password: opts.Password,
		DB:       opts.DB,
	}
	cli := redis.NewClient(rOpt)
	// ping result
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err = cli.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	cache = &Cache{Client: cli, prefix: appName}

	return
}
