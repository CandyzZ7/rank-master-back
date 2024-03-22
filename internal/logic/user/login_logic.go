package user

import (
	"context"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"

	"rank-master-back/infrastructure/e"
	"rank-master-back/internal/dao/gen/dal"
	"rank-master-back/internal/pkg/encrypt"
	"rank-master-back/internal/pkg/jwt"
	"rank-master-back/internal/svc"
	"rank-master-back/internal/types"
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
	rankMasterAccount := strings.TrimSpace(req.RankMasterAccount)
	password := strings.TrimSpace(req.Password)
	userDB := dal.Use(l.svcCtx.DB).User
	// 检查账号是否已经注册
	isExist, err := userDB.FindLockWithRankMasterAccount(rankMasterAccount)
	if err != nil {
		return nil, err
	}
	if isExist == 0 {
		return nil, e.ErrLoginMobileNotExist
	}
	// 检查密码是否正确
	// 从数据库中获取用户信息
	userEntity, err := userDB.Where(userDB.RankMasterAccount.Eq(rankMasterAccount)).First()
	if err != nil {
		return nil, err
	}
	// 密码加盐加密
	passwordWithSalt := password + userEntity.CryptSalt
	// 比较密码是否相同
	isSame := encrypt.EqualsEncryption(passwordWithSalt, userEntity.Password)
	if !isSame {
		return nil, e.ErrLoginPasswd
	}
	// 生成token
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
		UserId: userEntity.Id,
		Token: types.Token{
			AccessToken:  token.AccessToken,
			AccessExpire: token.AccessExpire,
		},
	}, nil
}
