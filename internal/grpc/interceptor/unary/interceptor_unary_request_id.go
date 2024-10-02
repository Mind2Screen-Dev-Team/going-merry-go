package interceptor_unary

import (
	"context"

	"github.com/rs/xid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func RequestIDInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		var (
			inMd, _ = metadata.FromIncomingContext(ctx)
			md      = metadata.Join(inMd, metadata.New(map[string]string{"traceId": xid.New().String()}))
		)

		resp, err = handler(
			metadata.NewIncomingContext(ctx, md),
			req,
		)

		return resp, err
	}
}
