package ctxkey

type CtxKeyHttpServer string

func (c CtxKeyHttpServer) String() string {
	return string(c)
}

const (
	CTX_KEY_HTTP_SERVER_APP_CONFIG     = CtxKeyHttpServer("CTX_KEY_HTTP_SERVER_APP_CONFIG")
	CTX_KEY_HTTP_SERVER_APP_DEPENDENCY = CtxKeyHttpServer("CTX_KEY_HTTP_SERVER_APP_DEPENDENCY")
	CTX_KEY_HTTP_SERVER_APP_REPOSITORY = CtxKeyHttpServer("CTX_KEY_HTTP_SERVER_APP_REPOSITORY")
	CTX_KEY_HTTP_SERVER_APP_SERVICE    = CtxKeyHttpServer("CTX_KEY_HTTP_SERVER_APP_SERVICE")
)
