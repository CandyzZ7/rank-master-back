// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	template "rank-master-back/internal/handler/template"
	test "rank-master-back/internal/handler/test"
	user "rank-master-back/internal/handler/user"
	"rank-master-back/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ping",
				Handler: test.PingHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/email/code",
				Handler: user.GetEmailCodeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:rank_master_account",
				Handler: user.GetRankMasterAccountHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: template.AddTemplateHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/template"),
	)
}
