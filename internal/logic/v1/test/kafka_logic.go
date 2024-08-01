package test

import (
	"context"

	"github.com/zeromicro/go-zero/core/threading"

	"rank-master-back/internal/svc"
	"rank-master-back/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type KafkaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// kafka
func NewKafkaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KafkaLogic {
	return &KafkaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KafkaLogic) Kafka() (resp *types.KafkaResp, err error) {
	threading.GoSafe(func() {
		if err := l.svcCtx.KqPusherClient.Push(l.ctx, "kafka_test"); err != nil {
			logx.Errorf("KqPusherClient Push Error , err :%v", err)
		}
	})
	return
}
