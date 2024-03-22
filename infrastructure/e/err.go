package e

var (
	OK          = New(OKCode, OKCode.String())
	ServerError = New(ServerErrorCode, ServerErrorCode.String())
)

var (
	ErrRegisterMobileExist  = New(ErrRegisterMobileExistCode, ErrRegisterMobileExistCode.String())
	ErrRegisterAccountExist = New(ErrRegisterAccountExistCode, ErrRegisterAccountExistCode.String())
	ErrLoginPasswd          = New(ErrLoginPasswdCode, ErrLoginPasswdCode.String())
	ErrEmailCodeFail        = New(ErrEmailCodeFailCode, ErrEmailCodeFailCode.String())
	ErrLoginMobileNotExist  = New(ErrLoginMobileNotExistCode, ErrLoginMobileNotExistCode.String())
)
