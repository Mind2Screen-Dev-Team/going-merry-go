package handler

import (
	"net/http"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/config"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/restkey"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/http/dto"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/http/interceptor"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xhttputil"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xresponse"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xvalidate"
)

type HandlerAuth struct {
	interceptor.ExampleInterceptor
}

func (h HandlerAuth) Login(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := xhttputil.LoadInput[dto.AuthLoginReqDTO](ctx)

	// # Example Basic Response Builder
	resp := xresponse.NewRestResponse[any, any](rw)

	// # Example Response Builder With Interceptor:
	// resp := xresponse.NewRestResponseWithInterceptor(
	// 	rw,
	// 	r,
	// 	h.ExampleInterceptor,
	// )

	if err := data.ValidateWithContext(ctx); err != nil {
		if errs, ok := err.(xvalidate.Errors); ok {
			resp.StatusCode(http.StatusUnprocessableEntity).Code(restkey.INVALID_ARGUMENT).Error(errs).Msg("invalid validation request data").JSON()
			return
		}
		config.Logger(ctx).Error().Err(err)
		resp.StatusCode(http.StatusInternalServerError).Code(restkey.INTERNAL).Error(err.Error()).Msg("internal server error").JSON()
		return
	}

	resp.Code(restkey.SUCCESS).Msg("auth login success").Data(data).JSON()
}

// 1. logger must has service name, host+port, trace id, message, other fields
// 2. that default fields send to logger
// 3. controller just send a message and additional fields
