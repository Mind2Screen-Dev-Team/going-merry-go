package bootstrap

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
)

type LoaderRepositoryFn interface {
	Loader(ctx context.Context, dep *registry.AppDependency, prov *registry.AppProvider, repo *registry.AppRepository)
}

func LoadRepository(ctx context.Context, dep *registry.AppDependency, prov *registry.AppProvider, loaders ...LoaderRepositoryFn) *registry.AppRepository {
	var repo registry.AppRepository

	if loaders == nil {
		return &repo
	}

	for _, l := range loaders {
		l.Loader(ctx, dep, prov, &repo)
	}

	return &repo
}
