package log

import (
	"context"
	"fmt"

	"github.com/ynjgit/erpc-go/protocol"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Setup the log setup
func Setup(cfgs []WriterConfig) error {
	var cores []zapcore.Core
	for _, cfg := range cfgs {
		writer := GetWriter(cfg.Writer)
		if writer == nil {
			return fmt.Errorf("writer:%s not regitered", cfg.Writer)
		}

		core, err := writer.Setup(cfg)
		if err != nil {
			return fmt.Errorf("writer:%s setup err:%s", cfg.Writer, err)
		}

		cores = append(cores, core)
	}

	if len(cores) == 0 {
		fmt.Println("not log writer configured")
		return nil
	}

	logger := zap.New(zapcore.NewTee(cores...),
		zap.AddCallerSkip(2),
		zap.AddCaller())

	z := &zapLogger{l: logger}
	setLogger(z)
	return nil
}

// Debug ...
func Debug(args ...interface{}) {
	getLogger().Debug(args...)
}

// DebugContext ...
func DebugContext(ctx context.Context, args ...interface{}) {
	erpcCtx := protocol.GetCtx(ctx)
	if l, ok := erpcCtx.Logger.(Logger); ok {
		l.Debug(args...)
		return
	}

	getLogger().Debug(args...)
}

// Debugf ...
func Debugf(format string, args ...interface{}) {
	getLogger().Debugf(format, args...)
}

// DebugContextf ...
func DebugContextf(ctx context.Context, format string, args ...interface{}) {
	erpcCtx := protocol.GetCtx(ctx)
	if l, ok := erpcCtx.Logger.(Logger); ok {
		l.Debugf(format, args...)
		return
	}

	getLogger().Debugf(format, args...)
}

// Info ...
func Info(args ...interface{}) {
	getLogger().Info(args...)
}

// InfoContext ...
func InfoContext(ctx context.Context, args ...interface{}) {
	erpcCtx := protocol.GetCtx(ctx)
	if l, ok := erpcCtx.Logger.(Logger); ok {
		l.Info(args...)
		return
	}

	getLogger().Info(args...)
}

// Infof ...
func Infof(format string, args ...interface{}) {
	getLogger().Infof(format, args...)
}

// InfoContextf ...
func InfoContextf(ctx context.Context, format string, args ...interface{}) {
	erpcCtx := protocol.GetCtx(ctx)
	if l, ok := erpcCtx.Logger.(Logger); ok {
		l.Infof(format, args...)
		return
	}

	getLogger().Infof(format, args...)
}

// Warn ...
func Warn(args ...interface{}) {
	getLogger().Warn(args...)
}

// WarnContext ...
func WarnContext(ctx context.Context, args ...interface{}) {
	erpcCtx := protocol.GetCtx(ctx)
	if l, ok := erpcCtx.Logger.(Logger); ok {
		l.Warn(args...)
		return
	}

	getLogger().Warn(args...)
}

// Warnf ...
func Warnf(format string, args ...interface{}) {
	getLogger().Warnf(format, args...)

}

// WarnContextf ...
func WarnContextf(ctx context.Context, format string, args ...interface{}) {
	erpcCtx := protocol.GetCtx(ctx)
	if l, ok := erpcCtx.Logger.(Logger); ok {
		l.Warnf(format, args...)
		return
	}

	getLogger().Warnf(format, args...)
}

// Error ...
func Error(args ...interface{}) {
	getLogger().Error(args...)
}

// ErrorContext ...
func ErrorContext(ctx context.Context, args ...interface{}) {
	erpcCtx := protocol.GetCtx(ctx)
	if l, ok := erpcCtx.Logger.(Logger); ok {
		l.Error(args...)
		return
	}

	getLogger().Error(args...)
}

// Errorf ...
func Errorf(format string, args ...interface{}) {
	getLogger().Errorf(format, args...)
}

// ErrorContextf ...
func ErrorContextf(ctx context.Context, format string, args ...interface{}) {
	erpcCtx := protocol.GetCtx(ctx)
	if l, ok := erpcCtx.Logger.(Logger); ok {
		l.Errorf(format, args...)
		return
	}

	getLogger().Errorf(format, args...)
}

// Fatal ...
func Fatal(args ...interface{}) {
	getLogger().Fatal(args...)
}

// FatalContext ...
func FatalContext(ctx context.Context, args ...interface{}) {
	erpcCtx := protocol.GetCtx(ctx)
	if l, ok := erpcCtx.Logger.(Logger); ok {
		l.Fatal(args...)
		return
	}

	getLogger().Fatal(args...)
}

// Fatalf ...
func Fatalf(format string, args ...interface{}) {
	getLogger().Fatalf(format, args...)
}

// FatalContextf ...
func FatalContextf(ctx context.Context, format string, args ...interface{}) {
	erpcCtx := protocol.GetCtx(ctx)
	if l, ok := erpcCtx.Logger.(Logger); ok {
		l.Fatalf(format, args...)
		return
	}

	getLogger().Fatalf(format, args...)
}

// WithContextFields ...
func WithContextFields(ctx context.Context, fields ...string) context.Context {
	erpcCtx := protocol.GetCtx(ctx)
	logger, ok := erpcCtx.Logger.(Logger)
	if ok {
		logger = logger.WithFields(fields...)
	} else {
		logger = getLogger().WithFields(fields...)
	}
	erpcCtx.Logger = logger
	return protocol.SetCtx(ctx, erpcCtx)
}

// Sync ...
func Sync() {
	getLogger().Sync()
}
