package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"rank-master-back/internal/logic/user"
	"rank-master-back/internal/svc"
	"rank-master-back/internal/types"
)

// 获取邮箱验证码
func GetEmailCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetEmailCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetEmailCodeLogic(r.Context(), svcCtx)
		resp, err := l.GetEmailCode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
