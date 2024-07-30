package build

import (
	"github.com/jinzhu/copier"

	"rank-master-back/infrastructure/pkg/snowflake"
	"rank-master-back/internal/model/entity"
	"rank-master-back/internal/types"
)

func UserTypes2Entity(user *types.User) (*entity.User, error) {
	userEntity := &entity.User{}
	err := copier.Copy(userEntity, user)
	if err != nil {
		return nil, err
	}
	userEntity.ID = snowflake.GenerateDefaultSnowflakeID()
	return userEntity, nil
}

func UserInfoTypes2Entity(user *types.User) (*entity.User, error) {
	userEntity := &entity.User{}
	err := copier.Copy(userEntity, user)
	if err != nil {
		return nil, err
	}
	return userEntity, nil
}

func UserEntity2Types(user *entity.User) (*types.User, error) {
	userTypes := &types.User{}
	err := copier.Copy(userTypes, user)
	if err != nil {
		return nil, err
	}
	return userTypes, nil
}

func UserEntityList2Types(userList []*entity.User) ([]*types.User, error) {
	userListEntity := make([]*types.User, len(userList))
	for i, user := range userList {
		userTypes, err := UserEntity2Types(user)
		if err != nil {
			return nil, err
		}
		userListEntity[i] = userTypes
	}
	return userListEntity, nil
}

func UserTypesList2Entity(userList []*types.User) ([]*entity.User, error) {
	userListEntity := make([]*entity.User, len(userList))
	for i, user := range userList {
		userTypes, err := UserInfoTypes2Entity(user)
		if err != nil {
			return nil, err
		}
		userListEntity[i] = userTypes
	}
	return userListEntity, nil
}
