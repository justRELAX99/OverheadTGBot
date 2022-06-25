package errors

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type Type uint
type Status uint

const (
	NoType = Type(iota)
	NotFound
	BadRequest
	Forbidden
	Timeout
	PaymentRequired
)

const (
	StatusDefault Status = iota
	StatusAlreadyExist
)

var statuses = map[Status]string{
	StatusAlreadyExist: "exist",
}

func (t Type) StatusCode() int {
	switch t {
	case NotFound:
		return http.StatusNotFound
	case BadRequest:
		return http.StatusBadRequest
	case Forbidden:
		return http.StatusForbidden
	case Timeout:
		return http.StatusRequestTimeout
	case PaymentRequired:
		return http.StatusPaymentRequired
	default:
		return http.StatusInternalServerError
	}
}

type customError struct {
	message    string
	errType    Type
	original   error
	status     Status
	stackTrace errors.StackTrace
}

func (err customError) Error() string {
	return err.original.Error()
}

func (err customError) Cause() error {
	return err.original
}

func (err customError) StackTrace() errors.StackTrace {
	return err.stackTrace
}

func (et Type) NewWithStatus(status Status, msg string) error {
	err := errors.New(msg)
	return customError{
		status:     status,
		errType:    et,
		original:   err,
		stackTrace: err.(stackTracer).StackTrace(),
	}
}

func (et Type) New(msg string) error {
	err := errors.New(msg)
	return customError{
		errType:    et,
		original:   err,
		stackTrace: err.(stackTracer).StackTrace(),
	}
}

func (et Type) NewfWithStatus(status Status, msg string, args ...interface{}) error {
	err := errors.Errorf(msg, args...)
	return customError{
		status:     status,
		errType:    et,
		original:   err,
		stackTrace: err.(stackTracer).StackTrace(),
	}
}

func (et Type) Newf(msg string, args ...interface{}) error {
	err := errors.Errorf(msg, args...)
	return customError{
		errType:    et,
		original:   err,
		stackTrace: err.(stackTracer).StackTrace(),
	}
}

func (et Type) Wrap(err error, msg string) error {
	return et.Wrapf(err, msg)
}

func (et Type) Wrapf(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	newErr := errors.Wrapf(err, msg, args...)
	ce := customError{
		errType:    et,
		original:   newErr,
		stackTrace: newErr.(stackTracer).StackTrace(),
		status:     GetStatus(err),
	}
	return ce
}

func GetType(err error) Type {
	if customErr, ok := err.(customError); ok {
		return customErr.errType
	}
	return NoType
}

func GetStatusStr(err error) string {
	if customErr, ok := err.(customError); ok {
		return statuses[customErr.status]
	}
	return ""
}

func GetStatus(err error) Status {
	if customErr, ok := err.(customError); ok {
		return customErr.status
	}
	return StatusDefault
}

func (et Type) NewWithMessage(err error, msg string) error {
	cr := customError{
		message:  msg,
		original: err,
		errType:  et,
	}
	return cr
}

func (et Type) NewMessage(msg string) error {
	cr := customError{
		message:  msg,
		errType:  et,
		original: errors.New(msg),
	}
	return cr
}

func (et Type) NewMessagef(msg string, args ...interface{}) error {
	cr := customError{
		message:  fmt.Sprintf(msg, args...),
		errType:  et,
		original: errors.New(fmt.Sprintf(msg, args...)),
	}
	return cr
}
