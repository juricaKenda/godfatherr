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
func (e *Error) WithCtx(ctx ...string) *Error {
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

// Error returns a std error
func (e *Error) Error() error {
	return e.err
}

// Message returns the Error message string
func (e *Error) Message() string {
	return e.err.Error()
}

// String returns the Error message string representation
func (e *Error) String() string {
	return fmt.Sprintf("%s %v", e.err, e.ctx)
}

// Print prints out the Error
func (e *Error) Print() {
	if !e.IsPresent() {
		fmt.Println("[empty]")
		return
	}
	fmt.Println(e)
}
