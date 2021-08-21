package interceptor

import (
	"context"
)

// Handler the final handle function
type Handler func(ctx context.Context, req interface{}, rsp interface{}) error

// Interceptor ...
type Interceptor func(ctx context.Context, req interface{}, rsp interface{}, f Handler) error

// Chain the interceptor chain, process interceptor befor the final handler
type Chain []Interceptor

// Handle handle the interceptor chain
func (c Chain) Handle(ctx context.Context, req interface{}, rsp interface{}, f Handler) error {
	if len(c) == 0 {
		return f(ctx, req, rsp)
	}

	return c[0](ctx, req, rsp, func(ctx context.Context, req interface{}, rsp interface{}) error {
		return c[1:].Handle(ctx, req, rsp, f)
	})
}

var (
	srvInterceptors = make(map[string]Interceptor)
	cliInterceptors = make(map[string]Interceptor)
)

// Register register interceptors
func Register(name string, serverInterceptor Interceptor, clientInterceptor Interceptor) {
	srvInterceptors[name] = serverInterceptor
	cliInterceptors[name] = clientInterceptor
}

// GetServerInterceptor ...
func GetServerInterceptor(name string) Interceptor {
	return srvInterceptors[name]
}

// GetClientInterceptor ...
func GetClientInterceptor(name string) Interceptor {
	return cliInterceptors[name]
}
