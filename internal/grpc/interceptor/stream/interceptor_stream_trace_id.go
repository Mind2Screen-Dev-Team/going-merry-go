package interceptor_stream

import (
	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/grpc/interceptor/util"
	"github.com/rs/xid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TraceIDInterceptor() grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		md, _ := metadata.FromIncomingContext(ss.Context())

		// Create and set metadata from interceptor to server.
		md.Append("traceId", xid.New().String())
		ctx := metadata.NewIncomingContext(ss.Context(), md)

		// Call the handler to complete the normal execution of the RPC.
		return handler(srv, &util.WrappedStream{ServerStream: ss, Ctx: ctx})
	}
}
