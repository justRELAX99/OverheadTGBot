package errors

import (
	"github.com/pkg/errors"
)

func New(msg string) error {
	return NoType.New(msg)
}

func Newf(msg string, args ...interface{}) error {
	return NoType.Newf(msg, args...)
}

func Wrapf(err error, msg string, args ...interface{}) error {
	if customErr, ok := err.(customError); ok {
		wrappedErr := errors.Wrapf(err, msg, args...)
		return customError{
			errType:    customErr.errType,
			original:   wrappedErr,
			stackTrace: customErr.stackTrace,
			status:     customErr.status,
			message:    customErr.message,
		}
	}
	return NoType.Wrapf(err, msg, args...)
}

func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

func Cause(err error) error {
	return errors.Cause(err)
}

// Compare является shorthand'ом для Cause(orig).Error() == compared.Error()
func Compare(orig, compared error) bool {
	return Cause(orig).Error() == compared.Error()
}
