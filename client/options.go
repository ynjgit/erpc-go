package client

import (
	"time"

	"github.com/ynjgit/erpc-go/interceptor"
)

// Options the client options
type Options struct {
	SVCName     string
	RPCName     string
	Target      string
	TimeOut     time.Duration
	Intercetors interceptor.Chain
}

// Option with set options
type Option func(*Options)

// WithSVCName set remoute service name
func WithSVCName(svc string) Option {
	return func(o *Options) {
		o.SVCName = svc
	}
}

// WithRPCName set remoute rpc name
func WithRPCName(rpc string) Option {
	return func(o *Options) {
		o.RPCName = rpc
	}
}

// WithTarget set client target
func WithTarget(t string) Option {
	return func(o *Options) {
		o.Target = t
	}
}

// WithTimeOut set client timeout
func WithTimeOut(t time.Duration) Option {
	return func(o *Options) {
		o.TimeOut = t
	}
}
