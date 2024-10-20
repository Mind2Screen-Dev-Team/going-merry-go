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
			ctx       = r.Context()
			logger    = xlogger.FromReqCtx(ctx)
			requestId = ctx.Value(ctxkey.RequestIDKey)

			ww   = middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			resp = xresponse.NewRestResponse[any, any](r, ww)
		)

		defer func() {
			if r := recover(); r != nil && r != http.ErrAbortHandler {
				stacks := debug.Stack()
				parsed, _ := gostackparse.Parse(bytes.NewReader(stacks))
				logger.Error(
					// msg
					"incoming request panic",
					"requestId", requestId,

					// fields
					"recover", r,
					"stack", parsed,
				)

				resp.StatusCode(http.StatusInternalServerError).Code(restkey.INTERNAL).Msg("internal server error").JSON()
			}

			logger.Info(
				// msg
				"incoming request",

				// fields
				"requestId", requestId,
				"remoteAddr", r.RemoteAddr,
				"path", r.URL.Path,
				"proto", r.Proto,
				"method", r.Method,
				"userAgent", r.UserAgent(),
				"status", http.StatusText(ww.Status()),
				"statusCode", ww.Status(),

				"bytesIn", r.ContentLength,
				"bytesOut", ww.BytesWritten(),
			)

		}()

		next.ServeHTTP(ww, r)
	}

	return http.HandlerFunc(fn)
}
