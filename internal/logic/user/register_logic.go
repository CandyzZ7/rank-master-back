package user

import (
	"context"
	"rank-master-back/internal/e"
	"rank-master-back/internal/gen/dal"
	"rank-master-back/internal/model"
	"rank-master-back/internal/pkg/encrypt"
	"rank-master-back/internal/pkg/jwt"
	"rank-master-back/internal/pkg/snowflake"
	"rank-master-back/internal/pkg/upload_file/local"
	"rank-master-back/internal/svc"
	"rank-master-back/internal/types"
	"strings"

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
	// 检查账号是否重复
	isExist, err := dal.Use(l.svcCtx.DB).User.FindLockWithRankMasterAccount(req.RankMasterAccount)
	if err != nil {
		return nil, errors.Wrap(err, "查询失败")
	}
	if isExist == 1 {
		return nil, errors.Wrapf(e.ErrRegisterAccountExist, "账号: %s", req.RankMasterAccount)
	}
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
	req.Avatar = strings.TrimSpace(req.Avatar)
	// 上传头像
	key, err := local.Upload(l.svcCtx.Config, req.Avatar)
	if err != nil {
		return nil, errors.Wrapf(err, "上传头像失败: %s", req.Avatar)
	}
	// 密码加密
	// 生成随机盐
	cryptSalt, err := encrypt.RandomString(encrypt.RandomNumberLen)
	if err != nil {
		return nil, err
	}
	// 加密密码
	encPassword := encrypt.Encryption(req.Password, cryptSalt)
	// 加密手机号
	mobile := encrypt.Encryption(req.Mobile, cryptSalt)
	userModel := &model.User{
		Id:                snowflake.GetSnowflakeID(),
		RankMasterAccount: req.RankMasterAccount,
		Name:              req.Name,
		Avatar:            key,
		Mobile:            mobile,
		Password:          encPassword,
		CryptSalt:         cryptSalt,
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
