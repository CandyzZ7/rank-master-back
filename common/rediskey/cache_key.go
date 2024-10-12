package rediskey

type RedisKey string

const (
	SendSmsCodeKey          RedisKey = "send_sms_code"
	SendSmsLimitKey         RedisKey = "send_sms_limit"
	SendSmsLimitEverydayKey RedisKey = "send_sms_limit_everyday"
	SendSmsCodeErrorKey     RedisKey = "send_sms_code_error"
	SmsLockKey              RedisKey = "sms_lock"
)

const (
	UserPasswordErrorCountKey RedisKey = "user_password_error_count"
)

const (
	TokenKey RedisKey = "token"
)

const (
	SendEmailCodeKey          RedisKey = "send_email_code"
	SendEmailLimitKey         RedisKey = "send_email_limit"
	SendEmailLimitEverydayKey RedisKey = "send_email_limit_everyday"
	SendEmailCodeErrorKey     RedisKey = "send_email_code_error"
	EmailLockKey              RedisKey = "email_lock"
)

func (key RedisKey) WithSymbol(symbol string) string {
	return string(key) + "_" + symbol
}

func (key RedisKey) WithParams(params ...string) string {
	if len(params) == 0 {
		return string(key)
	}
	k := string(key)
	for _, v := range params {
		k += ":" + v
	}
	return k
}
