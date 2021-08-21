package main

import (
	"context"
	"fmt"

	rpc "test/helloworld"
)

type greeterImpl struct{}

// SayHello
func (s *greeterImpl) SayHello(ctx context.Context, req *rpc.HelloRequest, rsp *rpc.HelloReply) error {
	// your business code ...

	rsp.Msg = "this is SayHello"
	return nil
}

// SayHi
func (s *greeterImpl) SayHi(ctx context.Context, req *rpc.HelloRequest, rsp *rpc.HelloReply) error {
	// your business code ...

	c := rpc.NewGreeterClient()
	sumReq := &rpc.SumReq{
		A: 1,
		B: 2,
	}
	sumRsp, err := c.Sum(ctx, sumReq)
	if err != nil {
		return err
	}

	rsp.Msg = fmt.Sprintf("SayHi call sum return:%d", sumRsp.GetSum())
	return nil
}

// Sum
func (s *greeterImpl) Sum(ctx context.Context, req *rpc.SumReq, rsp *rpc.SumRsp) error {
	// your business code ...

	rsp.Sum = req.A + req.B
	return nil
}

// TestPanic
func (s *greeterImpl) TestPanic(ctx context.Context, req *rpc.HelloRequest, rsp *rpc.HelloReply) error {
	// your business code ...

	return nil
}
