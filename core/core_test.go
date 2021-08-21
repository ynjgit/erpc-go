package core_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ynjgit/erpc-go/core"
)

type service struct{}

func (s *service) Handle(ctx context.Context, path string, reqBody []byte) (rspBody []byte, err error) {
	return reqBody, nil
}

func TestEchoCore(t *testing.T) {
	c := core.GetCore("echo")
	require.NotNil(t, c)

	test(c, t)
}

func TestStdCore(t *testing.T) {
	c := core.GetCore("std")
	require.NotNil(t, c)

	test(c, t)
}

func test(c core.Core, t *testing.T) {
	ctx := context.Background()
	s := &service{}
	go func() {
		c.Serve(ctx, core.WithListenAddress(":1323"), core.WithHandler(s))
	}()

	// wait core start
	time.Sleep(1 * time.Second)
	want := []byte("ok")
	rsp, err := http.Post("http://127.0.0.1:1323", "application/json", bytes.NewReader(want))
	require.Nil(t, err)
	require.Nil(t, err)
	require.EqualValues(t, 200, rsp.StatusCode)
	body, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	require.Nil(t, err)
	require.EqualValues(t, want, body)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	c.Shutdown(ctx)
}
