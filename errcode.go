// Package errcode provides for errors with codes. This is handly for
// HTTP errors among other things.
package errcode

import (
	"fmt"
)

// An Error and a Code.
type Error struct {
	code int
	err  error
}

// Return the underlying error.
func (e Error) Error() string {
	return e.err.Error()
}

// Return the associated code.
func (e Error) Code() int {
	return e.code
}

// Create a new Coded Error.
func New(code int, f string, args ...interface{}) Error {
	return Error{
		code: code,
		err:  fmt.Errorf(f, args...),
	}
}

// Add a Code to an existing Error.
func Add(code int, err error) Error {
	return Error{code: code, err: err}
}
