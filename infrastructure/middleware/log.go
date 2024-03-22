package middleware

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logc"
)

// logMiddleware 日志中间件
func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logc.Info(context.Background(), "Method:", r.Method, ",", "Path:", r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
