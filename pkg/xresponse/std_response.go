package xresponse

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// A Standard Response For API
type ResponseSTD[D any, E any] struct {
	statusCode     int                 `json:"-"`
	request        *http.Request       `json:"-"`
	responseHeader http.Header         `json:"-"`
	responseWriter http.ResponseWriter `json:"-"`

	interceptHandler InterceptHandler[D, E] `json:"-"`

	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data D      `json:"data"`
	Err  E      `json:"err"`
}

func (r *ResponseSTD[D, E]) SetMsg(msg string) *ResponseSTD[D, E] {
	r.Msg = msg
	return r
}

func (r *ResponseSTD[D, E]) SetCode(code string) *ResponseSTD[D, E] {
	r.Code = code
	return r
}

func (r *ResponseSTD[D, E]) SetData(data D) *ResponseSTD[D, E] {
	r.Data = data
	return r
}

func (r *ResponseSTD[D, E]) SetError(err E) *ResponseSTD[D, E] {
	r.Err = err
	return r
}

func (r *ResponseSTD[D, E]) JSONText() (string, error) {
	var buf bytes.Buffer

	defer buf.Reset()

	if err := json.NewEncoder(&buf).Encode(r); err != nil {
		return "", nil
	}

	return buf.String(), nil
}

func (r *ResponseSTD[D, E]) JSON(w io.Writer) {
	json.NewEncoder(w).Encode(r)
}

func (r *ResponseSTD[D, E]) JSONOrErr(w io.Writer) error {
	return json.NewEncoder(w).Encode(r)
}

// # HTTP Rest API

func (r *ResponseSTD[D, E]) SetStatusCode(code int) *ResponseSTD[D, E] {
	r.statusCode = code
	return r
}

func (r *ResponseSTD[D, E]) SetHeader(key string, value string) *ResponseSTD[D, E] {
	r.responseWriter.Header().Add(key, value)
	return r
}

func (r *ResponseSTD[D, E]) DelHeader(key string) *ResponseSTD[D, E] {
	r.responseWriter.Header().Del(key)
	return r
}

func (r *ResponseSTD[D, E]) RestJSON() {
	r.responseWriter.Header().Add("Accept", "application/json")
	r.responseWriter.Header().Add("Content-Type", "application/json")
	r.responseWriter.WriteHeader(r.statusCode)
	r.JSON(r.responseWriter)
}

func (r *ResponseSTD[D, E]) RestJSONOrErr() error {
	r.responseWriter.Header().Add("Accept", "application/json")
	r.responseWriter.Header().Add("Content-Type", "application/json")
	r.responseWriter.WriteHeader(r.statusCode)
	return r.JSONOrErr(r.responseWriter)
}

// # Add More Master Implementation on Here ...
