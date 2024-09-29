package app

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
)

func AppProviderLoader(cfg *appconfig.AppConfig, dep *registry.AppDependency) *registry.AppProvider {
	// # Load All Provider
	return bootstrap.LoadProvider(
		context.Background(),

		// Link Dependency
		dep,

		// # List Of Provider

		// add more on here...
	)
}
