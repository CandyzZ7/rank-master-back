//go:build wireinject
// +build wireinject

//go:generate wire
package main

import (
	"github.com/google/wire"

	"rank-master-back/infrastructure/pkg/uploadfile/oss"
	"rank-master-back/internal/config"

	"rank-master-back/infrastructure/pkg/ormengine"
	"rank-master-back/infrastructure/pkg/rdb"
	"rank-master-back/internal/svc"
)

func InitializeServiceContext(c config.Config) (*svc.ServiceContext, error) {
	panic(wire.Build(
		wire.Struct(new(svc.ServiceContext), "*"),
		ormengine.NewGormEngine,
		rdb.NewRdbClient,
		oss.NewOssClient,
	))
}
