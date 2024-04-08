package test

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"rank-master-back/internal/svc"
	"rank-master-back/internal/types"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping() (resp *types.PingRes, err error) {
	return &types.PingRes{
		Msg: "pong",
	}, nil
}
