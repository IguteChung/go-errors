package errors

import (
	"runtime"
)

// stack defines the program counters.
type stack []uintptr

func recordStack() *stack {
	s := make(stack, 64)
	n := runtime.Callers(3, s)
	s = s[:n]
	return &s
}

// Format formats the stack trace with the layout registered in FormatTracer.
func (s *stack) Format() (ret string) {
	// check nil.
	if s == nil {
		return
	}
	v := *s

	// handle if no frames.
	if len(v) == 0 {
		return
	}

	// get frames for current stack.
	frames := runtime.CallersFrames(v)
	for {
		frame, more := frames.Next()

		// apply the layout formatter.
		ret = ret + tracerFormatter.format(frame)

		if !more {
			return
		}
	}
}
