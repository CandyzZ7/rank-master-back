package main

import (
	_ "embed"
	"flag"
	"fmt"
	"net/http"
	"strings"

	"rank-master-back/internal/config"
	"rank-master-back/internal/handler"
	"rank-master-back/internal/svc"

	"github.com/swaggest/swgui/v5emb"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

//go:embed doc/swagger/app.json
var spec []byte

var configFile = flag.String("f", "etc/app.yaml", "the config file")

const (
	swaggerAPI     = "/api/doc"
	SwaggerJsonAPI = "/api/doc/app.json"
	Title          = "title"
)

var swaggerHandle = v5emb.New(
	Title,
	SwaggerJsonAPI,
	swaggerAPI,
)

func Notfound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, swaggerAPI) {
			swaggerHandle.ServeHTTP(w, r)
			return
		}
		http.NotFound(w, r)
	}
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithNotFoundHandler(Notfound()))
	defer server.Stop()
	// swagger  json file
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   SwaggerJsonAPI,
		Handler: func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write(spec)
		},
	})
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	fmt.Println("doc: http://localhost:8888/api/doc")

	server.Start()
}
