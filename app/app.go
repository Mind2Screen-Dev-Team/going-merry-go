package app

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
)

func LoadRegistry(ctx context.Context, cfg *appconfig.AppConfig, logFileName string) (dep *registry.AppDependency, repo *registry.AppRepository, service *registry.AppService) {

	// # Load All Dependency
	dep = AppDependencyLoader(ctx, cfg, logFileName)

	// # Load All Repository
	repo = AppRepositoryLoader(cfg, dep)

	// # Load All Service
	service = AppServiceLoader(cfg, dep, repo)

	return
}
