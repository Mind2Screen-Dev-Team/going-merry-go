package app

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
)

func LoadRegistry(ctx context.Context, cfg *appconfig.AppConfig, param DependencyRegistryLoaderParams) *registry.AppRegistry {
	// # Initiated Registry
	reg := registry.NewAppRegistry(cfg)

	// # Load All Dependency
	if err := DependencyRegistryLoader(ctx, reg, param); err != nil {
		panic(err)
	}

	// # Load All Provider
	if err := ProviderRegistryLoader(reg); err != nil {
		panic(err)
	}

	// # Load All Repository
	if err := RepositoryRegistryLoader(reg); err != nil {
		panic(err)
	}

	// # Load All Service
	if err := ServiceRegistryLoader(reg); err != nil {
		panic(err)
	}

	return reg
}
