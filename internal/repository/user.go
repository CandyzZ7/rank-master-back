package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	cachePkg "rank-master-back/infrastructure/pkg/cache"
	"rank-master-back/infrastructure/repository/generate/dal"
	"rank-master-back/internal/cache"
	"rank-master-back/internal/types"

	"rank-master-back/internal/model/entity"
)

var _ IUser = (*UserDao)(nil)

type UserDao struct {
	cache cache.IUserCache
	Query *dal.Query
	sfg   *singleflight.Group
}

func NewUserDao(cache cache.IUserCache) IUser {
	return &UserDao{
		cache: cache,
		Query: dal.Q,
		sfg:   new(singleflight.Group),
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

func (d *UserDao) FindUserByID(ctx context.Context, id string) (*entity.User, error) {
	record, err := d.cache.Get(ctx, id)
	if err == nil {
		return record, nil
	}
	if errors.Is(err, redis.Nil) {
		val, err, _ := d.sfg.Do(id, func() (interface{}, error) {
			user, err := d.Query.User.WithContext(ctx).Where(dal.User.ID.Eq(id)).First()
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					err = d.cache.SetCacheWithNotFound(ctx, id)
					if err != nil {
						return nil, err
					}
					return nil, gorm.ErrRecordNotFound
				}
				return nil, err
			}
			// set cache
			err = d.cache.Set(ctx, id, user, cachePkg.DefaultExpireTime)
			if err != nil {
				return nil, err
			}
			return user, nil
		})
		if err != nil {
			return nil, err
		}
		user, ok := val.(*entity.User)
		if !ok {
			return nil, gorm.ErrRecordNotFound
		}
		return user, err
	} else if errors.Is(err, cachePkg.ErrPlaceholder) {
		return nil, gorm.ErrRecordNotFound
	}
	return nil, err
}

func (d *UserDao) Create(ctx context.Context, user *entity.User) error {
	err := d.Query.User.WithContext(ctx).Create(user)
	if err != nil {
		return err
	}
	_ = d.cache.Del(ctx, user.ID)
	return nil
}

func (d *UserDao) FindLockWithRankMasterAccountExist(ctx context.Context, rankMasterAccount string) (int64, error) {
	return d.Query.User.WithContext(ctx).FindLockWithRankMasterAccountExist(rankMasterAccount)
}

func (d *UserDao) FindUserByRankMasterAccount(ctx context.Context, rankMasterAccount string) (*entity.User, error) {
	return d.Query.User.WithContext(ctx).Where(dal.User.RankMasterAccount.Eq(rankMasterAccount)).First()
}

func (d *UserDao) DeleteUserByID(ctx context.Context, id string) error {
	_, err := d.Query.User.WithContext(ctx).Where(dal.User.ID.Eq(id)).Delete()
	if err != nil {
		return err
	}
	_ = d.cache.Del(ctx, id)
	return nil
}

func (d *UserDao) DeleteByIDs(ctx context.Context, ids []string) error {
	_, err := d.Query.User.WithContext(ctx).Where(dal.User.ID.In(ids...)).Delete()
	if err != nil {
		return err
	}

	for _, id := range ids {
		_ = d.cache.Del(ctx, id)
	}

	return nil
}

func (d *UserDao) Page(ctx context.Context, pagination types.Pagination) ([]*entity.User, int64, error) {
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
