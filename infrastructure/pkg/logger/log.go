package logger

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
	infoStr      = "%s\n[info] "
	warnStr      = "%s\n[warn] "
	errStr       = "%s\n[error] "
	traceStr     = "%s\n[%.3fms] [rows:%v] %s"
	traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
	traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	Default      = NewLogger(logx.WithContext(context.TODO()).WithCallerSkip(1), 200*time.Millisecond)
)

type Logger struct {
	level         logger.LogLevel
	SlowThreshold time.Duration
	// logger        logx.Logger
}

func NewLogger(logger logx.Logger, slowThreshold time.Duration) *Logger {
	return &Logger{SlowThreshold: slowThreshold}
}

func (l *Logger) LogMode(level logger.LogLevel) logger.Interface {
	l.level = level
	return l
}

func (l *Logger) Info(ctx context.Context, s string, i ...interface{}) {

	if l.level == 0 || l.level >= logger.Info {
		logx.WithContext(ctx).Infof(infoStr+s, append([]interface{}{utils.FileWithLineNum()}, i...)...)
	}
}

func (l *Logger) Warn(ctx context.Context, s string, i ...interface{}) {
	if l.level == 0 || l.level >= logger.Warn {
		logx.WithContext(ctx).Errorf(warnStr+s, append([]interface{}{utils.FileWithLineNum()}, i...)...)
	}
}

func (l *Logger) Error(ctx context.Context, s string, i ...interface{}) {
	if l.level == 0 || l.level >= logger.Error {
		logx.WithContext(ctx).Errorf(errStr+s, append([]interface{}{utils.FileWithLineNum()}, i...)...)
	}
}

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
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
			logx.Slowf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			logx.Slowf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
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
