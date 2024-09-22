package xresponse

import (
	"context"
	"net/http"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/ctxkey"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/restkey"
	"github.com/ggicci/httpin"
)

// A Wrapper HTTP Rest API Response Builder Initiator
func NewRestResponse[D any, E any](rw http.ResponseWriter) RestResponseSTD[D, E] {
	return &restResponseSTD[D, E]{
		ResponseSTD: ResponseSTD[D, E]{
			responseWriter: rw,
		},
	}
}

// A Wrapper HTTP Rest API Response Builder Initiator With Interceptor, i.e:
//
//	// Interceptor
//	type SomeInterceptHandler struct{}
//
//	func NewSomeInterceptHandler() SomeInterceptHandler {
//		return SomeInterceptHandler{}
//	}
//
//	func (SomeInterceptHandler) SomeInterceptorHandler(req *http.Request, res RestResponseValue[dto.SomeResDTO, dto.SomeResErrorDTO]) {
//		// do something on here
//	}
//
//	// Controller / Handler
//	func SomeHandler(rw http.ResponseWriter, r *http.Request) {
//		ctx := r.Context()
//		xRes := xresponse.NewRestResponseWithInterceptor(rw, r, interceptor.NewSomeInterceptHandler())
//
//		xReqDto := xhttputil.LoadInput[dto.SomeReqDTO](ctx)
//		if err := xReqDto.ValidateWithContext(ctx); err != nil {
//			return xRes.StatusCode(http.StatusUnprocessableEntity).Code(restkey.INVALID_ARGUMENT).Msg("invalid request data").JSON()
//		}
//
//		// continue your bussines logic ...
//	}
func NewRestResponseWithInterceptor[D any, E any](rw http.ResponseWriter, r *http.Request, handler InterceptHandler[D, E], ctxKeys ...ctxkey.CtxKey) RestResponseSTD[D, E] {
	requestCtx := r.Context()

	nextCtx := context.Background()
	nextCtx = context.WithValue(nextCtx, httpin.Input, requestCtx.Value(httpin.Input))
	nextCtx = copyCtxValue(nextCtx, requestCtx, ctxkey.HTTP_SERVER_APP_CONFIG)
	nextCtx = copyCtxValue(nextCtx, requestCtx, ctxkey.HTTP_SERVER_APP_DEPENDENCY)
	nextCtx = copyCtxValue(nextCtx, requestCtx, ctxkey.HTTP_SERVER_APP_REPOSITORY)
	nextCtx = copyCtxValue(nextCtx, requestCtx, ctxkey.HTTP_SERVER_APP_SERVICE)

	for _, ctxKey := range ctxKeys {
		nextCtx = copyCtxValue(nextCtx, requestCtx, ctxKey)
	}

	return &restResponseSTD[D, E]{
		ResponseSTD: ResponseSTD[D, E]{
			responseWriter:   rw,
			request:          r,
			interceptHandler: handler,
		},
	}
}

// A Async Function For Interceptor HTTP REST API
type InterceptHandler[D any, E any] interface {
	Handler(req *http.Request, res RestResponseValue[D, E])
}

// A Wrapper Standard Response For HTTP REST API
//
// HTTP Rest API Response Getter For Interceptor
type RestResponseValue[D any, E any] interface {
	GetMsg() string
	GetCode() restkey.RestKey
	GetData() D
	GetError() E
	GetStatusCode() int
	GetResponseHeader() http.Header
	JSONText() (string, error)
}

// A Wrapper Standard Response For HTTP REST API
//
// HTTP Rest API Response Builder
type RestResponseSTD[D any, E any] interface {
	Msg(msg string) RestResponseSTD[D, E]
	Code(code restkey.RestKey) RestResponseSTD[D, E]
	Data(data D) RestResponseSTD[D, E]
	Error(err E) RestResponseSTD[D, E]

	// Setter HTTP Response Status Code
	//
	// HTTP status codes as registered with IANA.
	// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
	StatusCode(code int) RestResponseSTD[D, E]

	// Add HTTP Response Header, this is equal function with:
	//	responseWriter.Header().Add(key, value)
	AddHeader(key string, value string) RestResponseSTD[D, E]

	// Delete HTTP Response Header by Key, this is equal function with:
	//	responseWriter.Header().Del(key)
	DelHeader(key string) RestResponseSTD[D, E]

	// A Method that Usefull for Run Interceptor Only,
	//
	// When you not need a response api that only need run interceptor.
	Done()

	// A JSON Response Encoder for HTTP Response Writer, this is also auto set header and status code
	//	responseWriter.Header().Add("Accept", "application/json")
	//	responseWriter.Header().Add("Content-Type", "application/json")
	//
	//	// Default httpStatusCode is 200
	//	responseWriter.WriteHeader(httpStatusCode)
	JSON()

	// A JSON Response Encoder for HTTP Response Writer With Error, this is also auto set header and status code
	//	responseWriter.Header().Add("Accept", "application/json")
	//	responseWriter.Header().Add("Content-Type", "application/json")
	//
	//	// Default httpStatusCode is 200
	//	responseWriter.WriteHeader(httpStatusCode)
	JSONOrErr() error

	// A General Purpose JSON Response Encoder to text
	JSONText() (string, error)
}

type restResponseSTD[D any, E any] struct {
	ResponseSTD[D, E]
}

// # GETTER

func (r *restResponseSTD[D, E]) GetMsg() string {
	return r.ResponseSTD.Msg
}

func (r *restResponseSTD[D, E]) GetCode() restkey.RestKey {
	code, ok := r.ResponseSTD.Code.(restkey.RestKey)
	if !ok {
		return restkey.UNKNOWN
	}
	return code
}

func (r *restResponseSTD[D, E]) GetData() D {
	return r.ResponseSTD.Data
}

func (r *restResponseSTD[D, E]) GetError() E {
	return r.ResponseSTD.Err
}

func (r *restResponseSTD[D, E]) GetStatusCode() int {
	return r.ResponseSTD.statusCode
}

func (r *restResponseSTD[D, E]) GetResponseHeader() http.Header {
	return r.ResponseSTD.responseHeader
}

// # SETTER

func (r *restResponseSTD[D, E]) Msg(msg string) RestResponseSTD[D, E] {
	r.ResponseSTD.SetMsg(msg)
	return r
}

func (r *restResponseSTD[D, E]) Code(code restkey.RestKey) RestResponseSTD[D, E] {
	r.ResponseSTD.SetCode(code)
	return r
}

func (r *restResponseSTD[D, E]) Data(data D) RestResponseSTD[D, E] {
	r.ResponseSTD.SetData(data)
	return r
}

func (r *restResponseSTD[D, E]) Error(err E) RestResponseSTD[D, E] {
	r.ResponseSTD.SetError(err)
	return r
}

// # HTTP SETTER

func (r *restResponseSTD[D, E]) StatusCode(code int) RestResponseSTD[D, E] {
	r.ResponseSTD.SetStatusCode(code)
	return r
}

func (r *restResponseSTD[D, E]) AddHeader(key string, value string) RestResponseSTD[D, E] {
	r.ResponseSTD.SetHeader(key, value)
	return r
}

func (r *restResponseSTD[D, E]) DelHeader(key string) RestResponseSTD[D, E] {
	r.ResponseSTD.DelHeader(key)
	return r
}

// # BUILDER

func (r *restResponseSTD[D, E]) Done() {
	if r.statusCode == 0 {
		r.statusCode = 200 // OK as Default
	}

	defer func() {
		// is only allow call once
		r.interceptOnceFn.Do(func() {
			if r.interceptHandler == nil {
				return
			}

			// async function
			go r.interceptHandler.Handler(r.request, r)
		})
	}()
}

func (r *restResponseSTD[D, E]) JSON() {
	if r.statusCode == 0 {
		r.statusCode = 200 // OK as Default
	}

	defer func() {
		// is only allow call once
		r.interceptOnceFn.Do(func() {
			if r.interceptHandler == nil {
				return
			}

			// async function
			go r.interceptHandler.Handler(r.request, r)
		})
	}()

	r.ResponseSTD.RestJSON()
}

func (r *restResponseSTD[D, E]) JSONOrErr() error {
	if r.statusCode == 0 {
		r.statusCode = 200 // OK as Default
	}

	defer func() {
		// is only allow call once
		r.interceptOnceFn.Do(func() {
			if r.interceptHandler == nil {
				return
			}

			// async function
			go r.interceptHandler.Handler(r.request, r)
		})
	}()

	return r.ResponseSTD.RestJSONOrErr()
}

func (r *restResponseSTD[D, E]) JSONText() (string, error) {
	return r.ResponseSTD.JSONText()
}
