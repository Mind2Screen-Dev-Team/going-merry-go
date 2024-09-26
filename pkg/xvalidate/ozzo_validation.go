package xvalidate

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Errors map[string]any

func (e Errors) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func WrapperValidation(err error) error {
	if err == nil {
		return nil
	}

	if e, ok := err.(validation.InternalError); ok {
		return e.InternalError()
	}

	if e, ok := err.(validation.Errors); ok {
		err = e.Filter()
	}

	var errs Errors
	if es, ok := err.(validation.Errors); ok {
		jsonBytes, _ := es.MarshalJSON()
		json.Unmarshal(jsonBytes, &errs)
		return errs
	}

	return nil
}

func IsErrors(err error) (errs Errors, ok bool) {
	errs, ok = err.(Errors)
	return
}
