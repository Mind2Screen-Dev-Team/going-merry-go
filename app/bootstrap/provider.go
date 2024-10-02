package bootstrap

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
)

type LoaderProviderFn interface {
	Loader(ctx context.Context, reg *registry.AppRegistry)
}

func LoadProvider(ctx context.Context, reg *registry.AppRegistry, loaders ...LoaderProviderFn) {
	if loaders == nil {
		return
	}

	for _, l := range loaders {
		l.Loader(ctx, reg)
	}
}
