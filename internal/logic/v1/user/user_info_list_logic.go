package user

import (
	"context"

	"rank-master-back/internal/build"
	"rank-master-back/internal/svc"
	"rank-master-back/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUserInfoListLogic 用户信息列表
func NewUserInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoListLogic {
	return &UserInfoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoListLogic) UserInfoList(req *types.GetUserInfoListReq) (resp *types.GetUserInfoListResp, err error) {
	userListEntity, count, err := l.svcCtx.UserDao.Page(l.ctx, req.Pagination)
	if err != nil {
		return nil, err
	}
	userTypesList, err := build.UserEntityList2Types(userListEntity)
	if err != nil {
		return nil, err
	}
	return &types.GetUserInfoListResp{
		UserList: userTypesList,
		Count:    count,
	}, nil
}
