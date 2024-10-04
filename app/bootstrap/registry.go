package bootstrap

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
)

type LoaderRegistryFunc interface {
	Loader(ctx context.Context, dep *registry.AppRegistry) error
}

func LoadRegistry(ctx context.Context, reg *registry.AppRegistry, loaders ...LoaderRegistryFunc) error {
	if loaders == nil {
		return nil
	}

	for _, l := range loaders {
		if err := l.Loader(ctx, reg); err != nil {
			return err
		}
	}

	return nil
}
