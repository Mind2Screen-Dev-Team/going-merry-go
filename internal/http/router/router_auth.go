package router

import (
	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/http/dto"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/http/handler"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xhttputil"

	"github.com/go-chi/chi/v5"
)

type routerAuthImpl struct {
	router  chi.Router
	handler handler.HandlerAuth
}

func NewRouterAuth() *routerAuthImpl {
	return &routerAuthImpl{}
}

func (r *routerAuthImpl) Loader(router chi.Router) bootstrap.LoaderRouter {
	r.router = router
	return r
}

func (r *routerAuthImpl) Route() {
	inputOption := xhttputil.NewInputOption()
	r.router.
		With(
			xhttputil.WithInput[dto.AuthLoginReqDTO](
				// Max Memory Allocation: 15 MB
				inputOption.WithMaxMemory(15*1024*1024),
			),
		).
		Post("/api/v1/auth/login", r.handler.Login)
	// r.router.Post("/api/v1/auth/login", r.handler.Login)
}
