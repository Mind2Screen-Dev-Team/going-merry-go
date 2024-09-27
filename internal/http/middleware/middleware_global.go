package middleware

import (
	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
	"github.com/go-chi/chi/v5"
)

func Global(
	cfg *appconfig.AppConfig,
	dep *registry.AppDependency,
	repo *registry.AppRepository,
	service *registry.AppService,
	router chi.Router,
) {
	// Assign your global middleware on here ...
}
