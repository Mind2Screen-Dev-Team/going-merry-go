package service

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/grpc/greating"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xtracer"
	"google.golang.org/grpc"
)

type GrpcServiceGreating struct {
	greating.UnimplementedGreatingServiceServer
}

func NewGrpcServiceGreating() *GrpcServiceGreating {
	return &GrpcServiceGreating{}
}

func (h *GrpcServiceGreating) Loader(ctx context.Context, server *grpc.Server, reg *registry.AppRegistry) error {
	// # Register greating service server
	greating.RegisterGreatingServiceServer(server, h)
	// # add implemenation here...
	return nil
}

func (h *GrpcServiceGreating) Say(ctx context.Context, r *greating.GreatingRequest) (*greating.GreatingResponse, error) {
	ctx, span := xtracer.Start(ctx, "grpc.handler.greating.say")
	defer span.End()

	return &greating.GreatingResponse{Msg: r.Msg}, nil
}
