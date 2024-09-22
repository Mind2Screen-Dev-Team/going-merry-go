package xvalidate

import (
	"encoding/json"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var (
	ErrInvalidJsonMarshal   = errors.New("invalid json marshal error validation")
	ErrInvalidJsonUnmarshal = errors.New("invalid json unmarshal error validation")
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
	if e, ok := err.(validation.Errors); ok {
		b, err := e.MarshalJSON()
		if err != nil {
			return ErrInvalidJsonMarshal
		}

		if err := json.Unmarshal(b, &errs); err != nil {
			return ErrInvalidJsonUnmarshal
		}

		return errs
	}

	return nil
}
