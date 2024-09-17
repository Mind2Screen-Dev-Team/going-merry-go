package bootstrap

import (
	"context"
)

type LoaderRepositoryFn interface {
	Loader(ctx context.Context, dep *AppDependency, repo *AppRepository)
}

func LoadRepository(ctx context.Context, dep *AppDependency, loaders ...LoaderRepositoryFn) *AppRepository {
	var repo AppRepository

	if loaders == nil {
		return &repo
	}

	for _, l := range loaders {
		l.Loader(ctx, dep, &repo)
	}

	return &repo
}
