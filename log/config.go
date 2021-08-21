package log

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
)

// WriterConfig the writer config
type WriterConfig struct {
	Writer string `yaml:"writer"`
	Level  string `yaml:"level"`
	Format string `yaml:"format"` // the log output format eg: console, json

	// the different write config
	WriteConfig yaml.Node `yaml:"write_config"`
}

// LevelMap level string to zap level
var LevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}
