package bootstrap

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
)

type LoaderServiceFn interface {
	Loader(ctx context.Context, dep *registry.AppDependency, repo *registry.AppRepository, serv *registry.AppService)
}

func LoadService(ctx context.Context, dep *registry.AppDependency, repo *registry.AppRepository, loaders ...LoaderServiceFn) *registry.AppService {
	var serv registry.AppService

	if loaders == nil {
		return &serv
	}

	for _, l := range loaders {
		l.Loader(ctx, dep, repo, &serv)
	}

	return &serv
}
