package svc

import (
	"rank-master-back/internal/config"
	"rank-master-back/internal/pkg/orm_engine"
	"rank-master-back/internal/pkg/rdb"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
	Oss    *oss.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	OssClient, err := oss.New(c.UploadFile.AliYunOss.Endpoint, c.UploadFile.AliYunOss.AccessKeyId, c.UploadFile.AliYunOss.AccessKeySecret,
		oss.Timeout(c.UploadFile.AliYunOss.ConnectTimeout, c.UploadFile.AliYunOss.ReadWriteTimeout))
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		DB:     orm_engine.NewGormEngine(c.DataSource),
		RDB:    rdb.NewRdbClient(c.Redis),
		Oss:    OssClient,
	}
}
