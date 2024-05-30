package crontask

import "github.com/zeromicro/go-zero/core/logx"

type CLog struct{}

func (l *CLog) Info(msg string, keysAndValues ...interface{}) {
	logx.Infow(msg, l.extractFields(keysAndValues)...)
}

func (l *CLog) Error(err error, msg string, keysAndValues ...interface{}) {
	logx.Errorw(err.Error()+msg, l.extractFields(keysAndValues)...)
}

func (l *CLog) extractFields(keysAndValues []interface{}) []logx.LogField {
	logFields := make([]logx.LogField, len(keysAndValues)/2)
	for i := 0; i < len(keysAndValues); i += 2 {
		key := keysAndValues[i]
		value := keysAndValues[i+1]
		logFields[i/2] = logx.LogField{Key: key.(string), Value: value}
	}
	return logFields
}
