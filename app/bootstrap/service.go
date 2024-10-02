package bootstrap

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
)

type LoaderServiceFn interface {
	Loader(ctx context.Context, reg *registry.AppRegistry)
}

func LoadService(ctx context.Context, reg *registry.AppRegistry, loaders ...LoaderServiceFn) {
	if loaders == nil {
		return
	}

	for _, l := range loaders {
		l.Loader(
			ctx,
			reg,
		)
	}
}
