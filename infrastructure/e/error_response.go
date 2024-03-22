package e

import (
	"errors"
)

type CodeError struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

// 实现error的接口  然后CodeError继承一下Error方法  CodeError就为error类型的返回值
func (e *CodeError) Error() string {
	return e.Message
}

func (e *CodeError) GetCode() Code {
	return e.Code
}

func (e *CodeError) GetMessage() string {
	return e.Message
}

// ErrorResponse 返回给前端的数据
func (e *CodeError) ErrorResponse() *CodeError {
	return &CodeError{
		Code:    e.Code,
		Message: e.Message,
	}
}

// New 提供new方法，任意地方传递参数返回CodeError类型的数据
func New(code Code, msg string) *CodeError {
	return &CodeError{
		Code:    code,
		Message: msg,
	}
}

// DefaultErrHandler 默认异常状态码函数，只需传递错误信息即可，默认返回code-10001
func DefaultErrHandler(msg string) error {
	return &CodeError{
		Code:    ServerError.Code,
		Message: msg,
	}
}

// ErrHandler 自定义错误返回函数 错误函数主入口
func ErrHandler(err error) interface{} {
	var codeError *CodeError
	switch {
	// 如果错误类型为CodeError，就返回错误类型的结构体
	case errors.As(err, &codeError):
		return err
	default:
		// 系统错误，500 错误提示
		return DefaultErrHandler(err.Error())
	}
}
