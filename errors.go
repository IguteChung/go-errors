package errors

import "fmt"

// traceable indicates an error whose stacktrace is provided.
type traceable interface {
	// Cause returns original error.
	Cause() error
	// Stack returns stack trace of this error.
	Stack() *stack
}

// traceableError implements the basic traceable error.
type traceableError struct {
	err   error
	stack *stack
}

func (t traceableError) Error() string {
	if t.err == nil {
		return ""
	}
	return t.err.Error()
}

func (t traceableError) Cause() error {
	return t.err
}

func (t traceableError) Stack() *stack {
	return t.stack
}

// StackTrace prints the stack trace with given error by the formatter.
// If the error is not traceable, empty string is returned.
func StackTrace(err error) string {
	if tracer, ok := err.(traceable); ok {
		return tracer.Stack().Format()
	}
	return ""
}

// Cause returns the original error of a traceable error.
// If the error is not traceable, return itself.
func Cause(err error) error {
	if tracer, ok := err.(traceable); ok {
		return tracer.Cause()
	}
	return err
}

// New creates a traceable error with message.
func New(msg string) error {
	return traceableError{
		err:   fmt.Errorf(msg),
		stack: recordStack(),
	}
}

// Errorf creates a traceable error with formatted message.
// If args contains a traceable error, apply the stack of it.
// If args doesn't contain any traceable error, record the stack trace.
func Errorf(msg string, args ...interface{}) error {
	e := fmt.Errorf(msg, args...)
	for _, arg := range args {
		if t, ok := arg.(traceable); ok {
			return traceableError{
				err:   e,
				stack: t.Stack(),
			}
		}
	}
	return traceableError{
		err:   e,
		stack: recordStack(),
	}
}

// Wrap wraps an error into a traceable error.
// If the wrapped error is traceable, do nothing.
// If the wrapped error is not traceable, record the stack trace.
func Wrap(err error) error {
	if _, ok := err.(traceable); ok {
		return err
	}
	return traceableError{
		err:   err,
		stack: recordStack(),
	}
}

// WrapMessage wraps an error into a traceable error with message.
// If the wrapped error is traceable, only append the error message.
// If the wrapped error is not traceable, record the stack trace.
func WrapMessage(err error, msg string) error {
	e := fmt.Errorf("%s: %v", msg, err)
	if tracer, ok := err.(traceable); ok {
		return traceableError{
			err:   e,
			stack: tracer.Stack(),
		}
	}
	return traceableError{
		err:   e,
		stack: recordStack(),
	}
}

// WrapMessagef wraps an error into a traceable error with formatted message.
// If the wrapped error is traceable, only append the error message.
// If the wrapped error is not traceable, record the stack trace.
func WrapMessagef(err error, msg string, args ...interface{}) error {
	formattedMsg := fmt.Sprintf(msg, args...)
	e := fmt.Errorf("%s: %v", formattedMsg, err)
	if tracer, ok := err.(traceable); ok {
		return traceableError{
			err:   e,
			stack: tracer.Stack(),
		}
	}
	return traceableError{
		err:   e,
		stack: recordStack(),
	}
}
