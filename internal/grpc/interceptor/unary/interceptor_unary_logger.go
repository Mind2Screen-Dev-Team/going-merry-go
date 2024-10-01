package unary

import (
	"bytes"
	"context"
	"net/http"
	"runtime/debug"

	"github.com/DataDog/gostackparse"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/ctxkey"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xlogger"

	"google.golang.org/grpc"
)

type Logger struct {
	cfg  *appconfig.AppConfig
	dep  *registry.AppDependency
	repo *registry.AppRepository
	prov *registry.AppProvider
	serv *registry.AppService
}

func (l *Logger) Loader(
	cfg *appconfig.AppConfig,
	dep *registry.AppDependency,
	repo *registry.AppRepository,
	prov *registry.AppProvider,
	serv *registry.AppService,
) {
	l.cfg = cfg
	l.dep = dep
	l.repo = repo
	l.prov = prov
	l.serv = serv
}

func (l *Logger) Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		logger := xlogger.NewZeroLogger(&l.dep.ZeroLogger)
		traceId := ctx.Value(ctxkey.RequestIDKey)
		defer func() {
			isPanic := false
			if r := recover(); r != nil && r != http.ErrAbortHandler {
				stacks := debug.Stack()
				parsed, _ := gostackparse.Parse(bytes.NewReader(stacks))
				logger.Error(
					// msg
					"incoming request panic",
					"traceId", traceId,

					// fields
					"recover", r,
					"stack", parsed,
				)
				isPanic = true
			}

			logger.Info("incoming request", "traceId", traceId, "req", req, "info", info)
			if isPanic {
			}
		}()

		resp, err = handler(ctx, req)
		return resp, err
	}
}
