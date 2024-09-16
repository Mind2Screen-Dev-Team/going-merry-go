package bootstrap

import (
	"context"

	service_api "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/service/api"
)

// # SERVICE

type AppService struct {
	User service_api.UserServiceAPI
	// register your service on here
}

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
