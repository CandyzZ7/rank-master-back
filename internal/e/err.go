package e

import "github.com/pkg/errors"

var (
	ErrLoginMobileNotExist = errors.New(ErrLoginMobileNotExistCode.String())
	ErrLoginPasswd         = errors.New(ErrLoginPasswdCode.String())
	ErrRegisterMobileExist = errors.New(ErrRegisterMobileExistCode.String())
)
