package icpt_stream

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/ctxkey"
	icpt_util "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/grpc/interceptor/util"
	"google.golang.org/grpc"
)

func RegisterRegistry(reg *registry.AppRegistry) grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		// Call the handler to complete the normal execution of the RPC.
		ctx := ss.Context()

		// # Assign a Value To Context
		ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP_CONFIG, reg.Config)
		ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP_DEPENDENCY, reg.Dependency)
		ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP_REPOSITORY, reg.Repository)
		ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP_PROVIDER, reg.Provider)
		ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP_SERVICE, reg.Service)
		ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP_LOGGER, &reg.Dependency.ZeroLogger)
		ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP_TRACER, reg.Dependency.Tracer)

		return handler(srv, &icpt_util.WrappedStream{ServerStream: ss, Ctx: ctx})
	}
}
