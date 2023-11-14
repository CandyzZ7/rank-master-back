package user

import (
	"context"
	"rank-master-back/internal/pkg/verification_code"
	"time"

	"rank-master-back/internal/svc"
	"rank-master-back/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
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

func (l *GetEmailCodeLogic) GetEmailCode(req *types.GetEmailCodeReq) (resp *types.GetEmailCodeRes, err error) {
	code := verification_code.GetRand(verification_code.Six)
	err = l.svcCtx.RDB.Set(l.ctx, req.Email, code, time.Minute*verification_code.CodeValidityTime).Err()
	if err != nil {
		return nil, errors.WithMessage(err, "redis set error")
	}
	err = verification_code.SendEmailCode(l.svcCtx.Config, req.Email, code)
	if err != nil {
		return nil, errors.Wrapf(err, "邮箱: %s", req.Email)
	}
	return
}
