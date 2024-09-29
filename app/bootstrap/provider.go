package bootstrap

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
)

type LoaderProviderFn interface {
	Loader(ctx context.Context, dep *registry.AppDependency, repo *registry.AppProvider)
}

func LoadProvider(ctx context.Context, dep *registry.AppDependency, loaders ...LoaderProviderFn) *registry.AppProvider {
	var repo registry.AppProvider

	if loaders == nil {
		return &repo
	}

	for _, l := range loaders {
		l.Loader(ctx, dep, &repo)
	}

	return &repo
}
