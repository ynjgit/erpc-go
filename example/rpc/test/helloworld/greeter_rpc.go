package helloworld

import (
	"context"

	"github.com/ynjgit/erpc-go/client"
	"github.com/ynjgit/erpc-go/server"
)

// Greeter  Greeter
type Greeter interface {
    
    // SayHello 
    SayHello(ctx context.Context, req *HelloRequest, rsp *HelloReply) error
    
    // SayHi 
    SayHi(ctx context.Context, req *HelloRequest, rsp *HelloReply) error
    
    // Sum 
    Sum(ctx context.Context, req *SumReq, rsp *SumRsp) error
    
    // TestPanic 
    TestPanic(ctx context.Context, req *HelloRequest, rsp *HelloReply) error
    
}


func wrapperHandleSayHello(ctx context.Context, svcImpl interface{}, reqBody []byte, f server.InterceptorFunc) (interface{}, error) {
    req := &HelloRequest{}
    rsp := &HelloReply{}

    chain, err := f(reqBody, req)
	if err != nil {
		return nil, err
	}

	handler := func(ctx context.Context, req interface{}, rsp interface{}) error {
		return svcImpl.(Greeter).SayHello(ctx, req.(*HelloRequest), rsp.(*HelloReply))
	}

	err = chain.Handle(ctx, req, rsp, handler)
	return rsp, err
}

func wrapperHandleSayHi(ctx context.Context, svcImpl interface{}, reqBody []byte, f server.InterceptorFunc) (interface{}, error) {
    req := &HelloRequest{}
    rsp := &HelloReply{}

    chain, err := f(reqBody, req)
	if err != nil {
		return nil, err
	}

	handler := func(ctx context.Context, req interface{}, rsp interface{}) error {
		return svcImpl.(Greeter).SayHi(ctx, req.(*HelloRequest), rsp.(*HelloReply))
	}

	err = chain.Handle(ctx, req, rsp, handler)
	return rsp, err
}

func wrapperHandleSum(ctx context.Context, svcImpl interface{}, reqBody []byte, f server.InterceptorFunc) (interface{}, error) {
    req := &SumReq{}
    rsp := &SumRsp{}

    chain, err := f(reqBody, req)
	if err != nil {
		return nil, err
	}

	handler := func(ctx context.Context, req interface{}, rsp interface{}) error {
		return svcImpl.(Greeter).Sum(ctx, req.(*SumReq), rsp.(*SumRsp))
	}

	err = chain.Handle(ctx, req, rsp, handler)
	return rsp, err
}

func wrapperHandleTestPanic(ctx context.Context, svcImpl interface{}, reqBody []byte, f server.InterceptorFunc) (interface{}, error) {
    req := &HelloRequest{}
    rsp := &HelloReply{}

    chain, err := f(reqBody, req)
	if err != nil {
		return nil, err
	}

	handler := func(ctx context.Context, req interface{}, rsp interface{}) error {
		return svcImpl.(Greeter).TestPanic(ctx, req.(*HelloRequest), rsp.(*HelloReply))
	}

	err = chain.Handle(ctx, req, rsp, handler)
	return rsp, err
}


var GreeterMethods = []server.ServiceMethod{
    
    server.ServiceMethod{
        RPCName: "/erpc.app.helloworld.Greeter/SayHello",
        Handle: wrapperHandleSayHello,
    },
    
    server.ServiceMethod{
        RPCName: "/erpc.app.helloworld.Greeter/SayHi",
        Handle: wrapperHandleSayHi,
    },
    
    server.ServiceMethod{
        RPCName: "/erpc.app.helloworld.Greeter/Sum",
        Handle: wrapperHandleSum,
    },
    
    server.ServiceMethod{
        RPCName: "/erpc.app.helloworld.Greeter/TestPanic",
        Handle: wrapperHandleTestPanic,
    },
    
}

func RegisterGreeterRPC(s *server.Server, svcImpl interface{}) {
    for _, svcMethod := range GreeterMethods {
        svcMethod.SvcImpl = svcImpl
        s.AddServiceMethod(svcMethod)
    }    
}

type GreeterClient interface {
    
    // SayHello 
    SayHello(ctx context.Context, req *HelloRequest, opts ...client.Option) (*HelloReply, error)
    
    // SayHi 
    SayHi(ctx context.Context, req *HelloRequest, opts ...client.Option) (*HelloReply, error)
    
    // Sum 
    Sum(ctx context.Context, req *SumReq, opts ...client.Option) (*SumRsp, error)
    
    // TestPanic 
    TestPanic(ctx context.Context, req *HelloRequest, opts ...client.Option) (*HelloReply, error)
    
}

func NewGreeterClient(opts ...client.Option) GreeterClient {
    return &greeterClientImpl{
		opts: opts,
	}
}

type greeterClientImpl struct {
	opts []client.Option
}


func (c *greeterClientImpl) SayHello(ctx context.Context, req *HelloRequest, opts ...client.Option) (*HelloReply, error) {
    rsp := &HelloReply{}
    callopts := make([]client.Option, 0, len(c.opts)+len(opts)+2)
    callopts = append(callopts, c.opts...)
    callopts = append(callopts, client.WithSVCName("erpc.app.helloworld.Greeter"))
    callopts = append(callopts, client.WithRPCName("/erpc.app.helloworld.Greeter/SayHello"))
    callopts = append(callopts, opts...)
    err := client.DefaultClient.Call(ctx, req, rsp, callopts...)
    return rsp, err
}

func (c *greeterClientImpl) SayHi(ctx context.Context, req *HelloRequest, opts ...client.Option) (*HelloReply, error) {
    rsp := &HelloReply{}
    callopts := make([]client.Option, 0, len(c.opts)+len(opts)+2)
    callopts = append(callopts, c.opts...)
    callopts = append(callopts, client.WithSVCName("erpc.app.helloworld.Greeter"))
    callopts = append(callopts, client.WithRPCName("/erpc.app.helloworld.Greeter/SayHi"))
    callopts = append(callopts, opts...)
    err := client.DefaultClient.Call(ctx, req, rsp, callopts...)
    return rsp, err
}

func (c *greeterClientImpl) Sum(ctx context.Context, req *SumReq, opts ...client.Option) (*SumRsp, error) {
    rsp := &SumRsp{}
    callopts := make([]client.Option, 0, len(c.opts)+len(opts)+2)
    callopts = append(callopts, c.opts...)
    callopts = append(callopts, client.WithSVCName("erpc.app.helloworld.Greeter"))
    callopts = append(callopts, client.WithRPCName("/erpc.app.helloworld.Greeter/Sum"))
    callopts = append(callopts, opts...)
    err := client.DefaultClient.Call(ctx, req, rsp, callopts...)
    return rsp, err
}

func (c *greeterClientImpl) TestPanic(ctx context.Context, req *HelloRequest, opts ...client.Option) (*HelloReply, error) {
    rsp := &HelloReply{}
    callopts := make([]client.Option, 0, len(c.opts)+len(opts)+2)
    callopts = append(callopts, c.opts...)
    callopts = append(callopts, client.WithSVCName("erpc.app.helloworld.Greeter"))
    callopts = append(callopts, client.WithRPCName("/erpc.app.helloworld.Greeter/TestPanic"))
    callopts = append(callopts, opts...)
    err := client.DefaultClient.Call(ctx, req, rsp, callopts...)
    return rsp, err
}
