package user

import (
	"context"
	"rank-master-back/internal/e"
	"rank-master-back/internal/gen/dal"
	"rank-master-back/internal/pkg/encrypt"
	"rank-master-back/internal/pkg/jwt"
	"strings"

	"rank-master-back/internal/svc"
	"rank-master-back/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	req.Mobile = strings.TrimSpace(req.Mobile)
	req.Password = strings.TrimSpace(req.Password)
	userDB := dal.Use(l.svcCtx.DB).User
	// 检查手机号是否已经注册
	isExist, err := userDB.FindWithMobile(req.Mobile)
	if err != nil {
		return nil, err
	}
	if isExist == 0 {
		return nil, errors.New(e.ErrLoginMobileNotExist.String())
	}
	// 检查密码是否正确
	userEntity, err := userDB.Where(userDB.Mobile.Eq(req.Mobile)).First()
	if err != nil {
		return nil, err
	}
	password := req.Password + userEntity.CryptSalt
	isSame := encrypt.EqualsPassword(password, userEntity.Password)
	if !isSame {
		return nil, errors.New(e.ErrLoginPasswd.String())
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
	return &types.LoginRes{
		Token: types.Token{
			AccessToken:  token.AccessToken,
			AccessExpire: token.AccessExpire,
		},
	}, nil
}
