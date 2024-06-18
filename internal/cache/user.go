package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"

	"rank-master-back/infrastructure/pkg/cache"
	"rank-master-back/infrastructure/pkg/encoding"
	"rank-master-back/infrastructure/pkg/rdb"
	"rank-master-back/internal/config"
	"rank-master-back/internal/model/entity"
)

const (
	// PrefixUserCacheKey cache prefix
	PrefixUserCacheKey = "user"
)

type UserCache struct {
	RDB   *redis.Client
	cache cache.ICache
}

func NewUserCache(config config.Config) IUserCache {
	jsonEncoding := encoding.JSONEncoding{}
	rdbClient := rdb.NewRdbClient(config)
	c := cache.NewRedisCache(rdbClient, PrefixUserCacheKey, jsonEncoding, func() interface{} {
		return &entity.User{}
	})
	return &UserCache{
		RDB:   rdbClient,
		cache: c,
	}
}

type IUserCache interface {
	Get(ctx context.Context, id string) (*entity.User, error)
	Set(ctx context.Context, id string, data *entity.User, duration time.Duration) error
	Del(ctx context.Context, id string) error
	SetCacheWithNotFound(ctx context.Context, id string) error
}

func (c *UserCache) Get(ctx context.Context, id string) (*entity.User, error) {
	var data *entity.User
	err := c.cache.Get(ctx, id, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *UserCache) Set(ctx context.Context, id string, data *entity.User, duration time.Duration) error {
	if data == nil || len(id) == 0 {
		return nil
	}
	err := c.cache.Set(ctx, id, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Del delete cache
func (c *UserCache) Del(ctx context.Context, id string) error {
	err := c.cache.Del(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *UserCache) SetCacheWithNotFound(ctx context.Context, id string) error {
	err := c.cache.SetCacheWithNotFound(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
