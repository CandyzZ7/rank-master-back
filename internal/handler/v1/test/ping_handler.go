package test

import (
	"net/http"

	"rank-master-back/internal/logic/v1/test"
	"rank-master-back/internal/svc"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := test.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping()
		if err != nil {
			logc.Error(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
