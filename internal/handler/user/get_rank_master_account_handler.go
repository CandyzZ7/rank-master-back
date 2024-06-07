package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"rank-master-back/internal/logic/user"
	"rank-master-back/internal/svc"
	"rank-master-back/internal/types"
)

// 验证用户账号是否存在
func GetRankMasterAccountHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRankMasterAccountReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetRankMasterAccountLogic(r.Context(), svcCtx)
		resp, err := l.GetRankMasterAccount(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
