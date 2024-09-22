package app

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
)

func LoadRegistry(ctx context.Context, cfg *appconfig.AppConfig) (dep *registry.AppDependency, repo *registry.AppRepository, service *registry.AppService) {

	// # Load All Dependency
	dep = AppDependencyLoader(ctx, cfg)

	// # Load All Repository
	repo = AppRepositoryLoader(cfg, dep)

	// # Load All Service
	service = AppServiceLoader(cfg, dep, repo)

	return
}
