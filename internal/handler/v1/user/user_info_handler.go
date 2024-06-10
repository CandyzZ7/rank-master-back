package user

import (
	"net/http"

	"rank-master-back/infrastructure/response"
	"rank-master-back/internal/logic/v1/user"
	"rank-master-back/internal/svc"

	"github.com/zeromicro/go-zero/core/logc"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		if err != nil {
			logc.Error(r.Context(), err)
			response.Handler(w, nil, err)
		} else {
			response.Handler(w, resp, err)
		}
	}
}
