package middleware

import (
	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
	"github.com/go-chi/chi/v5"
)

func Global(
	cfg *appconfig.AppConfig,
	dep *bootstrap.AppDependency,
	repo *bootstrap.AppRepository,
	service *bootstrap.AppService,
	router chi.Router,
) {
	// Assign your global middleware on here ...
}
