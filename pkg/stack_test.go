package errors

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackTrace(t *testing.T) {
	caller := func() {
		frames := strings.Split(recordStack().Format(""), "\n")
		assert.True(t, strings.HasSuffix(frames[0], "stack_test.go:16"))
	}

	caller()
}
