package config

import (
	"fmt"
	"io/ioutil"
	"sync/atomic"

	"github.com/ynjgit/erpc-go/interceptor"

	"gopkg.in/yaml.v3"
)

const (
	defaultConfigFile = "./erpc_go.yaml"
)

var (
	globalConfig atomic.Value
)

// ParseFile parse the config file
func ParseFile() error {
	data, err := ioutil.ReadFile(defaultConfigFile)
	if err != nil {
		return err
	}

	return Parse(data)
}

// Parse parse from bytes
func Parse(b []byte) error {
	c := &Config{}
	err := yaml.Unmarshal(b, c)
	if err != nil {
		return err
	}

	err = repaireConfig(c)
	if err != nil {
		return err
	}

	SetGlobalConfig(c)
	return nil
}

func repaireConfig(c *Config) error {
	srvConf := c.Server
	if srvConf.App == "" {
		return fmt.Errorf("config: server app is empty")
	}
	if srvConf.Server == "" {
		return fmt.Errorf("config: server server is empty")
	}
	if srvConf.Address == "" {
		return fmt.Errorf("config: server address is empty")
	}
	if srvConf.Core == "" {
		return fmt.Errorf("config: server core is empty")
	}

	cliConfig := c.Client
	for _, name := range cliConfig.Interceptor {
		i := interceptor.GetClientInterceptor(name)
		if i == nil {
			return fmt.Errorf("config: client interceptor:%s is not registered", name)

		}
	}

	for _, remote := range cliConfig.Remote {
		if remote.Target == "" {
			return fmt.Errorf("config: client %s target is empty", remote.Name)
		}

		if remote.Timeout == 0 && cliConfig.Timeout > 0 {
			remote.Timeout = cliConfig.Timeout
		}
	}
	return nil
}
