package bootstrap

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"google.golang.org/grpc"
)

type LoaderGRPCHandlerFn interface {
	Loader(ctx context.Context, server *grpc.Server, reg *registry.AppRegistry)
}

func LoadGRPCHandler(ctx context.Context, server *grpc.Server, reg *registry.AppRegistry, loaders ...LoaderGRPCHandlerFn) {
	if loaders == nil {
		return
	}

	for _, l := range loaders {
		l.Loader(ctx, server, reg)
	}
}
