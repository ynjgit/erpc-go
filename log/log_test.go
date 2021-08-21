package log_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/ynjgit/erpc-go/log"
	_ "github.com/ynjgit/erpc-go/log/console"
	_ "github.com/ynjgit/erpc-go/log/file"
)

const (
	s = `
log:
  - writer: console
    level: debug
  - writer: file
    level: info
    write_config:
      log_path: ./logs
      filename: test_server.log
      max_size: 10
      max_age: 7
      max_backups: 10`
)

func TestBuiltinConsole(t *testing.T) {
	configs := []log.WriterConfig{
		log.WriterConfig{
			Writer: "console",
			Level:  "debug",
			Format: "console",
		},
	}

	err := log.Setup(configs)
	require.Nil(t, err)
	log.Debug("1231231", "12312312")
}

func TestBuiltinFile(t *testing.T) {
	configs := []log.WriterConfig{
		log.WriterConfig{
			Writer: "file",
			Level:  "debug",
			Format: "console",
		},
	}

	err := log.Setup(configs)
	require.Nil(t, err)
	log.Debug("1231231", "12312312")
	log.Sync()
}

func TestConfig(t *testing.T) {
	type logConfig struct {
		Log []log.WriterConfig `yaml:"log"`
	}

	lc := &logConfig{}
	err := yaml.Unmarshal([]byte(s), lc)
	require.Nil(t, err)

	err = log.Setup(lc.Log)
	require.Nil(t, err)
	log.Debug("this is debug log")
	log.Debugf("this is debugf log %s", "debug")
	log.Info("this is info log")
	log.Infof("this is infof log %s", "info")
	log.Warn("this is warn log")
	log.Warnf("this is warnf log %s", "warn")
	log.Error("this is error log")
	log.Errorf("this is errorf log %s", "error")

	ctx := context.Background()
	log.DebugContext(ctx, "yes", 1, 2, 3)
	ctx = log.WithContextFields(ctx, "with", "me")
	log.DebugContext(ctx, "after with")
	log.DebugContext(ctx, "yes", 1, 2, 3)
}
