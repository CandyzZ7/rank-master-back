package ormengine

import (
	"sync"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"rank-master-back/infrastructure/pkg/xlogger"
	"rank-master-back/internal/config"
)

var (
	gormEngine *gorm.DB
	once       sync.Once
)
var ErrFieldNotFound = errors.New("field not found")

func NewGormEngine(c config.Config) (*gorm.DB, error) {
	var err error
	once.Do(func() {
		gormEngine, err = gorm.Open(mysql.Open(c.DataSource), &gorm.Config{
			Logger: xlogger.DefaultGormLogger, // 调整日志级别，根据需要修改
		})
		if err != nil {
			return
		}
	})
	return gormEngine, err
}
