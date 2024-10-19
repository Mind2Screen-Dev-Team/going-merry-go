package xfilter

func conv[T any](in []T) []any {
	args := make([]any, len(in))
	for i, v := range in {
		args[i] = v
	}
	return args
}

func configChecker(cfg []Config, field string, f Filter) (column string, xtype string, ok bool) {
	for _, c := range cfg {

		if c.Disabled || field != c.Field || f.Type != c.Type {
			continue
		}

		column = c.Column
		xtype = c.Type
	}

	if len(column) <= 0 || len(xtype) <= 0 {
		return "", "", false
	}

	return column, xtype, true
}
