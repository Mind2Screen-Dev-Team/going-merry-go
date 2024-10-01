package handler

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/grpc/health"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
	"google.golang.org/grpc"
)

type HandlerHealth struct {
	health.UnimplementedHealthServiceServer
}

func NewHandlerHealth() *HandlerHealth {
	return &HandlerHealth{}
}

func (h *HandlerHealth) Loader(ctx context.Context, server *grpc.Server, cfg *appconfig.AppConfig, dep *registry.AppDependency, repo *registry.AppRepository, prov *registry.AppProvider, serv *registry.AppService) {
	// # Register health service server
	health.RegisterHealthServiceServer(server, h)

	// # add implemenation here...
}

func (h *HandlerHealth) Check(context.Context, *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}, nil
}
