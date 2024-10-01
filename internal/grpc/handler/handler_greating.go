package handler

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/grpc/greating"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
	"google.golang.org/grpc"
)

type HandlerGreating struct {
	greating.UnimplementedGreatingServiceServer
}

func NewHandlerGreating() *HandlerGreating {
	return &HandlerGreating{}
}

func (h *HandlerGreating) Loader(ctx context.Context, server *grpc.Server, cfg *appconfig.AppConfig, dep *registry.AppDependency, repo *registry.AppRepository, prov *registry.AppProvider, serv *registry.AppService) {
	// # Register greating service server
	greating.RegisterGreatingServiceServer(server, h)

	// # add implemenation here...
}

func (h *HandlerGreating) Say(ctx context.Context, req *greating.GreatingRequest) (*greating.GreatingResponse, error) {
	return &greating.GreatingResponse{
		Msg: req.Msg,
	}, nil
}
