package bootstrap

import (
	"context"

	repo_api "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/repo/api"
)

// # REPOSITORY

type AppRepository struct {
	User repo_api.UserRepoAPI
	// register your repository on here
}

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
