package ctxkey

const (
	REGISTRY_APP_CONFIG     = CtxKey("REGISTRY:APP_CONFIG")
	REGISTRY_APP_DEPENDENCY = CtxKey("REGISTRY:APP_DEPENDENCY")
	REGISTRY_APP_REPOSITORY = CtxKey("REGISTRY:APP_REPOSITORY")
	REGISTRY_APP_SERVICE    = CtxKey("REGISTRY:APP_SERVICE")
	REGISTRY_APP_PROVIDER   = CtxKey("REGISTRY:APP_PROVIDER")
	REGISTRY_APP_LOGGER     = CtxKey("REGISTRY:APP_LOGGER")
	REGISTRY_APP_TRACER     = CtxKey("REGISTRY:APP_TRACER")
)