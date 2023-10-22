package svc

import (
	"rank-master-back/internal/config"
	"rank-master-back/internal/pkg/orm_engine"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     orm_engine.NewGormEngine(c.DataSource),
	}
}
