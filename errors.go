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

// Consumerr is any func which consumes the Error
type Consumerr func(err *Error)

// Producerr is any func which produces the Error
type Producerr func() *Error

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
func (e *Error) ContainsCtx(ctx string) bool {
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

func (e *Error) Is(other *Error) bool {
	return false
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

// IfPresent will dispatch a consumer func
// otherwise, it will do nothing
func (e *Error) IfPresent(consumerr Consumerr) {
	if e.IsPresent() {
		consumerr(e)
	}
}

// IfAbsent executes a given Producer func if error is absent
func (e *Error) IfAbsent(producerr Producerr) *Error {
	if !e.IsPresent() {
		return producerr()
	}
	return e
}

// WhileAbsent executes each Producer func until it encounters the Error
// If the Error is not encountered, it will finish when it iterates through all Producers
func WhileAbsent(prod ...Producerr) *Error {
	for _, p := range prod {
		if err := p(); err.IsPresent() {
			return err
		}
	}
	return Empty()
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

//Panic delegation
func (e *Error) Panic() {
	if !e.IsPresent() {
		return
	}
	panic(e)
}

//Watch with the Watchdog
func (e *Error) Watch() {
	e.WithCtx(WATCHED)
	e.Panic()
}

func (e *Error) isWatched() bool {
	return e.ContainsCtx(WATCHED)
}

func (e *Error) unwatch() *Error {
	for i, other := range e.ctx {
		if other == WATCHED {
			e.ctx = append(e.ctx[:i], e.ctx[i+1:]...)
		}
	}
	return e
}
