package orm_engine

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	gormEngine *gorm.DB
	once       sync.Once
)

func NewGormEngine(dataSource string) *gorm.DB {
	once.Do(func() {
		gormEngine, _ = gorm.Open(mysql.Open(dataSource), &gorm.Config{})

	})
	return gormEngine
}
