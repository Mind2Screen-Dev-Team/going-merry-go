package registry

import "github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"

type AppRegistry struct {
	Config     *appconfig.AppConfig
	Dependency *AppDependency
	Provider   *AppProvider
	Repository *AppRepository
	Service    *AppService
}

func NewAppRegistry(cfg *appconfig.AppConfig) *AppRegistry {
	return &AppRegistry{
		Config:     cfg,
		Dependency: NewAppDependency(),
		Provider:   NewAppProvider(),
		Repository: NewAppRepository(),
		Service:    NewAppService(),
	}
}
