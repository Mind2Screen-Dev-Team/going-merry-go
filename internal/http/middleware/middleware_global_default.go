package middleware

import (
	"net/http"
	"runtime/debug"
	"time"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xlogger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

func Logger(logger xlogger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			reqId, _ := r.Context().Value(middleware.RequestIDKey).(string)
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			defer func() {
				if r := recover(); r != nil && r != http.ErrAbortHandler {
					logger.Error(
						// msg
						"incoming_request_panic",
						"trace_id", reqId,
						// fields
						"recover", r,
						"stack", debug.Stack(),
					)
					ww.WriteHeader(http.StatusInternalServerError)
				}

				logger.Info(
					// msg
					"incoming request",
					// fields
					"trace_id", reqId,
					"remote_addr", r.RemoteAddr,
					"path", r.URL.Path,
					"proto", r.Proto,
					"method", r.Method,
					"user_agent", r.UserAgent(),
					"status", http.StatusText(ww.Status()),
					"status_code", ww.Status(),
					"bytes_in", r.ContentLength,
					"bytes_out", ww.BytesWritten(),
				)
			}()
			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
