package service

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/grpc/health"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xtracer"
	"google.golang.org/grpc"
)

type GrpcServiceHealth struct {
	health.UnimplementedHealthServiceServer
}

func NewGrpcServiceHealth() *GrpcServiceHealth {
	return &GrpcServiceHealth{}
}

func (h *GrpcServiceHealth) Loader(ctx context.Context, server *grpc.Server, reg *registry.AppRegistry) error {
	// # Register health service server
	health.RegisterHealthServiceServer(server, h)
	// # add implemenation here...
	return nil
}

func (h *GrpcServiceHealth) Check(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	ctx, span := xtracer.Start(ctx, "grpc.handler.health.check")
	defer span.End()

	return &health.HealthCheckResponse{Status: health.HealthCheckResponse_SERVING}, nil
}
