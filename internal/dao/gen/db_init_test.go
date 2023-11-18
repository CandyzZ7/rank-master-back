package gen

import (
	"rank-master-back/internal/model/entity"
	"testing"

	"gorm.io/gen"
)

//go:generate go test .

func TestDBInit(t *testing.T) {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./dal", // output directory, default value is ./query
		Mode:          gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})
	g.ApplyBasic(entity.User{})
	g.ApplyInterface(func(IUser) {}, entity.User{}, entity.Template{})
	g.Execute()
}
