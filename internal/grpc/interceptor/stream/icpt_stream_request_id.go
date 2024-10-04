package icpt_stream

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/ctxkey"
	icpt_util "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/grpc/interceptor/util"

	"github.com/rs/xid"
	"google.golang.org/grpc"
)

func RequestIDInterceptor() grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		// Call the handler to complete the normal execution of the RPC.
		return handler(srv, &icpt_util.WrappedStream{ServerStream: ss, Ctx: context.WithValue(ss.Context(), ctxkey.RequestIDKey, xid.New().String())})
	}
}
