package user

import (
	"net/http"

	"rank-master-back/infrastructure/response"
	"rank-master-back/internal/logic/v1/user"
	"rank-master-back/internal/svc"
	"rank-master-back/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetEmailCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetEmailCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Handler(w, nil, err)
			return
		}

		err := validator.New().StructCtx(r.Context(), req)
		if err != nil {
			response.Handler(w, nil, err)
			return
		}

		l := user.NewGetEmailCodeLogic(r.Context(), svcCtx)
		resp, err := l.GetEmailCode(&req)
		if err != nil {
			logc.Error(r.Context(), err)
			response.Handler(w, nil, err)
		} else {
			response.Handler(w, resp, err)
		}
	}
}
