//go:build wireinject
// +build wireinject

//go:generate wire
package main

import (
	"github.com/google/wire"

	"rank-master-back/infrastructure/pkg/upload_file/oss"
	"rank-master-back/internal/config"

	"rank-master-back/infrastructure/pkg/orm_engine"
	"rank-master-back/infrastructure/pkg/rdb"
	"rank-master-back/internal/svc"
)

func InitializeServiceContext(c config.Config) (*svc.ServiceContext, error) {
	panic(wire.Build(
		wire.Struct(new(svc.ServiceContext), "*"),
		orm_engine.NewGormEngine,
		rdb.NewRdbClient,
		oss.NewOssClient,
	))
}
