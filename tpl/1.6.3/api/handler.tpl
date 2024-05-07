package {{.PkgName}}

import (
	"net/http"

    "rank-master-back/infrastructure/response"
	{{.ImportPackages}}

	{{if .HasRequest}}"github.com/go-playground/validator/v10"{{end}}
   	"github.com/zeromicro/go-zero/core/logc"
    {{if .HasRequest}}"github.com/zeromicro/go-zero/rest/httpx"{{end}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			response.Handler(w, nil, err)
			return
		}

		err := validator.New().StructCtx(r.Context(), req)
		if err != nil {
			response.Handler(w, nil, err)
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if err != nil {
			logc.Error(r.Context(), err)
			response.Handler(w, nil, err)
		} else {
			response.Handler(w, resp, err)
		}
	}
}
