package server

import (
	"time"

	"github.com/ynjgit/erpc-go/core"
	"github.com/ynjgit/erpc-go/interceptor"
)

// Options server options
type Options struct {
	address string            // listen address
	core    core.Core         // the core imple the network
	chain   interceptor.Chain // the interceptor chain
	timeout time.Duration     // the server handle timeout
}

// Option the opt function
type Option func(*Options)

// WithAddress set the address
func WithAddress(addr string) Option {
	return func(o *Options) {
		o.address = addr
	}
}

// WithTimeout set the timeout
func WithTimeout(t time.Duration) Option {
	return func(o *Options) {
		o.timeout = t
	}
}

// WithCore set the core
func WithCore(c core.Core) Option {
	return func(o *Options) {
		o.core = c
	}
}

// WithInterceptor add the interceptor
func WithInterceptor(i interceptor.Interceptor) Option {
	return func(o *Options) {
		o.chain = append(o.chain, i)
	}
}

// WithInterceptors add the interceptors
func WithInterceptors(i ...interceptor.Interceptor) Option {
	return func(o *Options) {
		o.chain = append(o.chain, i...)
	}
}
