package e

//go:generate stringer -type Code -linecomment

type Code int64

const (
	OKCode          Code = 200 // OK
	ServerErrorCode Code = 500 // Server Error
)

const (
	ErrRegisterMobileExistCode  Code = 2000 + iota // mobile already exists
	ErrRegisterAccountExistCode                    // account already exists
)

const (
	ErrLoginPasswdCode Code = 1000 + iota // password error
)

const (
	ErrEmailCodeFailCode Code = 3000 + iota // email code error
)

const (
	ErrLoginMobileNotExistCode Code = 4000 + iota // mobile not exists
)
