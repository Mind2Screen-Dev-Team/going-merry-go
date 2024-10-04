package app

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"

	service_impl "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/service/impl"
)

func ServiceRegistryLoader(reg *registry.AppRegistry) error {
	// # Load All Service
	return bootstrap.LoadRegistry(
		context.Background(),

		// Link Dependency and Repository
		reg,

		// # List Of Service
		service_impl.NewUserServiceImpl(),

		// add more on here...
	)
}
