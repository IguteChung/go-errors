package errors

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStackError(t *testing.T) {
	e := New("something wrong")
	assert.EqualError(t, e, "something wrong")

	frames := strings.Split(e.(ErrorTracer).Stack().Format(""), "\n")
	assert.True(t, strings.HasSuffix(frames[0], "errors_test.go:12"))
}

func TestErrorfError(t *testing.T) {
	e := Errorf("%s", "something wrong")
	assert.EqualError(t, e, "something wrong")

	frames := strings.Split(e.(ErrorTracer).Stack().Format(""), "\n")
	assert.True(t, strings.HasSuffix(frames[0], "errors_test.go:20"))
}

func TestWrapError(t *testing.T) {
	e := fmt.Errorf("something wrong")
	wrap := Wrap(e)
	assert.EqualError(t, wrap, "something wrong")

	frames := strings.Split(wrap.(ErrorTracer).Stack().Format(""), "\n")
	assert.True(t, strings.HasSuffix(frames[0], "errors_test.go:29"))
}

func TestWrapStackError(t *testing.T) {
	e := New("something wrong")
	wrap := Wrap(e)
	assert.EqualError(t, wrap, "something wrong")

	frames := strings.Split(wrap.(ErrorTracer).Stack().Format(""), "\n")
	assert.True(t, strings.HasSuffix(frames[0], "errors_test.go:37"))
}

func TestWrapMessageError(t *testing.T) {
	e := fmt.Errorf("something wrong")
	wrap := WrapMessage(e, "At someplace")
	assert.EqualError(t, wrap, "At someplace: something wrong")

	frames := strings.Split(wrap.(ErrorTracer).Stack().Format(""), "\n")
	assert.True(t, strings.HasSuffix(frames[0], "errors_test.go:47"))
}

func TestWrapMessageStackError(t *testing.T) {
	e := New("something wrong")
	wrap := WrapMessage(e, "At someplace")
	assert.EqualError(t, wrap, "At someplace: something wrong")

	frames := strings.Split(wrap.(ErrorTracer).Stack().Format(""), "\n")
	assert.True(t, strings.HasSuffix(frames[0], "errors_test.go:55"))
}

func TestWrapMessagefError(t *testing.T) {
	e := fmt.Errorf("something wrong")
	wrap := WrapMessagef(e, "At someplace %s", "here")
	assert.EqualError(t, wrap, "At someplace here: something wrong")

	frames := strings.Split(wrap.(ErrorTracer).Stack().Format(""), "\n")
	assert.True(t, strings.HasSuffix(frames[0], "errors_test.go:65"))
}

func TestWrapMessagefStackError(t *testing.T) {
	e := New("something wrong")
	wrap := WrapMessagef(e, "At someplace %s", "here")
	assert.EqualError(t, wrap, "At someplace here: something wrong")

	frames := strings.Split(wrap.(ErrorTracer).Stack().Format(""), "\n")
	assert.True(t, strings.HasSuffix(frames[0], "errors_test.go:73"))
}

func TestErrorfStackError(t *testing.T) {
	e := New("something wrong")
	wrap := Errorf("%v", e)
	assert.EqualError(t, wrap, "something wrong")

	frames := strings.Split(wrap.(ErrorTracer).Stack().Format(""), "\n")
	assert.True(t, strings.HasSuffix(frames[0], "errors_test.go:82"))
}

func TestComparable(t *testing.T) {
	e1 := New("something wrong")
	e2 := New("something wrong")
	assert.True(t, e1 == e1)
	assert.False(t, e1 == e2)
}
