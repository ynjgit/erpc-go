package interceptor_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ynjgit/erpc-go/interceptor"
)

func testHandle(ctx context.Context, req interface{}, rsp interface{}) error {
	fmt.Println("in testHandle")
	r := req.(*string)
	p := rsp.(*string)

	// echo back
	*p = *r
	return nil
}

func errorHandle(ctx context.Context, req interface{}, rsp interface{}) error {
	return fmt.Errorf("errorHanle")
}

func TestHanle(t *testing.T) {
	var req string
	var rsp string

	req = "hello"

	ctx := context.Background()
	var chain interceptor.Chain
	err := chain.Handle(ctx, &req, &rsp, testHandle)
	require.Nil(t, err)
	require.EqualValues(t, "hello", rsp)

	err = chain.Handle(ctx, &req, &rsp, errorHandle)
	require.NotNil(t, err)
	require.EqualValues(t, "errorHanle", err.Error())
}

func passInterceptor(ctx context.Context, req interface{}, rsp interface{}, f interceptor.Handler) error {
	fmt.Println("before passInterceptor")
	err := f(ctx, req, rsp)
	fmt.Println("after passInterceptor")
	return err
}

func changeInterceptor(ctx context.Context, req interface{}, rsp interface{}, f interceptor.Handler) error {
	fmt.Println("before changeInterceptor")
	err := f(ctx, req, rsp)
	fmt.Println("after changeInterceptor")

	if err != nil {
		return err
	}
	p := rsp.(*string)
	*p = "changed"
	return nil
}

func blockInterceptor(ctx context.Context, req interface{}, rsp interface{}, f interceptor.Handler) error {
	fmt.Println("before blockInterceptor")
	err := fmt.Errorf("blocked")
	fmt.Println("after blockInterceptor")

	return err
}

func TestChain(t *testing.T) {
	var req string
	var rsp string

	req = "hello"

	var chain interceptor.Chain
	chain = append(chain, passInterceptor)

	ctx := context.Background()
	err := chain.Handle(ctx, &req, &rsp, testHandle)
	require.Nil(t, err)
	require.EqualValues(t, "hello", rsp)

	chain = append(chain, changeInterceptor)
	err = chain.Handle(ctx, &req, &rsp, testHandle)
	require.Nil(t, err)
	require.EqualValues(t, "changed", rsp)

	chain = append(chain, blockInterceptor)
	err = chain.Handle(ctx, &req, &rsp, testHandle)
	require.NotNil(t, err)
	require.EqualValues(t, "blocked", err.Error())
}
