// Code generated by goctl. DO NOT EDIT.
package types

type PingRes struct {
	Msg string `json:"msg"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
}

type RegisterReq struct {
	Name     string `json:"name" validate:"required"`
	Mobile   string `json:"mobile" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRes struct {
	UserId string `json:"user_id"`
	Token  Token  `json:"token"`
}
