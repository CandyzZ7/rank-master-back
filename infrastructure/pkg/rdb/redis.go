package rdb

import (
	"github.com/redis/go-redis/v9"

	"rank-master-back/internal/config"
)

var (
	rdbClient *redis.Client
)

func NewRdbClient(c config.Config) *redis.Client {
	rdbClient = redis.NewClient(&redis.Options{
		Addr:     c.Redis.Address,
		Password: c.Redis.Password, // no password set
		DB:       0,                // use default DB
	})
	return rdbClient
}
