package template

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"rank-master-back/internal/build"
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

func (l *AddTemplateLogic) AddTemplate(req *types.AddTemplateReq) (resp *types.AddTemplateResp, err error) {
	templateEntity, err := build.TemplateTypes2Entity(req.Template)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.TemplateDao.Create(l.ctx, templateEntity)
	if err != nil {
		return nil, errors.WithMessage(err, "创建模板失败")
	}
	return &types.AddTemplateResp{
		Id: templateEntity.ID,
	}, nil
}
