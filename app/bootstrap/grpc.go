package bootstrap

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"google.golang.org/grpc"
)

type LoaderGrpcFunc interface {
	Loader(ctx context.Context, server *grpc.Server, reg *registry.AppRegistry) error
}

func LoadGrpcService(ctx context.Context, server *grpc.Server, reg *registry.AppRegistry, loaders ...LoaderGrpcFunc) error {
	if loaders == nil {
		return nil
	}

	for _, l := range loaders {
		if err := l.Loader(ctx, server, reg); err != nil {
			return err
		}
	}

	return nil
}
