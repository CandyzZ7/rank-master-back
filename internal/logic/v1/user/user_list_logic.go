package user

import (
	"context"

	"rank-master-back/internal/build"
	"rank-master-back/internal/svc"
	"rank-master-back/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUserListLogic 更新用户信息列表
func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UpdateUserListReq) (resp *types.UpdateUserListResp, err error) {
	users, err := build.UserTypesList2Entity(req.UserList)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.UserDao.UpdateBatchByFields(l.ctx, users, []string{"name"})
	if err != nil {
		return nil, err
	}
	return
}
