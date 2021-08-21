package server_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ynjgit/erpc-go/client"
	"github.com/ynjgit/erpc-go/core"
	"github.com/ynjgit/erpc-go/interceptor"
	_ "github.com/ynjgit/erpc-go/interceptor/debuglog"
	_ "github.com/ynjgit/erpc-go/interceptor/recovery"
	"github.com/ynjgit/erpc-go/log"
	_ "github.com/ynjgit/erpc-go/log/console"
	_ "github.com/ynjgit/erpc-go/log/file"
	"github.com/ynjgit/erpc-go/server"
)

func init() {
	configs := []log.WriterConfig{
		log.WriterConfig{
			Writer: "console",
			Level:  "debug",
			Format: "console",
		},
	}
	err := log.Setup(configs)
	if err != nil {
		panic(err)
	}
}

type testReq struct {
	A int
	B int
}

type testRsp struct {
	Num int
}

func SvcAddWrapper(ctx context.Context, svcImpl interface{}, reqBody []byte, f server.InterceptorFunc) (interface{}, error) {
	req := &testReq{}
	rsp := &testRsp{}

	chain, err := f(reqBody, req)
	if err != nil {
		return nil, err
	}

	handler := func(ctx context.Context, req interface{}, rsp interface{}) error {
		return svcImpl.(*testSvc).Add(ctx, req.(*testReq), rsp.(*testRsp))
	}

	err = chain.Handle(ctx, req, rsp, handler)
	return rsp, err
}

func SvcSubWrapper(ctx context.Context, svcImpl interface{}, reqBody []byte, f server.InterceptorFunc) (interface{}, error) {
	req := &testReq{}
	rsp := &testRsp{}

	chain, err := f(reqBody, req)
	if err != nil {
		return nil, err
	}

	handler := func(ctx context.Context, req interface{}, rsp interface{}) error {
		return svcImpl.(*testSvc).Sub(ctx, req.(*testReq), rsp.(*testRsp))
	}

	err = chain.Handle(ctx, req, rsp, handler)
	return rsp, err
}

type testSvc struct{}

func (s *testSvc) Add(ctx context.Context, req *testReq, rsp *testRsp) error {
	rsp.Num = req.A + req.B
	return nil
}

func (s *testSvc) Sub(ctx context.Context, req *testReq, rsp *testRsp) error {
	log.DebugContext(ctx, "in sub service")
	time.Sleep(100 * time.Millisecond)
	rsp.Num = req.A - req.B
	return nil
}

func printUinInterceptor(ctx context.Context, req interface{}, rsp interface{}, f interceptor.Handler) error {
	ctx = log.WithContextFields(ctx, "uin", "1000")
	return f(ctx, req, rsp)
}
func TestServer(t *testing.T) {
	s, err := server.New(
		server.WithAddress(":1323"),
		server.WithCore(core.GetCore("echo")),

		server.WithInterceptor(interceptor.GetServerInterceptor("recovery")),
		server.WithInterceptor(interceptor.GetServerInterceptor("debuglog")),
		server.WithInterceptor(printUinInterceptor))
	require.Nil(t, err)

	svc := &testSvc{}
	method := server.ServiceMethod{
		RPCName: "/add",
		Handle:  SvcAddWrapper,
		SvcImpl: svc,
	}
	s.AddServiceMethod(method)
	method = server.ServiceMethod{
		RPCName: "/sub",
		Handle:  SvcSubWrapper,
		SvcImpl: svc,
	}
	s.AddServiceMethod(method)

	go func() {
		err := s.Start()
		require.Nil(t, err)
	}()

	time.Sleep(1 * time.Second)

	testCaseList := []struct {
		rpc    string
		req    *testReq
		expect *testRsp
	}{
		{"/add", &testReq{A: 2, B: 1}, &testRsp{Num: 3}},
		{"/sub", &testReq{A: 2, B: 1}, &testRsp{Num: 1}},
	}

	// orginal http post
	for _, testCase := range testCaseList {
		reqBody, _ := json.Marshal(testCase.req)
		rsp, err := http.Post("http://127.0.0.1:1323"+testCase.rpc, "application/json", bytes.NewReader(reqBody))
		require.Nil(t, err)
		require.EqualValues(t, 200, rsp.StatusCode)
		rspBody, _ := ioutil.ReadAll(rsp.Body)
		defer rsp.Body.Close()

		tmp := &testRsp{}
		err = json.Unmarshal(rspBody, tmp)
		require.Nil(t, err)
		require.Equal(t, testCase.expect, tmp)
	}

	// client.DefaultClient Call
	for _, testCase := range testCaseList {
		ctx := context.Background()
		opts := []client.Option{
			client.WithRPCName(testCase.rpc),
			client.WithTarget("http://127.0.0.1:1323"),
			client.WithTimeOut(200 * time.Millisecond),
		}
		tmp := &testRsp{}
		err := client.DefaultClient.Call(ctx, testCase.req, tmp, opts...)
		require.Nil(t, err)
		require.Equal(t, testCase.expect, tmp)
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	err = s.Shutdown(ctx)
	require.Nil(t, err)
}
