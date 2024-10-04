package registry

import "github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"

type AppRegistry struct {
	Config     *appconfig.AppConfig
	Dependency *DependencyRegistry
	Provider   *ProviderRegistry
	Repository *RepositoryRegistry
	Service    *ServiceRegistry
}

func NewAppRegistry(cfg *appconfig.AppConfig) *AppRegistry {
	return &AppRegistry{
		Config:     cfg,
		Dependency: NewDependencyRegistry(),
		Provider:   NewProviderRegistry(),
		Repository: NewRepositoryRegistry(),
		Service:    NewServiceRegistry(),
	}
}
