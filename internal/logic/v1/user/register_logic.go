package user

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"rank-master-back/infrastructure/e"
	"rank-master-back/infrastructure/pkg/encrypt"
	"rank-master-back/infrastructure/pkg/jwt"
	"rank-master-back/infrastructure/pkg/snowflake"
	"rank-master-back/infrastructure/pkg/uploadfile/local"
	"rank-master-back/internal/model/entity"
	"rank-master-back/internal/svc"
	"rank-master-back/internal/types"
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

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// 检查账号是否重复
	isExist, err := l.svcCtx.UserDao.FindLockWithRankMasterAccountExist(l.ctx, req.User.RankMasterAccount)
	if err != nil {
		return nil, errors.Wrap(err, "查询失败")
	}
	if isExist == 1 {
		return nil, errors.Wrapf(e.ErrRegisterAccountExist, "账号: %s", req.User.RankMasterAccount)
	}
	code, err := l.svcCtx.RDB.Get(l.ctx, req.User.Email).Result()
	if err != nil {
		return nil, errors.WithMessage(err, "redis get error")
	}
	if code != req.User.Code {
		return nil, errors.Wrapf(e.ErrEmailCodeFail, "邮箱: %s", req.User.Email)
	}

	// 上传头像
	key, err := local.Upload(l.svcCtx.Config, req.User.Avatar)
	if err != nil {
		return nil, errors.Wrapf(err, "上传头像失败: %s", req.User.Avatar)
	}
	// 密码加密
	// 生成随机盐
	cryptSalt, err := encrypt.RandomString(encrypt.RandomNumberLen)
	if err != nil {
		return nil, err
	}
	// 加密密码
	encPassword := encrypt.EncryptMD5(req.User.Password + cryptSalt)
	// 加密手机号
	mobile := encrypt.EncryptMD5(req.User.Mobile + cryptSalt)

	userEntity := &entity.User{
		ID:                snowflake.GenerateDefaultSnowflakeID(),
		RankMasterAccount: req.User.RankMasterAccount,
		Name:              req.User.Name,
		Avatar:            &key,
		Mobile:            &mobile,
		Password:          encPassword,
		CryptSalt:         cryptSalt,
	}
	err = l.svcCtx.UserDao.Create(l.ctx, userEntity)
	if err != nil {
		return nil, errors.Wrap(err, "注册失败")
	}
	token, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		Fields: map[string]interface{}{
			"userId": userEntity.ID,
		},
	})
	if err != nil {
		logx.Errorf("BuildTokens error: %v", err)
		return nil, err
	}

	return &types.RegisterResp{
		UserId: userEntity.ID,
		Token: types.Token{
			AccessToken:  token.AccessToken,
			AccessExpire: token.AccessExpire,
		},
	}, nil
}
