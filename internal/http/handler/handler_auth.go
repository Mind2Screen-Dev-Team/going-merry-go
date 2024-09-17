package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/http/dto"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/httputil"
)

type HandlerAuth struct{}

func (HandlerAuth) Login(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	input := httputil.LoadInput[dto.AuthLoginReqDTO](ctx)

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(map[string]any{
		"status":  true,
		"message": "success",
		"data":    input,
	})
}
