package user

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"rank-master-back/infrastructure/pkg/verificationcode"
	"rank-master-back/internal/svc"
	"rank-master-back/internal/types"
)

type GetEmailCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailCodeLogic {
	return &GetEmailCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEmailCodeLogic) GetEmailCode(req *types.GetEmailCodeReq) (resp *types.GetEmailCodeResp, err error) {
	code := verificationcode.GetRand(verificationcode.Six)
	err = l.svcCtx.RDB.SetexCtx(l.ctx, req.Email, code, 60*verificationcode.CodeValidityTime)
	if err != nil {
		return nil, errors.WithMessage(err, "redis set error")
	}
	err = verificationcode.SendEmailCode(l.svcCtx.Config, req.Email, code)
	if err != nil {
		return nil, errors.Wrapf(err, "邮箱: %s", req.Email)
	}
	return
}
