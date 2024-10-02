package app

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/grpc/handler"

	"google.golang.org/grpc"
)

func AppGRPCHandlerLoader(server *grpc.Server, reg *registry.AppRegistry) {
	// # Load All Provider
	bootstrap.LoadGRPCHandler(
		context.Background(),

		// Link Dependency
		server,
		reg,

		// # List Of GRPC Handler
		handler.NewHandlerHealth(),
		handler.NewHandlerGreating(),

		// add more on here...
	)
}
