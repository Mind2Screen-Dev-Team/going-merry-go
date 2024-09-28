package app

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/config"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
)

type AppDependencyLoaderParams struct {
	LogFilename      string
	LogDefaultFields map[string]any
}

func AppDependencyLoader(ctx context.Context, cfg *appconfig.AppConfig, param AppDependencyLoaderParams) *registry.AppDependency {
	// # Load All Dependency
	return bootstrap.LoadDependency(
		ctx,
		cfg,

		// # List of Dependency
		config.NewLumberJackConfig(param.LogFilename),
		config.NewZeroLogConfig(param.LogDefaultFields),
		config.NewMinioClient(),
		config.NewHttpinCore(),
		config.NewMySqlX(),
		config.NewNatsClient(),
		config.NewRedisClient(),

		// add more on here...
	)
}
