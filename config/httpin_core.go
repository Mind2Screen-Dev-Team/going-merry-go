package config

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/restkey"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xhttpin"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xlogger"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xresponse"

	"github.com/ggicci/httpin/core"
	"github.com/ggicci/httpin/integration"
	"github.com/go-chi/chi/v5"
)

type httpinCore struct{}

func NewHttpinCore() *httpinCore {
	return &httpinCore{}
}

func (httpinCore) Loader(ctx context.Context, reg *registry.AppRegistry) {
	// # Go-Chi URL Param integrations
	integration.UseGochiURLParam("path", chi.URLParam)

	// # Register Error Handler
	core.RegisterErrorHandler(func(rw http.ResponseWriter, r *http.Request, err error) {
		logger := xlogger.FromReqCtx(r.Context())
		// status: 422
		resp := xresponse.NewRestResponse[any, map[string]any](r, rw)
		var invalidFieldError *core.InvalidFieldError
		if errors.As(err, &invalidFieldError) {
			logger.Error("invalid format, parse request dto", "error", err)
			resp.StatusCode(http.StatusUnprocessableEntity).Code(restkey.INVALID_ARGUMENT).Msg("failed parse request dto")
			if invalidFieldError.Directive != "body" {
				errs := map[string]any{
					invalidFieldError.Field: invalidFieldError.ErrorMessage,
				}

				resp.Error(errs).JSON()
				return
			}

			resp.JSON()
			return
		}

		// status: 500
		logger.Error("internal server error, parse request dto", "error", err)
		resp.StatusCode(http.StatusInternalServerError).Code(restkey.INTERNAL).Msg("internal server error, failed parse request dto").JSON()
	})

	// # Register Named Coder
	core.RegisterNamedCoder("any_time", func(t *time.Time) (core.Stringable, error) {
		return (*xhttpin.AnyTime)(t), nil
	})
}
