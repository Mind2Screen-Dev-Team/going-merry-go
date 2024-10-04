package app

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/grpc/service"

	"google.golang.org/grpc"
)

func AppGrpcServiceLoader(server *grpc.Server, reg *registry.AppRegistry) error {
	// # Load All Grpc Service
	return bootstrap.LoadGrpcService(
		context.Background(),

		// Link Dependency
		server,
		reg,

		// # List Of GRPC Handler
		service.NewGrpcServiceHealth(),
		service.NewGrpcServiceGreating(),

		// add more on here...
	)
}
