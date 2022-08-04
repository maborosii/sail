package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	// code    int      `json:"code"`
	// msg     string   `json:"msg"`
	// details []string `json:"details"`
	code    int
	msg     string
	details []string
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("this code %d is existed, please choose another one", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("msg: %s", e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	newError.details = append(newError.details, details...)
	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case RequestTypeNotSupport.Code(), RequestTypeNotAdapt.Code(), RequestSourceNotSupport.Code(), RequestExpired.Code():
		return http.StatusRequestedRangeNotSatisfiable
	case BadRequest.Code():
		return http.StatusBadRequest
	case NotFound.Code():
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
