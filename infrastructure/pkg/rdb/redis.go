package rdb

import (
	"sync"

	"github.com/go-redis/redis/v8"

	"rank-master-back/internal/config"
)

var (
	rdbClient *redis.Client
	once      sync.Once
)

func NewRdbClient(c config.Config) *redis.Client {
	once.Do(func() {
		rdbClient = redis.NewClient(&redis.Options{
			Addr:     c.Redis.Address,
			Password: c.Redis.Password, // no password set
			DB:       0,                // use default DB
		})
	})
	return rdbClient
}
