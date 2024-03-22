package orm_engine

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	gormEngine *gorm.DB
	once       sync.Once
)

func NewGormEngine(dataSource string) *gorm.DB {
	once.Do(func() {
		gormEngine, _ = gorm.Open(mysql.Open(dataSource), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), // 配置日志级别，打印出所有的sql
		})
	})
	return gormEngine
}
