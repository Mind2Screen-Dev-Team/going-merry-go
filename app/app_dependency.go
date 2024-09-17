package app

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/config"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
)

func AppDependencyLoader(appConfig *appconfig.AppConfig) *bootstrap.AppDependency {
	// # Load All Dependency
	return bootstrap.LoadDependency(
		context.Background(),

		// # List of Dependency
		config.NewMySQLX(appConfig),
		config.NewNatsClient(appConfig),

		// add more on here...
	)
}
