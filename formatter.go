package errors

import (
	"fmt"
	"runtime"
	"strings"
)

// Formatter defines how stack trace is formatted.
type Formatter string

// constants for Stack formatters.
const (
	DefaultFormatter    Formatter = "file.go:152\n"
	JavaLikeFormatter             = "at foo(file.go:152)\n"
	GoLikeFormatter               = "foo\n\tfile.go:152\n"
	PythonLikeFormatter           = "File file.go, line 152, in foo\n"
)

var tracerFormatter = DefaultFormatter

// ApplyFormatter specifies the formatter for errorTracer.
// Apply DefaultFormatter if not specified.
// The template of formatter: "foo" for function name, "file.go" for file name,
// "152" for line number.
func ApplyFormatter(formatter Formatter) {
	tracerFormatter = formatter
}

func (f Formatter) format(frame runtime.Frame) string {
	formatted := string(tracerFormatter)
	formatted = strings.Replace(formatted, "foo", frame.Function, 1)
	formatted = strings.Replace(formatted, "file.go", frame.File, 1)
	formatted = strings.Replace(formatted, "152", fmt.Sprint(frame.Line), 1)
	return formatted
}
