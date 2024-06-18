package repository

import (
	"context"

	"github.com/go-redis/redis/v8"

	"rank-master-back/infrastructure/pkg/rdb"
	"rank-master-back/infrastructure/repository/generate/dal"
	"rank-master-back/internal/config"
	"rank-master-back/internal/types"

	"rank-master-back/internal/model/entity"
)

var _ IUser = (*userDao)(nil)

type userDao struct {
	RDB   *redis.Client
	Query *dal.Query
}

func NewUserDao(c config.Config) IUser {
	return &userDao{
		RDB:   rdb.NewRdbClient(c),
		Query: dal.Q,
	}
}

type IUser interface {
	Create(ctx context.Context, template *entity.User) error
	FindLockWithRankMasterAccountExist(ctx context.Context, rankMasterAccount string) (int64, error)
	FindUserByRankMasterAccount(ctx context.Context, rankMasterAccount string) (*entity.User, error)
	FindUserByID(ctx context.Context, id string) (*entity.User, error)
	DeleteUserByID(ctx context.Context, id string) error
	DeleteByIDs(ctx context.Context, ids []string) error
	Page(ctx context.Context, pagination types.Pagination) ([]*entity.User, int64, error)
}

func (d *userDao) FindUserByID(ctx context.Context, id string) (*entity.User, error) {
	return d.Query.User.WithContext(ctx).Where(dal.User.ID.Eq(id)).First()
}

func (d *userDao) Create(ctx context.Context, user *entity.User) error {
	err := d.Query.User.WithContext(ctx).Create(user)
	if err != nil {
		return err
	}
	_ = d.RDB.Del(ctx, user.ID)
	return nil
}

func (d *userDao) FindLockWithRankMasterAccountExist(ctx context.Context, rankMasterAccount string) (int64, error) {
	return d.Query.User.WithContext(ctx).FindLockWithRankMasterAccountExist(rankMasterAccount)
}

func (d *userDao) FindUserByRankMasterAccount(ctx context.Context, rankMasterAccount string) (*entity.User, error) {
	return d.Query.User.WithContext(ctx).Where(dal.User.RankMasterAccount.Eq(rankMasterAccount)).First()
}

func (d *userDao) DeleteUserByID(ctx context.Context, id string) error {
	_, err := d.Query.User.WithContext(ctx).Where(dal.User.ID.Eq(id)).Delete()
	if err != nil {
		return err
	}
	_ = d.RDB.Del(ctx, id)
	return nil
}

func (d *userDao) DeleteByIDs(ctx context.Context, ids []string) error {
	_, err := d.Query.User.WithContext(ctx).Where(dal.User.ID.In(ids...)).Delete()
	if err != nil {
		return err
	}

	for _, id := range ids {
		_ = d.RDB.Del(ctx, id)
	}

	return nil
}

func (d *userDao) Page(ctx context.Context, pagination types.Pagination) ([]*entity.User, int64, error) {
	query := d.Query.User.WithContext(ctx)
	if pagination.Page <= 0 {
		pagination.Page = 1
	}
	if pagination.PageSize <= 0 {
		pagination.PageSize = 10
	}
	col, ok := d.Query.User.GetFieldByName(pagination.SortBy)
	if ok {
		if len(pagination.SortBy) > 0 {
			if pagination.SortOrder == "asc" {
				query = d.Query.User.WithContext(ctx).Order(col.Asc())
			}
			if pagination.SortOrder == "desc" {
				query = d.Query.User.WithContext(ctx).Order(col.Desc())
			}
		} else {
			query = d.Query.User.WithContext(ctx).Order(col)
		}
	}
	result, count, err := query.FindByPage((pagination.Page-1)*pagination.PageSize, pagination.PageSize)
	if err != nil {
		return nil, 0, err
	}
	return result, count, nil
}
