package template

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"rank-master-back/internal/logic/template"
	"rank-master-back/internal/svc"
	"rank-master-back/internal/types"
)

func AddTemplateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddTemplateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := template.NewAddTemplateLogic(r.Context(), svcCtx)
		resp, err := l.AddTemplate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
