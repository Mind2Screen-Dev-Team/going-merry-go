package unary

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/ctxkey"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"

	"github.com/rs/xid"
	"google.golang.org/grpc"
)

type RequestID struct {
	cfg  *appconfig.AppConfig
	dep  *registry.AppDependency
	repo *registry.AppRepository
	prov *registry.AppProvider
	serv *registry.AppService
}

func (l *RequestID) Loader(
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

func (l *RequestID) Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		resp, err = handler(context.WithValue(ctx, ctxkey.RequestIDKey, xid.New().String()), req)
		return resp, err
	}
}
