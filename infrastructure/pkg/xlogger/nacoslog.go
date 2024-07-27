package xlogger

import (
	"github.com/nacos-group/nacos-sdk-go/v2/common/logger"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm/utils"
)

var _ logger.Logger = (*NacosLogger)(nil)

type NacosLogger struct {
}

func NewNacosLogger() *NacosLogger {
	return &NacosLogger{}
}

func (n *NacosLogger) Info(args ...interface{}) {
	logx.Infof("", append([]interface{}{utils.FileWithLineNum()}, args...)...)
}

func (n *NacosLogger) Warn(args ...interface{}) {
	logx.Errorf("", append([]interface{}{utils.FileWithLineNum()}, args...)...)
}

func (n *NacosLogger) Error(args ...interface{}) {
	logx.Errorf("", append([]interface{}{utils.FileWithLineNum()}, args...)...)
}

func (n *NacosLogger) Debug(args ...interface{}) {
	logx.Debug(append([]interface{}{utils.FileWithLineNum()}, args...)...)
}

func (n *NacosLogger) Infof(fmt string, args ...interface{}) {
	logx.Infof(fmt, append([]interface{}{utils.FileWithLineNum()}, args...)...)
}

func (n *NacosLogger) Warnf(fmt string, args ...interface{}) {
	logx.Errorf(fmt, append([]interface{}{utils.FileWithLineNum()}, args...)...)
}

func (n *NacosLogger) Errorf(fmt string, args ...interface{}) {
	logx.Errorf(fmt, append([]interface{}{utils.FileWithLineNum()}, args...)...)
}

func (n *NacosLogger) Debugf(fmt string, args ...interface{}) {
	logx.Debugf(fmt, append([]interface{}{utils.FileWithLineNum()}, args...)...)
}
