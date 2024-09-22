package middleware

import (
	"net/http"
	"runtime/debug"
	"time"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

func DefaultGlobal(cfg *appconfig.AppConfig, dep *registry.AppDependency, r chi.Router) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(Logger(dep.Logger))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(
		time.Duration(cfg.App.Http.HandlerTimeout) * time.Second,
	))
	r.Use(middleware.Heartbeat("/api/v1/health"))
}

func Logger(logger *zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			defer func() {
				if r := recover(); r != nil && r != http.ErrAbortHandler {
					logger.Error().Interface("recover", r).Bytes("stack", debug.Stack()).Msg("incoming_request_panic")
					ww.WriteHeader(http.StatusInternalServerError)
				}
				logger.Info().Fields(map[string]interface{}{
					"remote_addr": r.RemoteAddr,
					"path":        r.URL.Path,
					"proto":       r.Proto,
					"method":      r.Method,
					"user_agent":  r.UserAgent(),
					"status":      http.StatusText(ww.Status()),
					"status_code": ww.Status(),
					"bytes_in":    r.ContentLength,
					"bytes_out":   ww.BytesWritten(),
				}).Msg("incoming_request")
			}()
			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
