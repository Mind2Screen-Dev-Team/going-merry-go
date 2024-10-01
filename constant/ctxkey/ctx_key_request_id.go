package ctxkey

// Key to use when setting the request ID.
type RequestID int

// RequestIDKey is the key that holds the unique request ID in a request context.
const RequestIDKey RequestID = 0
