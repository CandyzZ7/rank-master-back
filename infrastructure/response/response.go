package response

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"

	"rank-master-back/infrastructure/e"
)

type Body struct {
	Code    e.Code `json:"code"`
	Message string `json:"message,omitempty"`
	Result  any    `json:"data,omitempty"`
}

// Handler 统一返回入口，
func Handler(w http.ResponseWriter, resp any, err error) {
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

func ErrHandler(ctx context.Context, err error) (int, any) {
	if e.IsGrpcError(err) {
		return e.CodeFromGrpcError(err), Body{
			Code:    e.Code(e.CodeFromGrpcError(err)),
			Message: err.Error(),
		}
	} else {
		var codeError *e.StatusCode
		switch {
		// 如果错误类型为CodeError，就返回错误类型的结构体
		case errors.As(err, &codeError):
			return http.StatusOK, Body{
				Code:    codeError.Code,
				Message: codeError.Message,
			}
		default:
			return http.StatusBadRequest, Body{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
		}
	}
}

func OKHandler(ctx context.Context, data any) any {
	return Body{
		Code:    e.OK.Code,
		Message: e.OK.Message,
		Result:  data,
	}
}
