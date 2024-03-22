package template

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"rank-master-back/infrastructure/pkg/snowflake"
	"rank-master-back/internal/dao/gen/dal"
	"rank-master-back/internal/model/entity"
	"rank-master-back/internal/svc"
	"rank-master-back/internal/types"
)

type AddTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTemplateLogic {
	return &AddTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTemplateLogic) AddTemplate(req *types.AddTemplateReq) (resp *types.AddTemplateRes, err error) {
	templateDB := dal.Use(l.svcCtx.DB).Template
	templateEntity := &entity.Template{
		BaseEntity: entity.BaseEntity{
			Id: snowflake.GetSnowflakeID(),
		},
		Function: req.Function,
		Type:     req.Type,
		Topic:    req.Topic,
		Content:  req.Content,
		Remark:   req.Remark,
	}
	err = templateDB.Create(templateEntity)
	if err != nil {
		return nil, errors.WithMessage(err, "创建模板失败")
	}
	return &types.AddTemplateRes{
		Id: templateEntity.Id,
	}, nil
}
