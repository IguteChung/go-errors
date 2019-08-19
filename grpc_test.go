package errors

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGrpcErrorfError(t *testing.T) {
	e := fmt.Errorf("something wrong")
	grpc := GrpcErrorf(codes.InvalidArgument, "At someplace: %v", e)
	assert.EqualError(t, grpc, "rpc error: code = InvalidArgument desc = At someplace: something wrong")

	frames := strings.Split(StackTrace(grpc), "\n")
	assert.True(t, strings.HasSuffix(frames[0], "grpc_test.go:15"))

	assert.Equal(t, codes.InvalidArgument, status.Code(grpc))
}

func TestGrpcErrorfStackError(t *testing.T) {
	e := New("something wrong")
	grpc := GrpcErrorf(codes.InvalidArgument, "At someplace: %v", e)
	assert.EqualError(t, grpc, "rpc error: code = InvalidArgument desc = At someplace: something wrong")

	frames := strings.Split(StackTrace(grpc), "\n")
	assert.True(t, strings.HasSuffix(frames[0], "grpc_test.go:25"))

	assert.Equal(t, codes.InvalidArgument, status.Code(grpc))
}

func TestGrpcError(t *testing.T) {
	grpc := GrpcError(codes.InvalidArgument, "something wrong")
	assert.EqualError(t, grpc, "rpc error: code = InvalidArgument desc = something wrong")

	frames := strings.Split(StackTrace(grpc), "\n")
	assert.True(t, strings.HasSuffix(frames[0], "grpc_test.go:36"))

	assert.Equal(t, codes.InvalidArgument, status.Code(grpc))
}
