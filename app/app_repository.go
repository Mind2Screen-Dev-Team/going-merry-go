package app

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"

	repo_impl "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/repo/impl"
)

func RepositoryRegistryLoader(reg *registry.AppRegistry) error {
	// # Load All Repository
	return bootstrap.LoadRegistry(
		context.Background(),

		// Link Dependency
		reg,

		// # List Of Repository
		repo_impl.NewUserRepoImpl(),

		// add more on here...
	)
}
