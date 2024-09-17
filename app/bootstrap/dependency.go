package bootstrap

import (
	"context"
)

type LoaderDependencyFn interface {
	Loader(ctx context.Context, dep *AppDependency)
}

func LoadDependency(ctx context.Context, loaders ...LoaderDependencyFn) *AppDependency {
	var dep AppDependency

	if loaders == nil {
		return &dep
	}

	for _, l := range loaders {
		l.Loader(ctx, &dep)
	}

	return &dep
}
