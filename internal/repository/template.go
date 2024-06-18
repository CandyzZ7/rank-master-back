package repository

import (
	"context"

	"rank-master-back/infrastructure/repository/generate/dal"

	"rank-master-back/internal/model/entity"
)

var _ ITemplate = (*TemplateDao)(nil)

type ITemplate interface {
	Create(ctx context.Context, template *entity.Template) error
}

type TemplateDao struct {
}

func NewTemplateDao() ITemplate {
	return &TemplateDao{}
}

func (d *TemplateDao) Create(ctx context.Context, template *entity.Template) error {
	err := dal.Template.WithContext(ctx).Create(template)
	return err
}
