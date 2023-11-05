package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth       Auth
	DataSource string
	LogConf    logx.LogConf
	Email      Email
	Redis      Redis
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

type Redis struct {
	Address  string
	Password string
	DB       int
}
