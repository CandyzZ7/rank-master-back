package mqs

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"rank-master-back/internal/svc"
)

type TestKafka struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestKafka(ctx context.Context, svcCtx *svc.ServiceContext) *TestKafka {
	return &TestKafka{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestKafka) Consume(ctx context.Context, key, val string) error {
	logx.Infof("test kafka key :%s , val :%s", key, val)
	return nil
}
