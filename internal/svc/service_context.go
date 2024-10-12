package svc

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"rank-master-back/infrastructure/pkg/es"
	"rank-master-back/infrastructure/pkg/mq"
	"rank-master-back/infrastructure/pkg/ormengine"
	"rank-master-back/infrastructure/pkg/snowflake"
	"rank-master-back/infrastructure/pkg/uploadfile/xoss"
	"rank-master-back/infrastructure/pkg/xrdb"
	"rank-master-back/infrastructure/repository/generate/dal"
	"rank-master-back/internal/cache"
	"rank-master-back/internal/config"
	"rank-master-back/internal/repository"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config
	DB             *gorm.DB
	RDB            *redis.Redis
	KqPusherClient *kq.Pusher
	Es             *es.Es
	Oss            *oss.Client
	TemplateDao    repository.ITemplate
	UserDao        repository.IUser
}

func Init(ctx *ServiceContext) {
	// 雪花算法
	snowflake.InitNode(ctx.Config)
	// 数据库
	dal.SetDefault(ctx.DB)
}

func NewServiceContext(c config.Config) (*ServiceContext, error) {
	db := ormengine.MustNewGormEngine(c.MysqlConf)
	rdb := xrdb.NewRdbClientMust(c)
	pusher := mq.NewPusher(c)
	esEs, err := es.NewEs(c)
	if err != nil {
		return nil, err
	}
	client, err := xoss.NewOssClient(c)
	if err != nil {
		return nil, err
	}
	iTemplate := repository.NewTemplateDao()
	iUserCache := cache.NewUserCache(c)
	iUser := repository.NewUserDao(iUserCache)
	return &ServiceContext{
		Config:         c,
		DB:             db,
		RDB:            rdb,
		KqPusherClient: pusher,
		Es:             esEs,
		Oss:            client,
		TemplateDao:    iTemplate,
		UserDao:        iUser,
	}, nil
}
