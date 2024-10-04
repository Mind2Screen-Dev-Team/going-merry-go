package app

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
)

func ProviderRegistryLoader(reg *registry.AppRegistry) error {
	// # Load All Provider
	return bootstrap.LoadRegistry(
		context.Background(),

		// Link Dependency
		reg,

		// # List Of Provider

		// add more on here...
	)
}
