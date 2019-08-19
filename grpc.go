package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// grpcTracer implements a gRPC status error which is also an errorTracer.
type grpcTracer struct {
	tracer
}

// GRPCStatus implements gRPC error.
func (g grpcTracer) GRPCStatus() *status.Status {
	s, ok := status.FromError(g.err)
	if !ok {
		return nil
	}
	return s
}

// GrpcError creates an errorTracer with error message and gRPC status code.
func GrpcError(code codes.Code, msg string) error {
	return grpcTracer{tracer{
		err:   status.Error(code, msg),
		stack: recordStack(),
	}}
}

// GrpcErrorf creates an errorTracer with formatted message and gRPC status code.
// If args contains an errorTracer, apply the stack of it.
// If args doesn't contain an errorTracer, record the stack trace.
func GrpcErrorf(code codes.Code, msg string, args ...interface{}) error {
	e := status.Errorf(code, msg, args...)
	for _, arg := range args {
		if t, ok := arg.(errorTracer); ok {
			return grpcTracer{tracer{
				err:   e,
				stack: t.Stack(),
			}}
		}
	}
	return grpcTracer{tracer{
		err:   e,
		stack: recordStack(),
	}}
}
