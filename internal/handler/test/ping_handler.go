package test

import (
	"net/http"

	"rank-master-back/infrastructure/response"
	"rank-master-back/internal/logic/test"
	"rank-master-back/internal/svc"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
)

func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := test.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping()
		if err != nil {
			logc.Error(r.Context(), errors.Cause(err))
			response.Handler(w, nil, err)
		} else {
			response.Handler(w, resp, err)
		}
	}
}
