package xrdb

import (
	"github.com/zeromicro/go-zero/core/stores/redis"

	"rank-master-back/internal/config"
)

func NewRdbClient(c config.Config) (*redis.Redis, error) {
	return redis.NewRedis(c.Redis)
}

func NewRdbClientMust(c config.Config) *redis.Redis {
	return redis.MustNewRedis(c.Redis)
}
