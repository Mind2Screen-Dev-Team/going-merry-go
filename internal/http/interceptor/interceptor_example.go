package interceptor

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/http/dto"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xhttputil"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xresponse"
)

type ExampleInterceptor struct{}

func (ExampleInterceptor) Handler(req *http.Request, res xresponse.RestResponseValue[map[string]any, any]) {
	data := xhttputil.LoadInput[dto.AuthLoginReqDTO](req.Context())
	dump, err := httputil.DumpRequest(req, false)
	if err != nil {
		log.Printf("Example Interceptor Handler Dump Req, got err: %+v", err)
		return
	}

	body, _ := json.Marshal(data)
	resp, _ := json.Marshal(res)

	log.Printf("Example Interceptor Handler Dump Req: \n%s\nBody: \n%s\n\nResp: \n%s\n", string(dump), string(body), string(resp))
}
