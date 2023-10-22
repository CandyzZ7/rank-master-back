package user

import (
	"context"
	"rank-master-back/internal/e"
	"rank-master-back/internal/gen/dal"
	"rank-master-back/internal/model"
	"rank-master-back/internal/pkg/encrypt"
	"rank-master-back/internal/pkg/jwt"
	"rank-master-back/internal/pkg/snowflake"
	"strings"

	"rank-master-back/internal/svc"
	"rank-master-back/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRes, err error) {
	// 去除前后空格
	req.Name = strings.TrimSpace(req.Name)
	req.Mobile = strings.TrimSpace(req.Mobile)
	req.Password = strings.TrimSpace(req.Password)
	// 密码加密
	req.Password = encrypt.EncPassword(req.Password)
	userDB := dal.Use(l.svcCtx.DB).User
	// 检查手机号是否已经注册
	isExist, err := userDB.FindWithMobile(req.Mobile)
	if err != nil {
		return nil, errors.Wrap(err, "注册失败")
	}
	if isExist == 1 {
		return nil, errors.New(e.ErrRegisterMobileExist.String())
	}
	userEntity := &model.User{
		Id:       snowflake.GetSnowflakeID(),
		Name:     req.Name,
		Mobile:   req.Mobile,
		Password: req.Password,
	}
	err = userDB.Create(userEntity)
	if err != nil {
		return nil, errors.Wrap(err, "注册失败")
	}
	token, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		Fields: map[string]interface{}{
			"userId": userEntity.Id,
		},
	})
	if err != nil {
		logx.Errorf("BuildTokens error: %v", err)
		return nil, err
	}

	return &types.RegisterRes{
		UserId: userEntity.Id,
		Token: types.Token{
			AccessToken:  token.AccessToken,
			AccessExpire: token.AccessExpire,
		},
	}, nil
}
