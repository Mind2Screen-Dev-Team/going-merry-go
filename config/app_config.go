package config

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/ctxkey"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

// # All

func LoadRegistry(ctx context.Context) *registry.AppRegistry {
	if v, ok := ctx.Value(ctxkey.REGISTRY_APP).(*registry.AppRegistry); ok {
		return v
	}
	return nil
}

// # Details

func LoadConfig(ctx context.Context) *appconfig.AppConfig {
	if v, ok := ctx.Value(ctxkey.REGISTRY_APP_CONFIG).(*appconfig.AppConfig); ok {
		return v
	}
	return nil
}

func LoadDependencyRegistry(ctx context.Context) *registry.DependencyRegistry {
	if v, ok := ctx.Value(ctxkey.REGISTRY_APP_DEPENDENCY).(*registry.DependencyRegistry); ok {
		return v
	}
	return nil
}

func LoadRepositoryRegistry(ctx context.Context) *registry.RepositoryRegistry {
	if v, ok := ctx.Value(ctxkey.REGISTRY_APP_REPOSITORY).(*registry.RepositoryRegistry); ok {
		return v
	}
	return nil
}

func LoadServiceRegistry(ctx context.Context) *registry.ServiceRegistry {
	if v, ok := ctx.Value(ctxkey.REGISTRY_APP_SERVICE).(*registry.ServiceRegistry); ok {
		return v
	}
	return nil
}

func LoadProviderRegistry(ctx context.Context) *registry.ProviderRegistry {
	if v, ok := ctx.Value(ctxkey.REGISTRY_APP_PROVIDER).(*registry.ProviderRegistry); ok {
		return v
	}
	return nil
}

// # Spesificts

func LoadLogger(ctx context.Context) *zerolog.Logger {
	if v, ok := ctx.Value(ctxkey.REGISTRY_APP_LOGGER).(*zerolog.Logger); ok {
		return v
	}
	return nil
}

func LoadTracer(ctx context.Context) trace.Tracer {
	if v, ok := ctx.Value(ctxkey.REGISTRY_APP_TRACER).(trace.Tracer); ok {
		return v
	}
	return nil
}
