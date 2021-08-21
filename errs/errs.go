package errs

import (
	"fmt"
)

const (
	CodeOK = 0

	CodeServerUnmarshalFail = 11
	CodeServerMarshalFail   = 12

	CodeClientUnmarshalFail = 51
	CodeClientMarshalFail   = 52
	CodeClientCallFail      = 53

	// 100 ... 500, the http code

	CodeUnknown = 999

	// business err code start from 1000
)

// Err the err
type Err struct {
	Code int
	Msg  string
}

// Error the error interface
func (e Err) Error() string {
	return fmt.Sprintf("code:%d, msg:%s", e.Code, e.Msg)
}

// New ...
func New(code int, msg string) error {
	return &Err{Code: code, Msg: msg}
}

// Code get the err code
func Code(err error) int {
	if err == nil {
		return CodeOK
	}

	e, ok := err.(*Err)
	if !ok {
		return CodeUnknown
	}

	if e == nil {
		return CodeOK
	}

	return e.Code
}

// Msg get err msg
func Msg(err error) string {
	if err == nil {
		return "OK"
	}

	e, ok := err.(*Err)
	if !ok {
		return err.Error()
	}

	if e == nil {
		return "OK"
	}

	return e.Msg
}
