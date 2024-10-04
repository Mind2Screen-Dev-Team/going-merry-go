package config

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/ctxkey"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

func LoadAppConfig(ctx context.Context) *appconfig.AppConfig {
	if v, ok := ctx.Value(ctxkey.REGISTRY_APP_CONFIG).(*appconfig.AppConfig); ok {
		return v
	}
	return nil
}

func LoadAppDependency(ctx context.Context) *registry.AppDependency {
	if v, ok := ctx.Value(ctxkey.REGISTRY_APP_DEPENDENCY).(*registry.AppDependency); ok {
		return v
	}
	return nil
}

func LoadAppRepository(ctx context.Context) *registry.AppRepository {
	if v, ok := ctx.Value(ctxkey.REGISTRY_APP_REPOSITORY).(*registry.AppRepository); ok {
		return v
	}
	return nil
}

func LoadAppService(ctx context.Context) *registry.AppService {
	if v, ok := ctx.Value(ctxkey.REGISTRY_APP_SERVICE).(*registry.AppService); ok {
		return v
	}
	return nil
}

func LoadAppProvider(ctx context.Context) *registry.AppProvider {
	if v, ok := ctx.Value(ctxkey.REGISTRY_APP_PROVIDER).(*registry.AppProvider); ok {
		return v
	}
	return nil
}

func LoadAppLogger(ctx context.Context) *zerolog.Logger {
	if v, ok := ctx.Value(ctxkey.REGISTRY_APP_LOGGER).(*zerolog.Logger); ok {
		return v
	}
	return nil
}

func LoadAppTracer(ctx context.Context) trace.Tracer {
	if v, ok := ctx.Value(ctxkey.REGISTRY_APP_TRACER).(trace.Tracer); ok {
		return v
	}
	return nil
}
