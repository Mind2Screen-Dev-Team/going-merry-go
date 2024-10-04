package handler

import (
	"net/http"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/restkey"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/http/dto"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/http/interceptor"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xhttputil"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xlogger"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xresponse"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xtracer"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xvalidate"
	"go.opentelemetry.io/otel/trace"
)

type HandlerAuth struct {
	interceptor.ExampleInterceptor
}

func (h HandlerAuth) Login(rw http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		span trace.Span
	)

	ctx, span = xtracer.Start(ctx, "login.handler.process")
	defer span.End()

	data := xhttputil.LoadInput[dto.AuthLoginReqDTO](ctx)

	// # Example Basic Response Builder
	resp := xresponse.NewRestResponse[any, any](r, rw)

	// # Example Response Builder With Interceptor:
	// resp := xresponse.NewRestResponseWithInterceptor(
	// 	rw,
	// 	r,
	// 	h.ExampleInterceptor,
	// )

	if err := data.ValidateWithContext(ctx); err != nil {
		span.RecordError(err)
		if errs, ok := xvalidate.IsErrors(err); ok {
			resp.StatusCode(http.StatusUnprocessableEntity).Code(restkey.INVALID_ARGUMENT).Error(errs).Msg("invalid validation request data").JSON()
			return
		}

		xlogger.FromReqCtx(ctx).Error("validation internal server error", "error", err)
		resp.StatusCode(http.StatusInternalServerError).Code(restkey.INTERNAL).Msg("internal server error").JSON()
		return
	}

	resp.Code(restkey.SUCCESS).Msg("auth login success").Data(data).JSON()
}
