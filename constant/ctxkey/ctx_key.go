package ctxkey

type CtxKey string

func (c CtxKey) String() string {
	return string(c)
}
