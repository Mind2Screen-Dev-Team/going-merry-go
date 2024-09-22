package config

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/ctxkey"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
	"github.com/rs/zerolog"
)

var (
	nopLogger = zerolog.Nop()
)

func LoadAppConfig(ctx context.Context) *appconfig.AppConfig {
	v, ok := ctx.Value(ctxkey.HTTP_SERVER_APP_CONFIG).(*appconfig.AppConfig)
	if !ok {
		return nil
	}
	return v
}

func LoadAppDependency(ctx context.Context) *registry.AppDependency {
	v, ok := ctx.Value(ctxkey.HTTP_SERVER_APP_DEPENDENCY).(*registry.AppDependency)
	if !ok {
		return nil
	}
	return v
}

func LoadAppRepository(ctx context.Context) *registry.AppRepository {
	v, ok := ctx.Value(ctxkey.HTTP_SERVER_APP_REPOSITORY).(*registry.AppRepository)
	if !ok {
		return nil
	}
	return v
}

func LoadAppService(ctx context.Context) *registry.AppService {
	v, ok := ctx.Value(ctxkey.HTTP_SERVER_APP_SERVICE).(*registry.AppService)
	if !ok {
		return nil
	}
	return v
}

func Logger(ctx context.Context) *zerolog.Logger {
	v, ok := ctx.Value(ctxkey.HTTP_SERVER_APP_LOGGER).(*zerolog.Logger)
	if !ok {
		v = &nopLogger
	}
	return v
}
