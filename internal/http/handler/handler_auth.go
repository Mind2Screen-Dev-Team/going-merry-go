package handler

import (
	"net/http"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/http/dto"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/http/interceptor"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xhttputil"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xresponse"
)

type HandlerAuth struct {
	interceptor.ExampleInterceptor
}

func (h HandlerAuth) Login(rw http.ResponseWriter, r *http.Request) {
	data := xhttputil.LoadInput[dto.AuthLoginReqDTO](r.Context())

	// # Example Basic Response Builder
	resp := xresponse.NewRestResponse[map[string]any, any](rw)

	// # Example Response Builder With Interceptor:
	// resp := xresponse.NewRestResponseWithInterceptor(
	// 	rw,
	// 	r,
	// 	h.ExampleInterceptor,
	// )

	resp.Code("SUCCESS").Msg("auth login success").Data(map[string]any{"input": data}).JSON()
}
