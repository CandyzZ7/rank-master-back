package e

//go:generate stringer -type ErrCode -linecomment

type ErrCode int64

const (
	ErrRegisterMobileExistCode ErrCode = 2000 + iota // 手机号已存在
)

const (
	ErrLoginPasswdCode ErrCode = 1000 + iota // 密码错误
)

const (
	ErrEmailCodeFailCode ErrCode = 3000 + iota // 邮箱验证码错误
)

const (
	ErrLoginMobileNotExistCode ErrCode = 4000 + iota // 手机号不存在
)
