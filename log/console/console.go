package console

import (
	"os"

	"github.com/ynjgit/erpc-go/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	log.RegisterWriter("console", &consoleWriter{})
}

type consoleWriter struct {
}

func (w *consoleWriter) Setup(wc log.WriterConfig) (zapcore.Core, error) {
	level, ok := log.LevelMap[wc.Level]
	if !ok {
		level = zap.InfoLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	if wc.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(os.Stdout),
		level,
	)

	return core, nil
}
