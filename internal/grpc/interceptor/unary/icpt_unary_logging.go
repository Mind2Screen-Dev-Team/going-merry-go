package icpt_unary

import (
	icpt_util "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/grpc/interceptor/util"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

func Logging(l zerolog.Logger) grpc.UnaryServerInterceptor {
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		// Add any other option (check functions starting with logging.With).
	}
	return logging.UnaryServerInterceptor(icpt_util.Logger(l), opts...)
}
