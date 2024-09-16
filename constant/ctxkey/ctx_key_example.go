package ctxkey

type ExampleCtx string

func (h ExampleCtx) Value() string {
	return string(h)
}
