package errors

import "fmt"

// ErrorTracer indicates a tracable error.
type ErrorTracer interface {
	error
	// Cause returns original error.
	Cause() error
	// Stack returns stack trace of this error.
	Stack() *Stack
}

// tracer implements the basic ErrorTracer.
type tracer struct {
	err   error
	stack *Stack
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

func (t tracer) Stack() *Stack {
	return t.stack
}

// New creates an ErrorTracer with error message.
func New(msg string) error {
	return tracer{
		err:   fmt.Errorf(msg),
		stack: recordStack(),
	}
}

// Errorf creates an ErrorTracer with formatted message.
// If args contains an ErrorTracer, apply the stack of it.
// If args doesn't contain an ErrorTracer, record the stack trace.
func Errorf(msg string, args ...interface{}) error {
	e := fmt.Errorf(msg, args...)
	for _, arg := range args {
		if t, ok := arg.(ErrorTracer); ok {
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

// Wrap wraps an error into an ErrorTracer.
// If the wrapped error is an ErrorTracer, do nothing.
// If the wrapped error is not an ErrorTracer, record the stack trace.
func Wrap(err error) error {
	if _, ok := err.(ErrorTracer); ok {
		return err
	}
	return tracer{
		err:   err,
		stack: recordStack(),
	}
}

// WrapMessage wraps an error into an ErrorTracer with message.
// If the wrapped error is an ErrorTracer, only append the error message.
// If the wrapped error is not an ErrorTracer, record the stack trace.
func WrapMessage(err error, msg string) error {
	e := fmt.Errorf("%s: %v", msg, err)
	if ErrorTracer, ok := err.(ErrorTracer); ok {
		return tracer{
			err:   e,
			stack: ErrorTracer.Stack(),
		}
	}
	return tracer{
		err:   e,
		stack: recordStack(),
	}
}

// WrapMessagef wraps an error into an ErrorTracer with formatted message.
// If the wrapped error is an ErrorTracer, only append the error message.
// If the wrapped error is not an ErrorTracer, record the stack trace.
func WrapMessagef(err error, msg string, args ...interface{}) error {
	formattedMsg := fmt.Sprintf(msg, args...)
	e := fmt.Errorf("%s: %v", formattedMsg, err)
	if ErrorTracer, ok := err.(ErrorTracer); ok {
		return tracer{
			err:   e,
			stack: ErrorTracer.Stack(),
		}
	}
	return tracer{
		err:   e,
		stack: recordStack(),
	}
}
