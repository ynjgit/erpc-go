package client_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ynjgit/erpc-go/client"
)

func TestOptions(t *testing.T) {
	o := &client.Options{}

	rpcName := "test_service_name"
	opt := client.WithRPCName(rpcName)
	opt(o)
	require.EqualValues(t, rpcName, o.RPCName)

	target := "http://127.0.0.1:1323"
	opt = client.WithTarget(target)
	opt(o)
	require.EqualValues(t, target, o.Target)

	timeout := 1000 * time.Millisecond
	opt = client.WithTimeOut(timeout)
	opt(o)
	require.EqualValues(t, timeout, o.TimeOut)
}
