package middleware

import (
	"net/http"
	"strings"

	"github.com/swaggest/swgui/v5emb"
)

const (
	swaggerAPI     = "/api/doc"
	SwaggerJsonAPI = "/api/doc/app.json"
	Title          = "title"
)

var SwaggerHandle = v5emb.New(
	Title,
	SwaggerJsonAPI,
	swaggerAPI,
)

func Notfound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, swaggerAPI) {
			SwaggerHandle.ServeHTTP(w, r)
			return
		}
		http.NotFound(w, r)
	}
}
