package errors

import (
	"github.com/pkg/errors"
	"strings"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func StackTrace(err error) errors.StackTrace {
	if stErr, ok := err.(stackTracer); ok {
		// Если реализует, возвращаем ее stacktrace
		return stErr.StackTrace()
	}
	return nil
}
func GetStackTraceArr(err error) map[string]interface{} {
	st := StackTrace(err)
	stArray := make([]string, 0, len(st))
	for _, frame := range st {
		frBytes, _ := frame.MarshalText()
		stArray = append(stArray, string(frBytes))
	}
	return map[string]interface{}{"stack_trace": strings.Join(stArray, "\n")}
}
