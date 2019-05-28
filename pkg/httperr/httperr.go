package httperr

import (
	"fmt"
	"net/http"
)

type httpErr struct {
	message    string
	statusCode int
}

func (e httpErr) Error() string {
	return e.message
}

// StatusCode returns a HTTP status code for the given error.
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

// Write writes the given error to the HTTP response writewr.
func Write(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), StatusCode(err))
}
