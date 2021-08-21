package core

import (
	"context"
)

// Handler the service handler callback
type Handler interface {
	Handle(ctx context.Context, path string, req []byte) (rsp []byte, err error)
}

// Core the http impl core
type Core interface {
	Serve(ctx context.Context, opts ...ServerOpt) error
	Shutdown(ctx context.Context) error
}

var (
	cores = make(map[string]Core)
)

// Register core register
func Register(name string, core Core) {
	cores[name] = core
}

// GetCore ...
func GetCore(name string) Core {
	return cores[name]
}

// ErrRsp the default err response
type ErrRsp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
