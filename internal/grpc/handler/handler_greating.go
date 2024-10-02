package handler

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/grpc/greating"
	"google.golang.org/grpc"
)

type HandlerGreating struct {
	greating.UnimplementedGreatingServiceServer
}

func NewHandlerGreating() *HandlerGreating {
	return &HandlerGreating{}
}

func (h *HandlerGreating) Loader(ctx context.Context, server *grpc.Server, reg *registry.AppRegistry) {
	// # Register greating service server
	greating.RegisterGreatingServiceServer(server, h)

	// # add implemenation here...
}

func (h *HandlerGreating) Say(ctx context.Context, r *greating.GreatingRequest) (*greating.GreatingResponse, error) {
	return &greating.GreatingResponse{
		Msg: r.Msg,
	}, nil
}
