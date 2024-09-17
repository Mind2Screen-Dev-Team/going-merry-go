package app

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"

	repo_impl "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/repo/impl"
)

func AppRepositoryLoader(appConfig *appconfig.AppConfig, dep *bootstrap.AppDependency) *bootstrap.AppRepository {
	// # Load All Repository
	return bootstrap.LoadRepository(
		context.Background(),

		// Link Dependency
		dep,

		// # List Of Repository
		repo_impl.NewUserRepoImpl(),

		// add more on here...
	)
}
