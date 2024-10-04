package app

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/config"
)

type DependencyRegistryLoaderParams struct {
	Module           string
	ServerName       string
	ServerAddr       string
	LogFilename      string
	LogDefaultFields map[string]any
}

func DependencyRegistryLoader(ctx context.Context, reg *registry.AppRegistry, param DependencyRegistryLoaderParams) error {
	// # Load All Dependency
	return bootstrap.LoadRegistry(
		ctx,
		reg,

		// # List of Dependency
		config.NewLumberJackConfig(param.LogFilename),
		config.NewZeroLogConfig(param.LogDefaultFields),
		config.NewOtelClient(config.OtelParam{
			Module:        param.Module,
			ServerName:    param.ServerName,
			ServerAddress: param.ServerAddr,
		}),

		config.NewMinioClient(),
		config.NewHttpinCore(),
		config.NewMySqlX(),
		config.NewNatsClient(),
		config.NewRedisClient(),

		// add more on here...
	)
}
