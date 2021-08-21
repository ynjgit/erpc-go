package erpc

import (
	"fmt"

	"github.com/ynjgit/erpc-go/config"
	"github.com/ynjgit/erpc-go/core"
	"github.com/ynjgit/erpc-go/interceptor"
	"github.com/ynjgit/erpc-go/log"
	"github.com/ynjgit/erpc-go/server"
)

// NewServer new server
func NewServer(opts ...server.Option) (*server.Server, error) {
	err := config.ParseFile()
	if err != nil {
		return nil, err
	}
	logConfig := config.GetGlobalConfig().Log
	err = log.Setup(logConfig)
	if err != nil {
		return nil, err
	}

	svrConf := config.GetServerConfig()
	var interceptors []interceptor.Interceptor
	for _, name := range svrConf.Interceptor {
		i := interceptor.GetServerInterceptor(name)
		if i == nil {
			return nil, fmt.Errorf("server interceptor:%s not registered", name)
		}

		interceptors = append(interceptors, i)
	}

	svrCore := core.GetCore(svrConf.Core)
	if svrCore == nil {
		return nil, fmt.Errorf("server core:%s not registered", svrConf.Core)
	}

	newOpts := []server.Option{
		server.WithAddress(svrConf.Address),
		server.WithCore(svrCore),
		server.WithInterceptors(interceptors...),
	}
	newOpts = append(newOpts, opts...)
	return server.New(newOpts...)
}
