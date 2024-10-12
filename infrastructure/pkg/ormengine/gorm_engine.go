package ormengine

import (
	"sync"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConf struct {
	DataSource    string
	MaxOpenConns  int
	MaxIdleConns  int
	MaxLifetime   int
	SlowThreshold int64
}

var (
	gormEngine *gorm.DB
	once       sync.Once
)
var ErrFieldNotFound = errors.New("field not found")

func NewGormEngine(mysqlConf MysqlConf) (*gorm.DB, error) {

	if mysqlConf.MaxIdleConns == 0 {
		mysqlConf.MaxIdleConns = 10
	}
	if mysqlConf.MaxOpenConns == 0 {
		mysqlConf.MaxOpenConns = 100
	}
	if mysqlConf.MaxLifetime == 0 {
		mysqlConf.MaxLifetime = 3600
	}
	if mysqlConf.SlowThreshold == 0 {
		mysqlConf.SlowThreshold = int64(500 * time.Millisecond)
	}
	var err error
	once.Do(func() {
		gormEngine, err = gorm.Open(mysql.Open(mysqlConf.DataSource), &gorm.Config{
			Logger: NewGormLogger(time.Millisecond * time.Duration(mysqlConf.SlowThreshold)), // 调整日志级别，根据需要修改
		})
		if err != nil {
			return
		}
		db, err := gormEngine.DB()
		if err != nil {
			return
		}
		db.SetMaxIdleConns(mysqlConf.MaxIdleConns)
		db.SetMaxOpenConns(mysqlConf.MaxOpenConns)
		db.SetConnMaxLifetime(time.Second * time.Duration(mysqlConf.MaxLifetime))
		err = gormEngine.Use(NewCustomePlugin())
		if err != nil {
			return
		}
	})
	return gormEngine, err
}

func MustNewGormEngine(mysqlConf MysqlConf) *gorm.DB {
	db, err := NewGormEngine(mysqlConf)
	if err != nil {
		panic(err)
	}
	return db
}
