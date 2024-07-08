package xlogger

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

var (
	infoStr           = "%s\n[info] "
	warnStr           = "%s\n[warn] "
	errStr            = "%s\n[error] "
	debugStr          = "%s\n[debug] "
	traceStr          = "%s\n[%.3fms] [rows:%v] %s"
	traceWarnStr      = "%s %s\n[%.3fms] [rows:%v] %s"
	traceErrStr       = "%s %s\n[%.3fms] [rows:%v] %s"
	DefaultGormLogger = NewGormLogger(logx.WithContext(context.TODO()).WithCallerSkip(1), 200*time.Millisecond)
)

type GormLogger struct {
	level         logger.LogLevel
	SlowThreshold time.Duration
	logger        logx.Logger
}

func NewGormLogger(logger logx.Logger, slowThreshold time.Duration) *GormLogger {
	return &GormLogger{logger: logger, SlowThreshold: slowThreshold}
}

func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.level = level
	return l
}

func (l *GormLogger) Info(ctx context.Context, s string, i ...interface{}) {

	if l.level == 0 || l.level >= logger.Info {
		l.logger.WithContext(ctx).Infof(infoStr+s, append([]interface{}{utils.FileWithLineNum()}, i...)...)
	}
}

func (l *GormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	if l.level == 0 || l.level >= logger.Warn {
		l.logger.WithContext(ctx).Errorf(warnStr+s, append([]interface{}{utils.FileWithLineNum()}, i...)...)
	}
}

func (l *GormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	if l.level == 0 || l.level >= logger.Error {
		l.logger.WithContext(ctx).Errorf(errStr+s, append([]interface{}{utils.FileWithLineNum()}, i...)...)
	}
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil && !errors.Is(err, gorm.ErrRecordNotFound):
		sql, rows := fc()
		if rows == -1 {
			l.Info(ctx, traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Info(ctx, traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.logger.Slowf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.logger.Slowf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	default:
		sql, rows := fc()
		if rows == -1 {
			l.Info(ctx, traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Info(ctx, traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
