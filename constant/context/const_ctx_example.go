package const_ctx

type ExampleCtx string

func (h ExampleCtx) Value() string {
	return string(h)
}
