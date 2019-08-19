package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// grpcTraceableError implements a traceable gRPC status error.
type grpcTraceableError struct {
	traceableError
}

// GRPCStatus implements gRPC error.
func (g grpcTraceableError) GRPCStatus() *status.Status {
	s, ok := status.FromError(g.err)
	if !ok {
		return nil
	}
	return s
}

// GrpcError creates a traceable error with message and gRPC status code.
func GrpcError(code codes.Code, msg string) error {
	return grpcTraceableError{traceableError{
		err:   status.Error(code, msg),
		stack: recordStack(),
	}}
}

// GrpcErrorf creates a traceable error with formatted message and gRPC status code.
// If args contains a traceable error, apply the stack of it.
// If args doesn't contain a traceable error, record the stack trace.
func GrpcErrorf(code codes.Code, msg string, args ...interface{}) error {
	e := status.Errorf(code, msg, args...)
	for _, arg := range args {
		if t, ok := arg.(traceableError); ok {
			return grpcTraceableError{traceableError{
				err:   e,
				stack: t.Stack(),
			}}
		}
	}
	return grpcTraceableError{traceableError{
		err:   e,
		stack: recordStack(),
	}}
}
