package core

// ServerOption ...
type ServerOption struct {
	ListenAddress string
	Handler       Handler
}

// ServerOpt ...
type ServerOpt func(o *ServerOption)

// WithListenAddress ...
func WithListenAddress(addr string) ServerOpt {
	return func(o *ServerOption) {
		o.ListenAddress = addr
	}
}

// WithHandler ...
func WithHandler(h Handler) ServerOpt {
	return func(o *ServerOption) {
		o.Handler = h
	}
}
