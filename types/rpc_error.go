package types

import (
	"errors"
	"fmt"
)

type RPCError interface {
	ErrCode() int32
	ErrMessage() string
}

// IntoError Into RPC Error
type IntoError interface {
	Into(code int32, message string)
}

type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func NewError(code int32, message string) *Error {
	return &Error{Code: code, Message: message}
}

func (e *Error) Error() string {
	return fmt.Sprintf("code=%d, message=%q", e.Code, e.Message)
}

func (e *Error) Into(code int32, message string) {
	e.Code, e.Message = code, message
}

func (e *Error) ErrCode() int32 {
	return e.Code
}

func (e *Error) ErrMessage() string {
	return e.Message
}

func (e *Error) As(v any) bool {
	if v, ok := v.(IntoError); ok {
		v.Into(e.Code, e.Message)

		return true
	}

	return false
}

func (e *Error) Is(err error) bool {
	if ee := FromError(err); ee != nil {
		return e.Code == ee.Code && e.Message == ee.Message
	}

	if ee, ok := err.(RPCError); ok {
		return e.Code == ee.ErrCode() && e.Message == ee.ErrMessage()
	}

	return false
}

func FromError(err error) *Error {
	if err == nil {
		return nil
	}

	if e, ok := err.(*Error); ok {
		return e
	}

	if e := new(Error); errors.As(err, &e) {
		return e
	}

	return nil
}
