package errors

import (
	"strconv"
)

func NewError(code int, msg string) *CodeError {
	return &CodeError{code, msg, nil}
}

func NewErrorMsg(msg string) *CodeError {
	return &CodeError{0, msg, nil}
}

func NewErrorData(code int, msg string, data interface{}) *CodeError {
	return &CodeError{code, msg, data}
}

func FromError(err error) *CodeError {
	if err == nil {
		return nil
	}
	return &CodeError{0, err.Error(), nil}
}

type CodeError struct {
	Code    int
	Message string
	Data    interface{}
}

func (e *CodeError) Error() string {
	return strconv.Itoa(e.Code) + ": " + e.Message
}
