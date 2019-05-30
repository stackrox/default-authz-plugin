package httperr

import (
	"fmt"
	"net/http"
)

// httpErr is an error type that includes an HTTP status code in addition to the error message.
type httpErr struct {
	message    string
	statusCode int
}

// Error implements the builtin `error` interface.
func (e httpErr) Error() string {
	return e.message
}

// StatusCode returns a HTTP status code for the given error. If the error is not an HTTP error created via `New` or
// `Newf` from this package, the default status code to be used is `Internal Server Error (500)`.
func StatusCode(err error) int {
	if he, ok := err.(httpErr); ok {
		return he.statusCode
	}
	return http.StatusInternalServerError
}

// New returns a new HTTP error with the given code and message.
func New(code int, msg string) error {
	return httpErr{
		message:    msg,
		statusCode: code,
	}
}

// Newf returns a new HTTP error with the given code and a message constructed from the format string and args.
func Newf(code int, format string, args ...interface{}) error {
	return New(code, fmt.Sprintf(format, args...))
}

// Write writes the given error to the HTTP response writer.
func Write(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), StatusCode(err))
}
