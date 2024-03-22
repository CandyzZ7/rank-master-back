package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"rank-master-back/infrastructure/e"
)

type Body struct {
	Code    e.Code      `json:"code"`
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"data,omitempty"`
}

// Handler 统一返回入口，
func Handler(w http.ResponseWriter, resp interface{}, err error) {
	if err != nil {
		httpx.OkJson(w, e.ErrHandler(err))
		// 如果err不为空的话，走错误处理函数，将err传递过去
	} else {
		// 如果err为空的话，走正常返回函数，将resp传递过去
		httpx.OkJson(w, Body{
			Code:    e.OK.Code,
			Message: e.OK.Message,
			Result:  resp,
		})
	}
}
