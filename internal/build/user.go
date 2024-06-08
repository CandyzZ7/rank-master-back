package build

import (
	"github.com/jinzhu/copier"

	"rank-master-back/infrastructure/pkg/snowflake"
	"rank-master-back/internal/model/entity"
	"rank-master-back/internal/types"
)

func UserTypes2Entity(user types.User) (*entity.User, error) {
	userEntity := &entity.User{}
	err := copier.Copy(userEntity, user)
	if err != nil {
		return nil, err
	}
	userEntity.ID = snowflake.GenerateDefaultSnowflakeID()
	return userEntity, nil
}
