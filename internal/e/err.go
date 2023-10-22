package e

//go:generate stringer -type ErrCode -linecomment

type ErrCode int64

const (
	ErrRegisterMobileEmpty ErrCode = 3000 + iota // 手机号不能为空
	ErrRegisterPasswdEmpty                       // 密码不能为空
	ErrRegisterNameEmpty                         // 用户名不能为空
)

const (
	ErrRegisterMobileExist ErrCode = 2000 + iota // 手机号已存在
)
