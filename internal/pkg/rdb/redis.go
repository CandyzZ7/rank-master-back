package rdb

import (
	"rank-master-back/internal/config"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	rdbClient *redis.Client
	once      sync.Once
)

func NewRdbClient(rdb config.Redis) *redis.Client {
	once.Do(func() {
		rdbClient = redis.NewClient(&redis.Options{
			Addr:     rdb.Address,
			Password: rdb.Password, // no password set
			DB:       0,            // use default DB
		})
	})
	return rdbClient
}
