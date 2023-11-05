package svc

import (
	"rank-master-back/internal/config"
	"rank-master-back/internal/pkg/orm_engine"
	"rank-master-back/internal/pkg/rdb"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     orm_engine.NewGormEngine(c.DataSource),
		RDB:    rdb.NewRdbClient(c.Redis),
	}
}
