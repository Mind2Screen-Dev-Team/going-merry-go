package app

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"

	service_impl "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/service/impl"
)

func AppServiceLoader(
	cfg *appconfig.AppConfig,
	dep *registry.AppDependency,
	repo *registry.AppRepository,
	prov *registry.AppProvider,
) *registry.AppService {
	// # Load All Service
	return bootstrap.LoadService(
		context.Background(),

		// Link Dependency and Repository
		dep,
		repo,
		prov,

		// # List Of Service
		service_impl.NewUserServiceImpl(),

		// add more on here...
	)
}
