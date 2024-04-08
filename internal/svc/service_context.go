package svc

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"rank-master-back/internal/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
	Oss    *oss.Client
}
