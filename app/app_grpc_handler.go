package app

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/grpc/handler"
	"google.golang.org/grpc"
)

func AppGRPCHandlerLoader(cfg *appconfig.AppConfig, server *grpc.Server, dep *registry.AppDependency, repo *registry.AppRepository, prov *registry.AppProvider, serv *registry.AppService) {
	// # Load All Provider
	bootstrap.LoadGRPCHandler(
		context.Background(),

		// Link Dependency
		server,
		cfg,
		dep,
		repo,
		prov,
		serv,

		// # List Of GRPC Handler
		handler.NewHealthHandler(),

		// add more on here...
	)
}
