package bootstrap

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
)

type LoaderDependencyFn interface {
	Loader(ctx context.Context, dep *registry.AppRegistry)
}

func LoadDependency(ctx context.Context, reg *registry.AppRegistry, loaders ...LoaderDependencyFn) {
	if loaders == nil {
		return
	}

	for _, l := range loaders {
		l.Loader(ctx, reg)
	}
}
