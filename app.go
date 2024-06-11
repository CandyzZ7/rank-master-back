package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"

	"rank-master-back/infrastructure/middleware"
	"rank-master-back/infrastructure/response"
	"rank-master-back/internal/config"
	"rank-master-back/internal/handler"
	"rank-master-back/internal/svc"
)

//go:embed doc/swagger/app.json
var spec []byte

var configFile = flag.String("f", "etc/app.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors(), rest.WithNotFoundHandler(middleware.Notfound()))
	defer server.Stop()

	if c.Mode == service.DevMode || c.Mode == service.TestMode {
		// 新增swagger json接口
		server.AddRoute(rest.Route{
			Method: http.MethodGet,
			Path:   "/api/doc",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				if strings.HasPrefix(r.URL.Path, middleware.SwaggerJsonAPI) {
					_, _ = w.Write(spec)
				}
				middleware.SwaggerHandle.ServeHTTP(w, r)

			},
		})
		fmt.Println("doc: http://localhost:8888/api/doc")
	}

	ctx, err := InitializeServiceContext(c)
	if err != nil {
		logc.Error(context.Background(), errors.Cause(err))
	}
	// 初始化
	svc.Init(ctx)
	// 注册路由
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	// 自定义错误处理方法
	httpx.SetErrorHandlerCtx(response.ErrHandler)
	// 自定义返回成功方法
	// httpx.SetOkHandler(response.OKHandler)
	server.Start()
}
