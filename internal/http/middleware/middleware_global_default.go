package middleware

import (
	"bytes"
	"context"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/ctxkey"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/restkey"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xhttputil"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xlogger"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xresponse"

	"github.com/DataDog/gostackparse"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/xid"
)

func DefaultGlobal(req *registry.AppRegistry, r chi.Router) {
	r.Use(middleware.RealIP)
	r.Use(RequestID)
	r.Use(Logger)
	r.Use(middleware.Timeout(
		time.Duration(req.Config.Http.HandlerTimeout) * time.Second,
	))
	r.Use(middleware.Heartbeat("/health"))
}

func RequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, ctxkey.RequestIDKey, xid.New().String())
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var (
			reqtime = time.Now().UnixMilli()
			rw      = bytes.Buffer{}
		)

		var (
			rCopy = xhttputil.DeepCopyRequest(r, true)
			ww    = middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			resp  = xresponse.NewRestResponse[any, any](r, ww)
		)

		defer func() {
			var (
				panicked     bool
				recorverErr  any
				parsedStacks = make([]*gostackparse.Goroutine, 0)
			)

			if r := recover(); r != nil && r != http.ErrAbortHandler {
				// assign
				panicked = true
				recorverErr = r

				stacks := debug.Stack()
				parsedStacks, _ = gostackparse.Parse(bytes.NewReader(stacks))

				resp.StatusCode(http.StatusInternalServerError).Code(restkey.INTERNAL).Msg("internal server error").JSON()
			}

			go IncomingLogging(rCopy, ww, &rw, reqtime, panicked, recorverErr, parsedStacks)
		}()

		ww.Tee(&rw)
		next.ServeHTTP(ww, r)
	}

	return http.HandlerFunc(fn)
}

func IncomingLogging(
	// general params
	r *http.Request,
	ww middleware.WrapResponseWriter,
	resbody *bytes.Buffer,
	reqtime int64,

	// panic params
	panicked bool,
	recorverErr any,
	stacks []*gostackparse.Goroutine,
) {
	var (
		ctx       = r.Context()
		logger    = xlogger.FromReqCtx(ctx)
		requestId = ctx.Value(ctxkey.RequestIDKey)
		fields    = []any{
			"request.id", requestId,
			"request.time", reqtime,
			"request.remote.address", r.RemoteAddr,
			"request.path", r.URL.Path,
			"request.proto", r.Proto,
			"request.method", r.Method,
			"request.user.agent", r.UserAgent(),
			"request.body", ctx.Value(xhttputil.STR_REQ_BODY),
			"response.status", http.StatusText(ww.Status()),
			"response.status.code", ww.Status(),
			"response.bytes.in", r.ContentLength,
			"response.bytes.out", ww.BytesWritten(),
			"response.body", resbody.String(),
		}
	)

	defer resbody.Reset()

	if panicked {
		fields = append([]any{"panic.recover", recorverErr, "panic.stack", stacks}, fields...)
	}

	logger.Info(
		// msg
		"incoming request",

		// fields
		fields...,
	)
}
