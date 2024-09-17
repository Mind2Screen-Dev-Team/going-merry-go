package bootstrap

import (
	"context"
)

type LoaderServiceFn interface {
	Loader(ctx context.Context, dep *AppDependency, repo *AppRepository, serv *AppService)
}

func LoadService(ctx context.Context, dep *AppDependency, repo *AppRepository, loaders ...LoaderServiceFn) *AppService {
	var serv AppService

	if loaders == nil {
		return &serv
	}

	for _, l := range loaders {
		l.Loader(ctx, dep, repo, &serv)
	}

	return &serv
}
