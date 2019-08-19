package errors

import "fmt"

// errorTracer indicates a tracable error.
type errorTracer interface {
	error
	// Cause returns original error.
	Cause() error
	// Stack returns stack trace of this error.
	Stack() *stack
}

// tracer implements the basic errorTracer.
type tracer struct {
	err   error
	stack *stack
}

func (t tracer) Error() string {
	if t.err == nil {
		return ""
	}
	return t.err.Error()
}

func (t tracer) Cause() error {
	return t.err
}

func (t tracer) Stack() *stack {
	return t.stack
}

// StackTrace prints the stack trace with given error by the formatter.
// If the error is not an errorTracer, empty string is returned.
func StackTrace(err error) string {
	if tracer, ok := err.(errorTracer); ok {
		return tracer.Stack().Format()
	}
	return ""
}

// New creates an errorTracer with error message.
func New(msg string) error {
	return tracer{
		err:   fmt.Errorf(msg),
		stack: recordStack(),
	}
}

// Errorf creates an errorTracer with formatted message.
// If args contains an errorTracer, apply the stack of it.
// If args doesn't contain an errorTracer, record the stack trace.
func Errorf(msg string, args ...interface{}) error {
	e := fmt.Errorf(msg, args...)
	for _, arg := range args {
		if t, ok := arg.(errorTracer); ok {
			return tracer{
				err:   e,
				stack: t.Stack(),
			}
		}
	}
	return tracer{
		err:   e,
		stack: recordStack(),
	}
}

// Wrap wraps an error into an errorTracer.
// If the wrapped error is an errorTracer, do nothing.
// If the wrapped error is not an errorTracer, record the stack trace.
func Wrap(err error) error {
	if _, ok := err.(errorTracer); ok {
		return err
	}
	return tracer{
		err:   err,
		stack: recordStack(),
	}
}

// WrapMessage wraps an error into an errorTracer with message.
// If the wrapped error is an errorTracer, only append the error message.
// If the wrapped error is not an errorTracer, record the stack trace.
func WrapMessage(err error, msg string) error {
	e := fmt.Errorf("%s: %v", msg, err)
	if errorTracer, ok := err.(errorTracer); ok {
		return tracer{
			err:   e,
			stack: errorTracer.Stack(),
		}
	}
	return tracer{
		err:   e,
		stack: recordStack(),
	}
}

// WrapMessagef wraps an error into an errorTracer with formatted message.
// If the wrapped error is an errorTracer, only append the error message.
// If the wrapped error is not an errorTracer, record the stack trace.
func WrapMessagef(err error, msg string, args ...interface{}) error {
	formattedMsg := fmt.Sprintf(msg, args...)
	e := fmt.Errorf("%s: %v", formattedMsg, err)
	if errorTracer, ok := err.(errorTracer); ok {
		return tracer{
			err:   e,
			stack: errorTracer.Stack(),
		}
	}
	return tracer{
		err:   e,
		stack: recordStack(),
	}
}
