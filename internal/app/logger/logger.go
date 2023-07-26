package logger

import (
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func LoggingMiddleware(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			next.ServeHTTP(ww, r)

			duration := time.Since(start)

			logger.Info("Request",
				zap.String("uri", r.RequestURI),
				zap.String("method", r.Method),
				zap.Int("status_code", ww.Status()),
				zap.Int("response_size", ww.BytesWritten()),
				zap.Duration("duration", duration),
			)
		})
	}
}
