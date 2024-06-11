package e

//go:generate stringer -type Code -linecomment

type Code int

const (
	// OKCode 成功
	OKCode Code = 200 // OK
	// ServerErrorCode 服务器错误
	ServerErrorCode Code = 500 // Server Error
)

const (
	// ErrLoginPasswdCode 密码错误
	ErrLoginPasswdCode Code = 1000 + iota // password error
)

const (
	// ErrRegisterMobileExistCode 手机号已存在
	ErrRegisterMobileExistCode Code = 2000 + iota // mobile already exists
	// ErrRegisterAccountExistCode 账号已存在
	ErrRegisterAccountExistCode // account already exists
)

const (
	// ErrEmailCodeFailCode 邮箱验证码错误
	ErrEmailCodeFailCode Code = 3000 + iota // email code error
)

const (
	// ErrLoginMobileNotExistCode 手机号不存在
	ErrLoginMobileNotExistCode Code = 4000 + iota // mobile not exists
)
