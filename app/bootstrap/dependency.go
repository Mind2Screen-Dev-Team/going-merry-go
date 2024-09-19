package bootstrap

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
)

type LoaderDependencyFn interface {
	Loader(ctx context.Context, dep *registry.AppDependency)
}

func LoadDependency(ctx context.Context, loaders ...LoaderDependencyFn) *registry.AppDependency {
	var dep registry.AppDependency

	if loaders == nil {
		return &dep
	}

	for _, l := range loaders {
		l.Loader(ctx, &dep)
	}

	return &dep
}
