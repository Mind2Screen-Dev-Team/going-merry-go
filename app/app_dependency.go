package app

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/config"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
)

func AppDependencyLoader(cfg *appconfig.AppConfig) *bootstrap.AppDependency {
	// # Load All Dependency
	return bootstrap.LoadDependency(
		context.Background(),

		// # List of Dependency
		config.NewMySQLX(cfg),
		config.NewNatsClient(cfg),
		config.NewRedisClient(cfg),

		// add more on here...
	)
}
