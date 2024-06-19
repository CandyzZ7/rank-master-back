package template

import (
	"net/http"

	"rank-master-back/internal/logic/v1/template"
	"rank-master-back/internal/svc"
	"rank-master-back/internal/types"

	"rank-master-back/infrastructure/pkg/validator"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddTemplateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddTemplateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		err := validator.New().StructCtx(r.Context(), req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := template.NewAddTemplateLogic(r.Context(), svcCtx)
		resp, err := l.AddTemplate(&req)
		if err != nil {
			logc.Error(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
