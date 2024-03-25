package e

var (
	OK          = newStatusCode(OKCode, OKCode.String())
	ServerError = newStatusCode(ServerErrorCode, ServerErrorCode.String())
)

var (
	ErrRegisterMobileExist  = newStatusCode(ErrRegisterMobileExistCode, ErrRegisterMobileExistCode.String())
	ErrRegisterAccountExist = newStatusCode(ErrRegisterAccountExistCode, ErrRegisterAccountExistCode.String())
	ErrLoginPasswd          = newStatusCode(ErrLoginPasswdCode, ErrLoginPasswdCode.String())
	ErrEmailCodeFail        = newStatusCode(ErrEmailCodeFailCode, ErrEmailCodeFailCode.String())
	ErrLoginMobileNotExist  = newStatusCode(ErrLoginMobileNotExistCode, ErrLoginMobileNotExistCode.String())
)
