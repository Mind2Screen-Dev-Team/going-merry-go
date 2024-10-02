package bootstrap

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
)

type LoaderRepositoryFn interface {
	Loader(ctx context.Context, reg *registry.AppRegistry)
}

func LoadRepository(ctx context.Context, reg *registry.AppRegistry, loaders ...LoaderRepositoryFn) {
	if loaders == nil {
		return
	}

	for _, l := range loaders {
		l.Loader(ctx, reg)
	}
}
