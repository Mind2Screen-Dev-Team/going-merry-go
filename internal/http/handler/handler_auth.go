package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/ctxkey"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/http/dto"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xhttputil"
)

type HandlerAuth struct{}

func (HandlerAuth) Login(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	input := xhttputil.LoadInput[dto.AuthLoginReqDTO](ctx)

	cfg, ok := ctx.Value(ctxkey.CTX_KEY_HTTP_SERVER_APP_CONFIG).(*appconfig.AppConfig)

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(map[string]any{
		"status":  true,
		"message": "success",
		"data":    input,
		"ok":      ok,
		"cfg":     cfg,
	})
}
