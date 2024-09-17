package app

import (
	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
)

func LoadRegistry(appConfig *appconfig.AppConfig) (dep *bootstrap.AppDependency, repo *bootstrap.AppRepository, service *bootstrap.AppService) {

	// # Load All Dependency
	dep = AppDependencyLoader(appConfig)

	// # Load All Repository
	repo = AppRepositoryLoader(appConfig, dep)

	// # Load All Service
	service = AppServiceLoader(appConfig, dep, repo)

	return
}
