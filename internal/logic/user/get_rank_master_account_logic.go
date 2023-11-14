package user

import (
	"context"
	"rank-master-back/internal/dao/gen/dal"
	"rank-master-back/internal/e"
	"strings"

	"rank-master-back/internal/svc"
	"rank-master-back/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *GetRankMasterAccountLogic) GetRankMasterAccount(req *types.GetRankMasterAccountReq) (resp *types.GetRankMasterAccountRes, err error) {
	userDB := dal.Use(l.svcCtx.DB).User
	rankMasterAccount := strings.TrimSpace(req.RankMasterAccount)
	isExist, err := userDB.FindLockWithRankMasterAccount(rankMasterAccount)
	if err != nil {
		return nil, err
	}
	if isExist == 0 {
		return nil, e.ErrLoginMobileNotExist
	}
	return
}
