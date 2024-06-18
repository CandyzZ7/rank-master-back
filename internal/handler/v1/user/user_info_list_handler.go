package user

import (
	"net/http"

	"rank-master-back/internal/logic/v1/user"
	"rank-master-back/internal/svc"
	"rank-master-back/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserInfoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserInfoListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		err := validator.New().StructCtx(r.Context(), req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserInfoListLogic(r.Context(), svcCtx)
		resp, err := l.UserInfoList(&req)
		if err != nil {
			logc.Error(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
