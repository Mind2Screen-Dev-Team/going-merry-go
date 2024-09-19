package app

import (
	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
)

func LoadRegistry(cfg *appconfig.AppConfig) (dep *bootstrap.AppDependency, repo *bootstrap.AppRepository, service *bootstrap.AppService) {

	// # Load All Dependency
	dep = AppDependencyLoader(cfg)

	// # Load All Repository
	repo = AppRepositoryLoader(cfg, dep)

	// # Load All Service
	service = AppServiceLoader(cfg, dep, repo)

	return
}
