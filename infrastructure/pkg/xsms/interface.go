package xsms

import (
	"github.com/pkg/errors"
)

const (
	CloudServiceTencent = "tencent"
)

type SmsConf struct {
	Enabled             bool
	Type                string
	TencentSms          *TencentSmsConfig
	SmsDailyLimit       int64
	SmsRateLimitSeconds int64
	SmsErrorMaxCount    int64
	SmsErrorPeriod      int64
}
type TencentSmsConfig struct {
	SecretId   string
	SecretKey  string
	AppId      string
	Sign       string
	TemplateId string
	Region     string
}

type ISmsSender interface {
	SendMsg(mobiles []string, args ...interface{}) error
}

type EmptySmsSender struct {
}

func (e *EmptySmsSender) SendMsg(mobiles []string, args ...interface{}) error {
	return nil
}

func NewSmsSender(config SmsConf) (ISmsSender, error) {
	if !config.Enabled {
		return &EmptySmsSender{}, nil
	}
	switch config.Type {
	case CloudServiceTencent:
		smsSender, err := NewTencentSmsSender(config.TencentSms)
		if err != nil {
			return nil, err
		}
		return smsSender, nil
	}
	return nil, errors.New("not support sms type")
}
