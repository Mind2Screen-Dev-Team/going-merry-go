package xhttputil

import (
	"bytes"
	"io"
	"net/http"
)

func DeepCopyRequest(r *http.Request) *http.Request {
	// Read the body if it's non-nil
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = io.ReadAll(r.Body)
		// Refill the original request body to preserve it for further usage.
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	// Create a shallow copy of the request
	rCopy := r.Clone(r.Context())

	// Replace the body of the new request with a new reader wrapping the copied bytes
	if bodyBytes != nil {
		rCopy.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	return rCopy
}
