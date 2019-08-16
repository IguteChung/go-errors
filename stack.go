package errors

import (
	"fmt"
	"runtime"
	"strings"
)

// constants for Stack formatters.
const (
	DefaultFormatter    = "file.go:152\n"
	JavaLikeFormatter   = "at foo(file.go:152)\n"
	GoLikeFormatter     = "foo\n\tfile.go:152\n"
	PythonLikeFormatter = "File file.go, line 152, in foo\n"
)

// Stack defines the program counters.
type Stack []uintptr

func recordStack() *Stack {
	s := make(Stack, 64)
	n := runtime.Callers(3, s)
	s = s[:n]
	return &s
}

// Format formats the stack trace with customization.
// Layout subsistution keywords:
// foo: function name
// file.go: file name
// 152: line number
func (s *Stack) Format(layout string) (ret string) {
	// check nil.
	if s == nil {
		return
	}
	v := *s

	// check formatter
	if layout == "" {
		layout = DefaultFormatter
	}

	// handle if no frames.
	if len(v) == 0 {
		return
	}

	// get frames for current stack.
	frames := runtime.CallersFrames(v)
	for {
		frame, more := frames.Next()

		// apply the layout formatter.
		formatted := layout
		formatted = strings.Replace(formatted, "foo", frame.Function, 1)
		formatted = strings.Replace(formatted, "file.go", frame.File, 1)
		formatted = strings.Replace(formatted, "152", fmt.Sprint(frame.Line), 1)
		ret = ret + formatted

		if !more {
			return
		}
	}
}
