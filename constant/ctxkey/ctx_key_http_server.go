package ctxkey

type CtxKeyHttpServer string

func (c CtxKeyHttpServer) String() string {
	return string(c)
}

const (
	CTX_KEY_HTTP_SERVER_APP_CONFIG = CtxKeyHttpServer("CTX_KEY_HTTP_SERVER_APP_CONFIG")
)
