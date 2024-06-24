//go:build wireinject
// +build wireinject

//go:generate wire
package main

import (
	"github.com/google/wire"

	"rank-master-back/infrastructure/pkg/es"
	"rank-master-back/infrastructure/pkg/mq"
	"rank-master-back/infrastructure/pkg/ormengine"
	"rank-master-back/infrastructure/pkg/rdb"
	"rank-master-back/infrastructure/pkg/uploadfile/oss"
	"rank-master-back/internal/cache"
	"rank-master-back/internal/config"
	"rank-master-back/internal/repository"
	"rank-master-back/internal/svc"
)

func InitializeServiceContext(c config.Config) (*svc.ServiceContext, error) {
	panic(wire.Build(
		wire.Struct(new(svc.ServiceContext), "*"),
		ormengine.NewGormEngine,
		rdb.NewRdbClient,
		oss.NewOssClient,
		mq.NewPusher,
		es.NewEs,
		RepositorySet,
		CacheSet,
	))
}

var RepositorySet = wire.NewSet(
	repository.NewTemplateDao,
	// wire.Bind(new(repository.ITemplate), new(*repository.TemplateDao)),
	repository.NewUserDao,
	// wire.Bind(new(repository.IUser), new(*repository.UserDao)),
)

var CacheSet = wire.NewSet(
	cache.NewUserCache,
	// wire.Bind(new(cache.IUserCache), new(*cache.UserCache)),
)
