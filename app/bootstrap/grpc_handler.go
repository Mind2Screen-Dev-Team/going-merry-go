package bootstrap

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
	"google.golang.org/grpc"
)

type LoaderGRPCHandlerFn interface {
	Loader(ctx context.Context, server *grpc.Server, cfg *appconfig.AppConfig, dep *registry.AppDependency, repo *registry.AppRepository, prov *registry.AppProvider, serv *registry.AppService)
}

func LoadGRPCHandler(ctx context.Context, server *grpc.Server, cfg *appconfig.AppConfig, dep *registry.AppDependency, repo *registry.AppRepository, prov *registry.AppProvider, serv *registry.AppService, loaders ...LoaderGRPCHandlerFn) {
	if loaders == nil {
		return
	}

	for _, l := range loaders {
		l.Loader(ctx, server, cfg, dep, repo, prov, serv)
	}
}
