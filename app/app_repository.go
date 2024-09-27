package app

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"

	repo_impl "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/repo/impl"
)

func AppRepositoryLoader(cfg *appconfig.AppConfig, dep *registry.AppDependency) *registry.AppRepository {
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
