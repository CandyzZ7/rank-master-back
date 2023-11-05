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
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.WithMessage(err, "redis get error")
	}
	if code != req.Code {
		return nil, errors.Wrapf(e.ErrEmailCodeFail, "邮箱: %s", req.Email)
	}
	// 去除前后空格
	req.Name = strings.TrimSpace(req.Name)
	req.Mobile = strings.TrimSpace(req.Mobile)
	req.Password = strings.TrimSpace(req.Password)
	// 密码加密
	// 生成随机盐
	cryptSalt, err := encrypt.RandomString(encrypt.RandomNumberLen)
	if err != nil {
		return nil, err
	}
	// 加密密码
	encPassword := encrypt.Encryption(req.Password, cryptSalt)
	// 检查手机号是否已经注册
	isExist, err := dal.Use(l.svcCtx.DB).User.FindWithMobile(req.Mobile)
	if err != nil {
		return nil, errors.Wrap(err, "注册失败")
	}
	if isExist == 1 {
		return nil, e.ErrRegisterMobileExist
	}
	userModel := &model.User{
		Id:        snowflake.GetSnowflakeID(),
		Name:      req.Name,
		Mobile:    req.Mobile,
		Password:  encPassword,
		CryptSalt: cryptSalt,
	}
	err = dal.Use(l.svcCtx.DB).User.Create(userModel)
	if err != nil {
		return nil, errors.Wrap(err, "注册失败")
	}
	token, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		Fields: map[string]interface{}{
			"userId": userModel.Id,
		},
	})
	if err != nil {
		logx.Errorf("BuildTokens error: %v", err)
		return nil, err
	}

	return &types.RegisterRes{
		UserId: userModel.Id,
		Token: types.Token{
			AccessToken:  token.AccessToken,
			AccessExpire: token.AccessExpire,
		},
	}, nil
}
