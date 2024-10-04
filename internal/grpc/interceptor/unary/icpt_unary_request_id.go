package icpt_unary

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/ctxkey"

	"github.com/rs/xid"
	"google.golang.org/grpc"
)

func RequestIDInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		resp, err = handler(
			context.WithValue(
				ctx,
				ctxkey.RequestIDKey,
				xid.New().String(),
			),
			req,
		)

		return resp, err
	}
}
