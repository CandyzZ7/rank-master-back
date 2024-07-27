package config

import (
	"time"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth           Auth
	Mysql          Mysql
	LogConf        logx.LogConf
	Email          Email
	Redis          redis.RedisConf
	UploadFile     UploadFile
	WorkerId       int64
	KqPusherConf   KqPusherConf
	KqConsumerConf kq.KqConf
	Es             Es
}

type KqPusherConf struct {
	Brokers []string
	Topic   string
}

type Es struct {
	Addresses  []string
	Username   string
	Password   string
	MaxRetries int
}

type Auth struct {
	AccessSecret  string
	AccessExpire  int64
	RefreshSecret string
	RefreshExpire int64
	RefreshAfter  int64
}

type Email struct {
	AuthorizationPassword string
}

type UploadFile struct {
	AliYunOss AliYunOss
	Path      string
}

type AliYunOss struct {
	Endpoint         string
	AccessKeyId      string
	AccessKeySecret  string
	BucketName       string
	ConnectTimeout   int64 `json:",optional"`
	ReadWriteTimeout int64 `json:",optional"`
}

type Mysql struct {
	DataSource    string
	MaxOpenConns  int
	MaxIdleConns  int
	MaxLifetime   int
	SlowThreshold time.Duration
}
