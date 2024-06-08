package build

import (
	"github.com/jinzhu/copier"

	"rank-master-back/infrastructure/pkg/snowflake"
	"rank-master-back/internal/model/entity"
	"rank-master-back/internal/types"
)

func TemplateTypes2Entity(template types.Template) (*entity.Template, error) {
	templateEntity := &entity.Template{}
	err := copier.Copy(templateEntity, template)
	if err != nil {
		return nil, err
	}
	templateEntity.ID = snowflake.GenerateDefaultSnowflakeID()
	return templateEntity, nil
}
