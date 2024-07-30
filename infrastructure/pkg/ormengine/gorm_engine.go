package ormengine

import (
	"sync"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"rank-master-back/internal/config"
)

var (
	gormEngine *gorm.DB
	once       sync.Once
)
var ErrFieldNotFound = errors.New("field not found")

func NewGormEngine(c config.Config) (*gorm.DB, error) {

	if c.Mysql.MaxIdleConns == 0 {
		c.Mysql.MaxIdleConns = 10
	}
	if c.Mysql.MaxOpenConns == 0 {
		c.Mysql.MaxOpenConns = 100
	}
	if c.Mysql.MaxLifetime == 0 {
		c.Mysql.MaxLifetime = 3600
	}
	if c.Mysql.SlowThreshold == 0 {
		c.Mysql.SlowThreshold = int64(200 * time.Millisecond)
	}
	var err error
	once.Do(func() {
		gormEngine, err = gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
			Logger: NewGormLogger(time.Duration(c.Mysql.SlowThreshold)), // 调整日志级别，根据需要修改
		})
		if err != nil {
			return
		}
		db, err := gormEngine.DB()
		if err != nil {
			return
		}
		db.SetMaxIdleConns(c.Mysql.MaxIdleConns)
		db.SetMaxOpenConns(c.Mysql.MaxOpenConns)
		db.SetConnMaxLifetime(time.Second * time.Duration(c.Mysql.MaxLifetime))
		err = gormEngine.Use(NewCustomePlugin())
		if err != nil {
			return
		}
	})
	return gormEngine, err
}

func MustNewGormEngine(c config.Config) *gorm.DB {
	db, err := NewGormEngine(c)
	if err != nil {
		panic(err)
	}
	return db
}
