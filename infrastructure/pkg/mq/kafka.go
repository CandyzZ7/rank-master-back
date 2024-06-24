package mq

import (
	"github.com/zeromicro/go-queue/kq"

	"rank-master-back/internal/config"
)

func NewPusher(c config.Config) *kq.Pusher {
	return kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic)
}
