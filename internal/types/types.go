// Code generated by goctl. DO NOT EDIT.
package types

type AddTemplateReq struct {
	Template Template `json:"template"`
}

type AddTemplateResp struct {
	Id string `json:"id"`
}

type GetEmailCodeReq struct {
	Email string `json:"email" validate:"required,email"` // 邮箱
}

type GetEmailCodeResp struct {
}

type GetRankMasterAccountReq struct {
	RankMasterAccount string `path:"rank_master_account" validate:"required"` // RankMaster账号
}

type GetRankMasterAccountResp struct {
}

type GetUserInfoListReq struct {
	Pagination Pagination `json:"pagination"` // 分页信息
}

type GetUserInfoListResp struct {
	UserList []*User `json:"user_list"`
	Count    int64   `json:"count"` // 总数
}

type GetUserInfoReq struct {
}

type GetUserInfoResp struct {
	User User `json:"user"`
}

type LoginReq struct {
	RankMasterAccount string `json:"rank_master_account" validate:"required"` // RankMaster账号
	Password          string `json:"password" validate:"required"`            // 密码
}

type LoginResp struct {
	UserId string `json:"user_id"` // 用户ID
	Token  Token  `json:"token"`   // token
}

type Pagination struct {
	Page      int         `json:"page"`      // 当前页码
	PageSize  int         `json:"pageSize"`  // 每页条数
	SortBy    string      `json:"sortBy"`    // 排序字段
	SortOrder string      `json:"sortOrder"` // 排序顺序：asc 或 desc
	Filter    interface{} `json:"filter"`    // 过滤条件，可以是一个结构体
}

type PingResp struct {
	Msg string `json:"msg"`
}

type RegisterReq struct {
	User User `json:"user"`
}

type RegisterResp struct {
	UserId string `json:"user_id"` // 用户ID
	Token  Token  `json:"token"`   // token
}

type Template struct {
	Function string `json:"function" validate:"required"`
	Type     string `json:"type" validate:"required"`
	Topic    string `json:"topic" validate:"required"`
	Content  string `json:"content" validate:"required"`
	Remark   string `json:"remark" validate:"required"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
}

type User struct {
	Name              string `json:"name" validate:"required"`                // 昵称
	RankMasterAccount string `json:"rank_master_account" validate:"required"` // RankMaster账号
	Mobile            string `json:"mobile" validate:"required,len=11,phone"` // 手机号
	Avatar            string `json:"avatar" validate:"required"`              // 头像
	Email             string `json:"email" validate:"required,email"`         // 邮箱
	Code              string `json:"code" validate:"required"`                // 邮箱验证码
	Password          string `json:"password" validate:"required"`            // 密码
}
