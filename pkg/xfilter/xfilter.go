package xfilter

import (
	"slices"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

// Filter

type Filter struct {
	Type      string   `json:"type"`
	Operation string   `json:"operation"`
	Values    []string `json:"values"`
	Disabled  bool     `json:"disabled"`
}

// Configuration

const (
	Text    = "text"
	Number  = "number"
	Select  = "select"
	Boolean = "boolean"
	Date    = "date"
)

var (
	TextOperation = []string{
		"is",
		"is_not",
		"contains",
		"does_not_contain",
		"star_with",
		"end_with",
		"is_empty",
		"is_not_empty",
	}

	NumberOperation = []string{
		"is",
		"is_not",
		"is_equal",
		"is_not_equal",
		"is_greater_than",
		"is_less_than",
		"is_greater_than_or_equal",
		"is_less_than_or_equal",
		"is_empty",
		"is_not_empty",
	}

	SelectOperation = []string{
		"is",
		"is_not",
		"is_empty",
		"is_not_empty",
	}

	BooleanOperation = []string{
		"is",
		"is_empty",
		"is_not_empty",
	}

	DateOperation = []string{
		"is",
		"is_before",
		"is_after",
		"is_on_or_before",
		"is_on_or_after",
		"is_between",
		"is_empty",
		"is_not_empty",
	}
)

type Config struct {
	Column        string         `json:"column"`
	Label         string         `json:"label"`
	Field         string         `json:"field"`
	Type          string         `json:"type"`
	Description   string         `json:"description"`
	Suggestion    bool           `json:"suggestion"`
	Disabled      bool           `json:"disabled"`
	DefaultValues []DefaultValue `json:"default_values"`
	Operations    []string       `json:"operations"`
}

type DefaultValue struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// Builder

type BuilderFn = func(string, Filter) Builder

var (
	_FilterTypes = [...]string{
		"text",
		"number",
		"select",
		"boolean",
		"date",
	}

	_FilterOperations = [...]string{
		"is",
		"is_not",
		"contains",
		"does_not_contain",
		"star_with",
		"end_with",
		"is_equal",
		"is_not_equal",
		"is_greater_than",
		"is_less_than",
		"is_greater_than_or_equal",
		"is_less_than_or_equal",
		"is_before",
		"is_after",
		"is_on_or_before",
		"is_on_or_after",
		"is_between",
		"is_empty",
		"is_not_empty",
	}

	_FilterEmptyOperation = [...]string{
		"is_empty",
		"is_not_empty",
	}

	_FilterTypesFn = map[string]BuilderFn{
		"text":    NewBuildText,
		"number":  NewBuildNumber,
		"select":  NewBuildSelect,
		"boolean": NewBuildBool,
		"date":    NewBuildDate,
	}
)

type Builder interface {
	Build() exp.Expression
}

type Build struct {
	filterBy map[string]Filter
	configs  []Config
}

func NewBuild(filterBy map[string]Filter, configs []Config) *Build {
	return &Build{filterBy, configs}
}

func (b *Build) ToExpression() (exps []exp.Expression) {
	for k, f := range b.filterBy {
		if len(k) <= 0 {
			continue
		}

		if f.Disabled {
			continue
		}

		if !slices.Contains(_FilterTypes[:], f.Type) {
			continue
		}

		if !slices.Contains(_FilterOperations[:], f.Operation) {
			continue
		}

		if !slices.Contains(_FilterEmptyOperation[:], f.Operation) {
			if len(f.Values) <= 0 {
				continue
			}

			if slices.ContainsFunc(f.Values, func(v string) bool { return len(v) <= 0 }) {
				continue
			}

			if f.Operation == "is_between" && len(f.Values) < 2 {
				continue
			}
		}

		c, t, ok := configChecker(b.configs, k, f)
		if !ok {
			continue
		}

		fn, ok := _FilterTypesFn[t]
		if !ok {
			continue
		}

		e := fn(c, f).Build()
		exps = append(exps, e)
	}

	return
}

// TEXT BUILD IMPLEMENTATION

type BuildText struct {
	column string
	filter Filter
}

func NewBuildText(column string, filter Filter) Builder {
	return &BuildText{column, filter}
}

func (b *BuildText) Build() exp.Expression {
	var (
		l = len(b.filter.Values)
		c = goqu.L(b.column)
	)

	if l <= 0 {
		return nil
	}

	switch b.filter.Operation {
	case "is":
		return c.In(conv(b.filter.Values)...)

	case "is_not":
		return c.NotIn(conv(b.filter.Values)...)

	case "contains":
		var e = make([]exp.Expression, l)
		for i, v := range b.filter.Values {
			e[i] = c.Like("%" + v + "%")
		}

		if l == 1 {
			return e[0]
		}

		return goqu.Or(e...)

	case "does_not_contain":
		var e = make([]exp.Expression, l)
		for i, v := range b.filter.Values {
			e[i] = c.NotLike("%" + v + "%")
		}

		if l == 1 {
			return e[0]
		}

		return goqu.Or(e...)

	case "star_with":
		var e = make([]exp.Expression, l)
		for i, v := range b.filter.Values {
			e[i] = c.Like(v + "%")
		}

		if l == 1 {
			return e[0]
		}

		return goqu.Or(e...)

	case "end_with":
		var e = make([]exp.Expression, l)
		for i, v := range b.filter.Values {
			e[i] = c.Like("%" + v)
		}

		if l == 1 {
			return e[0]
		}

		return goqu.Or(e...)

	case "is_empty":
		return c.IsNull()

	case "is_not_empty":
		return c.IsNotNull()
	}

	return nil
}

// NUMBER BUILD IMPLEMENTATION

type BuildNumber struct {
	column string
	filter Filter
}

func NewBuildNumber(column string, filter Filter) Builder {
	return &BuildNumber{column, filter}
}

func (b *BuildNumber) Build() exp.Expression {
	var (
		l = len(b.filter.Values)
		c = goqu.L(b.column)
	)

	if l <= 0 {
		return nil
	}

	switch b.filter.Operation {
	case "is":
		return c.In(conv(b.filter.Values)...)

	case "is_not":
		return c.NotIn(conv(b.filter.Values)...)

	case "is_equal":
		return c.Eq(b.filter.Values[0])

	case "is_not_equal":
		return c.Neq(b.filter.Values[0])

	case "is_greater_than":
		return c.Gt(b.filter.Values[0])

	case "is_less_than":
		return c.Lt(b.filter.Values[0])

	case "is_greater_than_or_equal":
		return c.Gte(b.filter.Values[0])

	case "is_less_than_or_equal":
		return c.Lte(b.filter.Values[0])

	case "is_empty":
		return c.IsNull()

	case "is_not_empty":
		return c.IsNotNull()
	}

	return nil
}

// SELECT BUILD IMPLEMENTATION

type BuildSelect struct {
	column string
	filter Filter
}

func NewBuildSelect(column string, filter Filter) Builder {
	return &BuildSelect{column, filter}
}

func (b *BuildSelect) Build() exp.Expression {
	var (
		l = len(b.filter.Values)
		c = goqu.L(b.column)
	)

	if l <= 0 {
		return nil
	}

	switch b.filter.Operation {
	case "is":
		return c.In(conv(b.filter.Values)...)

	case "is_not":
		return c.NotIn(conv(b.filter.Values)...)

	case "is_empty":
		return c.IsNull()

	case "is_not_empty":
		return c.IsNotNull()
	}

	return nil
}

// BOOL BUILD IMPLEMENTATION

type BuildBool struct {
	column string
	filter Filter
}

func NewBuildBool(column string, filter Filter) Builder {
	return &BuildBool{column, filter}
}

func (b *BuildBool) Build() exp.Expression {
	var (
		l = len(b.filter.Values)
		c = goqu.L(b.column)
	)

	if l <= 0 {
		return nil
	}

	switch b.filter.Operation {
	case "is":
		if b.filter.Values[0] == "active" {
			return c.IsTrue()
		}
		return c.IsFalse()

	case "is_empty":
		return c.IsNull()

	case "is_not_empty":
		return c.IsNotNull()
	}

	return nil
}

// DATE BUILD IMPLEMENTATION

type BuildDate struct {
	column string
	filter Filter
}

func NewBuildDate(column string, filter Filter) Builder {
	return &BuildDate{column, filter}
}

func (b *BuildDate) Build() exp.Expression {
	var (
		l = len(b.filter.Values)
		c = goqu.L(b.column)
	)

	if l <= 0 {
		return nil
	}

	switch b.filter.Operation {
	case "is":
		return goqu.L("DATE(?) = DATE(?)", goqu.I(b.column), goqu.V(b.filter.Values[0]))

	case "is_before":
		return goqu.L("DATE(?) < DATE(?)", goqu.I(b.column), goqu.V(b.filter.Values[0]))

	case "is_after":
		return goqu.L("DATE(?) > DATE(?)", goqu.I(b.column), goqu.V(b.filter.Values[0]))

	case "is_on_or_before":
		return goqu.L("DATE(?) <= DATE(?)", goqu.I(b.column), goqu.V(b.filter.Values[0]))

	case "is_on_or_after":
		return goqu.L("DATE(?) >= DATE(?)", goqu.I(b.column), goqu.V(b.filter.Values[0]))

	case "is_between":
		return goqu.L("? BETWEEN ? AND ?", goqu.I(b.column), goqu.V(b.filter.Values[0]), goqu.V(b.filter.Values[1]))

	case "is_empty":
		return c.IsNull()

	case "is_not_empty":
		return c.IsNotNull()
	}

	return nil
}
