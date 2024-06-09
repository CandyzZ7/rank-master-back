package svc

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"rank-master-back/infrastructure/pkg/snowflake"
	"rank-master-back/infrastructure/repository/generate/dal"
	"rank-master-back/internal/config"
	"rank-master-back/internal/repository"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	DB          *gorm.DB
	RDB         *redis.Client
	Oss         *oss.Client
	TemplateDao repository.ITemplate
	UserDao     repository.IUser
}

func Init(ctx *ServiceContext) {
	// 雪花算法
	snowflake.InitNode(ctx.Config)
	// 数据库
	dal.SetDefault(ctx.DB)
}
