package app

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"

	service_impl "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/service/impl"
)

func AppServiceLoader(cfg *appconfig.AppConfig, dep *bootstrap.AppDependency, repo *bootstrap.AppRepository) *bootstrap.AppService {
	// # Load All Service
	return bootstrap.LoadService(
		context.Background(),

		// Link Dependency and Repository
		dep,
		repo,

		// # List Of Service
		service_impl.NewUserServiceImpl(),
	)
}
