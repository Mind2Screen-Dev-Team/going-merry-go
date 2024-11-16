package xhttputil

import (
	"bytes"
	"context"
	"io"
	"net/http"
)

type StrReqBody string

func (b StrReqBody) String() string {
	return string(b)
}

const (
	STR_REQ_BODY = StrReqBody("STR_REQ_BODY")
)

func DeepCopyRequest(r *http.Request, saveReqBody ...bool) *http.Request {
	// Read the body if it's non-nil
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = io.ReadAll(r.Body)
		// Refill the original request body to preserve it for further usage.
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	if len(saveReqBody) > 0 {
		if saveReqBody[0] && bodyBytes != nil {
			r = r.WithContext(context.WithValue(r.Context(), STR_REQ_BODY, string(bodyBytes)))
		}
	}

	// Create a shallow copy of the request
	rCopy := r.Clone(r.Context())

	if len(saveReqBody) > 0 {
		if saveReqBody[0] && bodyBytes != nil {
			rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), STR_REQ_BODY, string(bodyBytes)))
		}
	}

	// Replace the body of the new request with a new reader wrapping the copied bytes
	if bodyBytes != nil {
		rCopy.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	return rCopy
}
