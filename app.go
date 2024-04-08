package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest"

	"rank-master-back/infrastructure/middleware"
	"rank-master-back/infrastructure/pkg/snowflake"
	"rank-master-back/internal/config"
	"rank-master-back/internal/handler"
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
	// server.Use(logMiddleware)
	// swagger  json file
	// 新增swagger json接口
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   middleware.SwaggerJsonAPI,
		Handler: func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write(spec)
		},
	})
	// 初始化
	snowflake.InitNode(c)
	ctx, err := InitializeServiceContext(c)
	if err != nil {
		logc.Error(context.Background(), errors.Cause(err))
	}
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	fmt.Println("doc: http://localhost:8888/api/doc")

	server.Start()
}
