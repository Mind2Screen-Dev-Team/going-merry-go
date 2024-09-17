package middleware

import (
	"time"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func DefaultGlobal(cfg *appconfig.AppConfig, r chi.Router) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(
		time.Duration(cfg.AppHttp.HandlerTimeout) * time.Second,
	))
	r.Use(middleware.Heartbeat("/api/v1/health"))
}
