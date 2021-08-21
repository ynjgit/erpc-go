package config_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ynjgit/erpc-go/config"
)

var strConfig = `server:
app: app
server: helloworld
address: 127.0.0.1:5000
interceptors:
  - recovery
  - debuglog

client:
timeout: 1000
interceptors:
  - debuglog
remote:
  - name: erpc.app.helloworld.Test
	target: http://127.0.0.1:5000
	timeout: 500`

func TestConfig(t *testing.T) {
	err := config.Parse([]byte(strConfig))
	require.Nil(t, err)

	svr := config.GetServerConfig()
	require.EqualValues(t, "app", svr.App)
	require.EqualValues(t, "helloworld", svr.Server)
	require.EqualValues(t, "127.0.0.1:5000", svr.Address)
	require.Len(t, svr.Interceptor, 2)
	require.EqualValues(t, "recovery", svr.Interceptor[0])
	require.EqualValues(t, "debuglog", svr.Interceptor[1])

	cli := config.GetClientConfig()
	require.Len(t, cli.Interceptor, 1)
	require.EqualValues(t, "debuglog", cli.Interceptor[0])
}
