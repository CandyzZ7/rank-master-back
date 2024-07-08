package xlogger

import (
	"context"
	"time"

	"github.com/nacos-group/nacos-sdk-go/v2/common/logger"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm/utils"
)

var _ logger.Logger = (*NacosLogger)(nil)

var DefaultNacosLogger = NewNacosLogger(logx.WithContext(context.TODO()).WithCallerSkip(1), 200*time.Millisecond)

type NacosLogger struct {
	SlowThreshold time.Duration
	logger        logx.Logger
}

func NewNacosLogger(logger logx.Logger, slowThreshold time.Duration) *NacosLogger {
	return &NacosLogger{logger: logger, SlowThreshold: slowThreshold}
}

func (n *NacosLogger) Info(args ...interface{}) {
	n.logger.Infof(infoStr, append([]interface{}{utils.FileWithLineNum()}, args...)...)
}

func (n *NacosLogger) Warn(args ...interface{}) {
	n.logger.Errorf(warnStr, append([]interface{}{utils.FileWithLineNum()}, args...)...)
}

func (n *NacosLogger) Error(args ...interface{}) {
	n.logger.Errorf(errStr, append([]interface{}{utils.FileWithLineNum()}, args...)...)
}

func (n *NacosLogger) Debug(args ...interface{}) {
	n.logger.Debug(append([]interface{}{utils.FileWithLineNum()}, args...)...)
}

func (n *NacosLogger) Infof(fmt string, args ...interface{}) {
	n.logger.Infof(infoStr, append([]interface{}{utils.FileWithLineNum()}, args...)...)
}

func (n *NacosLogger) Warnf(fmt string, args ...interface{}) {
	n.logger.Errorf(warnStr+fmt, append([]interface{}{utils.FileWithLineNum()}, args...)...)
}

func (n *NacosLogger) Errorf(fmt string, args ...interface{}) {
	n.logger.Errorf(errStr+fmt, append([]interface{}{utils.FileWithLineNum()}, args...)...)
}

func (n *NacosLogger) Debugf(fmt string, args ...interface{}) {
	n.logger.Debugf(debugStr+fmt, append([]interface{}{utils.FileWithLineNum()}, args...)...)
}
