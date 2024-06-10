package user

import (
	"context"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"

	"rank-master-back/infrastructure/e"
	"rank-master-back/internal/svc"
	"rank-master-back/internal/types"
)

type GetRankMasterAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRankMasterAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRankMasterAccountLogic {
	return &GetRankMasterAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRankMasterAccountLogic) GetRankMasterAccount(req *types.GetRankMasterAccountReq) (resp *types.GetRankMasterAccountResp, err error) {
	rankMasterAccount := strings.TrimSpace(req.RankMasterAccount)
	isExist, err := l.svcCtx.UserDao.FindLockWithRankMasterAccountExist(l.ctx, rankMasterAccount)
	if err != nil {
		return nil, err
	}
	if isExist == 1 {
		return nil, e.ErrRegisterMobileExist
	}
	return
}
