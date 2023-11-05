// Code generated by "stringer -type ErrCode -linecomment"; DO NOT EDIT.

package e

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrRegisterMobileExistCode-2000]
	_ = x[ErrLoginPasswdCode-1000]
	_ = x[ErrEmailCodeFailCode-3000]
	_ = x[ErrLoginMobileNotExistCode-4000]
}

const (
	_ErrCode_name_0 = "密码错误"
	_ErrCode_name_1 = "手机号已存在"
	_ErrCode_name_2 = "邮箱验证码错误"
	_ErrCode_name_3 = "手机号不存在"
)

func (i ErrCode) String() string {
	switch {
	case i == 1000:
		return _ErrCode_name_0
	case i == 2000:
		return _ErrCode_name_1
	case i == 3000:
		return _ErrCode_name_2
	case i == 4000:
		return _ErrCode_name_3
	default:
		return "ErrCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
