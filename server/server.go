package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/ynjgit/erpc-go/core"
	"github.com/ynjgit/erpc-go/errs"
	"github.com/ynjgit/erpc-go/interceptor"
	"github.com/ynjgit/erpc-go/log"
	"github.com/ynjgit/erpc-go/protocol"
)

// InterceptorFunc get server interceptor chain
type InterceptorFunc func(body []byte, req interface{}) (interceptor.Chain, error)

// Server the server
type Server struct {
	svcMethods map[string]ServiceMethod
	options    *Options
}

// New new server
func New(opts ...Option) (*Server, error) {
	s := &Server{
		svcMethods: make(map[string]ServiceMethod),
		options:    &Options{},
	}

	for _, opt := range opts {
		opt(s.options)
	}

	if s.options.core == nil {
		return nil, fmt.Errorf("server core not set")
	}

	return s, nil
}

// Start start the server and serve
func (s *Server) Start() error {
	ctx := context.Background()
	return s.options.core.Serve(ctx,
		core.WithListenAddress(s.options.address),
		core.WithHandler(s))
}

// Shutdown shutdown the server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.options.core.Shutdown(ctx)
}

// AddServiceMethod add service method handler
func (s *Server) AddServiceMethod(svcMethod ServiceMethod) {
	s.svcMethods[svcMethod.RPCName] = svcMethod
}

// Handle the server handle callback
func (s *Server) Handle(ctx context.Context, path string, reqBody []byte) (rspBody []byte, err error) {
	svc, ok := s.svcMethods[path]
	if !ok {
		code := http.StatusNotFound
		return nil, errs.New(code, http.StatusText(code))
	}

	// initail the erpc Ctx
	erpcCtx := protocol.GetCtx(ctx)
	erpcCtx.ServerRPCName = path
	ctx = protocol.SetCtx(ctx, erpcCtx)
	ctx = log.WithContextFields(ctx, "trace_id", genTraceID(), "service", erpcCtx.ServerRPCName)

	if s.options.timeout > 0 {
		var cancle context.CancelFunc
		ctx, cancle = context.WithTimeout(ctx, s.options.timeout)
		defer cancle()
	}

	rsp, err := svc.Handle(ctx, svc.SvcImpl, reqBody, s.interceptorFunc)
	if err != nil {
		return nil, err
	}

	// marshal rsp
	rspBody, err = protocol.Marshal(rsp)
	if err != nil {
		return nil, errs.New(errs.CodeServerMarshalFail, "server marshal fail:"+err.Error())
	}

	return
}

func (s *Server) interceptorFunc(body []byte, req interface{}) (interceptor.Chain, error) {
	// unmarshal the reqbody
	err := protocol.Unmarshal(body, req)
	if err != nil {
		return nil, errs.New(errs.CodeServerUnmarshalFail, "server unmarshal fail:"+err.Error())
	}

	return s.options.chain, nil
}

func genTraceID() string {
	id := uuid.New().String()
	return strings.ReplaceAll(id, "-", "")
}
