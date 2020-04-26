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

func New(message string) *Error {
	return &Error{
		err: errors.New(message),
		ctx: make([]string, 0),
	}
}
func Empty() *Error {
	return &Error{}
}

func (e *Error) IsPresent() bool {
	return e.err != nil
}

func (e *Error) WithCtx(ctx ...string) *Error {
	e.ctx = append(e.ctx, ctx...)
	return e
}

func (e *Error) ContaincCtx(ctx string) bool {
	for _, element := range e.ctx {
		if element == ctx {
			return true
		}
	}
	return false
}

func (e *Error) Fatal() *Error {
	e.fatal = true
	return e
}

func (e *Error) IsFatal() bool {
	return e.fatal
}

func (e *Error) Error() error {
	return e.err
}

func (e *Error) Message() string {
	return e.err.Error()
}

func (e *Error) String() string {
	return fmt.Sprintf("%s %v", e.err, e.ctx)
}

func (e *Error) Print() {
	if !e.IsPresent() {
		fmt.Println("[empty]")
		return
	}
	fmt.Println(e)
}
