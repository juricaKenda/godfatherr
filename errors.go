package godfatherr

import (
	"errors"
	"fmt"
)

type Error struct {
	err   error
	ctx   []string
	fatal bool
}

// New creates a new Error containing a given message
func New(message string) *Error {
	return &Error{
		err: errors.New(message),
		ctx: make([]string, 0),
	}
}

// Empty creates a new Empty message, indicating a lack of error
func Empty() *Error {
	return &Error{}
}

// IsPresent returns true if the error is present, false otherwise
func (e *Error) IsPresent() bool {
	return e.err != nil
}

// WithCtx assigns all given context strings to the Error
// An empty Error will not be assigned any context
func (e *Error) WithCtx(ctx ...string) *Error {
	if !e.IsPresent() {
		return e
	}
	e.ctx = append(e.ctx, ctx...)
	return e
}

// ContainsCtx returns true if the Error contains a given ctx value, false otherwise
func (e *Error) ContaincCtx(ctx string) bool {
	for _, element := range e.ctx {
		if element == ctx {
			return true
		}
	}
	return false
}

// Fatal marks the Error fatal
func (e *Error) Fatal() *Error {
	e.fatal = true
	return e
}

// IsFatal returns true if the Error is fatal, false otherwise
func (e *Error) IsFatal() bool {
	return e.fatal
}

// Error returns a std error, nil if the Error is empty
func (e *Error) Error() error {
	return e.err
}

// Message returns the Error message string
// An empty Error returns an empty string
func (e *Error) Message() string {
	if !e.IsPresent() {
		return ""
	}
	return e.err.Error()
}

// Dispatchable is any func which returns the Error
type Dispatchable func(...interface{}) *Error

// IfAbsent will dispatch a func if the Error is NOT present and return the resulting Error of the dispatchable
// otherwise, it will return itself and will not dispatch a given func
func (e *Error) IfAbsent(d Dispatchable, args ...interface{}) *Error {
	if e.IsPresent() {
		return e
	}
	return d(args)
}

// OrElse will dispatch a func if the Error is present and return the resulting Error of the dispatchable
// // otherwise, it will return itself and will not dispatch a given func
// *OrElse is a polar opposite of IfAbsent*
func (e *Error) OrElse(d Dispatchable, args ...interface{}) *Error {
	if e.IsPresent() {
		return d(args)
	}
	return e
}

// String returns the Error message string representation
// An empty Error returns an empty err tag
func (e *Error) String() string {
	if !e.IsPresent() {
		return "[empty err]"
	}
	return fmt.Sprintf("%s %v", e.err, e.ctx)
}

// Print prints out the Error
func (e *Error) Print() {
	fmt.Println(e)
}
