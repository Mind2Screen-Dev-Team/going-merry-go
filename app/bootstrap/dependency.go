package bootstrap

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
)

type LoaderDependencyFn interface {
	Loader(ctx context.Context, cfg *appconfig.AppConfig, dep *registry.AppDependency)
}

func LoadDependency(ctx context.Context, cfg *appconfig.AppConfig, loaders ...LoaderDependencyFn) *registry.AppDependency {
	var dep registry.AppDependency

	if loaders == nil {
		return &dep
	}

	for _, l := range loaders {
		l.Loader(ctx, cfg, &dep)
	}

	return &dep
}
