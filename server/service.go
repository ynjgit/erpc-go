package server

import (
	"context"
)

// ServiceMethod the service method
type ServiceMethod struct {
	RPCName string
	Handle  func(ctx context.Context, svcImpl interface{}, body []byte, f InterceptorFunc) (rsp interface{}, err error)
	SvcImpl interface{}
}
