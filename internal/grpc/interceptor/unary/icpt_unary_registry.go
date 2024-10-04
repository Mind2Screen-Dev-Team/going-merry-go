package icpt_unary

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/ctxkey"
	"google.golang.org/grpc"
)

func RegisterRegistry(reg *registry.AppRegistry) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

		// # Assign a Value To Context
		ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP, reg)
		ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP_CONFIG, reg.Config)
		ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP_DEPENDENCY, reg.Dependency)
		ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP_REPOSITORY, reg.Repository)
		ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP_PROVIDER, reg.Provider)
		ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP_SERVICE, reg.Service)
		ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP_LOGGER, &reg.Dependency.ZeroLogger)
		ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP_TRACER, reg.Dependency.Tracer)

		return handler(ctx, req)
	}
}
