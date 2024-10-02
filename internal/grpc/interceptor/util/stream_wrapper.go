package util

import (
	"context"

	"google.golang.org/grpc"
)

type WrappedStream struct {
	grpc.ServerStream
	Ctx context.Context
}

func (s *WrappedStream) Context() context.Context {
	return s.Ctx
}
