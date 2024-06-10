package repository

import (
	"context"

	"rank-master-back/infrastructure/repository/generate/dal"

	"rank-master-back/internal/model/entity"
)

var _ IUser = (*UserDao)(nil)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

type IUser interface {
	Create(ctx context.Context, template *entity.User) error
	FindLockWithRankMasterAccountExist(ctx context.Context, rankMasterAccount string) (int64, error)
	FindUserByRankMasterAccount(ctx context.Context, rankMasterAccount string) (*entity.User, error)
	FindUserByID(ctx context.Context, id string) (*entity.User, error)
}

func (d *UserDao) FindUserByID(ctx context.Context, id string) (*entity.User, error) {
	return dal.User.WithContext(ctx).Where(dal.User.ID.Eq(id)).First()
}

func (d *UserDao) Create(ctx context.Context, user *entity.User) error {
	err := dal.User.WithContext(ctx).Create(user)
	return err
}

func (d *UserDao) FindLockWithRankMasterAccountExist(ctx context.Context, rankMasterAccount string) (int64, error) {
	return dal.User.WithContext(ctx).FindLockWithRankMasterAccountExist(rankMasterAccount)
}
func (d *UserDao) FindUserByRankMasterAccount(ctx context.Context, rankMasterAccount string) (*entity.User, error) {
	return dal.User.WithContext(ctx).Where(dal.User.RankMasterAccount.Eq(rankMasterAccount)).First()
}
