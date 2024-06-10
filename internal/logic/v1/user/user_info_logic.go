package user

import (
	"context"

	"rank-master-back/internal/build"
	"rank-master-back/internal/constant"
	"rank-master-back/internal/svc"
	"rank-master-back/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUserInfoLogic 用户信息
func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.GetUserInfoResp, err error) {
	userID := l.ctx.Value(constant.UserIdKey).(string)
	userEntity, err := l.svcCtx.UserDao.FindUserByID(l.ctx, userID)
	if err != nil {
		return nil, err
	}
	userTypes, err := build.UserEntity2Types(userEntity)
	if err != nil {
		return nil, err
	}
	return &types.GetUserInfoResp{
		User: *userTypes,
	}, nil
}
