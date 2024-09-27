package config

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/restkey"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xhttpin"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xresponse"

	"github.com/ggicci/httpin/core"
	"github.com/ggicci/httpin/integration"
	"github.com/go-chi/chi/v5"
)

type httpinCore struct{}

func NewHttpinCore() *httpinCore {
	return &httpinCore{}
}

func (httpinCore) Loader(ctx context.Context, cfg *appconfig.AppConfig, app *registry.AppDependency) {
	// # Go-Chi URL Param integrations
	integration.UseGochiURLParam("path", chi.URLParam)

	// # Register Error Handler
	core.RegisterErrorHandler(func(rw http.ResponseWriter, r *http.Request, err error) {
		// status: 422
		resp := xresponse.NewRestResponse[any, map[string]any](rw)
		var invalidFieldError *core.InvalidFieldError
		if errors.As(err, &invalidFieldError) {
			errs := map[string]any{
				invalidFieldError.Field: invalidFieldError.ErrorMessage,
			}
			resp.StatusCode(http.StatusUnprocessableEntity).Code(restkey.INVALID_ARGUMENT).Error(errs).JSON()
			return
		}

		// status: 500
		resp.StatusCode(http.StatusInternalServerError).Code(restkey.INTERNAL).Msg("failed to parse a request dto").JSON()
	})

	// # Register Named Coder
	core.RegisterNamedCoder("any_time", func(t *time.Time) (core.Stringable, error) {
		return (*xhttpin.AnyTime)(t), nil
	})
}
