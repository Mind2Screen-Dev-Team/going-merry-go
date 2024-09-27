package app

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/config"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
)

func AppDependencyLoader(ctx context.Context, cfg *appconfig.AppConfig, logFileName string) *registry.AppDependency {
	// # Load All Dependency
	return bootstrap.LoadDependency(
		ctx,
		cfg,

		// # List of Dependency
		config.NewLumberJackConfig(logFileName),
		config.NewZeroLogConfig(),
		config.NewHttpinCore(),
		config.NewMySqlX(),
		config.NewNatsClient(),
		config.NewRedisClient(),

		// add more on here...
	)
}
