package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// grpcTracer implements a gRPC status error which is also an ErrorTracer.
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

// GrpcError creates an ErrorTracer with error message and gRPC status code.
func GrpcError(code codes.Code, msg string) error {
	return grpcTracer{tracer{
		err:   status.Error(code, msg),
		stack: recordStack(),
	}}
}

// GrpcErrorf creates an ErrorTracer with formatted message and gRPC status code.
// If args contains an ErrorTracer, apply the stack of it.
// If args doesn't contain an ErrorTracer, record the stack trace.
func GrpcErrorf(code codes.Code, msg string, args ...interface{}) error {
	e := status.Errorf(code, msg, args...)
	for _, arg := range args {
		if t, ok := arg.(ErrorTracer); ok {
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
