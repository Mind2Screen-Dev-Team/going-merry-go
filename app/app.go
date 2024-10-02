package app

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
)

func LoadRegistry(ctx context.Context, cfg *appconfig.AppConfig, param AppDependencyLoaderParams) *registry.AppRegistry {
	// # Initiated Registry
	reg := registry.NewAppRegistry(cfg)

	// # Load All Dependency
	AppDependencyLoader(ctx, reg, param)

	// # Load All Provider
	AppProviderLoader(reg)

	// # Load All Repository
	AppRepositoryLoader(reg)

	// # Load All Service
	AppServiceLoader(reg)

	return reg
}
