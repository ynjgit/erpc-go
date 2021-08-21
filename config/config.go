package config

import (
	"github.com/ynjgit/erpc-go/log"
)

// Config the config define
type Config struct {
	Server ServerConfig
	Client ClientConfig
	Log    []log.WriterConfig
}

// ServerConfig the server config define
type ServerConfig struct {
	App         string
	Server      string
	Core        string   // which core to use eg: echo, std, or your own registered core
	Address     string   // the server listen address
	Timeout     int      // the server handle timeout, ms
	Interceptor []string // the interceptors
}

// ClientConfig the client config define
type ClientConfig struct {
	Interceptor []string        // the interceptors
	Timeout     int             // the client global timeout, ms
	Remote      []*RemoteConfig // the client remote config
}

// RemoteConfig the client remote config define
type RemoteConfig struct {
	Name    string
	Target  string
	Timeout int // one remote config timeout, ms
}

// SetGlobalConfig ...
func SetGlobalConfig(c *Config) {
	globalConfig.Store(c)
}

// GetGlobalConfig ...
func GetGlobalConfig() *Config {
	return globalConfig.Load().(*Config)
}

// GetServerConfig ...
func GetServerConfig() ServerConfig {
	if c, ok := globalConfig.Load().(*Config); ok {
		return c.Server
	}

	return ServerConfig{}
}

// GetClientConfig ...
func GetClientConfig() ClientConfig {
	if c, ok := globalConfig.Load().(*Config); ok {
		return c.Client
	}
	return ClientConfig{}
}
